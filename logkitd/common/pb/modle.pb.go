// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: modle.proto

package pb

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

type TypeName int32

const (
	TypeName_GetOs  TypeName = 0
	TypeName_GetPid TypeName = 1
)

// Enum value maps for TypeName.
var (
	TypeName_name = map[int32]string{
		0: "GetOs",
		1: "GetPid",
	}
	TypeName_value = map[string]int32{
		"GetOs":  0,
		"GetPid": 1,
	}
)

func (x TypeName) Enum() *TypeName {
	p := new(TypeName)
	*p = x
	return p
}

func (x TypeName) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TypeName) Descriptor() protoreflect.EnumDescriptor {
	return file_modle_proto_enumTypes[0].Descriptor()
}

func (TypeName) Type() protoreflect.EnumType {
	return &file_modle_proto_enumTypes[0]
}

func (x TypeName) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TypeName.Descriptor instead.
func (TypeName) EnumDescriptor() ([]byte, []int) {
	return file_modle_proto_rawDescGZIP(), []int{0}
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TypeName `protobuf:"varint,1,opt,name=type,proto3,enum=pb.TypeName" json:"type,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_modle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_modle_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetType() TypeName {
	if x != nil {
		return x.Type
	}
	return TypeName_GetOs
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  TypeName   `protobuf:"varint,1,opt,name=type,proto3,enum=pb.TypeName" json:"type,omitempty"`
	Os    *OsInfoRsp `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Pid   *PidRsp    `protobuf:"bytes,3,opt,name=pid,proto3" json:"pid,omitempty"`
	Error string     `protobuf:"bytes,10,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_modle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_modle_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetType() TypeName {
	if x != nil {
		return x.Type
	}
	return TypeName_GetOs
}

func (x *Response) GetOs() *OsInfoRsp {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *Response) GetPid() *PidRsp {
	if x != nil {
		return x.Pid
	}
	return nil
}

func (x *Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type OsInfoRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Os   string `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Arch string `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
}

func (x *OsInfoRsp) Reset() {
	*x = OsInfoRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OsInfoRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OsInfoRsp) ProtoMessage() {}

func (x *OsInfoRsp) ProtoReflect() protoreflect.Message {
	mi := &file_modle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OsInfoRsp.ProtoReflect.Descriptor instead.
func (*OsInfoRsp) Descriptor() ([]byte, []int) {
	return file_modle_proto_rawDescGZIP(), []int{2}
}

func (x *OsInfoRsp) GetOs() string {
	if x != nil {
		return x.Os
	}
	return ""
}

func (x *OsInfoRsp) GetArch() string {
	if x != nil {
		return x.Arch
	}
	return ""
}

type PidRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *PidRsp) Reset() {
	*x = PidRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PidRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PidRsp) ProtoMessage() {}

func (x *PidRsp) ProtoReflect() protoreflect.Message {
	mi := &file_modle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PidRsp.ProtoReflect.Descriptor instead.
func (*PidRsp) Descriptor() ([]byte, []int) {
	return file_modle_proto_rawDescGZIP(), []int{3}
}

func (x *PidRsp) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

var File_modle_proto protoreflect.FileDescriptor

var file_modle_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x2b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7f,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x02,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x1c, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69,
	0x64, 0x52, 0x73, 0x70, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22,
	0x2f, 0x0a, 0x09, 0x4f, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x63, 0x68,
	0x22, 0x1a, 0x0a, 0x06, 0x50, 0x69, 0x64, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x2a, 0x21, 0x0a, 0x08,
	0x54, 0x79, 0x70, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x65, 0x74, 0x4f,
	0x73, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x50, 0x69, 0x64, 0x10, 0x01, 0x42,
	0x18, 0x5a, 0x16, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x6b, 0x69, 0x74, 0x64, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_modle_proto_rawDescOnce sync.Once
	file_modle_proto_rawDescData = file_modle_proto_rawDesc
)

func file_modle_proto_rawDescGZIP() []byte {
	file_modle_proto_rawDescOnce.Do(func() {
		file_modle_proto_rawDescData = protoimpl.X.CompressGZIP(file_modle_proto_rawDescData)
	})
	return file_modle_proto_rawDescData
}

var file_modle_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_modle_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_modle_proto_goTypes = []interface{}{
	(TypeName)(0),     // 0: pb.TypeName
	(*Request)(nil),   // 1: pb.Request
	(*Response)(nil),  // 2: pb.Response
	(*OsInfoRsp)(nil), // 3: pb.OsInfoRsp
	(*PidRsp)(nil),    // 4: pb.PidRsp
}
var file_modle_proto_depIdxs = []int32{
	0, // 0: pb.Request.type:type_name -> pb.TypeName
	0, // 1: pb.Response.type:type_name -> pb.TypeName
	3, // 2: pb.Response.os:type_name -> pb.OsInfoRsp
	4, // 3: pb.Response.pid:type_name -> pb.PidRsp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_modle_proto_init() }
func file_modle_proto_init() {
	if File_modle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_modle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_modle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OsInfoRsp); i {
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
		file_modle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PidRsp); i {
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
			RawDescriptor: file_modle_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_modle_proto_goTypes,
		DependencyIndexes: file_modle_proto_depIdxs,
		EnumInfos:         file_modle_proto_enumTypes,
		MessageInfos:      file_modle_proto_msgTypes,
	}.Build()
	File_modle_proto = out.File
	file_modle_proto_rawDesc = nil
	file_modle_proto_goTypes = nil
	file_modle_proto_depIdxs = nil
}
