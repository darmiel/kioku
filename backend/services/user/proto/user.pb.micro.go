// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user.proto

package user

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for User service

func NewUserEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for User service

type UserService interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*NameIDResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*NameIDResponse, error)
	GetUserIDFromEmail(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*UserID, error)
	GetUserInformation(ctx context.Context, in *UserInformationRequest, opts ...client.CallOption) (*UserInformationResponse, error)
	GetUserProfileInformation(ctx context.Context, in *UserID, opts ...client.CallOption) (*UserProfileInformationResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Register(ctx context.Context, in *RegisterRequest, opts ...client.CallOption) (*NameIDResponse, error) {
	req := c.c.NewRequest(c.name, "User.Register", in)
	out := new(NameIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*NameIDResponse, error) {
	req := c.c.NewRequest(c.name, "User.Login", in)
	out := new(NameIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserIDFromEmail(ctx context.Context, in *UserIDRequest, opts ...client.CallOption) (*UserID, error) {
	req := c.c.NewRequest(c.name, "User.GetUserIDFromEmail", in)
	out := new(UserID)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserInformation(ctx context.Context, in *UserInformationRequest, opts ...client.CallOption) (*UserInformationResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserInformation", in)
	out := new(UserInformationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetUserProfileInformation(ctx context.Context, in *UserID, opts ...client.CallOption) (*UserProfileInformationResponse, error) {
	req := c.c.NewRequest(c.name, "User.GetUserProfileInformation", in)
	out := new(UserProfileInformationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for User service

type UserHandler interface {
	Register(context.Context, *RegisterRequest, *NameIDResponse) error
	Login(context.Context, *LoginRequest, *NameIDResponse) error
	GetUserIDFromEmail(context.Context, *UserIDRequest, *UserID) error
	GetUserInformation(context.Context, *UserInformationRequest, *UserInformationResponse) error
	GetUserProfileInformation(context.Context, *UserID, *UserProfileInformationResponse) error
}

func RegisterUserHandler(s server.Server, hdlr UserHandler, opts ...server.HandlerOption) error {
	type user interface {
		Register(ctx context.Context, in *RegisterRequest, out *NameIDResponse) error
		Login(ctx context.Context, in *LoginRequest, out *NameIDResponse) error
		GetUserIDFromEmail(ctx context.Context, in *UserIDRequest, out *UserID) error
		GetUserInformation(ctx context.Context, in *UserInformationRequest, out *UserInformationResponse) error
		GetUserProfileInformation(ctx context.Context, in *UserID, out *UserProfileInformationResponse) error
	}
	type User struct {
		user
	}
	h := &userHandler{hdlr}
	return s.Handle(s.NewHandler(&User{h}, opts...))
}

type userHandler struct {
	UserHandler
}

func (h *userHandler) Register(ctx context.Context, in *RegisterRequest, out *NameIDResponse) error {
	return h.UserHandler.Register(ctx, in, out)
}

func (h *userHandler) Login(ctx context.Context, in *LoginRequest, out *NameIDResponse) error {
	return h.UserHandler.Login(ctx, in, out)
}

func (h *userHandler) GetUserIDFromEmail(ctx context.Context, in *UserIDRequest, out *UserID) error {
	return h.UserHandler.GetUserIDFromEmail(ctx, in, out)
}

func (h *userHandler) GetUserInformation(ctx context.Context, in *UserInformationRequest, out *UserInformationResponse) error {
	return h.UserHandler.GetUserInformation(ctx, in, out)
}

func (h *userHandler) GetUserProfileInformation(ctx context.Context, in *UserID, out *UserProfileInformationResponse) error {
	return h.UserHandler.GetUserProfileInformation(ctx, in, out)
}
