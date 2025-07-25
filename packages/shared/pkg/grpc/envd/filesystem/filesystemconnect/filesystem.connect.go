// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: filesystem/filesystem.proto

package filesystemconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	filesystem "github.com/e2b-dev/infra/packages/shared/pkg/grpc/envd/filesystem"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// FilesystemName is the fully-qualified name of the Filesystem service.
	FilesystemName = "filesystem.Filesystem"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// FilesystemStatProcedure is the fully-qualified name of the Filesystem's Stat RPC.
	FilesystemStatProcedure = "/filesystem.Filesystem/Stat"
	// FilesystemMakeDirProcedure is the fully-qualified name of the Filesystem's MakeDir RPC.
	FilesystemMakeDirProcedure = "/filesystem.Filesystem/MakeDir"
	// FilesystemMoveProcedure is the fully-qualified name of the Filesystem's Move RPC.
	FilesystemMoveProcedure = "/filesystem.Filesystem/Move"
	// FilesystemListDirProcedure is the fully-qualified name of the Filesystem's ListDir RPC.
	FilesystemListDirProcedure = "/filesystem.Filesystem/ListDir"
	// FilesystemRemoveProcedure is the fully-qualified name of the Filesystem's Remove RPC.
	FilesystemRemoveProcedure = "/filesystem.Filesystem/Remove"
	// FilesystemWatchDirProcedure is the fully-qualified name of the Filesystem's WatchDir RPC.
	FilesystemWatchDirProcedure = "/filesystem.Filesystem/WatchDir"
	// FilesystemCreateWatcherProcedure is the fully-qualified name of the Filesystem's CreateWatcher
	// RPC.
	FilesystemCreateWatcherProcedure = "/filesystem.Filesystem/CreateWatcher"
	// FilesystemGetWatcherEventsProcedure is the fully-qualified name of the Filesystem's
	// GetWatcherEvents RPC.
	FilesystemGetWatcherEventsProcedure = "/filesystem.Filesystem/GetWatcherEvents"
	// FilesystemRemoveWatcherProcedure is the fully-qualified name of the Filesystem's RemoveWatcher
	// RPC.
	FilesystemRemoveWatcherProcedure = "/filesystem.Filesystem/RemoveWatcher"
)

// FilesystemClient is a client for the filesystem.Filesystem service.
type FilesystemClient interface {
	Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error)
	MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error)
	Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error)
	ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error)
	Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error)
	WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest]) (*connect.ServerStreamForClient[filesystem.WatchDirResponse], error)
	// Non-streaming versions of WatchDir
	CreateWatcher(context.Context, *connect.Request[filesystem.CreateWatcherRequest]) (*connect.Response[filesystem.CreateWatcherResponse], error)
	GetWatcherEvents(context.Context, *connect.Request[filesystem.GetWatcherEventsRequest]) (*connect.Response[filesystem.GetWatcherEventsResponse], error)
	RemoveWatcher(context.Context, *connect.Request[filesystem.RemoveWatcherRequest]) (*connect.Response[filesystem.RemoveWatcherResponse], error)
}

