// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.ext.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type RegistAccountReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistAccountReq) Reset()         { *m = RegistAccountReq{} }
func (m *RegistAccountReq) String() string { return proto.CompactTextString(m) }
func (*RegistAccountReq) ProtoMessage()    {}
func (*RegistAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{0}
}

func (m *RegistAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistAccountReq.Unmarshal(m, b)
}
func (m *RegistAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistAccountReq.Marshal(b, m, deterministic)
}
func (m *RegistAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistAccountReq.Merge(m, src)
}
func (m *RegistAccountReq) XXX_Size() int {
	return xxx_messageInfo_RegistAccountReq.Size(m)
}
func (m *RegistAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistAccountReq proto.InternalMessageInfo

func (m *RegistAccountReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegistAccountReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegistAccountReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegistAccountRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistAccountRes) Reset()         { *m = RegistAccountRes{} }
func (m *RegistAccountRes) String() string { return proto.CompactTextString(m) }
func (*RegistAccountRes) ProtoMessage()    {}
func (*RegistAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{1}
}

func (m *RegistAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistAccountRes.Unmarshal(m, b)
}
func (m *RegistAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistAccountRes.Marshal(b, m, deterministic)
}
func (m *RegistAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistAccountRes.Merge(m, src)
}
func (m *RegistAccountRes) XXX_Size() int {
	return xxx_messageInfo_RegistAccountRes.Size(m)
}
func (m *RegistAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_RegistAccountRes proto.InternalMessageInfo

type LoginAccountReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAccountReq) Reset()         { *m = LoginAccountReq{} }
func (m *LoginAccountReq) String() string { return proto.CompactTextString(m) }
func (*LoginAccountReq) ProtoMessage()    {}
func (*LoginAccountReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{2}
}

func (m *LoginAccountReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAccountReq.Unmarshal(m, b)
}
func (m *LoginAccountReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAccountReq.Marshal(b, m, deterministic)
}
func (m *LoginAccountReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAccountReq.Merge(m, src)
}
func (m *LoginAccountReq) XXX_Size() int {
	return xxx_messageInfo_LoginAccountReq.Size(m)
}
func (m *LoginAccountReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAccountReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAccountReq proto.InternalMessageInfo

func (m *LoginAccountReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginAccountReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginAccountRes struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginAccountRes) Reset()         { *m = LoginAccountRes{} }
func (m *LoginAccountRes) String() string { return proto.CompactTextString(m) }
func (*LoginAccountRes) ProtoMessage()    {}
func (*LoginAccountRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{3}
}

func (m *LoginAccountRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAccountRes.Unmarshal(m, b)
}
func (m *LoginAccountRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAccountRes.Marshal(b, m, deterministic)
}
func (m *LoginAccountRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAccountRes.Merge(m, src)
}
func (m *LoginAccountRes) XXX_Size() int {
	return xxx_messageInfo_LoginAccountRes.Size(m)
}
func (m *LoginAccountRes) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAccountRes.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAccountRes proto.InternalMessageInfo

func (m *LoginAccountRes) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *LoginAccountRes) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *LoginAccountRes) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginAccountRes) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type WantScoreReq struct {
	UserId               int64    `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MovieId              int64    `protobuf:"varint,2,opt,name=movieId,proto3" json:"movieId,omitempty"`
	Score                int64    `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WantScoreReq) Reset()         { *m = WantScoreReq{} }
func (m *WantScoreReq) String() string { return proto.CompactTextString(m) }
func (*WantScoreReq) ProtoMessage()    {}
func (*WantScoreReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{4}
}

func (m *WantScoreReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WantScoreReq.Unmarshal(m, b)
}
func (m *WantScoreReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WantScoreReq.Marshal(b, m, deterministic)
}
func (m *WantScoreReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WantScoreReq.Merge(m, src)
}
func (m *WantScoreReq) XXX_Size() int {
	return xxx_messageInfo_WantScoreReq.Size(m)
}
func (m *WantScoreReq) XXX_DiscardUnknown() {
	xxx_messageInfo_WantScoreReq.DiscardUnknown(m)
}

