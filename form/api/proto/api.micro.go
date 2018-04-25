// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/examples/form/api/proto/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/examples/form/api/proto/api.proto

It has these top-level messages:
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_api "github.com/micro/go-api/proto"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = go_api.Response{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Form service

type FormService interface {
	// regular form
	Submit(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
	// multipart form
	Multipart(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error)
}

type formService struct {
	c           client.Client
	serviceName string
}

func NewFormService(serviceName string, c client.Client) FormService {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "form"
	}
	return &formService{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *formService) Submit(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Form.Submit", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *formService) Multipart(ctx context.Context, in *go_api.Request, opts ...client.CallOption) (*go_api.Response, error) {
	req := c.c.NewRequest(c.serviceName, "Form.Multipart", in)
	out := new(go_api.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Form service

type FormHandler interface {
	// regular form
	Submit(context.Context, *go_api.Request, *go_api.Response) error
	// multipart form
	Multipart(context.Context, *go_api.Request, *go_api.Response) error
}

func RegisterFormHandler(s server.Server, hdlr FormHandler, opts ...server.HandlerOption) {
	type form interface {
		Submit(ctx context.Context, in *go_api.Request, out *go_api.Response) error
		Multipart(ctx context.Context, in *go_api.Request, out *go_api.Response) error
	}
	type Form struct {
		form
	}
	h := &formHandler{hdlr}
	s.Handle(s.NewHandler(&Form{h}, opts...))
}

type formHandler struct {
	FormHandler
}

func (h *formHandler) Submit(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.FormHandler.Submit(ctx, in, out)
}

func (h *formHandler) Multipart(ctx context.Context, in *go_api.Request, out *go_api.Response) error {
	return h.FormHandler.Multipart(ctx, in, out)
}
