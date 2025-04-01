// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: config.proto

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

// A config for the messages limitation and data size
type SysConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum data size for a message
	MaxDataSize uint64 `protobuf:"varint,1,opt,name=max_data_size,json=maxDataSize,proto3" json:"max_data_size,omitempty"`
	// The transmit message queue size
	TxQueueSize uint64 `protobuf:"varint,2,opt,name=tx_queue_size,json=txQueueSize,proto3" json:"tx_queue_size,omitempty"`
	// The received message queue size
	RxQueueSize uint64 `protobuf:"varint,3,opt,name=rx_queue_size,json=rxQueueSize,proto3" json:"rx_queue_size,omitempty"`
	// What is the switch port for this configuration
	VnetPort uint32 `protobuf:"varint,4,opt,name=vnet_port,json=vnetPort,proto3" json:"vnet_port,omitempty"`
	// The local uuid
	LocalUuid string `protobuf:"bytes,5,opt,name=local_uuid,json=localUuid,proto3" json:"local_uuid,omitempty"`
	// The remote uuid
	RemoteUuid string `protobuf:"bytes,6,opt,name=remote_uuid,json=remoteUuid,proto3" json:"remote_uuid,omitempty"`
	// The address of the connection initiator, regardless of the side.
	Address string `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	// force external mode in case two switches are on the same machine
	// woth different ports
	ForceExternal bool `protobuf:"varint,8,opt,name=force_external,json=forceExternal,proto3" json:"force_external,omitempty"`
	// Custom local alias for this vnic
	LocalAlias string `protobuf:"bytes,9,opt,name=local_alias,json=localAlias,proto3" json:"local_alias,omitempty"`
	// Custom remote alias for this vnic
	RemoteAlias string `protobuf:"bytes,10,opt,name=remote_alias,json=remoteAlias,proto3" json:"remote_alias,omitempty"`
	// Provided services
	Services *Services `protobuf:"bytes,11,opt,name=services,proto3" json:"services,omitempty"`
	// Keep Alive interval in Seconds
	KeepAliveIntervalSeconds int64 `protobuf:"varint,12,opt,name=keep_alive_interval_seconds,json=keepAliveIntervalSeconds,proto3" json:"keep_alive_interval_seconds,omitempty"`
}

func (x *SysConfig) Reset() {
	*x = SysConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysConfig) ProtoMessage() {}

func (x *SysConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysConfig.ProtoReflect.Descriptor instead.
func (*SysConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *SysConfig) GetMaxDataSize() uint64 {
	if x != nil {
		return x.MaxDataSize
	}
	return 0
}

func (x *SysConfig) GetTxQueueSize() uint64 {
	if x != nil {
		return x.TxQueueSize
	}
	return 0
}

func (x *SysConfig) GetRxQueueSize() uint64 {
	if x != nil {
		return x.RxQueueSize
	}
	return 0
}

func (x *SysConfig) GetVnetPort() uint32 {
	if x != nil {
		return x.VnetPort
	}
	return 0
}

func (x *SysConfig) GetLocalUuid() string {
	if x != nil {
		return x.LocalUuid
	}
	return ""
}

func (x *SysConfig) GetRemoteUuid() string {
	if x != nil {
		return x.RemoteUuid
	}
	return ""
}

func (x *SysConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *SysConfig) GetForceExternal() bool {
	if x != nil {
		return x.ForceExternal
	}
	return false
}

func (x *SysConfig) GetLocalAlias() string {
	if x != nil {
		return x.LocalAlias
	}
	return ""
}

func (x *SysConfig) GetRemoteAlias() string {
	if x != nil {
		return x.RemoteAlias
	}
	return ""
}

func (x *SysConfig) GetServices() *Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *SysConfig) GetKeepAliveIntervalSeconds() int64 {
	if x != nil {
		return x.KeepAliveIntervalSeconds
	}
	return 0
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x03, 0x0a, 0x09, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x74, 0x78, 0x5f, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x74, 0x78, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72,
	0x78, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x72, 0x78, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x76, 0x6e, 0x65, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x76, 0x6e, 0x65, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x66, 0x6f, 0x72, 0x63, 0x65, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x21,
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61,
	0x73, 0x12, 0x2b, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x3d,
	0x0a, 0x1b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x18, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x23, 0x0a,
	0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_proto_goTypes = []interface{}{
	(*SysConfig)(nil), // 0: types.SysConfig
	(*Services)(nil),  // 1: types.Services
}
var file_config_proto_depIdxs = []int32{
	1, // 0: types.SysConfig.services:type_name -> types.Services
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	file_services_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysConfig); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}