var xxx_messageInfo_WantScoreReq proto.InternalMessageInfo

func (m *WantScoreReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *WantScoreReq) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func (m *WantScoreReq) GetScore() int64 {
	if m != nil {
		return m.Score
	}
	return 0
}

type WantScoreRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WantScoreRes) Reset()         { *m = WantScoreRes{} }
func (m *WantScoreRes) String() string { return proto.CompactTextString(m) }
func (*WantScoreRes) ProtoMessage()    {}
func (*WantScoreRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{5}
}

func (m *WantScoreRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WantScoreRes.Unmarshal(m, b)
}
func (m *WantScoreRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WantScoreRes.Marshal(b, m, deterministic)
}
func (m *WantScoreRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WantScoreRes.Merge(m, src)
}
func (m *WantScoreRes) XXX_Size() int {
	return xxx_messageInfo_WantScoreRes.Size(m)
}
func (m *WantScoreRes) XXX_DiscardUnknown() {
	xxx_messageInfo_WantScoreRes.DiscardUnknown(m)
}

var xxx_messageInfo_WantScoreRes proto.InternalMessageInfo

type UpdateUserProfileReq struct {
	UserImage            string   `protobuf:"bytes,1,opt,name=userImage,proto3" json:"userImage,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail            string   `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserPhone            string   `protobuf:"bytes,4,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	UserId               int64    `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserProfileReq) Reset()         { *m = UpdateUserProfileReq{} }
func (m *UpdateUserProfileReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileReq) ProtoMessage()    {}
func (*UpdateUserProfileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{6}
}

func (m *UpdateUserProfileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileReq.Unmarshal(m, b)
}
func (m *UpdateUserProfileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserProfileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileReq.Merge(m, src)
}
func (m *UpdateUserProfileReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileReq.Size(m)
}
func (m *UpdateUserProfileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileReq proto.InternalMessageInfo

func (m *UpdateUserProfileReq) GetUserImage() string {
	if m != nil {
		return m.UserImage
	}
	return ""
}

func (m *UpdateUserProfileReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateUserProfileReq) GetUserEmail() string {
	if m != nil {
		return m.UserEmail
	}
	return ""
}

func (m *UpdateUserProfileReq) GetUserPhone() string {
	if m != nil {
		return m.UserPhone
	}
	return ""
}

func (m *UpdateUserProfileReq) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type UpdateUserProfileRes struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserProfileRes) Reset()         { *m = UpdateUserProfileRes{} }
func (m *UpdateUserProfileRes) String() string { return proto.CompactTextString(m) }
func (*UpdateUserProfileRes) ProtoMessage()    {}
func (*UpdateUserProfileRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_f85764c91792e488, []int{7}
}

func (m *UpdateUserProfileRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserProfileRes.Unmarshal(m, b)
}
func (m *UpdateUserProfileRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserProfileRes.Marshal(b, m, deterministic)
}
func (m *UpdateUserProfileRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserProfileRes.Merge(m, src)
}
func (m *UpdateUserProfileRes) XXX_Size() int {
	return xxx_messageInfo_UpdateUserProfileRes.Size(m)
}
func (m *UpdateUserProfileRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserProfileRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserProfileRes proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RegistAccountReq)(nil), "pb.RegistAccountReq")
	proto.RegisterType((*RegistAccountRes)(nil), "pb.RegistAccountRes")
	proto.RegisterType((*LoginAccountReq)(nil), "pb.LoginAccountReq")
	proto.RegisterType((*LoginAccountRes)(nil), "pb.LoginAccountRes")
	proto.RegisterType((*WantScoreReq)(nil), "pb.WantScoreReq")
	proto.RegisterType((*WantScoreRes)(nil), "pb.WantScoreRes")
	proto.RegisterType((*UpdateUserProfileReq)(nil), "pb.UpdateUserProfileReq")
	proto.RegisterType((*UpdateUserProfileRes)(nil), "pb.UpdateUserProfileRes")
}

