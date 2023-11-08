// https://protobuf.dev/programming-guides/style/

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: wg/cosmo/graphqlmetrics/v1/graphqlmetrics.proto

package graphqlmetricsv1

import (
	_ "github.com/wundergraph/cosmo/router/gen/proto/wg/cosmo/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OperationType int32

const (
	OperationType_QUERY        OperationType = 0
	OperationType_MUTATION     OperationType = 1
	OperationType_SUBSCRIPTION OperationType = 2
)

// Enum value maps for OperationType.
var (
	OperationType_name = map[int32]string{
		0: "QUERY",
		1: "MUTATION",
		2: "SUBSCRIPTION",
	}
	OperationType_value = map[string]int32{
		"QUERY":        0,
		"MUTATION":     1,
		"SUBSCRIPTION": 2,
	}
)

func (x OperationType) Enum() *OperationType {
	p := new(OperationType)
	*p = x
	return p
}

func (x OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_enumTypes[0].Descriptor()
}

func (OperationType) Type() protoreflect.EnumType {
	return &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_enumTypes[0]
}

func (x OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OperationType.Descriptor instead.
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{0}
}

type RequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32 `protobuf:"varint,1,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	Error      bool  `protobuf:"varint,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *RequestInfo) Reset() {
	*x = RequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestInfo) ProtoMessage() {}

func (x *RequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestInfo.ProtoReflect.Descriptor instead.
func (*RequestInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{0}
}

func (x *RequestInfo) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RequestInfo) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

type SchemaUsageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestDocument is the GraphQL request document
	RequestDocument string `protobuf:"bytes,1,opt,name=RequestDocument,proto3" json:"RequestDocument,omitempty"`
	// TypeFieldMetrics is the list of used fields in the request document
	TypeFieldMetrics []*TypeFieldUsageInfo `protobuf:"bytes,2,rep,name=TypeFieldMetrics,proto3" json:"TypeFieldMetrics,omitempty"`
	// OperationInfo is the operation info
	OperationInfo *OperationInfo `protobuf:"bytes,3,opt,name=OperationInfo,proto3" json:"OperationInfo,omitempty"`
	// SchemaInfo is the schema info
	SchemaInfo *SchemaInfo `protobuf:"bytes,4,opt,name=SchemaInfo,proto3" json:"SchemaInfo,omitempty"`
	// ClientInfo is the client info
	ClientInfo *ClientInfo `protobuf:"bytes,5,opt,name=ClientInfo,proto3" json:"ClientInfo,omitempty"`
	// RequestInfo is the request info
	RequestInfo *RequestInfo `protobuf:"bytes,6,opt,name=RequestInfo,proto3" json:"RequestInfo,omitempty"`
	// Attributes is a map of attributes that can be used to filter the metrics
	Attributes map[string]string `protobuf:"bytes,7,rep,name=Attributes,proto3" json:"Attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SchemaUsageInfo) Reset() {
	*x = SchemaUsageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaUsageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaUsageInfo) ProtoMessage() {}

func (x *SchemaUsageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaUsageInfo.ProtoReflect.Descriptor instead.
func (*SchemaUsageInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{1}
}

func (x *SchemaUsageInfo) GetRequestDocument() string {
	if x != nil {
		return x.RequestDocument
	}
	return ""
}

func (x *SchemaUsageInfo) GetTypeFieldMetrics() []*TypeFieldUsageInfo {
	if x != nil {
		return x.TypeFieldMetrics
	}
	return nil
}

func (x *SchemaUsageInfo) GetOperationInfo() *OperationInfo {
	if x != nil {
		return x.OperationInfo
	}
	return nil
}

func (x *SchemaUsageInfo) GetSchemaInfo() *SchemaInfo {
	if x != nil {
		return x.SchemaInfo
	}
	return nil
}

func (x *SchemaUsageInfo) GetClientInfo() *ClientInfo {
	if x != nil {
		return x.ClientInfo
	}
	return nil
}

func (x *SchemaUsageInfo) GetRequestInfo() *RequestInfo {
	if x != nil {
		return x.RequestInfo
	}
	return nil
}

func (x *SchemaUsageInfo) GetAttributes() map[string]string {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the GraphQL client name obtained from the request header
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	// Version is the GraphQL client version obtained from the request header
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{2}
}

func (x *ClientInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type OperationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash is the hash of the request document and the operation name
	Hash string `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	// Name is the operation name
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	// Type is the operation type
	Type OperationType `protobuf:"varint,3,opt,name=Type,proto3,enum=wg.cosmo.graphqlmetrics.v1.OperationType" json:"Type,omitempty"`
}

