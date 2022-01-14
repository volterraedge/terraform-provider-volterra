// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: ves.io/schema/views/terraform_parameters/public_custom_actionapi.proto

/*
Package terraform_parameters is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package terraform_parameters

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

func request_CustomActionAPI_Run_0(ctx context.Context, marshaler runtime.Marshaler, client CustomActionAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RunRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["view_kind"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_kind")
	}

	protoReq.ViewKind, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_kind", err)
	}

	val, ok = pathParams["view_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_name")
	}

	protoReq.ViewName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_name", err)
	}

	msg, err := client.Run(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomActionAPI_Run_0(ctx context.Context, marshaler runtime.Marshaler, server CustomActionAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq RunRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["view_kind"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_kind")
	}

	protoReq.ViewKind, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_kind", err)
	}

	val, ok = pathParams["view_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_name")
	}

	protoReq.ViewName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_name", err)
	}

	msg, err := server.Run(ctx, &protoReq)
	return msg, metadata, err

}

func request_CustomActionAPI_ForceDelete_0(ctx context.Context, marshaler runtime.Marshaler, client CustomActionAPIClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ForceDeleteRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["view_kind"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_kind")
	}

	protoReq.ViewKind, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_kind", err)
	}

	val, ok = pathParams["view_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_name")
	}

	protoReq.ViewName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_name", err)
	}

	msg, err := client.ForceDelete(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CustomActionAPI_ForceDelete_0(ctx context.Context, marshaler runtime.Marshaler, server CustomActionAPIServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ForceDeleteRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["namespace"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "namespace")
	}

	protoReq.Namespace, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "namespace", err)
	}

	val, ok = pathParams["view_kind"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_kind")
	}

	protoReq.ViewKind, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_kind", err)
	}

	val, ok = pathParams["view_name"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "view_name")
	}

	protoReq.ViewName, err = runtime.String(val)

	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "view_name", err)
	}

	msg, err := server.ForceDelete(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCustomActionAPIHandlerServer registers the http handlers for service CustomActionAPI to "mux".
// UnaryRPC     :call CustomActionAPIServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
func RegisterCustomActionAPIHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CustomActionAPIServer) error {

	mux.Handle("POST", pattern_CustomActionAPI_Run_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomActionAPI_Run_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomActionAPI_Run_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CustomActionAPI_ForceDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CustomActionAPI_ForceDelete_0(rctx, inboundMarshaler, server, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomActionAPI_ForceDelete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCustomActionAPIHandlerFromEndpoint is same as RegisterCustomActionAPIHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCustomActionAPIHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterCustomActionAPIHandler(ctx, mux, conn)
}

// RegisterCustomActionAPIHandler registers the http handlers for service CustomActionAPI to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCustomActionAPIHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCustomActionAPIHandlerClient(ctx, mux, NewCustomActionAPIClient(conn))
}

// RegisterCustomActionAPIHandlerClient registers the http handlers for service CustomActionAPI
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CustomActionAPIClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CustomActionAPIClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CustomActionAPIClient" to call the correct interceptors.
func RegisterCustomActionAPIHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CustomActionAPIClient) error {

	mux.Handle("POST", pattern_CustomActionAPI_Run_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomActionAPI_Run_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomActionAPI_Run_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_CustomActionAPI_ForceDelete_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CustomActionAPI_ForceDelete_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CustomActionAPI_ForceDelete_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CustomActionAPI_Run_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 2, 6}, []string{"public", "namespaces", "namespace", "terraform", "view_kind", "view_name", "run"}, "", runtime.AssumeColonVerbOpt(false)))

	pattern_CustomActionAPI_ForceDelete_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 1, 0, 4, 1, 5, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5, 2, 6}, []string{"public", "namespaces", "namespace", "terraform", "view_kind", "view_name", "force-delete"}, "", runtime.AssumeColonVerbOpt(false)))
)

var (
	forward_CustomActionAPI_Run_0 = runtime.ForwardResponseMessage

	forward_CustomActionAPI_ForceDelete_0 = runtime.ForwardResponseMessage
)