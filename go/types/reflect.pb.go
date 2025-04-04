// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: reflect.proto

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

type DecoratorType int32

const (
	DecoratorType_Invalid            DecoratorType = 0
	DecoratorType_Primary            DecoratorType = 1
	DecoratorType_Unique             DecoratorType = 2
	DecoratorType_NonUnique          DecoratorType = 3
	DecoratorType_Title              DecoratorType = 4
	DecoratorType_Editor             DecoratorType = 5
	DecoratorType_NoNestedInspection DecoratorType = 6
	DecoratorType_Ignore             DecoratorType = 7
	DecoratorType_IgnoreAttr         DecoratorType = 8
)

// Enum value maps for DecoratorType.
var (
	DecoratorType_name = map[int32]string{
		0: "Invalid",
		1: "Primary",
		2: "Unique",
		3: "NonUnique",
		4: "Title",
		5: "Editor",
		6: "NoNestedInspection",
		7: "Ignore",
		8: "IgnoreAttr",
	}
	DecoratorType_value = map[string]int32{
		"Invalid":            0,
		"Primary":            1,
		"Unique":             2,
		"NonUnique":          3,
		"Title":              4,
		"Editor":             5,
		"NoNestedInspection": 6,
		"Ignore":             7,
		"IgnoreAttr":         8,
	}
)

func (x DecoratorType) Enum() *DecoratorType {
	p := new(DecoratorType)
	*p = x
	return p
}

func (x DecoratorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DecoratorType) Descriptor() protoreflect.EnumDescriptor {
	return file_reflect_proto_enumTypes[0].Descriptor()
}

func (DecoratorType) Type() protoreflect.EnumType {
	return &file_reflect_proto_enumTypes[0]
}

func (x DecoratorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DecoratorType.Descriptor instead.
func (DecoratorType) EnumDescriptor() ([]byte, []int) {
	return file_reflect_proto_rawDescGZIP(), []int{0}
}

type RNode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type name, in case of a map, this is the value type.
	TypeName string `protobuf:"bytes,1,opt,name=type_name,json=typeName,proto3" json:"type_name,omitempty"`
	// In case this attribute is a cell in a map or a slice, this is the key type
	KeyTypeName string `protobuf:"bytes,2,opt,name=key_type_name,json=keyTypeName,proto3" json:"key_type_name,omitempty"`
	// The parent node, nil if root.
	Parent *RNode `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	// The attribute name in the parent
	FieldName string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	// In case this attribute is a struct, a map between the attribute name and the child registry node.
	Attributes map[string]*RNode `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If this attribute is a map
	IsMap bool `protobuf:"varint,6,opt,name=is_map,json=isMap,proto3" json:"is_map,omitempty"`
	// If this attribute is a slice
	IsSlice bool `protobuf:"varint,7,opt,name=is_slice,json=isSlice,proto3" json:"is_slice,omitempty"`
	// The cached key so we won't need to calculate it all the time.
	CachedKey string `protobuf:"bytes,8,opt,name=cached_key,json=cachedKey,proto3" json:"cached_key,omitempty"`
	// Decorators
	Decorators map[int32]string `protobuf:"bytes,9,rep,name=decorators,proto3" json:"decorators,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// If the attribute is struct or map/slice to struct.
	IsStruct bool `protobuf:"varint,10,opt,name=is_struct,json=isStruct,proto3" json:"is_struct,omitempty"`
}

func (x *RNode) Reset() {
	*x = RNode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflect_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RNode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RNode) ProtoMessage() {}

func (x *RNode) ProtoReflect() protoreflect.Message {
	mi := &file_reflect_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RNode.ProtoReflect.Descriptor instead.
func (*RNode) Descriptor() ([]byte, []int) {
	return file_reflect_proto_rawDescGZIP(), []int{0}
}

func (x *RNode) GetTypeName() string {
	if x != nil {
		return x.TypeName
	}
	return ""
}

func (x *RNode) GetKeyTypeName() string {
	if x != nil {
		return x.KeyTypeName
	}
	return ""
}

func (x *RNode) GetParent() *RNode {
	if x != nil {
		return x.Parent
	}
	return nil
}

func (x *RNode) GetFieldName() string {
	if x != nil {
		return x.FieldName
	}
	return ""
}

func (x *RNode) GetAttributes() map[string]*RNode {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *RNode) GetIsMap() bool {
	if x != nil {
		return x.IsMap
	}
	return false
}

func (x *RNode) GetIsSlice() bool {
	if x != nil {
		return x.IsSlice
	}
	return false
}

func (x *RNode) GetCachedKey() string {
	if x != nil {
		return x.CachedKey
	}
	return ""
}

func (x *RNode) GetDecorators() map[int32]string {
	if x != nil {
		return x.Decorators
	}
	return nil
}

func (x *RNode) GetIsStruct() bool {
	if x != nil {
		return x.IsStruct
	}
	return false
}

type TableView struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Table     *RNode   `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Columns   []*RNode `protobuf:"bytes,2,rep,name=columns,proto3" json:"columns,omitempty"`
	SubTables []*RNode `protobuf:"bytes,3,rep,name=sub_tables,json=subTables,proto3" json:"sub_tables,omitempty"`
}

