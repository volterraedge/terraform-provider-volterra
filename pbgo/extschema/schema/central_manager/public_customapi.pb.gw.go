// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ves.io/schema/central_manager/public_customapi.proto

/*
Package central_manager is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package central_manager

import (
	"context"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_CustomAPI_List_0(ctx context.Context, marshaler runtime.Marshaler, client CustomAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListRequest
	var metadata runtime.ServerMetadata

	msg, err := client.List(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomAPI_List_0(ctx context.Context, marshaler runtime.Marshaler, server CustomAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListRequest
	var metadata runtime.ServerMetadata

	msg, err := server.List(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCustomAPIHandlerServer registers the http handlers for service CustomAPI to "mux".
// UnaryRPC     :call CustomAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterCustomAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CustomAPIServer) error {

	mux.Handle("GET", pattern_CustomAPI_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomAPI_List_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomAPI_List_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCustomAPIHandlerFromEndpoint is same as RegisterCustomAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCustomAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCustomAPIHandler(ctx, mux, conn)
}

// RegisterCustomAPIHandler registers the http handlers for service CustomAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIClient(conn))
}

// RegisterCustomAPIHandlerClient registers the http handlers for service CustomAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CustomAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CustomAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CustomAPIClient" to call the correct interceptors.
func RegisterCustomAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CustomAPIClient) error {

	mux.Handle("GET", pattern_CustomAPI_List_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomAPI_List_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomAPI_List_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CustomAPI_List_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"public", "central_manager", "managers"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CustomAPI_List_0 = runtime.ForwardResponseMessage
)
