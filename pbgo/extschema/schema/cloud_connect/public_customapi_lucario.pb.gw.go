// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ves.io/schema/cloud_connect/public_customapi_lucario.proto

/*
Package cloud_connect is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package cloud_connect

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

func request_CloudDataCustomAPI_DiscoverVPC_0(ctx context.Context, marshaler runtime.Marshaler, client CloudDataCustomAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DiscoverVPCRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.DiscoverVPC(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudDataCustomAPI_DiscoverVPC_0(ctx context.Context, marshaler runtime.Marshaler, server CloudDataCustomAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq DiscoverVPCRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.DiscoverVPC(ctx, &protoReq)
	return msg, metadata, err

}

func request_CloudDataCustomAPI_ReApplyVPCAttachment_0(ctx context.Context, marshaler runtime.Marshaler, client CloudDataCustomAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReApplyVPCAttachmentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.ReApplyVPCAttachment(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CloudDataCustomAPI_ReApplyVPCAttachment_0(ctx context.Context, marshaler runtime.Marshaler, server CloudDataCustomAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ReApplyVPCAttachmentRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.ReApplyVPCAttachment(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCloudDataCustomAPIHandlerServer registers the http handlers for service CloudDataCustomAPI to "mux".
// UnaryRPC     :call CloudDataCustomAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterCloudDataCustomAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CloudDataCustomAPIServer) error {

	mux.Handle("POST", pattern_CloudDataCustomAPI_DiscoverVPC_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudDataCustomAPI_DiscoverVPC_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudDataCustomAPI_DiscoverVPC_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudDataCustomAPI_ReApplyVPCAttachment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CloudDataCustomAPI_ReApplyVPCAttachment_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudDataCustomAPI_ReApplyVPCAttachment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCloudDataCustomAPIHandlerFromEndpoint is same as RegisterCloudDataCustomAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCloudDataCustomAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCloudDataCustomAPIHandler(ctx, mux, conn)
}

// RegisterCloudDataCustomAPIHandler registers the http handlers for service CloudDataCustomAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCloudDataCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCloudDataCustomAPIHandlerClient(ctx, mux, NewCloudDataCustomAPIClient(conn))
}

// RegisterCloudDataCustomAPIHandlerClient registers the http handlers for service CloudDataCustomAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CloudDataCustomAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CloudDataCustomAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CloudDataCustomAPIClient" to call the correct interceptors.
func RegisterCloudDataCustomAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CloudDataCustomAPIClient) error {

	mux.Handle("POST", pattern_CloudDataCustomAPI_DiscoverVPC_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudDataCustomAPI_DiscoverVPC_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudDataCustomAPI_DiscoverVPC_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CloudDataCustomAPI_ReApplyVPCAttachment_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CloudDataCustomAPI_ReApplyVPCAttachment_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CloudDataCustomAPI_ReApplyVPCAttachment_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CloudDataCustomAPI_DiscoverVPC_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"public", "namespaces", "system", "discover_vpc"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_CloudDataCustomAPI_ReApplyVPCAttachment_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"public", "namespaces", "system", "cloud_connect_reapply_vpc_attachment"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CloudDataCustomAPI_DiscoverVPC_0 = runtime.ForwardResponseMessage

	forward_CloudDataCustomAPI_ReApplyVPCAttachment_0 = runtime.ForwardResponseMessage
)