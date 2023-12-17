// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: protobuffs/product.proto

package buffers

import (
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

type Status int32

const (
	Status_SUCCESS Status = 0
	Status_FAILED  Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	Status_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuffs_product_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_protobuffs_product_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{0}
}

type StockUpdate int32

const (
	StockUpdate_INC StockUpdate = 0
	StockUpdate_DEC StockUpdate = 1
)

// Enum value maps for StockUpdate.
var (
	StockUpdate_name = map[int32]string{
		0: "INC",
		1: "DEC",
	}
	StockUpdate_value = map[string]int32{
		"INC": 0,
		"DEC": 1,
	}
)

func (x StockUpdate) Enum() *StockUpdate {
	p := new(StockUpdate)
	*p = x
	return p
}

func (x StockUpdate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StockUpdate) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuffs_product_proto_enumTypes[1].Descriptor()
}

func (StockUpdate) Type() protoreflect.EnumType {
	return &file_protobuffs_product_proto_enumTypes[1]
}

func (x StockUpdate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StockUpdate.Descriptor instead.
func (StockUpdate) EnumDescriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{1}
}

type GetProductByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
}

func (x *GetProductByIdRequest) Reset() {
	*x = GetProductByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdRequest) ProtoMessage() {}

func (x *GetProductByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdRequest.ProtoReflect.Descriptor instead.
func (*GetProductByIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductByIdRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

type GetProductByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	DealerId    string  `protobuf:"bytes,5,opt,name=dealerId,proto3" json:"dealerId,omitempty"`
}

func (x *GetProductByIdResponse) Reset() {
	*x = GetProductByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductByIdResponse) ProtoMessage() {}

func (x *GetProductByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductByIdResponse.ProtoReflect.Descriptor instead.
func (*GetProductByIdResponse) Descriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductByIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetProductByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductByIdResponse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetProductByIdResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GetProductByIdResponse) GetDealerId() string {
	if x != nil {
		return x.DealerId
	}
	return ""
}

type UpdateStockByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string      `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Operation StockUpdate `protobuf:"varint,2,opt,name=operation,proto3,enum=productpb.StockUpdate" json:"operation,omitempty"`
	ByNumber  uint32      `protobuf:"varint,3,opt,name=byNumber,proto3" json:"byNumber,omitempty"`
}

func (x *UpdateStockByIdRequest) Reset() {
	*x = UpdateStockByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockByIdRequest) ProtoMessage() {}

func (x *UpdateStockByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockByIdRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockByIdRequest) Descriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateStockByIdRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *UpdateStockByIdRequest) GetOperation() StockUpdate {
	if x != nil {
		return x.Operation
	}
	return StockUpdate_INC
}

func (x *UpdateStockByIdRequest) GetByNumber() uint32 {
	if x != nil {
		return x.ByNumber
	}
	return 0
}

type UpdateStockByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=productpb.Status" json:"status,omitempty"`
}

func (x *UpdateStockByIdResponse) Reset() {
	*x = UpdateStockByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuffs_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockByIdResponse) ProtoMessage() {}

func (x *UpdateStockByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuffs_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockByIdResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockByIdResponse) Descriptor() ([]byte, []int) {
	return file_protobuffs_product_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateStockByIdResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_SUCCESS
}

var File_protobuffs_product_proto protoreflect.FileDescriptor

var file_protobuffs_product_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x66, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x70, 0x62, 0x22, 0x35, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x22, 0x90, 0x01, 0x0a,
	0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x88, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x62, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70,
	0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2a, 0x21, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45,
	0x44, 0x10, 0x01, 0x2a, 0x1f, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x43, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x44,
	0x45, 0x43, 0x10, 0x01, 0x32, 0xc1, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x58,
	0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuffs_product_proto_rawDescOnce sync.Once
	file_protobuffs_product_proto_rawDescData = file_protobuffs_product_proto_rawDesc
)

func file_protobuffs_product_proto_rawDescGZIP() []byte {
	file_protobuffs_product_proto_rawDescOnce.Do(func() {
		file_protobuffs_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuffs_product_proto_rawDescData)
	})
	return file_protobuffs_product_proto_rawDescData
}

var file_protobuffs_product_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_protobuffs_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_protobuffs_product_proto_goTypes = []interface{}{
	(Status)(0),                     // 0: productpb.Status
	(StockUpdate)(0),                // 1: productpb.StockUpdate
	(*GetProductByIdRequest)(nil),   // 2: productpb.GetProductByIdRequest
	(*GetProductByIdResponse)(nil),  // 3: productpb.GetProductByIdResponse
	(*UpdateStockByIdRequest)(nil),  // 4: productpb.UpdateStockByIdRequest
	(*UpdateStockByIdResponse)(nil), // 5: productpb.UpdateStockByIdResponse
}
var file_protobuffs_product_proto_depIdxs = []int32{
	1, // 0: productpb.UpdateStockByIdRequest.operation:type_name -> productpb.StockUpdate
	0, // 1: productpb.UpdateStockByIdResponse.status:type_name -> productpb.Status
	2, // 2: productpb.ProductService.GetProductById:input_type -> productpb.GetProductByIdRequest
	4, // 3: productpb.ProductService.StockUpdateById:input_type -> productpb.UpdateStockByIdRequest
	3, // 4: productpb.ProductService.GetProductById:output_type -> productpb.GetProductByIdResponse
	5, // 5: productpb.ProductService.StockUpdateById:output_type -> productpb.UpdateStockByIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuffs_product_proto_init() }
func file_protobuffs_product_proto_init() {
	if File_protobuffs_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuffs_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdRequest); i {
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
		file_protobuffs_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductByIdResponse); i {
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
		file_protobuffs_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockByIdRequest); i {
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
		file_protobuffs_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockByIdResponse); i {
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
			RawDescriptor: file_protobuffs_product_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuffs_product_proto_goTypes,
		DependencyIndexes: file_protobuffs_product_proto_depIdxs,
		EnumInfos:         file_protobuffs_product_proto_enumTypes,
		MessageInfos:      file_protobuffs_product_proto_msgTypes,
	}.Build()
	File_protobuffs_product_proto = out.File
	file_protobuffs_product_proto_rawDesc = nil
	file_protobuffs_product_proto_goTypes = nil
	file_protobuffs_product_proto_depIdxs = nil
}
