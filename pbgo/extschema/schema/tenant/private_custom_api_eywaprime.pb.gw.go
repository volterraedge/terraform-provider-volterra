// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ves.io/schema/tenant/private_custom_api_eywaprime.proto

/*
Package tenant is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package tenant

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

func request_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(ctx context.Context, marshaler runtime.Marshaler, client CustomPrivateAPIEywaprimeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PrivateDeactivateTenantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.RestrictedDeactivateTenant(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(ctx context.Context, marshaler runtime.Marshaler, server CustomPrivateAPIEywaprimeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PrivateDeactivateTenantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.RestrictedDeactivateTenant(ctx, &protoReq)
	return msg, metadata, err

}

func request_CustomPrivateAPIEywaprime_DeactivateTenant_0(ctx context.Context, marshaler runtime.Marshaler, client CustomPrivateAPIEywaprimeClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PrivateDeactivateTenantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DeactivateTenant(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomPrivateAPIEywaprime_DeactivateTenant_0(ctx context.Context, marshaler runtime.Marshaler, server CustomPrivateAPIEywaprimeServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq PrivateDeactivateTenantRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DeactivateTenant(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCustomPrivateAPIEywaprimeHandlerServer registers the http handlers for service CustomPrivateAPIEywaprime to "mux".
// UnaryRPC     :call CustomPrivateAPIEywaprimeServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterCustomPrivateAPIEywaprimeHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CustomPrivateAPIEywaprimeServer) error {

	mux.Handle("POST", pattern_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CustomPrivateAPIEywaprime_DeactivateTenant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomPrivateAPIEywaprime_DeactivateTenant_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomPrivateAPIEywaprime_DeactivateTenant_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCustomPrivateAPIEywaprimeHandlerFromEndpoint is same as RegisterCustomPrivateAPIEywaprimeHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCustomPrivateAPIEywaprimeHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCustomPrivateAPIEywaprimeHandler(ctx, mux, conn)
}

// RegisterCustomPrivateAPIEywaprimeHandler registers the http handlers for service CustomPrivateAPIEywaprime to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCustomPrivateAPIEywaprimeHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCustomPrivateAPIEywaprimeHandlerClient(ctx, mux, NewCustomPrivateAPIEywaprimeClient(conn))
}

// RegisterCustomPrivateAPIEywaprimeHandlerClient registers the http handlers for service CustomPrivateAPIEywaprime
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CustomPrivateAPIEywaprimeClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CustomPrivateAPIEywaprimeClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CustomPrivateAPIEywaprimeClient" to call the correct interceptors.
func RegisterCustomPrivateAPIEywaprimeHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CustomPrivateAPIEywaprimeClient) error {

	mux.Handle("POST", pattern_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CustomPrivateAPIEywaprime_DeactivateTenant_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomPrivateAPIEywaprime_DeactivateTenant_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomPrivateAPIEywaprime_DeactivateTenant_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4, 2, 5}, []string{"ves.io.schema", "introspect", "restricted", "write", "tenant", "deactivate"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_CustomPrivateAPIEywaprime_DeactivateTenant_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"ves.io.schema", "introspect", "write", "tenant", "deactivate"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CustomPrivateAPIEywaprime_RestrictedDeactivateTenant_0 = runtime.ForwardResponseMessage

	forward_CustomPrivateAPIEywaprime_DeactivateTenant_0 = runtime.ForwardResponseMessage
)
