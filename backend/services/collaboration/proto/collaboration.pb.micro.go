// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/collaboration.proto

package collaboration

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

// Api Endpoints for Collaboration service

func NewCollaborationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Collaboration service

type CollaborationService interface {
	CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*SuccessResponse, error)
	GetGroupUserRole(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupRoleResponse, error)
	GetUserGroups(ctx context.Context, in *UserGroupsRequest, opts ...client.CallOption) (*UserGroupsResponse, error)
	FindGroupByPublicID(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error)
}

type collaborationService struct {
	c    client.Client
	name string
}

func NewCollaborationService(name string, c client.Client) CollaborationService {
	return &collaborationService{
		c:    c,
		name: name,
	}
}

func (c *collaborationService) CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, opts ...client.CallOption) (*SuccessResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.CreateNewGroupWithAdmin", in)
	out := new(SuccessResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetGroupUserRole(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupRoleResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetGroupUserRole", in)
	out := new(GroupRoleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) GetUserGroups(ctx context.Context, in *UserGroupsRequest, opts ...client.CallOption) (*UserGroupsResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.GetUserGroups", in)
	out := new(UserGroupsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collaborationService) FindGroupByPublicID(ctx context.Context, in *GroupRequest, opts ...client.CallOption) (*GroupResponse, error) {
	req := c.c.NewRequest(c.name, "Collaboration.FindGroupByPublicID", in)
	out := new(GroupResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Collaboration service

type CollaborationHandler interface {
	CreateNewGroupWithAdmin(context.Context, *CreateGroupRequest, *SuccessResponse) error
	GetGroupUserRole(context.Context, *GroupRequest, *GroupRoleResponse) error
	GetUserGroups(context.Context, *UserGroupsRequest, *UserGroupsResponse) error
	FindGroupByPublicID(context.Context, *GroupRequest, *GroupResponse) error
}

func RegisterCollaborationHandler(s server.Server, hdlr CollaborationHandler, opts ...server.HandlerOption) error {
	type collaboration interface {
		CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, out *SuccessResponse) error
		GetGroupUserRole(ctx context.Context, in *GroupRequest, out *GroupRoleResponse) error
		GetUserGroups(ctx context.Context, in *UserGroupsRequest, out *UserGroupsResponse) error
		FindGroupByPublicID(ctx context.Context, in *GroupRequest, out *GroupResponse) error
	}
	type Collaboration struct {
		collaboration
	}
	h := &collaborationHandler{hdlr}
	return s.Handle(s.NewHandler(&Collaboration{h}, opts...))
}

type collaborationHandler struct {
	CollaborationHandler
}

func (h *collaborationHandler) CreateNewGroupWithAdmin(ctx context.Context, in *CreateGroupRequest, out *SuccessResponse) error {
	return h.CollaborationHandler.CreateNewGroupWithAdmin(ctx, in, out)
}

func (h *collaborationHandler) GetGroupUserRole(ctx context.Context, in *GroupRequest, out *GroupRoleResponse) error {
	return h.CollaborationHandler.GetGroupUserRole(ctx, in, out)
}

func (h *collaborationHandler) GetUserGroups(ctx context.Context, in *UserGroupsRequest, out *UserGroupsResponse) error {
	return h.CollaborationHandler.GetUserGroups(ctx, in, out)
}

func (h *collaborationHandler) FindGroupByPublicID(ctx context.Context, in *GroupRequest, out *GroupResponse) error {
	return h.CollaborationHandler.FindGroupByPublicID(ctx, in, out)
}
