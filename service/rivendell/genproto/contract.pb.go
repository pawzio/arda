// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: contract.proto

package genproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetProductsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*Product `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetProductsResponse) Reset() {
	*x = GetProductsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductsResponse) ProtoMessage() {}

func (x *GetProductsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductsResponse.ProtoReflect.Descriptor instead.
func (*GetProductsResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductsResponse) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x6c, 0x6c, 0x22, 0x45, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x72, 0x69, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x6c, 0x6c,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x22, 0x4f, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x4c, 0x0a, 0x09,
	0x52, 0x69, 0x76, 0x65, 0x6e, 0x64, 0x65, 0x6c, 0x6c, 0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x69, 0x76, 0x65, 0x6e,
	0x64, 0x65, 0x6c, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1e, 0x2e, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x64, 0x65, 0x6c, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x77, 0x7a, 0x69, 0x6f, 0x2f,
	0x61, 0x72, 0x64, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x72, 0x69, 0x76,
	0x65, 0x6e, 0x64, 0x65, 0x6c, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contract_proto_goTypes = []interface{}{
	(*GetProductsResponse)(nil), // 0: rivendell.GetProductsResponse
	(*Product)(nil),             // 1: rivendell.Product
	(*Empty)(nil),               // 2: rivendell.Empty
}
var file_contract_proto_depIdxs = []int32{
	1, // 0: rivendell.GetProductsResponse.products:type_name -> rivendell.Product
	2, // 1: rivendell.Rivendell.GetProducts:input_type -> rivendell.Empty
	0, // 2: rivendell.Rivendell.GetProducts:output_type -> rivendell.GetProductsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductsResponse); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
