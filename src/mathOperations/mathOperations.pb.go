// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: src/mathOperations/mathOperations.proto

package __

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

// Format of the input Message used for gRPC Server
type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_mathOperations_mathOperations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_src_mathOperations_mathOperations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_src_mathOperations_mathOperations_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Input) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

// Format of the Output Message used for gRPC Server
type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Output float64 `protobuf:"fixed64,1,opt,name=output,proto3" json:"output,omitempty"`
	Error  *Error  `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_mathOperations_mathOperations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_src_mathOperations_mathOperations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_src_mathOperations_mathOperations_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetOutput() float64 {
	if x != nil {
		return x.Output
	}
	return 0
}

func (x *Output) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

// Format of the Error Message used for gRPC Server
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorCode    string `protobuf:"bytes,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_mathOperations_mathOperations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_src_mathOperations_mathOperations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_src_mathOperations_mathOperations_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

func (x *Error) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_src_mathOperations_mathOperations_proto protoreflect.FileDescriptor

var file_src_mathOperations_mathOperations_proto_rawDesc = []byte{
	0x0a, 0x27, 0x73, 0x72, 0x63, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x74, 0x68, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x23, 0x0a, 0x05, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78,
	0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x22, 0x4d,
	0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x2b, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x49, 0x0a,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xff, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x74,
	0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x36, 0x0a, 0x03, 0x41,
	0x64, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x53, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12,
	0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x15, 0x2e, 0x6d,
	0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a,
	0x08, 0x44, 0x69, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_mathOperations_mathOperations_proto_rawDescOnce sync.Once
	file_src_mathOperations_mathOperations_proto_rawDescData = file_src_mathOperations_mathOperations_proto_rawDesc
)

func file_src_mathOperations_mathOperations_proto_rawDescGZIP() []byte {
	file_src_mathOperations_mathOperations_proto_rawDescOnce.Do(func() {
		file_src_mathOperations_mathOperations_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_mathOperations_mathOperations_proto_rawDescData)
	})
	return file_src_mathOperations_mathOperations_proto_rawDescData
}

var file_src_mathOperations_mathOperations_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_mathOperations_mathOperations_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: mathOperations.Input
	(*Output)(nil), // 1: mathOperations.Output
	(*Error)(nil),  // 2: mathOperations.Error
}
var file_src_mathOperations_mathOperations_proto_depIdxs = []int32{
	2, // 0: mathOperations.Output.error:type_name -> mathOperations.Error
	0, // 1: mathOperations.MathOperations.Add:input_type -> mathOperations.Input
	0, // 2: mathOperations.MathOperations.Subtract:input_type -> mathOperations.Input
	0, // 3: mathOperations.MathOperations.Multiply:input_type -> mathOperations.Input
	0, // 4: mathOperations.MathOperations.Division:input_type -> mathOperations.Input
	1, // 5: mathOperations.MathOperations.Add:output_type -> mathOperations.Output
	1, // 6: mathOperations.MathOperations.Subtract:output_type -> mathOperations.Output
	1, // 7: mathOperations.MathOperations.Multiply:output_type -> mathOperations.Output
	1, // 8: mathOperations.MathOperations.Division:output_type -> mathOperations.Output
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_src_mathOperations_mathOperations_proto_init() }
func file_src_mathOperations_mathOperations_proto_init() {
	if File_src_mathOperations_mathOperations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_mathOperations_mathOperations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_src_mathOperations_mathOperations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
		file_src_mathOperations_mathOperations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
			RawDescriptor: file_src_mathOperations_mathOperations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_mathOperations_mathOperations_proto_goTypes,
		DependencyIndexes: file_src_mathOperations_mathOperations_proto_depIdxs,
		MessageInfos:      file_src_mathOperations_mathOperations_proto_msgTypes,
	}.Build()
	File_src_mathOperations_mathOperations_proto = out.File
	file_src_mathOperations_mathOperations_proto_rawDesc = nil
	file_src_mathOperations_mathOperations_proto_goTypes = nil
	file_src_mathOperations_mathOperations_proto_depIdxs = nil
}