func init() { proto.RegisterFile("user.ext.proto", fileDescriptor_f85764c91792e488) }

var fileDescriptor_f85764c91792e488 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x85, 0xf6, 0xc1, 0x7b, 0xdc, 0xf0, 0x78, 0xbc, 0xb1, 0x21, 0x4d, 0xe3, 0xc2, 0xcc, 0xca,
	0x55, 0x13, 0x75, 0x67, 0xe2, 0xc2, 0x18, 0x16, 0x24, 0xc6, 0x90, 0x12, 0x74, 0x6b, 0x29, 0x57,
	0x6c, 0x42, 0x3b, 0x65, 0x66, 0x40, 0x7e, 0xc1, 0xff, 0xf0, 0x43, 0xcd, 0xcc, 0x40, 0x29, 0xa5,
	0x10, 0x97, 0xe7, 0xdc, 0xb9, 0xe7, 0xdc, 0xd3, 0x7b, 0x0b, 0x9d, 0xa5, 0x40, 0xee, 0xe3, 0x5a,
	0xfa, 0x19, 0x67, 0x92, 0x11, 0x2b, 0x9b, 0xd0, 0x57, 0xe8, 0x06, 0x38, 0x8b, 0x85, 0xbc, 0x8f,
	0x22, 0xb6, 0x4c, 0x65, 0x80, 0x0b, 0xe2, 0x40, 0x03, 0x93, 0x30, 0x9e, 0xbb, 0xf5, 0x8b, 0xfa,
	0x65, 0x2b, 0x30, 0x80, 0x78, 0xf0, 0x47, 0xf5, 0x3f, 0x85, 0x09, 0xba, 0x96, 0x2e, 0xe4, 0x58,
	0xd5, 0xb2, 0x50, 0x88, 0x0f, 0xc6, 0xa7, 0xae, 0x6d, 0x6a, 0x5b, 0x4c, 0xc9, 0x81, 0x83, 0xa0,
	0x0f, 0xf0, 0xef, 0x91, 0xcd, 0xe2, 0xf4, 0x27, 0xa6, 0xb9, 0xb0, 0x55, 0x12, 0x5e, 0x94, 0x45,
	0x04, 0xe9, 0x41, 0x53, 0xcd, 0x34, 0x98, 0x6a, 0x15, 0x3b, 0xd8, 0xa0, 0x93, 0xb3, 0xe7, 0xc6,
	0x76, 0xd1, 0xd8, 0x81, 0x46, 0xf6, 0xce, 0x52, 0x74, 0x7f, 0x19, 0x56, 0x03, 0xfa, 0x0c, 0xed,
	0x97, 0x30, 0x95, 0xa3, 0x88, 0x71, 0x54, 0x43, 0x1f, 0xf3, 0x73, 0xe1, 0x77, 0xc2, 0x56, 0x31,
	0x0e, 0xcc, 0xd4, 0x76, 0xb0, 0x85, 0x4a, 0x57, 0xa8, 0x6e, 0xed, 0x66, 0x07, 0x06, 0xd0, 0xce,
	0x9e, 0xae, 0xa0, 0x5f, 0x75, 0x70, 0xc6, 0xd9, 0x34, 0x94, 0x38, 0x16, 0xc8, 0x87, 0x9c, 0xbd,
	0xc5, 0x73, 0x6d, 0x78, 0x0e, 0x2d, 0x6d, 0x91, 0x84, 0x33, 0xdc, 0x7c, 0xa9, 0x1d, 0x71, 0x32,
	0xe6, 0xa6, 0xb3, 0x5f, 0x88, 0xba, 0x23, 0xb6, 0xd5, 0x61, 0x21, 0xf2, 0x8e, 0x50, 0x31, 0xc7,
	0x26, 0x66, 0xc3, 0xc4, 0x34, 0x88, 0xf6, 0x2a, 0xa7, 0x14, 0xd7, 0x9f, 0x16, 0x74, 0x14, 0x35,
	0x42, 0xbe, 0x8a, 0x23, 0xec, 0xaf, 0x25, 0xb9, 0x83, 0xbf, 0x7b, 0x57, 0x40, 0x1c, 0x3f, 0x9b,
	0xf8, 0xe5, 0xd3, 0xf3, 0xaa, 0x58, 0x41, 0x6b, 0xe4, 0x16, 0xda, 0xc5, 0x5d, 0x93, 0x33, 0xf5,
	0xae, 0x74, 0x42, 0x5e, 0x05, 0xa9, 0x7a, 0xaf, 0xa0, 0x95, 0x7f, 0x5c, 0xd2, 0x55, 0x6f, 0x8a,
	0x3b, 0xf4, 0xca, 0x8c, 0x6a, 0x19, 0xc0, 0xff, 0x83, 0x60, 0xc4, 0x55, 0x0f, 0xab, 0xb6, 0xe2,
	0x1d, 0xab, 0x08, 0x5a, 0x9b, 0x34, 0xf5, 0xbf, 0x76, 0xf3, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x0a,
	0x94, 0x68, 0x63, 0x7d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserServiceExt service

