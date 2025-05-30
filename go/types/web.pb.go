// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: web.proto

package types

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

type WebService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName    string  `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceArea    int32   `protobuf:"varint,2,opt,name=service_area,json=serviceArea,proto3" json:"service_area,omitempty"`
	PostBodyType   string  `protobuf:"bytes,3,opt,name=post_body_type,json=postBodyType,proto3" json:"post_body_type,omitempty"`
	PostRespType   string  `protobuf:"bytes,4,opt,name=post_resp_type,json=postRespType,proto3" json:"post_resp_type,omitempty"`
	PutBodyType    string  `protobuf:"bytes,5,opt,name=put_body_type,json=putBodyType,proto3" json:"put_body_type,omitempty"`
	PutRespType    string  `protobuf:"bytes,6,opt,name=put_resp_type,json=putRespType,proto3" json:"put_resp_type,omitempty"`
	PatchBodyType  string  `protobuf:"bytes,7,opt,name=patch_body_type,json=patchBodyType,proto3" json:"patch_body_type,omitempty"`
	PatchRespType  string  `protobuf:"bytes,8,opt,name=patch_resp_type,json=patchRespType,proto3" json:"patch_resp_type,omitempty"`
	DeleteBodyType string  `protobuf:"bytes,9,opt,name=delete_body_type,json=deleteBodyType,proto3" json:"delete_body_type,omitempty"`
	DeleteRespType string  `protobuf:"bytes,10,opt,name=delete_resp_type,json=deleteRespType,proto3" json:"delete_resp_type,omitempty"`
	GetBodyType    string  `protobuf:"bytes,11,opt,name=get_body_type,json=getBodyType,proto3" json:"get_body_type,omitempty"`
	GetRespType    string  `protobuf:"bytes,12,opt,name=get_resp_type,json=getRespType,proto3" json:"get_resp_type,omitempty"`
	Plugin         *Plugin `protobuf:"bytes,13,opt,name=plugin,proto3" json:"plugin,omitempty"`
}

func (x *WebService) Reset() {
	*x = WebService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebService) ProtoMessage() {}

func (x *WebService) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebService.ProtoReflect.Descriptor instead.
func (*WebService) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{0}
}

func (x *WebService) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *WebService) GetServiceArea() int32 {
	if x != nil {
		return x.ServiceArea
	}
	return 0
}

func (x *WebService) GetPostBodyType() string {
	if x != nil {
		return x.PostBodyType
	}
	return ""
}

func (x *WebService) GetPostRespType() string {
	if x != nil {
		return x.PostRespType
	}
	return ""
}

func (x *WebService) GetPutBodyType() string {
	if x != nil {
		return x.PutBodyType
	}
	return ""
}

func (x *WebService) GetPutRespType() string {
	if x != nil {
		return x.PutRespType
	}
	return ""
}

func (x *WebService) GetPatchBodyType() string {
	if x != nil {
		return x.PatchBodyType
	}
	return ""
}

func (x *WebService) GetPatchRespType() string {
	if x != nil {
		return x.PatchRespType
	}
	return ""
}

func (x *WebService) GetDeleteBodyType() string {
	if x != nil {
		return x.DeleteBodyType
	}
	return ""
}

func (x *WebService) GetDeleteRespType() string {
	if x != nil {
		return x.DeleteRespType
	}
	return ""
}

func (x *WebService) GetGetBodyType() string {
	if x != nil {
		return x.GetBodyType
	}
	return ""
}

func (x *WebService) GetGetRespType() string {
	if x != nil {
		return x.GetRespType
	}
	return ""
}

func (x *WebService) GetPlugin() *Plugin {
	if x != nil {
		return x.Plugin
	}
	return nil
}

type Plugin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Plugin) Reset() {
	*x = Plugin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_web_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugin) ProtoMessage() {}

func (x *Plugin) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugin.ProtoReflect.Descriptor instead.
func (*Plugin) Descriptor() ([]byte, []int) {
	return file_web_proto_rawDescGZIP(), []int{1}
}

func (x *Plugin) GetData() string {
	if x != nil {
		return x.Data
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
		mi := &file_web_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_web_proto_msgTypes[2]
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
	return file_web_proto_rawDescGZIP(), []int{2}
}

var File_web_proto protoreflect.FileDescriptor

var file_web_proto_rawDesc = []byte{
	0x0a, 0x09, 0x77, 0x65, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x22, 0xf9, 0x03, 0x0a, 0x0a, 0x57, 0x65, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x6f, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a,
	0x0e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x75, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x74, 0x42,
	0x6f, 0x64, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x70, 0x75, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x74, 0x63, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x64,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x22, 0x1c,
	0x0a, 0x06, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x07, 0x0a, 0x05,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x23, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_web_proto_rawDescOnce sync.Once
	file_web_proto_rawDescData = file_web_proto_rawDesc
)

func file_web_proto_rawDescGZIP() []byte {
	file_web_proto_rawDescOnce.Do(func() {
		file_web_proto_rawDescData = protoimpl.X.CompressGZIP(file_web_proto_rawDescData)
	})
	return file_web_proto_rawDescData
}

var file_web_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_web_proto_goTypes = []interface{}{
	(*WebService)(nil), // 0: types.WebService
	(*Plugin)(nil),     // 1: types.Plugin
	(*Empty)(nil),      // 2: types.Empty
}
var file_web_proto_depIdxs = []int32{
	1, // 0: types.WebService.plugin:type_name -> types.Plugin
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_web_proto_init() }
func file_web_proto_init() {
	if File_web_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_web_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebService); i {
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
		file_web_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugin); i {
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
		file_web_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_web_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_web_proto_goTypes,
		DependencyIndexes: file_web_proto_depIdxs,
		MessageInfos:      file_web_proto_msgTypes,
	}.Build()
	File_web_proto = out.File
	file_web_proto_rawDesc = nil
	file_web_proto_goTypes = nil
	file_web_proto_depIdxs = nil
}
