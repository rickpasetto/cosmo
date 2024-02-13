import { and, eq, inArray, or, sql } from 'drizzle-orm';
import { PostgresJsDatabase } from 'drizzle-orm/postgres-js';
import { NewSchemaChangeOperationUsage } from '../../db/models.js';
import * as schema from '../../db/schema.js';
import {
  schemaCheckChangeAction,
  schemaCheckChangeActionOperationUsage,
  schemaCheckComposition,
  schemaChecks,
} from '../../db/schema.js';
import { ComposedFederatedGraph } from '../composition/composer.js';
import { SchemaDiff } from '../composition/schemaCheck.js';
import { InspectorOperationResult } from '../services/SchemaUsageTrafficInspector.js';
import { FederatedGraphConfig } from './FederatedGraphRepository.js';

export class SchemaCheckRepository {
  constructor(private db: PostgresJsDatabase<typeof schema>) {}

  public async create(data: {
    targetId: string;
    isComposable?: boolean;
    isDeleted?: boolean;
    proposedSubgraphSchemaSDL: string;
  }): Promise<string> {
    const insertedSchemaCheck = await this.db
      .insert(schemaChecks)
      .values({
        targetId: data.targetId,
        isComposable: data.isComposable,
        isDeleted: data.isDeleted,
        proposedSubgraphSchemaSDL: data.proposedSubgraphSchemaSDL,
      })
      .returning()
      .execute();
    return insertedSchemaCheck[0].id;
  }

  public async update(data: {
    schemaCheckID: string;
    isComposable?: boolean;
    hasClientTraffic?: boolean;
    hasBreakingChanges?: boolean;
  }): Promise<string | undefined> {
    const updatedSchemaCheck = await this.db
      .update(schemaChecks)
      .set({
        isComposable: data.isComposable,
        hasBreakingChanges: data.hasBreakingChanges,
        hasClientTraffic: data.hasClientTraffic,
      })
      .where(eq(schemaChecks.id, data.schemaCheckID))
      .returning()
      .execute();
    return updatedSchemaCheck[0].id;
  }

  public createSchemaCheckChanges(data: { schemaCheckID: string; changes: SchemaDiff[] }) {
    if (data.changes.length === 0) {
      return [];
    }
    return this.db
      .insert(schemaCheckChangeAction)
      .values(
        data.changes.map((change) => ({
          schemaCheckId: data.schemaCheckID,
          changeType: change.changeType,
          changeMessage: change.message,
          path: change.path,
          isBreaking: change.isBreaking,
        })),
      )
      .returning();
  }

  public async createOperationUsage(
    schemaCheckActionOperations: Map<string, InspectorOperationResult[]>,
    federatedGraphId: string,
  ) {
    const values: NewSchemaChangeOperationUsage[] = [];

    for (const [schemaCheckChangeActionId, operations] of schemaCheckActionOperations.entries()) {
      values.push(
        ...operations.map(
          (op) =>
            ({
              schemaCheckChangeActionId,
              name: op.name,
              type: op.type,
              hash: op.hash,
              firstSeenAt: op.firstSeenAt,
              lastSeenAt: op.lastSeenAt,
              federatedGraphId,
              isSafeOverride: op.isSafeOverride,
            }) as NewSchemaChangeOperationUsage,
        ),
      );
    }

    if (values.length === 0) {
      return;
    }

    await this.db.insert(schemaCheckChangeActionOperationUsage).values(values).execute();
  }

  private mapChangesFromDriverValue = (val: any) => {
    if (typeof val === 'string' && val.length > 0 && val !== '{}') {
      const pairs = val.slice(2, -2).split('","');

      return pairs.map((pair) => {
        const [changeType, path] = pair.slice(1, -1).split(',');
        return { changeType, path };
      });
    }
    return [];
  };

