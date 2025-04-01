// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: notification.proto

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

type NotificationType int32

const (
	NotificationType_Invalid_Notify NotificationType = 0
	NotificationType_Add            NotificationType = 1
	NotificationType_Replace        NotificationType = 2
	NotificationType_Update         NotificationType = 3
	NotificationType_Delete         NotificationType = 4
)

// Enum value maps for NotificationType.
var (
	NotificationType_name = map[int32]string{
		0: "Invalid_Notify",
		1: "Add",
		2: "Replace",
		3: "Update",
		4: "Delete",
	}
	NotificationType_value = map[string]int32{
		"Invalid_Notify": 0,
		"Add":            1,
		"Replace":        2,
		"Update":         3,
		"Delete":         4,
	}
)

func (x NotificationType) Enum() *NotificationType {
	p := new(NotificationType)
	*p = x
	return p
}

func (x NotificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NotificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_notification_proto_enumTypes[0].Descriptor()
}

func (NotificationType) Type() protoreflect.EnumType {
	return &file_notification_proto_enumTypes[0]
}

func (x NotificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NotificationType.Descriptor instead.
func (NotificationType) EnumDescriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

type NotificationSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source           string           `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Time             int64            `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Sequence         uint32           `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	ServiceName      string           `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	ServiceArea      int32            `protobuf:"varint,5,opt,name=service_area,json=serviceArea,proto3" json:"service_area,omitempty"`
	ModelType        string           `protobuf:"bytes,6,opt,name=model_type,json=modelType,proto3" json:"model_type,omitempty"`
	Type             NotificationType `protobuf:"varint,7,opt,name=type,proto3,enum=types.NotificationType" json:"type,omitempty"`
	NotificationList []*Notification  `protobuf:"bytes,8,rep,name=notification_list,json=notificationList,proto3" json:"notification_list,omitempty"`
}

func (x *NotificationSet) Reset() {
	*x = NotificationSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationSet) ProtoMessage() {}

func (x *NotificationSet) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationSet.ProtoReflect.Descriptor instead.
func (*NotificationSet) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationSet) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *NotificationSet) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *NotificationSet) GetSequence() uint32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *NotificationSet) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *NotificationSet) GetServiceArea() int32 {
	if x != nil {
		return x.ServiceArea
	}
	return 0
}

func (x *NotificationSet) GetModelType() string {
	if x != nil {
		return x.ModelType
	}
	return ""
}

func (x *NotificationSet) GetType() NotificationType {
	if x != nil {
		return x.Type
	}
	return NotificationType_Invalid_Notify
}

func (x *NotificationSet) GetNotificationList() []*Notification {
	if x != nil {
		return x.NotificationList
	}
	return nil
}

type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId string `protobuf:"bytes,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	OldValue   []byte `protobuf:"bytes,2,opt,name=oldValue,proto3" json:"oldValue,omitempty"`
	NewValue   []byte `protobuf:"bytes,3,opt,name=newValue,proto3" json:"newValue,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notification_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetPropertyId() string {
	if x != nil {
		return x.PropertyId
	}
	return ""
}

func (x *Notification) GetOldValue() []byte {
	if x != nil {
		return x.OldValue
	}
	return nil
}

func (x *Notification) GetNewValue() []byte {
	if x != nil {
		return x.NewValue
	}
	return nil
}

var File_notification_proto protoreflect.FileDescriptor

var file_notification_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0xad, 0x02, 0x0a, 0x0f,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73,
	0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x1d, 0x0a,
	0x0a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x40, 0x0a, 0x11, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x67, 0x0a, 0x0c, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6f, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x2a, 0x54, 0x0a, 0x10, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03,
	0x41, 0x64, 0x64, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x03, 0x12, 0x0a,
	0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x04, 0x42, 0x23, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notification_proto_rawDescOnce sync.Once
	file_notification_proto_rawDescData = file_notification_proto_rawDesc
)

func file_notification_proto_rawDescGZIP() []byte {
	file_notification_proto_rawDescOnce.Do(func() {
		file_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_notification_proto_rawDescData)
	})
	return file_notification_proto_rawDescData
}

var file_notification_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notification_proto_goTypes = []interface{}{
	(NotificationType)(0),   // 0: types.NotificationType
	(*NotificationSet)(nil), // 1: types.NotificationSet
	(*Notification)(nil),    // 2: types.Notification
}
var file_notification_proto_depIdxs = []int32{
	0, // 0: types.NotificationSet.type:type_name -> types.NotificationType
	2, // 1: types.NotificationSet.notification_list:type_name -> types.Notification
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_notification_proto_init() }
func file_notification_proto_init() {
	if File_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationSet); i {
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
		file_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
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
			RawDescriptor: file_notification_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notification_proto_goTypes,
		DependencyIndexes: file_notification_proto_depIdxs,
		EnumInfos:         file_notification_proto_enumTypes,
		MessageInfos:      file_notification_proto_msgTypes,
	}.Build()
	File_notification_proto = out.File
	file_notification_proto_rawDesc = nil
	file_notification_proto_goTypes = nil
	file_notification_proto_depIdxs = nil
}
