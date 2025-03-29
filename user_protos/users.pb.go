// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.19.6
// source: user_protos/users.proto

package pachuco_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IdRequest) Reset() {
	*x = IdRequest{}
	mi := &file_user_protos_users_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdRequest) ProtoMessage() {}

func (x *IdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_protos_users_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdRequest.ProtoReflect.Descriptor instead.
func (*IdRequest) Descriptor() ([]byte, []int) {
	return file_user_protos_users_proto_rawDescGZIP(), []int{0}
}

func (x *IdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Email         string                 `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EmailRequest) Reset() {
	*x = EmailRequest{}
	mi := &file_user_protos_users_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailRequest) ProtoMessage() {}

func (x *EmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_protos_users_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailRequest.ProtoReflect.Descriptor instead.
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return file_user_protos_users_proto_rawDescGZIP(), []int{1}
}

func (x *EmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Nickname      string                 `protobuf:"bytes,1,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Balance       float32                `protobuf:"fixed32,5,opt,name=balance,proto3" json:"balance,omitempty"`
	Username      string                 `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	mi := &file_user_protos_users_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_protos_users_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_user_protos_users_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserRequest) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *UserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname      string                 `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	FirstName     string                 `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Balance       float32                `protobuf:"fixed32,6,opt,name=balance,proto3" json:"balance,omitempty"`
	Username      string                 `protobuf:"bytes,7,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	mi := &file_user_protos_users_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_protos_users_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_user_protos_users_proto_rawDescGZIP(), []int{3}
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserResponse) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *UserResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	User          *UserRequest           `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_user_protos_users_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_protos_users_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_protos_users_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateUserRequest) GetUser() *UserRequest {
	if x != nil {
		return x.User
	}
	return nil
}

var File_user_protos_users_proto protoreflect.FileDescriptor

const file_user_protos_users_proto_rawDesc = "" +
	"\n" +
	"\x17user_protos/users.proto\x12\rpachuco_proto\"\x1b\n" +
	"\tIdRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\"$\n" +
	"\fEmailRequest\x12\x14\n" +
	"\x05email\x18\x01 \x01(\tR\x05email\"\xb1\x01\n" +
	"\vUserRequest\x12\x1a\n" +
	"\bnickname\x18\x01 \x01(\tR\bnickname\x12\x14\n" +
	"\x05email\x18\x02 \x01(\tR\x05email\x12\x1d\n" +
	"\n" +
	"first_name\x18\x03 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x04 \x01(\tR\blastName\x12\x18\n" +
	"\abalance\x18\x05 \x01(\x02R\abalance\x12\x1a\n" +
	"\busername\x18\x06 \x01(\tR\busername\"\xc2\x01\n" +
	"\fUserResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1a\n" +
	"\bnickname\x18\x02 \x01(\tR\bnickname\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1d\n" +
	"\n" +
	"first_name\x18\x04 \x01(\tR\tfirstName\x12\x1b\n" +
	"\tlast_name\x18\x05 \x01(\tR\blastName\x12\x18\n" +
	"\abalance\x18\x06 \x01(\x02R\abalance\x12\x1a\n" +
	"\busername\x18\a \x01(\tR\busername\"S\n" +
	"\x11UpdateUserRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12.\n" +
	"\x04user\x18\x02 \x01(\v2\x1a.pachuco_proto.UserRequestR\x04user2\xf2\x02\n" +
	"\x04User\x12@\n" +
	"\aGetUser\x12\x18.pachuco_proto.IdRequest\x1a\x1b.pachuco_proto.UserResponse\x12J\n" +
	"\x0eGetUserByEmail\x12\x1b.pachuco_proto.EmailRequest\x1a\x1b.pachuco_proto.UserResponse\x12B\n" +
	"\aAddUser\x12\x1a.pachuco_proto.UserRequest\x1a\x1b.pachuco_proto.UserResponse\x12K\n" +
	"\n" +
	"UpdateUser\x12 .pachuco_proto.UpdateUserRequest\x1a\x1b.pachuco_proto.UserResponse\x12K\n" +
	"\n" +
	"DeleteUser\x12 .pachuco_proto.UpdateUserRequest\x1a\x1b.pachuco_proto.UserResponseB\"Z github.com/monguis/pachuco-protob\x06proto3"

var (
	file_user_protos_users_proto_rawDescOnce sync.Once
	file_user_protos_users_proto_rawDescData []byte
)

func file_user_protos_users_proto_rawDescGZIP() []byte {
	file_user_protos_users_proto_rawDescOnce.Do(func() {
		file_user_protos_users_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_protos_users_proto_rawDesc), len(file_user_protos_users_proto_rawDesc)))
	})
	return file_user_protos_users_proto_rawDescData
}

var file_user_protos_users_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_user_protos_users_proto_goTypes = []any{
	(*IdRequest)(nil),         // 0: pachuco_proto.IdRequest
	(*EmailRequest)(nil),      // 1: pachuco_proto.EmailRequest
	(*UserRequest)(nil),       // 2: pachuco_proto.UserRequest
	(*UserResponse)(nil),      // 3: pachuco_proto.UserResponse
	(*UpdateUserRequest)(nil), // 4: pachuco_proto.UpdateUserRequest
}
var file_user_protos_users_proto_depIdxs = []int32{
	2, // 0: pachuco_proto.UpdateUserRequest.user:type_name -> pachuco_proto.UserRequest
	0, // 1: pachuco_proto.User.GetUser:input_type -> pachuco_proto.IdRequest
	1, // 2: pachuco_proto.User.GetUserByEmail:input_type -> pachuco_proto.EmailRequest
	2, // 3: pachuco_proto.User.AddUser:input_type -> pachuco_proto.UserRequest
	4, // 4: pachuco_proto.User.UpdateUser:input_type -> pachuco_proto.UpdateUserRequest
	4, // 5: pachuco_proto.User.DeleteUser:input_type -> pachuco_proto.UpdateUserRequest
	3, // 6: pachuco_proto.User.GetUser:output_type -> pachuco_proto.UserResponse
	3, // 7: pachuco_proto.User.GetUserByEmail:output_type -> pachuco_proto.UserResponse
	3, // 8: pachuco_proto.User.AddUser:output_type -> pachuco_proto.UserResponse
	3, // 9: pachuco_proto.User.UpdateUser:output_type -> pachuco_proto.UserResponse
	3, // 10: pachuco_proto.User.DeleteUser:output_type -> pachuco_proto.UserResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_protos_users_proto_init() }
func file_user_protos_users_proto_init() {
	if File_user_protos_users_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_protos_users_proto_rawDesc), len(file_user_protos_users_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_protos_users_proto_goTypes,
		DependencyIndexes: file_user_protos_users_proto_depIdxs,
		MessageInfos:      file_user_protos_users_proto_msgTypes,
	}.Build()
	File_user_protos_users_proto = out.File
	file_user_protos_users_proto_goTypes = nil
	file_user_protos_users_proto_depIdxs = nil
}