type UserServiceExtClient interface {
	//注册用户
	RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRes, error)
	//用户登录
	LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRes, error)
	//评分
	WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRes, error)
	//修改用户信息
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...client.CallOption) (*UpdateUserProfileRes, error)
}

type userServiceExtClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceExtClient(serviceName string, c client.Client) UserServiceExtClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "pb"
	}
	return &userServiceExtClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceExtClient) RegistAccount(ctx context.Context, in *RegistAccountReq, opts ...client.CallOption) (*RegistAccountRes, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.RegistAccount", in)
	out := new(RegistAccountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) LoginAccount(ctx context.Context, in *LoginAccountReq, opts ...client.CallOption) (*LoginAccountRes, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.LoginAccount", in)
	out := new(LoginAccountRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) WantScore(ctx context.Context, in *WantScoreReq, opts ...client.CallOption) (*WantScoreRes, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.WantScore", in)
	out := new(WantScoreRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceExtClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, opts ...client.CallOption) (*UpdateUserProfileRes, error) {
	req := c.c.NewRequest(c.serviceName, "UserServiceExt.UpdateUserProfile", in)
	out := new(UpdateUserProfileRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserServiceExt service

type UserServiceExtHandler interface {
	//注册用户
	RegistAccount(context.Context, *RegistAccountReq, *RegistAccountRes) error
	//用户登录
	LoginAccount(context.Context, *LoginAccountReq, *LoginAccountRes) error
	//评分
	WantScore(context.Context, *WantScoreReq, *WantScoreRes) error
	//修改用户信息
	UpdateUserProfile(context.Context, *UpdateUserProfileReq, *UpdateUserProfileRes) error
}

func RegisterUserServiceExtHandler(s server.Server, hdlr UserServiceExtHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserServiceExt{hdlr}, opts...))
}

type UserServiceExt struct {
	UserServiceExtHandler
}

func (h *UserServiceExt) RegistAccount(ctx context.Context, in *RegistAccountReq, out *RegistAccountRes) error {
	return h.UserServiceExtHandler.RegistAccount(ctx, in, out)
}

func (h *UserServiceExt) LoginAccount(ctx context.Context, in *LoginAccountReq, out *LoginAccountRes) error {
	return h.UserServiceExtHandler.LoginAccount(ctx, in, out)
}

func (h *UserServiceExt) WantScore(ctx context.Context, in *WantScoreReq, out *WantScoreRes) error {
	return h.UserServiceExtHandler.WantScore(ctx, in, out)
}

func (h *UserServiceExt) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileReq, out *UpdateUserProfileRes) error {
	return h.UserServiceExtHandler.UpdateUserProfile(ctx, in, out)
}