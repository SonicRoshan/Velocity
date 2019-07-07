// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user-srv.proto

package user_srv

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Users service

type UsersService interface {
	Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserMain, error)
	GetExtra(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserExtra, error)
	GetByUsernameOrEmail(ctx context.Context, in *GetByUsernameOrEmailRequest, opts ...client.CallOption) (*UserMain, error)
	Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	UpdateExtra(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Activate(ctx context.Context, in *ActivateRequest, opts ...client.CallOption) (*ActivateResponse, error)
}

type usersService struct {
	c    client.Client
	name string
}

func NewUsersService(name string, c client.Client) UsersService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "users"
	}
	return &usersService{
		c:    c,
		name: name,
	}
}

func (c *usersService) Add(ctx context.Context, in *AddRequest, opts ...client.CallOption) (*AddResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Add", in)
	out := new(AddResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserMain, error) {
	req := c.c.NewRequest(c.name, "Users.Get", in)
	out := new(UserMain)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) GetExtra(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*UserExtra, error) {
	req := c.c.NewRequest(c.name, "Users.GetExtra", in)
	out := new(UserExtra)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) GetByUsernameOrEmail(ctx context.Context, in *GetByUsernameOrEmailRequest, opts ...client.CallOption) (*UserMain, error) {
	req := c.c.NewRequest(c.name, "Users.GetByUsernameOrEmail", in)
	out := new(UserMain)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Auth(ctx context.Context, in *AuthRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Auth", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) UpdateExtra(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.name, "Users.UpdateExtra", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersService) Activate(ctx context.Context, in *ActivateRequest, opts ...client.CallOption) (*ActivateResponse, error) {
	req := c.c.NewRequest(c.name, "Users.Activate", in)
	out := new(ActivateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Users service

type UsersHandler interface {
	Add(context.Context, *AddRequest, *AddResponse) error
	Get(context.Context, *GetRequest, *UserMain) error
	GetExtra(context.Context, *GetRequest, *UserExtra) error
	GetByUsernameOrEmail(context.Context, *GetByUsernameOrEmailRequest, *UserMain) error
	Auth(context.Context, *AuthRequest, *AuthResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	UpdateExtra(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Activate(context.Context, *ActivateRequest, *ActivateResponse) error
}

func RegisterUsersHandler(s server.Server, hdlr UsersHandler, opts ...server.HandlerOption) error {
	type users interface {
		Add(ctx context.Context, in *AddRequest, out *AddResponse) error
		Get(ctx context.Context, in *GetRequest, out *UserMain) error
		GetExtra(ctx context.Context, in *GetRequest, out *UserExtra) error
		GetByUsernameOrEmail(ctx context.Context, in *GetByUsernameOrEmailRequest, out *UserMain) error
		Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error
		Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		UpdateExtra(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Activate(ctx context.Context, in *ActivateRequest, out *ActivateResponse) error
	}
	type Users struct {
		users
	}
	h := &usersHandler{hdlr}
	return s.Handle(s.NewHandler(&Users{h}, opts...))
}

type usersHandler struct {
	UsersHandler
}

func (h *usersHandler) Add(ctx context.Context, in *AddRequest, out *AddResponse) error {
	return h.UsersHandler.Add(ctx, in, out)
}

func (h *usersHandler) Get(ctx context.Context, in *GetRequest, out *UserMain) error {
	return h.UsersHandler.Get(ctx, in, out)
}

func (h *usersHandler) GetExtra(ctx context.Context, in *GetRequest, out *UserExtra) error {
	return h.UsersHandler.GetExtra(ctx, in, out)
}

func (h *usersHandler) GetByUsernameOrEmail(ctx context.Context, in *GetByUsernameOrEmailRequest, out *UserMain) error {
	return h.UsersHandler.GetByUsernameOrEmail(ctx, in, out)
}

func (h *usersHandler) Auth(ctx context.Context, in *AuthRequest, out *AuthResponse) error {
	return h.UsersHandler.Auth(ctx, in, out)
}

func (h *usersHandler) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UsersHandler.Update(ctx, in, out)
}

func (h *usersHandler) UpdateExtra(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.UsersHandler.UpdateExtra(ctx, in, out)
}

func (h *usersHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.UsersHandler.Delete(ctx, in, out)
}

func (h *usersHandler) Activate(ctx context.Context, in *ActivateRequest, out *ActivateResponse) error {
	return h.UsersHandler.Activate(ctx, in, out)
}
