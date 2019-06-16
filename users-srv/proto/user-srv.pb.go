// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user-srv.proto

package user_srv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateRequest struct {
	UserID               string    `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Update               *UserMain `protobuf:"bytes,2,opt,name=update,proto3" json:"update,omitempty"`
	Collection           string    `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{0}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateRequest) GetUpdate() *UserMain {
	if m != nil {
		return m.Update
	}
	return nil
}

func (m *UpdateRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

type AddRequest struct {
	MainData             *UserMain  `protobuf:"bytes,1,opt,name=mainData,proto3" json:"mainData,omitempty"`
	ExtraData            *UserExtra `protobuf:"bytes,2,opt,name=extraData,proto3" json:"extraData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetMainData() *UserMain {
	if m != nil {
		return m.MainData
	}
	return nil
}

func (m *AddRequest) GetExtraData() *UserExtra {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

type UserMain struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserMain) Reset()         { *m = UserMain{} }
func (m *UserMain) String() string { return proto.CompactTextString(m) }
func (*UserMain) ProtoMessage()    {}
func (*UserMain) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{2}
}

func (m *UserMain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMain.Unmarshal(m, b)
}
func (m *UserMain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMain.Marshal(b, m, deterministic)
}
func (m *UserMain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMain.Merge(m, src)
}
func (m *UserMain) XXX_Size() int {
	return xxx_messageInfo_UserMain.Size(m)
}
func (m *UserMain) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMain.DiscardUnknown(m)
}

var xxx_messageInfo_UserMain proto.InternalMessageInfo

func (m *UserMain) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UserMain) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserMain) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserMain) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserMain) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserExtra struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Gender               string   `protobuf:"bytes,4,opt,name=gender,proto3" json:"gender,omitempty"`
	BirthdayUTC          int64    `protobuf:"varint,5,opt,name=birthdayUTC,proto3" json:"birthdayUTC,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserExtra) Reset()         { *m = UserExtra{} }
func (m *UserExtra) String() string { return proto.CompactTextString(m) }
func (*UserExtra) ProtoMessage()    {}
func (*UserExtra) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{3}
}

func (m *UserExtra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserExtra.Unmarshal(m, b)
}
func (m *UserExtra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserExtra.Marshal(b, m, deterministic)
}
func (m *UserExtra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserExtra.Merge(m, src)
}
func (m *UserExtra) XXX_Size() int {
	return xxx_messageInfo_UserExtra.Size(m)
}
func (m *UserExtra) XXX_DiscardUnknown() {
	xxx_messageInfo_UserExtra.DiscardUnknown(m)
}

var xxx_messageInfo_UserExtra proto.InternalMessageInfo

func (m *UserExtra) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserExtra) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserExtra) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UserExtra) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserExtra) GetBirthdayUTC() int64 {
	if m != nil {
		return m.BirthdayUTC
	}
	return 0
}

func (m *UserExtra) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddResponse struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{4}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *AddResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{5}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetByUsernameOrEmailRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByUsernameOrEmailRequest) Reset()         { *m = GetByUsernameOrEmailRequest{} }
func (m *GetByUsernameOrEmailRequest) String() string { return proto.CompactTextString(m) }
func (*GetByUsernameOrEmailRequest) ProtoMessage()    {}
func (*GetByUsernameOrEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{6}
}

