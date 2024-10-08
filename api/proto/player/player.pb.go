// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: player/player.proto

package playerpb

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

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CasesSolved int32  `protobuf:"varint,3,opt,name=cases_solved,json=casesSolved,proto3" json:"cases_solved,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{0}
}

func (x *Player) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Player) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Player) GetCasesSolved() int32 {
	if x != nil {
		return x.CasesSolved
	}
	return 0
}

type CreatePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreatePlayerRequest) Reset() {
	*x = CreatePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerRequest) ProtoMessage() {}

func (x *CreatePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePlayerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{2}
}

func (x *GetPlayerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UpdatePlayerProgressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CaseSolved bool   `protobuf:"varint,2,opt,name=case_solved,json=caseSolved,proto3" json:"case_solved,omitempty"`
}

func (x *UpdatePlayerProgressRequest) Reset() {
	*x = UpdatePlayerProgressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_player_player_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePlayerProgressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePlayerProgressRequest) ProtoMessage() {}

func (x *UpdatePlayerProgressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_player_player_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePlayerProgressRequest.ProtoReflect.Descriptor instead.
func (*UpdatePlayerProgressRequest) Descriptor() ([]byte, []int) {
	return file_player_player_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePlayerProgressRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdatePlayerProgressRequest) GetCaseSolved() bool {
	if x != nil {
		return x.CaseSolved
	}
	return false
}

var File_player_player_proto protoreflect.FileDescriptor

var file_player_player_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x22,
	0x4f, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x61, 0x73, 0x65, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x73, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x22, 0x29, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4e, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x63, 0x61, 0x73, 0x65, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x32,
	0xdc, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x1d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x12, 0x39, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12,
	0x1a, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x4f, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x25, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x1f,
	0x5a, 0x1d, 0x6f, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x3b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_player_player_proto_rawDescOnce sync.Once
	file_player_player_proto_rawDescData = file_player_player_proto_rawDesc
)

func file_player_player_proto_rawDescGZIP() []byte {
	file_player_player_proto_rawDescOnce.Do(func() {
		file_player_player_proto_rawDescData = protoimpl.X.CompressGZIP(file_player_player_proto_rawDescData)
	})
	return file_player_player_proto_rawDescData
}

var file_player_player_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_player_player_proto_goTypes = []any{
	(*Player)(nil),                      // 0: playerpb.Player
	(*CreatePlayerRequest)(nil),         // 1: playerpb.CreatePlayerRequest
	(*GetPlayerRequest)(nil),            // 2: playerpb.GetPlayerRequest
	(*UpdatePlayerProgressRequest)(nil), // 3: playerpb.UpdatePlayerProgressRequest
}
var file_player_player_proto_depIdxs = []int32{
	1, // 0: playerpb.PlayerService.CreatePlayer:input_type -> playerpb.CreatePlayerRequest
	2, // 1: playerpb.PlayerService.GetPlayer:input_type -> playerpb.GetPlayerRequest
	3, // 2: playerpb.PlayerService.UpdatePlayerProgress:input_type -> playerpb.UpdatePlayerProgressRequest
	0, // 3: playerpb.PlayerService.CreatePlayer:output_type -> playerpb.Player
	0, // 4: playerpb.PlayerService.GetPlayer:output_type -> playerpb.Player
	0, // 5: playerpb.PlayerService.UpdatePlayerProgress:output_type -> playerpb.Player
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_player_player_proto_init() }
func file_player_player_proto_init() {
	if File_player_player_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_player_player_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
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
		file_player_player_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreatePlayerRequest); i {
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
		file_player_player_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetPlayerRequest); i {
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
		file_player_player_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdatePlayerProgressRequest); i {
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
			RawDescriptor: file_player_player_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_player_player_proto_goTypes,
		DependencyIndexes: file_player_player_proto_depIdxs,
		MessageInfos:      file_player_player_proto_msgTypes,
	}.Build()
	File_player_player_proto = out.File
	file_player_player_proto_rawDesc = nil
	file_player_player_proto_goTypes = nil
	file_player_player_proto_depIdxs = nil
}
