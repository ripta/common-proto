// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.25.1
// source: geometry/geometry.proto

package geometry

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Box struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P1 *Point `protobuf:"bytes,1,opt,name=p1,proto3" json:"p1,omitempty"`
	P2 *Point `protobuf:"bytes,2,opt,name=p2,proto3" json:"p2,omitempty"`
}

func (x *Box) Reset() {
	*x = Box{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_geometry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Box) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Box) ProtoMessage() {}

func (x *Box) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_geometry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Box.ProtoReflect.Descriptor instead.
func (*Box) Descriptor() ([]byte, []int) {
	return file_geometry_geometry_proto_rawDescGZIP(), []int{0}
}

func (x *Box) GetP1() *Point {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *Box) GetP2() *Point {
	if x != nil {
		return x.P2
	}
	return nil
}

type Circle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Center *Point  `protobuf:"bytes,1,opt,name=center,proto3" json:"center,omitempty"`
	Radius float64 `protobuf:"fixed64,2,opt,name=radius,proto3" json:"radius,omitempty"`
}

func (x *Circle) Reset() {
	*x = Circle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_geometry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Circle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Circle) ProtoMessage() {}

func (x *Circle) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_geometry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Circle.ProtoReflect.Descriptor instead.
func (*Circle) Descriptor() ([]byte, []int) {
	return file_geometry_geometry_proto_rawDescGZIP(), []int{1}
}

func (x *Circle) GetCenter() *Point {
	if x != nil {
		return x.Center
	}
	return nil
}

func (x *Circle) GetRadius() float64 {
	if x != nil {
		return x.Radius
	}
	return 0
}

type Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P1 *Point `protobuf:"bytes,1,opt,name=p1,proto3" json:"p1,omitempty"`
	P2 *Point `protobuf:"bytes,2,opt,name=p2,proto3" json:"p2,omitempty"`
}

func (x *Line) Reset() {
	*x = Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_geometry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Line) ProtoMessage() {}

func (x *Line) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_geometry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Line.ProtoReflect.Descriptor instead.
func (*Line) Descriptor() ([]byte, []int) {
	return file_geometry_geometry_proto_rawDescGZIP(), []int{2}
}

func (x *Line) GetP1() *Point {
	if x != nil {
		return x.P1
	}
	return nil
}

func (x *Line) GetP2() *Point {
	if x != nil {
		return x.P2
	}
	return nil
}

type Point struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float64 `protobuf:"fixed64,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float64 `protobuf:"fixed64,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point) Reset() {
	*x = Point{}
	if protoimpl.UnsafeEnabled {
		mi := &file_geometry_geometry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point) ProtoMessage() {}

func (x *Point) ProtoReflect() protoreflect.Message {
	mi := &file_geometry_geometry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point.ProtoReflect.Descriptor instead.
func (*Point) Descriptor() ([]byte, []int) {
	return file_geometry_geometry_proto_rawDescGZIP(), []int{3}
}

func (x *Point) GetX() float64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point) GetY() float64 {
	if x != nil {
		return x.Y
	}
	return 0
}

var File_geometry_geometry_proto protoreflect.FileDescriptor

var file_geometry_geometry_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x22, 0x47, 0x0a, 0x03, 0x42, 0x6f, 0x78, 0x12, 0x1f, 0x0a, 0x02, 0x70, 0x31,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x70, 0x31, 0x12, 0x1f, 0x0a, 0x02, 0x70,
	0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x70, 0x32, 0x22, 0x49, 0x0a, 0x06,
	0x43, 0x69, 0x72, 0x63, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72,
	0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x65, 0x12,
	0x1f, 0x0a, 0x02, 0x70, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x65,
	0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x70, 0x31,
	0x12, 0x1f, 0x0a, 0x02, 0x70, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67,
	0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x02, 0x70,
	0x32, 0x22, 0x23, 0x0a, 0x05, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x01, 0x79, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x70, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_geometry_geometry_proto_rawDescOnce sync.Once
	file_geometry_geometry_proto_rawDescData = file_geometry_geometry_proto_rawDesc
)

func file_geometry_geometry_proto_rawDescGZIP() []byte {
	file_geometry_geometry_proto_rawDescOnce.Do(func() {
		file_geometry_geometry_proto_rawDescData = protoimpl.X.CompressGZIP(file_geometry_geometry_proto_rawDescData)
	})
	return file_geometry_geometry_proto_rawDescData
}

var file_geometry_geometry_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_geometry_geometry_proto_goTypes = []interface{}{
	(*Box)(nil),    // 0: geometry.Box
	(*Circle)(nil), // 1: geometry.Circle
	(*Line)(nil),   // 2: geometry.Line
	(*Point)(nil),  // 3: geometry.Point
}
var file_geometry_geometry_proto_depIdxs = []int32{
	3, // 0: geometry.Box.p1:type_name -> geometry.Point
	3, // 1: geometry.Box.p2:type_name -> geometry.Point
	3, // 2: geometry.Circle.center:type_name -> geometry.Point
	3, // 3: geometry.Line.p1:type_name -> geometry.Point
	3, // 4: geometry.Line.p2:type_name -> geometry.Point
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_geometry_geometry_proto_init() }
func file_geometry_geometry_proto_init() {
	if File_geometry_geometry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_geometry_geometry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Box); i {
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
		file_geometry_geometry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Circle); i {
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
		file_geometry_geometry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Line); i {
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
		file_geometry_geometry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point); i {
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
			RawDescriptor: file_geometry_geometry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_geometry_geometry_proto_goTypes,
		DependencyIndexes: file_geometry_geometry_proto_depIdxs,
		MessageInfos:      file_geometry_geometry_proto_msgTypes,
	}.Build()
	File_geometry_geometry_proto = out.File
	file_geometry_geometry_proto_rawDesc = nil
	file_geometry_geometry_proto_goTypes = nil
	file_geometry_geometry_proto_depIdxs = nil
}