func (x *OperationInfo) Reset() {
	*x = OperationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationInfo) ProtoMessage() {}

func (x *OperationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationInfo.ProtoReflect.Descriptor instead.
func (*OperationInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{3}
}

func (x *OperationInfo) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *OperationInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OperationInfo) GetType() OperationType {
	if x != nil {
		return x.Type
	}
	return OperationType_QUERY
}

type SchemaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version is the schema version
	Version string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *SchemaInfo) Reset() {
	*x = SchemaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchemaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchemaInfo) ProtoMessage() {}

func (x *SchemaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchemaInfo.ProtoReflect.Descriptor instead.
func (*SchemaInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{4}
}

func (x *SchemaInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type TypeFieldUsageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path is the path to the field in the request document but without the root type query, mutation, or subscription
	Path []string `protobuf:"bytes,1,rep,name=Path,proto3" json:"Path,omitempty"`
	// TypeNames is the list of type names that the field is used as
	TypeNames []string `protobuf:"bytes,2,rep,name=TypeNames,proto3" json:"TypeNames,omitempty"`
	// SubgraphIDs is the list of datasource IDs (e.g subgraph ID) that the field is used from
	SubgraphIDs []string `protobuf:"bytes,3,rep,name=SubgraphIDs,proto3" json:"SubgraphIDs,omitempty"`
	// Count is the number of times the field is used
	Count uint64 `protobuf:"varint,4,opt,name=Count,proto3" json:"Count,omitempty"`
	// NamedType is the underlying type of the field
	NamedType string `protobuf:"bytes,5,opt,name=NamedType,proto3" json:"NamedType,omitempty"`
}

func (x *TypeFieldUsageInfo) Reset() {
	*x = TypeFieldUsageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeFieldUsageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeFieldUsageInfo) ProtoMessage() {}

func (x *TypeFieldUsageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeFieldUsageInfo.ProtoReflect.Descriptor instead.
func (*TypeFieldUsageInfo) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{5}
}

func (x *TypeFieldUsageInfo) GetPath() []string {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *TypeFieldUsageInfo) GetTypeNames() []string {
	if x != nil {
		return x.TypeNames
	}
	return nil
}

func (x *TypeFieldUsageInfo) GetSubgraphIDs() []string {
	if x != nil {
		return x.SubgraphIDs
	}
	return nil
}

func (x *TypeFieldUsageInfo) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TypeFieldUsageInfo) GetNamedType() string {
	if x != nil {
		return x.NamedType
	}
	return ""
}

type PublishGraphQLRequestMetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SchemaUsage []*SchemaUsageInfo `protobuf:"bytes,1,rep,name=SchemaUsage,proto3" json:"SchemaUsage,omitempty"`
}

func (x *PublishGraphQLRequestMetricsRequest) Reset() {
	*x = PublishGraphQLRequestMetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishGraphQLRequestMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishGraphQLRequestMetricsRequest) ProtoMessage() {}

func (x *PublishGraphQLRequestMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishGraphQLRequestMetricsRequest.ProtoReflect.Descriptor instead.
func (*PublishGraphQLRequestMetricsRequest) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{6}
}

func (x *PublishGraphQLRequestMetricsRequest) GetSchemaUsage() []*SchemaUsageInfo {
	if x != nil {
		return x.SchemaUsage
	}
	return nil
}

type PublishOperationCoverageReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PublishOperationCoverageReportResponse) Reset() {
	*x = PublishOperationCoverageReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishOperationCoverageReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishOperationCoverageReportResponse) ProtoMessage() {}