// NewFilesystemClient constructs a client for the filesystem.Filesystem service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewFilesystemClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) FilesystemClient {
	baseURL = strings.TrimRight(baseURL, "/")
	filesystemMethods := filesystem.File_filesystem_filesystem_proto.Services().ByName("Filesystem").Methods()
	return &filesystemClient{
		stat: connect.NewClient[filesystem.StatRequest, filesystem.StatResponse](
			httpClient,
			baseURL+FilesystemStatProcedure,
			connect.WithSchema(filesystemMethods.ByName("Stat")),
			connect.WithClientOptions(opts...),
		),
		makeDir: connect.NewClient[filesystem.MakeDirRequest, filesystem.MakeDirResponse](
			httpClient,
			baseURL+FilesystemMakeDirProcedure,
			connect.WithSchema(filesystemMethods.ByName("MakeDir")),
			connect.WithClientOptions(opts...),
		),
		move: connect.NewClient[filesystem.MoveRequest, filesystem.MoveResponse](
			httpClient,
			baseURL+FilesystemMoveProcedure,
			connect.WithSchema(filesystemMethods.ByName("Move")),
			connect.WithClientOptions(opts...),
		),
		listDir: connect.NewClient[filesystem.ListDirRequest, filesystem.ListDirResponse](
			httpClient,
			baseURL+FilesystemListDirProcedure,
			connect.WithSchema(filesystemMethods.ByName("ListDir")),
			connect.WithClientOptions(opts...),
		),
		remove: connect.NewClient[filesystem.RemoveRequest, filesystem.RemoveResponse](
			httpClient,
			baseURL+FilesystemRemoveProcedure,
			connect.WithSchema(filesystemMethods.ByName("Remove")),
			connect.WithClientOptions(opts...),
		),
		watchDir: connect.NewClient[filesystem.WatchDirRequest, filesystem.WatchDirResponse](
			httpClient,
			baseURL+FilesystemWatchDirProcedure,
			connect.WithSchema(filesystemMethods.ByName("WatchDir")),
			connect.WithClientOptions(opts...),
		),
		createWatcher: connect.NewClient[filesystem.CreateWatcherRequest, filesystem.CreateWatcherResponse](
			httpClient,
			baseURL+FilesystemCreateWatcherProcedure,
			connect.WithSchema(filesystemMethods.ByName("CreateWatcher")),
			connect.WithClientOptions(opts...),
		),
		getWatcherEvents: connect.NewClient[filesystem.GetWatcherEventsRequest, filesystem.GetWatcherEventsResponse](
			httpClient,
			baseURL+FilesystemGetWatcherEventsProcedure,
			connect.WithSchema(filesystemMethods.ByName("GetWatcherEvents")),
			connect.WithClientOptions(opts...),
		),
		removeWatcher: connect.NewClient[filesystem.RemoveWatcherRequest, filesystem.RemoveWatcherResponse](
			httpClient,
			baseURL+FilesystemRemoveWatcherProcedure,
			connect.WithSchema(filesystemMethods.ByName("RemoveWatcher")),
			connect.WithClientOptions(opts...),
		),
	}
}

// filesystemClient implements FilesystemClient.
type filesystemClient struct {
	stat             *connect.Client[filesystem.StatRequest, filesystem.StatResponse]
	makeDir          *connect.Client[filesystem.MakeDirRequest, filesystem.MakeDirResponse]
	move             *connect.Client[filesystem.MoveRequest, filesystem.MoveResponse]
	listDir          *connect.Client[filesystem.ListDirRequest, filesystem.ListDirResponse]
	remove           *connect.Client[filesystem.RemoveRequest, filesystem.RemoveResponse]
	watchDir         *connect.Client[filesystem.WatchDirRequest, filesystem.WatchDirResponse]
	createWatcher    *connect.Client[filesystem.CreateWatcherRequest, filesystem.CreateWatcherResponse]
	getWatcherEvents *connect.Client[filesystem.GetWatcherEventsRequest, filesystem.GetWatcherEventsResponse]
	removeWatcher    *connect.Client[filesystem.RemoveWatcherRequest, filesystem.RemoveWatcherResponse]
}

// Stat calls filesystem.Filesystem.Stat.
func (c *filesystemClient) Stat(ctx context.Context, req *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error) {
	return c.stat.CallUnary(ctx, req)
}

// MakeDir calls filesystem.Filesystem.MakeDir.
func (c *filesystemClient) MakeDir(ctx context.Context, req *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error) {
	return c.makeDir.CallUnary(ctx, req)
}

// Move calls filesystem.Filesystem.Move.
func (c *filesystemClient) Move(ctx context.Context, req *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error) {
	return c.move.CallUnary(ctx, req)
}

