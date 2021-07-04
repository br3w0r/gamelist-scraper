// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: gamelist.proto

package proto

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

type GameProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Platforms    []*Named `protobuf:"bytes,2,rep,name=platforms,proto3" json:"platforms,omitempty"`
	YearReleased int32    `protobuf:"varint,3,opt,name=year_released,json=yearReleased,proto3" json:"year_released,omitempty"`
	ImageUrl     string   `protobuf:"bytes,4,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Genres       []*Named `protobuf:"bytes,5,rep,name=genres,proto3" json:"genres,omitempty"`
}

func (x *GameProperties) Reset() {
	*x = GameProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamelist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameProperties) ProtoMessage() {}

func (x *GameProperties) ProtoReflect() protoreflect.Message {
	mi := &file_gamelist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameProperties.ProtoReflect.Descriptor instead.
func (*GameProperties) Descriptor() ([]byte, []int) {
	return file_gamelist_proto_rawDescGZIP(), []int{0}
}

func (x *GameProperties) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GameProperties) GetPlatforms() []*Named {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *GameProperties) GetYearReleased() int32 {
	if x != nil {
		return x.YearReleased
	}
	return 0
}

func (x *GameProperties) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *GameProperties) GetGenres() []*Named {
	if x != nil {
		return x.Genres
	}
	return nil
}

type Named struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Named) Reset() {
	*x = Named{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gamelist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Named) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Named) ProtoMessage() {}

func (x *Named) ProtoReflect() protoreflect.Message {
	mi := &file_gamelist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Named.ProtoReflect.Descriptor instead.
func (*Named) Descriptor() ([]byte, []int) {
	return file_gamelist_proto_rawDescGZIP(), []int{1}
}

func (x *Named) GetName() string {
	if x != nil {
		return x.Name
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
		mi := &file_gamelist_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gamelist_proto_msgTypes[2]
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
	return file_gamelist_proto_rawDescGZIP(), []int{2}
}

var File_gamelist_proto protoreflect.FileDescriptor

var file_gamelist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x52,
	0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x79, 0x65,
	0x61, 0x72, 0x5f, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x79, 0x65, 0x61, 0x72, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x06,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x22, 0x1b, 0x0a, 0x05, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x42, 0x0a, 0x0a, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x47, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x30, 0x01, 0x42, 0x2b, 0x5a, 0x29,
	0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x62, 0x72,
	0x33, 0x77, 0x30, 0x72, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gamelist_proto_rawDescOnce sync.Once
	file_gamelist_proto_rawDescData = file_gamelist_proto_rawDesc
)

func file_gamelist_proto_rawDescGZIP() []byte {
	file_gamelist_proto_rawDescOnce.Do(func() {
		file_gamelist_proto_rawDescData = protoimpl.X.CompressGZIP(file_gamelist_proto_rawDescData)
	})
	return file_gamelist_proto_rawDescData
}

var file_gamelist_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gamelist_proto_goTypes = []interface{}{
	(*GameProperties)(nil), // 0: proto.GameProperties
	(*Named)(nil),          // 1: proto.Named
	(*Empty)(nil),          // 2: proto.Empty
}
var file_gamelist_proto_depIdxs = []int32{
	1, // 0: proto.GameProperties.platforms:type_name -> proto.Named
	1, // 1: proto.GameProperties.genres:type_name -> proto.Named
	2, // 2: proto.GameScrape.ScrapeGames:input_type -> proto.Empty
	0, // 3: proto.GameScrape.ScrapeGames:output_type -> proto.GameProperties
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gamelist_proto_init() }
func file_gamelist_proto_init() {
	if File_gamelist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gamelist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameProperties); i {
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
		file_gamelist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Named); i {
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
		file_gamelist_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_gamelist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gamelist_proto_goTypes,
		DependencyIndexes: file_gamelist_proto_depIdxs,
		MessageInfos:      file_gamelist_proto_msgTypes,
	}.Build()
	File_gamelist_proto = out.File
	file_gamelist_proto_rawDesc = nil
	file_gamelist_proto_goTypes = nil
	file_gamelist_proto_depIdxs = nil
}