func (x *PublishOperationCoverageReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishOperationCoverageReportResponse.ProtoReflect.Descriptor instead.
func (*PublishOperationCoverageReportResponse) Descriptor() ([]byte, []int) {
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP(), []int{7}
}

var File_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto protoreflect.FileDescriptor

var file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x77, 0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x77,
	0x67, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0xdf, 0x04, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x5a,
	0x0a, 0x10, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x4f, 0x0a, 0x0d, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x0a, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73,
	0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x49, 0x0a, 0x0b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x5b, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x77, 0x67, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x3a, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x76,
	0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x9c,
	0x01, 0x0a, 0x12, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x79, 0x70,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x49, 0x44, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x75,
	0x62, 0x67, 0x72, 0x61, 0x70, 0x68, 0x49, 0x44, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x74, 0x0a,
	0x23, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x67, 0x2e, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x26, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x3a, 0x0a,
	0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09,
	0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4d, 0x55, 0x54,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x55, 0x42, 0x53, 0x43,
	0x52, 0x49, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x32, 0xb8, 0x01, 0x0a, 0x15, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x51, 0x4c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x3f, 0x2e,
	0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x47, 0x72, 0x61, 0x70, 0x68, 0x51, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x42,
	0x2e, 0x77, 0x67, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71,
	0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x9b, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x67, 0x2e,
	0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x13, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x59,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x6e, 0x64, 0x65,
	0x72, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x77, 0x67,
	0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x43, 0x47, 0xaa,
	0x02, 0x1a, 0x57, 0x67, 0x2e, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1a, 0x57,
	0x67, 0x5c, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x26, 0x57, 0x67, 0x5c, 0x43,
	0x6f, 0x73, 0x6d, 0x6f, 0x5c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x1d, 0x57, 0x67, 0x3a, 0x3a, 0x43, 0x6f, 0x73, 0x6d, 0x6f, 0x3a, 0x3a,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescOnce sync.Once
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescData = file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDesc
)

func file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescGZIP() []byte {
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescOnce.Do(func() {
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescData)
	})
	return file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDescData
}

var file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_goTypes = []interface{}{
	(OperationType)(0),                             // 0: wg.cosmo.graphqlmetrics.v1.OperationType
	(*RequestInfo)(nil),                            // 1: wg.cosmo.graphqlmetrics.v1.RequestInfo
	(*SchemaUsageInfo)(nil),                        // 2: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo
	(*ClientInfo)(nil),                             // 3: wg.cosmo.graphqlmetrics.v1.ClientInfo
	(*OperationInfo)(nil),                          // 4: wg.cosmo.graphqlmetrics.v1.OperationInfo
	(*SchemaInfo)(nil),                             // 5: wg.cosmo.graphqlmetrics.v1.SchemaInfo
	(*TypeFieldUsageInfo)(nil),                     // 6: wg.cosmo.graphqlmetrics.v1.TypeFieldUsageInfo
	(*PublishGraphQLRequestMetricsRequest)(nil),    // 7: wg.cosmo.graphqlmetrics.v1.PublishGraphQLRequestMetricsRequest
	(*PublishOperationCoverageReportResponse)(nil), // 8: wg.cosmo.graphqlmetrics.v1.PublishOperationCoverageReportResponse
	nil, // 9: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.AttributesEntry
}
var file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_depIdxs = []int32{
	6, // 0: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.TypeFieldMetrics:type_name -> wg.cosmo.graphqlmetrics.v1.TypeFieldUsageInfo
	4, // 1: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.OperationInfo:type_name -> wg.cosmo.graphqlmetrics.v1.OperationInfo
	5, // 2: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.SchemaInfo:type_name -> wg.cosmo.graphqlmetrics.v1.SchemaInfo
	3, // 3: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.ClientInfo:type_name -> wg.cosmo.graphqlmetrics.v1.ClientInfo
	1, // 4: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.RequestInfo:type_name -> wg.cosmo.graphqlmetrics.v1.RequestInfo
	9, // 5: wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.Attributes:type_name -> wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo.AttributesEntry
	0, // 6: wg.cosmo.graphqlmetrics.v1.OperationInfo.Type:type_name -> wg.cosmo.graphqlmetrics.v1.OperationType
	2, // 7: wg.cosmo.graphqlmetrics.v1.PublishGraphQLRequestMetricsRequest.SchemaUsage:type_name -> wg.cosmo.graphqlmetrics.v1.SchemaUsageInfo
	7, // 8: wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishGraphQLMetrics:input_type -> wg.cosmo.graphqlmetrics.v1.PublishGraphQLRequestMetricsRequest
	8, // 9: wg.cosmo.graphqlmetrics.v1.GraphQLMetricsService.PublishGraphQLMetrics:output_type -> wg.cosmo.graphqlmetrics.v1.PublishOperationCoverageReportResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_init() }
func file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_init() {
	if File_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaUsageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperationInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchemaInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeFieldUsageInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishGraphQLRequestMetricsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishOperationCoverageReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_goTypes,
		DependencyIndexes: file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_depIdxs,
		EnumInfos:         file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_enumTypes,
		MessageInfos:      file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_msgTypes,
	}.Build()
	File_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto = out.File
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_rawDesc = nil
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_goTypes = nil
	file_wg_cosmo_graphqlmetrics_v1_graphqlmetrics_proto_depIdxs = nil
}