func (x *TableView) Reset() {
	*x = TableView{}
	if protoimpl.UnsafeEnabled {
		mi := &file_reflect_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableView) ProtoMessage() {}

func (x *TableView) ProtoReflect() protoreflect.Message {
	mi := &file_reflect_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableView.ProtoReflect.Descriptor instead.
func (*TableView) Descriptor() ([]byte, []int) {
	return file_reflect_proto_rawDescGZIP(), []int{1}
}

func (x *TableView) GetTable() *RNode {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *TableView) GetColumns() []*RNode {
	if x != nil {
		return x.Columns
	}
	return nil
}

func (x *TableView) GetSubTables() []*RNode {
	if x != nil {
		return x.SubTables
	}
	return nil
}

var File_reflect_proto protoreflect.FileDescriptor

var file_reflect_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x83, 0x04, 0x0a, 0x05, 0x52, 0x4e, 0x6f, 0x64, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x4d, 0x61, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x53, 0x6c, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64,
	0x5f, 0x6b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0a, 0x64, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x64, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x1a, 0x4b, 0x0a, 0x0f, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x4e, 0x6f,
	0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3d, 0x0a,
	0x0f, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x84, 0x01, 0x0a,
	0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x22, 0x0a, 0x05, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26,
	0x0a, 0x07, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x07, 0x63,
	0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x52, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x09, 0x73, 0x75, 0x62, 0x54, 0x61, 0x62,
	0x6c, 0x65, 0x73, 0x2a, 0x8f, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x4e,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x10,
	0x05, 0x12, 0x16, 0x0a, 0x12, 0x4e, 0x6f, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x49, 0x6e, 0x73,
	0x70, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x41,
	0x74, 0x74, 0x72, 0x10, 0x08, 0x42, 0x23, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x42, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50,
	0x01, 0x5a, 0x07, 0x2e, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_reflect_proto_rawDescOnce sync.Once
	file_reflect_proto_rawDescData = file_reflect_proto_rawDesc
)

func file_reflect_proto_rawDescGZIP() []byte {
	file_reflect_proto_rawDescOnce.Do(func() {
		file_reflect_proto_rawDescData = protoimpl.X.CompressGZIP(file_reflect_proto_rawDescData)
	})
	return file_reflect_proto_rawDescData
}

var file_reflect_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_reflect_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_reflect_proto_goTypes = []interface{}{
	(DecoratorType)(0), // 0: types.DecoratorType
	(*RNode)(nil),      // 1: types.RNode
	(*TableView)(nil),  // 2: types.TableView
	nil,                // 3: types.RNode.AttributesEntry
	nil,                // 4: types.RNode.DecoratorsEntry
}
var file_reflect_proto_depIdxs = []int32{
	1, // 0: types.RNode.parent:type_name -> types.RNode
	3, // 1: types.RNode.attributes:type_name -> types.RNode.AttributesEntry
	4, // 2: types.RNode.decorators:type_name -> types.RNode.DecoratorsEntry
	1, // 3: types.TableView.table:type_name -> types.RNode
	1, // 4: types.TableView.columns:type_name -> types.RNode
	1, // 5: types.TableView.sub_tables:type_name -> types.RNode
	1, // 6: types.RNode.AttributesEntry.value:type_name -> types.RNode
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_reflect_proto_init() }
func file_reflect_proto_init() {
	if File_reflect_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_reflect_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RNode); i {
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
		file_reflect_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableView); i {
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
			RawDescriptor: file_reflect_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_reflect_proto_goTypes,
		DependencyIndexes: file_reflect_proto_depIdxs,
		EnumInfos:         file_reflect_proto_enumTypes,
		MessageInfos:      file_reflect_proto_msgTypes,
	}.Build()
	File_reflect_proto = out.File
	file_reflect_proto_rawDesc = nil
	file_reflect_proto_goTypes = nil
	file_reflect_proto_depIdxs = nil
}
