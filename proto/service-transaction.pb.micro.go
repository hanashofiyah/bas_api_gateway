// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/service-transaction.proto

package proto

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

// Api Endpoints for ServiceTransaction service

func NewServiceTransactionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ServiceTransaction service

type ServiceTransactionService interface {
	Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error)
}

type serviceTransactionService struct {
	c    client.Client
	name string
}

func NewServiceTransactionService(name string, c client.Client) ServiceTransactionService {
	return &serviceTransactionService{
		c:    c,
		name: name,
	}
}

func (c *serviceTransactionService) Login(ctx context.Context, in *LoginRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "ServiceTransaction.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ServiceTransaction service

type ServiceTransactionHandler interface {
	Login(context.Context, *LoginRequest, *LoginResponse) error
}

func RegisterServiceTransactionHandler(s server.Server, hdlr ServiceTransactionHandler, opts ...server.HandlerOption) error {
	type serviceTransaction interface {
		Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error
	}
	type ServiceTransaction struct {
		serviceTransaction
	}
	h := &serviceTransactionHandler{hdlr}
	return s.Handle(s.NewHandler(&ServiceTransaction{h}, opts...))
}

type serviceTransactionHandler struct {
	ServiceTransactionHandler
}

func (h *serviceTransactionHandler) Login(ctx context.Context, in *LoginRequest, out *LoginResponse) error {
	return h.ServiceTransactionHandler.Login(ctx, in, out)
}
