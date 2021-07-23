// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: 68ba2c3132
// Version Date: Tue Jun  8 17:59:18 UTC 2021

package svc

// This file contains methods to make individual endpoints from services,
// request and response types to serve those endpoints, as well as encoders and
// decoders for those types, for all of our supported transport serialization
// formats.

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"

	pb "core/proto"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.
//
// In a server, it's useful for functions that need to operate on a per-endpoint
// basis. For example, you might pass an Endpoints to a function that produces
// an http.Handler, with each method (endpoint) wired up to a specific path. (It
// is probably a mistake in design to invoke the Service methods on the
// Endpoints struct in a server.)
//
// In a client, it's useful to collect individually constructed endpoints into a
// single type that implements the Service interface. For example, you might
// construct individual endpoints using transport/http.NewClient, combine them into an Endpoints, and return it to the caller as a Service.
type Endpoints struct {
	EchoEndpoint      endpoint.Endpoint
	LouderEndpoint    endpoint.Endpoint
	LouderGetEndpoint endpoint.Endpoint
}

// Endpoints

func (e Endpoints) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	response, err := e.EchoEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.EchoResponse), nil
}

func (e Endpoints) Louder(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	response, err := e.LouderEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.EchoResponse), nil
}

func (e Endpoints) LouderGet(ctx context.Context, in *pb.LouderRequest) (*pb.EchoResponse, error) {
	response, err := e.LouderGetEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.EchoResponse), nil
}

// Make Endpoints

func MakeEchoEndpoint(s pb.EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.EchoRequest)
		v, err := s.Echo(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeLouderEndpoint(s pb.EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LouderRequest)
		v, err := s.Louder(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func MakeLouderGetEndpoint(s pb.EchoServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.LouderRequest)
		v, err := s.LouderGet(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

// WrapAllExcept wraps each Endpoint field of struct Endpoints with a
// go-kit/kit/endpoint.Middleware.
// Use this for applying a set of middlewares to every endpoint in the service.
// Optionally, endpoints can be passed in by name to be excluded from being wrapped.
// WrapAllExcept(middleware, "Status", "Ping")
func (e *Endpoints) WrapAllExcept(middleware endpoint.Middleware, excluded ...string) {
	included := map[string]struct{}{
		"Echo":      {},
		"Louder":    {},
		"LouderGet": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Echo" {
			e.EchoEndpoint = middleware(e.EchoEndpoint)
		}
		if inc == "Louder" {
			e.LouderEndpoint = middleware(e.LouderEndpoint)
		}
		if inc == "LouderGet" {
			e.LouderGetEndpoint = middleware(e.LouderGetEndpoint)
		}
	}
}

// LabeledMiddleware will get passed the endpoint name when passed to
// WrapAllLabeledExcept, this can be used to write a generic metrics
// middleware which can send the endpoint name to the metrics collector.
type LabeledMiddleware func(string, endpoint.Endpoint) endpoint.Endpoint

// WrapAllLabeledExcept wraps each Endpoint field of struct Endpoints with a
// LabeledMiddleware, which will receive the name of the endpoint. See
// LabeldMiddleware. See method WrapAllExept for details on excluded
// functionality.
func (e *Endpoints) WrapAllLabeledExcept(middleware func(string, endpoint.Endpoint) endpoint.Endpoint, excluded ...string) {
	included := map[string]struct{}{
		"Echo":      {},
		"Louder":    {},
		"LouderGet": {},
	}

	for _, ex := range excluded {
		if _, ok := included[ex]; !ok {
			panic(fmt.Sprintf("Excluded endpoint '%s' does not exist; see middlewares/endpoints.go", ex))
		}
		delete(included, ex)
	}

	for inc := range included {
		if inc == "Echo" {
			e.EchoEndpoint = middleware("Echo", e.EchoEndpoint)
		}
		if inc == "Louder" {
			e.LouderEndpoint = middleware("Louder", e.LouderEndpoint)
		}
		if inc == "LouderGet" {
			e.LouderGetEndpoint = middleware("LouderGet", e.LouderGetEndpoint)
		}
	}
}
