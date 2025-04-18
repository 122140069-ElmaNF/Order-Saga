// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v6.30.2
// source: proto/shipping.proto

package shippingpb

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

type StartShippingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *StartShippingRequest) Reset() {
	*x = StartShippingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingRequest) ProtoMessage() {}

func (x *StartShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingRequest.ProtoReflect.Descriptor instead.
func (*StartShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *StartShippingRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type StartShippingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippingId string `protobuf:"bytes,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // SHIPPED, FAILED
}

func (x *StartShippingResponse) Reset() {
	*x = StartShippingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingResponse) ProtoMessage() {}

func (x *StartShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingResponse.ProtoReflect.Descriptor instead.
func (*StartShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *StartShippingResponse) GetShippingId() string {
	if x != nil {
		return x.ShippingId
	}
	return ""
}

func (x *StartShippingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CancelShippingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShippingId string `protobuf:"bytes,1,opt,name=shipping_id,json=shippingId,proto3" json:"shipping_id,omitempty"`
}

func (x *CancelShippingRequest) Reset() {
	*x = CancelShippingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingRequest) ProtoMessage() {}

func (x *CancelShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingRequest.ProtoReflect.Descriptor instead.
func (*CancelShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *CancelShippingRequest) GetShippingId() string {
	if x != nil {
		return x.ShippingId
	}
	return ""
}

type CancelShippingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // CANCELLED
}

func (x *CancelShippingResponse) Reset() {
	*x = CancelShippingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shipping_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingResponse) ProtoMessage() {}

func (x *CancelShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingResponse.ProtoReflect.Descriptor instead.
func (*CancelShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{3}
}

func (x *CancelShippingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_shipping_proto protoreflect.FileDescriptor

var file_proto_shipping_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x22, 0x31, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x22, 0x50, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x38, 0x0a, 0x15, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x22,
	0x30, 0x0a, 0x16, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0xb8, 0x01, 0x0a, 0x0f, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x69, 0x70, 0x70,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x68, 0x69,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x53, 0x68, 0x69, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19,
	0x67, 0x65, 0x6e, 0x2f, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x3b, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_shipping_proto_rawDescOnce sync.Once
	file_proto_shipping_proto_rawDescData = file_proto_shipping_proto_rawDesc
)

func file_proto_shipping_proto_rawDescGZIP() []byte {
	file_proto_shipping_proto_rawDescOnce.Do(func() {
		file_proto_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shipping_proto_rawDescData)
	})
	return file_proto_shipping_proto_rawDescData
}

var file_proto_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shipping_proto_goTypes = []interface{}{
	(*StartShippingRequest)(nil),   // 0: shipping.StartShippingRequest
	(*StartShippingResponse)(nil),  // 1: shipping.StartShippingResponse
	(*CancelShippingRequest)(nil),  // 2: shipping.CancelShippingRequest
	(*CancelShippingResponse)(nil), // 3: shipping.CancelShippingResponse
}
var file_proto_shipping_proto_depIdxs = []int32{
	0, // 0: shipping.ShippingService.StartShipping:input_type -> shipping.StartShippingRequest
	2, // 1: shipping.ShippingService.CancelShipping:input_type -> shipping.CancelShippingRequest
	1, // 2: shipping.ShippingService.StartShipping:output_type -> shipping.StartShippingResponse
	3, // 3: shipping.ShippingService.CancelShipping:output_type -> shipping.CancelShippingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_shipping_proto_init() }
func file_proto_shipping_proto_init() {
	if File_proto_shipping_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shipping_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartShippingRequest); i {
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
		file_proto_shipping_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartShippingResponse); i {
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
		file_proto_shipping_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelShippingRequest); i {
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
		file_proto_shipping_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelShippingResponse); i {
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
			RawDescriptor: file_proto_shipping_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shipping_proto_goTypes,
		DependencyIndexes: file_proto_shipping_proto_depIdxs,
		MessageInfos:      file_proto_shipping_proto_msgTypes,
	}.Build()
	File_proto_shipping_proto = out.File
	file_proto_shipping_proto_rawDesc = nil
	file_proto_shipping_proto_goTypes = nil
	file_proto_shipping_proto_depIdxs = nil
}