  public async checkClientTrafficAgainstOverrides(data: {
    changes: { id: string; changeType: string | null; path: string | null }[];
    inspectorResultsByChangeId: Map<string, InspectorOperationResult[]>;
    federatedGraphId: string;
  }) {
    let hasUnsafeClientTraffic = false;

    const result = new Map(data.inspectorResultsByChangeId);

    const changeActionsByOperationHash: Map<string, typeof data.changes> = new Map();

    for (const [schemaCheckChangeId, operationResults] of result.entries()) {
      for (const operationResult of operationResults) {
        const { hash } = operationResult;
        if (!changeActionsByOperationHash.has(hash)) {
          changeActionsByOperationHash.set(hash, []);
        }

        const change = data.changes.find((c) => c.id === schemaCheckChangeId);
        if (change) {
          changeActionsByOperationHash.get(hash)?.push(change);
        }
      }
    }

    for (const [hash, changes] of changeActionsByOperationHash) {
      const incomingChanges = `array[${changes.map((c) => `('${c.changeType}'::schema_change_type, '${c.path}'::text)`)}]`;
      const storedChanges = `array_agg(distinct (${schema.schemaCheckChangeAction.changeType.name}, ${schema.schemaCheckChangeAction.path.name}))`;
      const ignoreSql = `bool_or(${schema.operationOverrides.ignoreAll.name})`;

      const res = await this.db
        .select({
          ignoreAll: sql.raw(ignoreSql).mapWith(Boolean),
          unsafeChanges: sql
            .raw(`array(select unnest(${incomingChanges}) except select unnest(${storedChanges}))`)
            .mapWith({
              mapFromDriverValue: this.mapChangesFromDriverValue,
            }),
          safeChanges: sql
            .raw(`array(select unnest(${incomingChanges}) intersect select unnest(${storedChanges}))`)
            .mapWith({
              mapFromDriverValue: this.mapChangesFromDriverValue,
            }),
        })
        .from(schema.schemaCheckChangeAction)
        .leftJoin(
          schema.operationOverrides,
          eq(schema.schemaCheckChangeAction.schemaCheckId, schema.operationOverrides.schemaCheckId),
        )
        .leftJoin(
          schema.schemaCheckFederatedGraphs,
          eq(schema.schemaCheckChangeAction.schemaCheckId, schema.schemaCheckFederatedGraphs.checkId),
        )
        .where(
          and(
            eq(schema.operationOverrides.hash, hash),
            eq(schema.schemaCheckFederatedGraphs.federatedGraphId, data.federatedGraphId),
            eq(schema.schemaCheckChangeAction.isBreaking, true),
          ),
        )
        .groupBy(schema.operationOverrides.hash)
        .having(or(sql.raw(`${storedChanges} @> ${incomingChanges}`), sql.raw(ignoreSql)));

      if (res.length === 0) {
        // If no safe overrides are found, then mark traffic as unsafe
        hasUnsafeClientTraffic = true;
        continue;
      }

      const safeChanges = res[0].ignoreAll ? changes : res[0].safeChanges;

      for (const safeChange of safeChanges) {
        const change = changes.find((c) => c.changeType === safeChange.changeType && c.path === safeChange.path);
        if (!change) {
          continue;
        }

        const op = result.get(change.id)?.find((c) => c.hash === hash);
        if (!op) {
          continue;
        }

        op.isSafeOverride = true;
      }

      if (!res[0].ignoreAll && res[0].unsafeChanges.length > 0) {
        hasUnsafeClientTraffic = true;
      }
    }

    return { hasUnsafeClientTraffic, result };
  }

  public async createCheckedFederatedGraph(schemaCheckId: string, federatedGraphId: string, trafficCheckDays: number) {
    await this.db
      .insert(schema.schemaCheckFederatedGraphs)
      .values({
        checkId: schemaCheckId,
        federatedGraphId,
        trafficCheckDays,
      })
      .execute();
  }

  public async getAffectedOperationsByCheckId(checkId: string) {
    const changeActionIds = (
      await this.db.query.schemaCheckChangeAction.findMany({
        where: and(
          eq(schema.schemaCheckChangeAction.schemaCheckId, checkId),
          eq(schema.schemaCheckChangeAction.isBreaking, true),
        ),
        columns: {
          id: true,
        },
      })
    ).map((r) => r.id);

    if (changeActionIds.length > 0) {
      return await this.db
        .selectDistinctOn([schema.schemaCheckChangeActionOperationUsage.hash], {
          hash: schema.schemaCheckChangeActionOperationUsage.hash,
          name: schema.schemaCheckChangeActionOperationUsage.name,
          type: schema.schemaCheckChangeActionOperationUsage.type,
          firstSeenAt: sql<Date>`min(${schema.schemaCheckChangeActionOperationUsage.firstSeenAt})`,
          lastSeenAt: sql<Date>`max(${schema.schemaCheckChangeActionOperationUsage.lastSeenAt})`,
          schemaChangeIds: sql<
            string[]
          >`array_agg(${schema.schemaCheckChangeActionOperationUsage.schemaCheckChangeActionId})`,
          isSafe: sql<boolean>`true = all(array_agg(${schema.schemaCheckChangeActionOperationUsage.isSafeOverride}))`,
        })
        .from(schema.schemaCheckChangeActionOperationUsage)
        .where(inArray(schema.schemaCheckChangeActionOperationUsage.schemaCheckChangeActionId, changeActionIds))
        .groupBy(({ hash, name, type }) => [hash, name, type]);
    }

    return [];
  }

  public createSchemaCheckCompositions(data: { schemaCheckID: string; compositions: ComposedFederatedGraph[] }) {
    if (data.compositions.length === 0) {
      return;
    }

    return this.db.transaction(async (tx) => {
      // let's check if the subgraph change has produced any composition error.
      // In that case, we will mark all checks as not composable
      const hasCompositionErrors = data.compositions.some((composition) => composition.errors.length > 0);

      await tx
        .insert(schemaCheckComposition)
        .values(
          data.compositions.map((composition) => ({
            federatedTargetId: composition.targetID,
            schemaCheckId: data.schemaCheckID,
            composedSchemaSDL: composition.composedSchema,
            compositionErrors: composition.errors?.map((e) => e.toString()).join('\n'),
          })),
        )
        .execute();

      // update the isComposable column in schema_checks table
      await tx
        .update(schemaChecks)
        .set({
          isComposable: !hasCompositionErrors,
        })
        .where(eq(schemaChecks.id, data.schemaCheckID))
        .returning()
        .execute();
    });
  }

  public async getFederatedGraphConfigForCheckId(
    checkId: string,
    federatedGraphId: string,
  ): Promise<FederatedGraphConfig> {
    const result = await this.db.query.schemaCheckFederatedGraphs.findFirst({
      where: and(
        eq(schema.schemaCheckFederatedGraphs.checkId, checkId),
        eq(schema.schemaCheckFederatedGraphs.federatedGraphId, federatedGraphId),
      ),
    });

    return {
      trafficCheckDays: result?.trafficCheckDays ?? 7,
    };
  }
}