// ListDir calls filesystem.Filesystem.ListDir.
func (c *filesystemClient) ListDir(ctx context.Context, req *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error) {
	return c.listDir.CallUnary(ctx, req)
}

// Remove calls filesystem.Filesystem.Remove.
func (c *filesystemClient) Remove(ctx context.Context, req *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error) {
	return c.remove.CallUnary(ctx, req)
}

// WatchDir calls filesystem.Filesystem.WatchDir.
func (c *filesystemClient) WatchDir(ctx context.Context, req *connect.Request[filesystem.WatchDirRequest]) (*connect.ServerStreamForClient[filesystem.WatchDirResponse], error) {
	return c.watchDir.CallServerStream(ctx, req)
}

// CreateWatcher calls filesystem.Filesystem.CreateWatcher.
func (c *filesystemClient) CreateWatcher(ctx context.Context, req *connect.Request[filesystem.CreateWatcherRequest]) (*connect.Response[filesystem.CreateWatcherResponse], error) {
	return c.createWatcher.CallUnary(ctx, req)
}

// GetWatcherEvents calls filesystem.Filesystem.GetWatcherEvents.
func (c *filesystemClient) GetWatcherEvents(ctx context.Context, req *connect.Request[filesystem.GetWatcherEventsRequest]) (*connect.Response[filesystem.GetWatcherEventsResponse], error) {
	return c.getWatcherEvents.CallUnary(ctx, req)
}

// RemoveWatcher calls filesystem.Filesystem.RemoveWatcher.
func (c *filesystemClient) RemoveWatcher(ctx context.Context, req *connect.Request[filesystem.RemoveWatcherRequest]) (*connect.Response[filesystem.RemoveWatcherResponse], error) {
	return c.removeWatcher.CallUnary(ctx, req)
}

// FilesystemHandler is an implementation of the filesystem.Filesystem service.
type FilesystemHandler interface {
	Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error)
	MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error)
	Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error)
	ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error)
	Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error)
	WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest], *connect.ServerStream[filesystem.WatchDirResponse]) error
	// Non-streaming versions of WatchDir
	CreateWatcher(context.Context, *connect.Request[filesystem.CreateWatcherRequest]) (*connect.Response[filesystem.CreateWatcherResponse], error)
	GetWatcherEvents(context.Context, *connect.Request[filesystem.GetWatcherEventsRequest]) (*connect.Response[filesystem.GetWatcherEventsResponse], error)
	RemoveWatcher(context.Context, *connect.Request[filesystem.RemoveWatcherRequest]) (*connect.Response[filesystem.RemoveWatcherResponse], error)
}

// NewFilesystemHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewFilesystemHandler(svc FilesystemHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	filesystemMethods := filesystem.File_filesystem_filesystem_proto.Services().ByName("Filesystem").Methods()
	filesystemStatHandler := connect.NewUnaryHandler(
		FilesystemStatProcedure,
		svc.Stat,
		connect.WithSchema(filesystemMethods.ByName("Stat")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemMakeDirHandler := connect.NewUnaryHandler(
		FilesystemMakeDirProcedure,
		svc.MakeDir,
		connect.WithSchema(filesystemMethods.ByName("MakeDir")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemMoveHandler := connect.NewUnaryHandler(
		FilesystemMoveProcedure,
		svc.Move,
		connect.WithSchema(filesystemMethods.ByName("Move")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemListDirHandler := connect.NewUnaryHandler(
		FilesystemListDirProcedure,
		svc.ListDir,
		connect.WithSchema(filesystemMethods.ByName("ListDir")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemRemoveHandler := connect.NewUnaryHandler(
		FilesystemRemoveProcedure,
		svc.Remove,
		connect.WithSchema(filesystemMethods.ByName("Remove")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemWatchDirHandler := connect.NewServerStreamHandler(
		FilesystemWatchDirProcedure,
		svc.WatchDir,
		connect.WithSchema(filesystemMethods.ByName("WatchDir")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemCreateWatcherHandler := connect.NewUnaryHandler(
		FilesystemCreateWatcherProcedure,
		svc.CreateWatcher,
		connect.WithSchema(filesystemMethods.ByName("CreateWatcher")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemGetWatcherEventsHandler := connect.NewUnaryHandler(
		FilesystemGetWatcherEventsProcedure,
		svc.GetWatcherEvents,
		connect.WithSchema(filesystemMethods.ByName("GetWatcherEvents")),
		connect.WithHandlerOptions(opts...),
	)
	filesystemRemoveWatcherHandler := connect.NewUnaryHandler(
		FilesystemRemoveWatcherProcedure,
		svc.RemoveWatcher,
		connect.WithSchema(filesystemMethods.ByName("RemoveWatcher")),
		connect.WithHandlerOptions(opts...),
	)
	return "/filesystem.Filesystem/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case FilesystemStatProcedure:
			filesystemStatHandler.ServeHTTP(w, r)
		case FilesystemMakeDirProcedure:
			filesystemMakeDirHandler.ServeHTTP(w, r)
		case FilesystemMoveProcedure:
			filesystemMoveHandler.ServeHTTP(w, r)
		case FilesystemListDirProcedure:
			filesystemListDirHandler.ServeHTTP(w, r)
		case FilesystemRemoveProcedure:
			filesystemRemoveHandler.ServeHTTP(w, r)
		case FilesystemWatchDirProcedure:
			filesystemWatchDirHandler.ServeHTTP(w, r)
		case FilesystemCreateWatcherProcedure:
			filesystemCreateWatcherHandler.ServeHTTP(w, r)
		case FilesystemGetWatcherEventsProcedure:
			filesystemGetWatcherEventsHandler.ServeHTTP(w, r)
		case FilesystemRemoveWatcherProcedure:
			filesystemRemoveWatcherHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedFilesystemHandler returns CodeUnimplemented from all methods.
type UnimplementedFilesystemHandler struct{}

func (UnimplementedFilesystemHandler) Stat(context.Context, *connect.Request[filesystem.StatRequest]) (*connect.Response[filesystem.StatResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Stat is not implemented"))
}

func (UnimplementedFilesystemHandler) MakeDir(context.Context, *connect.Request[filesystem.MakeDirRequest]) (*connect.Response[filesystem.MakeDirResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.MakeDir is not implemented"))
}

func (UnimplementedFilesystemHandler) Move(context.Context, *connect.Request[filesystem.MoveRequest]) (*connect.Response[filesystem.MoveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Move is not implemented"))
}

func (UnimplementedFilesystemHandler) ListDir(context.Context, *connect.Request[filesystem.ListDirRequest]) (*connect.Response[filesystem.ListDirResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.ListDir is not implemented"))
}

func (UnimplementedFilesystemHandler) Remove(context.Context, *connect.Request[filesystem.RemoveRequest]) (*connect.Response[filesystem.RemoveResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.Remove is not implemented"))
}

func (UnimplementedFilesystemHandler) WatchDir(context.Context, *connect.Request[filesystem.WatchDirRequest], *connect.ServerStream[filesystem.WatchDirResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.WatchDir is not implemented"))
}

func (UnimplementedFilesystemHandler) CreateWatcher(context.Context, *connect.Request[filesystem.CreateWatcherRequest]) (*connect.Response[filesystem.CreateWatcherResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.CreateWatcher is not implemented"))
}

func (UnimplementedFilesystemHandler) GetWatcherEvents(context.Context, *connect.Request[filesystem.GetWatcherEventsRequest]) (*connect.Response[filesystem.GetWatcherEventsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.GetWatcherEvents is not implemented"))
}

func (UnimplementedFilesystemHandler) RemoveWatcher(context.Context, *connect.Request[filesystem.RemoveWatcherRequest]) (*connect.Response[filesystem.RemoveWatcherResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("filesystem.Filesystem.RemoveWatcher is not implemented"))
}
