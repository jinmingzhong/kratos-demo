// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v3.12.4
// source: calc_gateway/v1/calc_gateway.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCalcGatewaygetSum = "/api.calc_gateway.v1.CalcGateway/getSum"

type CalcGatewayHTTPServer interface {
	GetSum(context.Context, *GetSumRequest) (*GetSumReply, error)
}

func RegisterCalcGatewayHTTPServer(s *http.Server, srv CalcGatewayHTTPServer) {
	r := s.Route("/")
	r.GET("/sum/{a}/{b}", _CalcGateway_GetSum1_HTTP_Handler(srv))
}

func _CalcGateway_GetSum1_HTTP_Handler(srv CalcGatewayHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSumRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCalcGatewaygetSum)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSum(ctx, req.(*GetSumRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSumReply)
		return ctx.Result(200, reply)
	}
}

type CalcGatewayHTTPClient interface {
	GetSum(ctx context.Context, req *GetSumRequest, opts ...http.CallOption) (rsp *GetSumReply, err error)
}

type CalcGatewayHTTPClientImpl struct {
	cc *http.Client
}

func NewCalcGatewayHTTPClient(client *http.Client) CalcGatewayHTTPClient {
	return &CalcGatewayHTTPClientImpl{client}
}

func (c *CalcGatewayHTTPClientImpl) GetSum(ctx context.Context, in *GetSumRequest, opts ...http.CallOption) (*GetSumReply, error) {
	var out GetSumReply
	pattern := "/sum/{a}/{b}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCalcGatewaygetSum))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