func (m *GetByUsernameOrEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByUsernameOrEmailRequest.Unmarshal(m, b)
}
func (m *GetByUsernameOrEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByUsernameOrEmailRequest.Marshal(b, m, deterministic)
}
func (m *GetByUsernameOrEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByUsernameOrEmailRequest.Merge(m, src)
}
func (m *GetByUsernameOrEmailRequest) XXX_Size() int {
	return xxx_messageInfo_GetByUsernameOrEmailRequest.Size(m)
}
func (m *GetByUsernameOrEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByUsernameOrEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByUsernameOrEmailRequest proto.InternalMessageInfo

func (m *GetByUsernameOrEmailRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetByUsernameOrEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type AuthRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthRequest) Reset()         { *m = AuthRequest{} }
func (m *AuthRequest) String() string { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()    {}
func (*AuthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{7}
}

func (m *AuthRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthRequest.Unmarshal(m, b)
}
func (m *AuthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthRequest.Marshal(b, m, deterministic)
}
func (m *AuthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthRequest.Merge(m, src)
}
func (m *AuthRequest) XXX_Size() int {
	return xxx_messageInfo_AuthRequest.Size(m)
}
func (m *AuthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthRequest proto.InternalMessageInfo

func (m *AuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AuthRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	Valid                bool     `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserID               string   `protobuf:"bytes,3,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthResponse) Reset()         { *m = AuthResponse{} }
func (m *AuthResponse) String() string { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()    {}
func (*AuthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{8}
}

func (m *AuthResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthResponse.Unmarshal(m, b)
}
func (m *AuthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthResponse.Marshal(b, m, deterministic)
}
func (m *AuthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthResponse.Merge(m, src)
}
func (m *AuthResponse) XXX_Size() int {
	return xxx_messageInfo_AuthResponse.Size(m)
}
func (m *AuthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthResponse proto.InternalMessageInfo

func (m *AuthResponse) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *AuthResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *AuthResponse) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type UpdateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Delete Function
type DeleteRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *DeleteRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *DeleteRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type DeleteResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7318836fc8ef5002, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateRequest)(nil), "UpdateRequest")
	proto.RegisterType((*AddRequest)(nil), "AddRequest")
	proto.RegisterType((*UserMain)(nil), "UserMain")
	proto.RegisterType((*UserExtra)(nil), "UserExtra")
	proto.RegisterType((*AddResponse)(nil), "AddResponse")
	proto.RegisterType((*GetRequest)(nil), "GetRequest")
	proto.RegisterType((*GetByUsernameOrEmailRequest)(nil), "GetByUsernameOrEmailRequest")
	proto.RegisterType((*AuthRequest)(nil), "AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "AuthResponse")
	proto.RegisterType((*UpdateResponse)(nil), "UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "DeleteResponse")
}

func init() { proto.RegisterFile("user-srv.proto", fileDescriptor_7318836fc8ef5002) }

var fileDescriptor_7318836fc8ef5002 = []byte{
	// 553 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x4e, 0x9a, 0x35, 0x34, 0x27, 0x6d, 0x27, 0x59, 0x15, 0x8a, 0xc2, 0x04, 0xc5, 0x02, 0xa9,
	0xda, 0x84, 0x85, 0xc6, 0x03, 0xa0, 0x42, 0xa7, 0x8a, 0x0b, 0x98, 0x14, 0x51, 0x6e, 0x10, 0x42,
	0xde, 0x62, 0xb6, 0xa0, 0x34, 0x29, 0xb6, 0x3b, 0xd8, 0x2b, 0xf0, 0x36, 0x3c, 0x16, 0x6f, 0x81,
	0x1c, 0x27, 0x8d, 0x8d, 0x68, 0x8b, 0x76, 0x97, 0xf3, 0xff, 0xf9, 0xfb, 0xce, 0x09, 0x0c, 0xd7,
	0x82, 0xf1, 0x67, 0x82, 0xdf, 0x90, 0x15, 0x2f, 0x65, 0x89, 0xbf, 0xc2, 0x60, 0xb1, 0x4a, 0xa9,
	0x64, 0x09, 0xfb, 0xb6, 0x66, 0x42, 0xa2, 0xfb, 0xe0, 0xab, 0x94, 0x37, 0xb3, 0xc8, 0x1d, 0xbb,
	0x93, 0x20, 0xa9, 0x2d, 0xf4, 0x18, 0xfc, 0x75, 0x95, 0x18, 0x75, 0xc6, 0xee, 0x24, 0x3c, 0x0d,
	0xc8, 0x42, 0x30, 0xfe, 0x96, 0x66, 0x45, 0x52, 0x07, 0xd0, 0x43, 0x80, 0xcb, 0x32, 0xcf, 0xd9,
	0xa5, 0xcc, 0xca, 0x22, 0xf2, 0xaa, 0x72, 0xc3, 0x83, 0x3f, 0x01, 0x4c, 0xd3, 0xb4, 0x19, 0xf4,
	0x14, 0x7a, 0x4b, 0x9a, 0x15, 0x33, 0x2a, 0x69, 0x35, 0xca, 0x6a, 0xb9, 0x09, 0xa1, 0x09, 0x04,
	0xec, 0x87, 0xe4, 0xb4, 0xca, 0xd3, 0xa3, 0xa1, 0xca, 0x3b, 0x53, 0xde, 0xa4, 0x0d, 0xe2, 0x9f,
	0x2e, 0xf4, 0x9a, 0x06, 0x5b, 0x9f, 0x11, 0x43, 0x4f, 0x7d, 0x15, 0x74, 0xa9, 0x1f, 0x12, 0x24,
	0x1b, 0x1b, 0x8d, 0xa0, 0xcb, 0x96, 0x34, 0xcb, 0x6b, 0xe8, 0xda, 0x50, 0x15, 0x2b, 0x2a, 0xc4,
	0xf7, 0x92, 0xa7, 0xd1, 0x81, 0xae, 0x68, 0x6c, 0x14, 0xc1, 0xbd, 0x25, 0x13, 0x82, 0x5e, 0xb1,
	0xa8, 0x5b, 0x85, 0x1a, 0x13, 0xff, 0x72, 0x21, 0xd8, 0xa0, 0xdc, 0xa0, 0x49, 0x2d, 0x34, 0x29,
	0x3a, 0x82, 0xe0, 0x4b, 0xc6, 0x85, 0x7c, 0xd7, 0xc2, 0x69, 0x1d, 0x6a, 0x72, 0x4e, 0xeb, 0xa0,
	0x86, 0xb4, 0xb1, 0x55, 0xc7, 0x2b, 0x56, 0xa4, 0x8c, 0xd7, 0x98, 0x6a, 0x0b, 0x8d, 0x21, 0xbc,
	0xc8, 0xb8, 0xbc, 0x4e, 0xe9, 0xed, 0xe2, 0xfd, 0xeb, 0x0a, 0x95, 0x97, 0x98, 0x2e, 0x13, 0xb3,
	0x6f, 0x63, 0x7e, 0x09, 0x61, 0xa5, 0x8f, 0x58, 0x95, 0x85, 0x60, 0x5b, 0x29, 0x34, 0x1a, 0x74,
	0xec, 0x06, 0x4f, 0x00, 0xe6, 0x4c, 0xee, 0xd9, 0x24, 0x7c, 0x0e, 0x0f, 0xe6, 0x4c, 0xbe, 0xba,
	0x5d, 0xd4, 0xbc, 0x9f, 0xf3, 0x33, 0x45, 0x74, 0x53, 0x66, 0x2a, 0xe4, 0x6e, 0x53, 0xa8, 0x63,
	0x28, 0x84, 0x3f, 0x42, 0x38, 0x5d, 0xcb, 0xeb, 0x3b, 0x37, 0xb0, 0x24, 0xf6, 0x6c, 0x89, 0xf1,
	0x07, 0xe8, 0xeb, 0xe6, 0x35, 0x2b, 0x23, 0xe8, 0xde, 0xd0, 0x3c, 0xd3, 0x4a, 0xf6, 0x12, 0x6d,
	0x6c, 0xe7, 0xc4, 0x60, 0xc1, 0xb3, 0x58, 0x38, 0x86, 0x61, 0x73, 0x78, 0x75, 0x67, 0xa3, 0x87,
	0x6b, 0xf3, 0xfa, 0x19, 0x06, 0x33, 0x96, 0xb3, 0xfd, 0x47, 0xba, 0x6b, 0xbb, 0x77, 0x3d, 0xf2,
	0x18, 0x86, 0xcd, 0x80, 0x7d, 0x60, 0x4e, 0x7f, 0x77, 0xa0, 0xab, 0xa4, 0x13, 0x08, 0x83, 0x37,
	0x4d, 0x53, 0x14, 0x92, 0xf6, 0xaa, 0xe3, 0x3e, 0x31, 0x56, 0x08, 0x3b, 0xe8, 0x11, 0x78, 0x73,
	0x26, 0x51, 0x48, 0xda, 0xc5, 0x88, 0xdb, 0x3b, 0xc7, 0x8e, 0xfa, 0x0d, 0xcc, 0x99, 0xd4, 0x67,
	0x62, 0x65, 0x19, 0x57, 0x8e, 0x1d, 0x34, 0x85, 0xd1, 0xbf, 0x96, 0x06, 0x1d, 0x91, 0x1d, 0xbb,
	0xf4, 0xf7, 0xa4, 0x03, 0xa5, 0x24, 0xea, 0x13, 0x63, 0x5b, 0xe2, 0x01, 0x31, 0xe5, 0xc5, 0x0e,
	0x3a, 0x01, 0x5f, 0x0b, 0x83, 0x86, 0xc4, 0xfa, 0x35, 0xc6, 0x87, 0xc4, 0x56, 0x0c, 0x3b, 0xe8,
	0x39, 0x84, 0xda, 0xa7, 0x1f, 0xf0, 0x1f, 0x15, 0x27, 0xe0, 0x6b, 0xaa, 0xd1, 0x90, 0x58, 0xa2,
	0xc6, 0x87, 0xc4, 0xd6, 0x00, 0x3b, 0x17, 0x7e, 0xf5, 0x93, 0x7e, 0xf1, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0x5f, 0xfc, 0x5f, 0xb6, 0x05, 0x00, 0x00,
}
