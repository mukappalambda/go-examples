// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: note/note.proto

package note_server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NoteService_CreateNote_FullMethodName = "/note.NoteService/CreateNote"
	NoteService_GetNote_FullMethodName    = "/note.NoteService/GetNote"
	NoteService_ListNotes_FullMethodName  = "/note.NoteService/ListNotes"
)

// NoteServiceClient is the client API for NoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NoteServiceClient interface {
	CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error)
	GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error)
	ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Note], error)
}

type noteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNoteServiceClient(cc grpc.ClientConnInterface) NoteServiceClient {
	return &noteServiceClient{cc}
}

func (c *noteServiceClient) CreateNote(ctx context.Context, in *CreateNoteRequest, opts ...grpc.CallOption) (*CreateNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_CreateNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) GetNote(ctx context.Context, in *GetNoteRequest, opts ...grpc.CallOption) (*GetNoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNoteResponse)
	err := c.cc.Invoke(ctx, NoteService_GetNote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *noteServiceClient) ListNotes(ctx context.Context, in *ListNotesRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Note], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &NoteService_ServiceDesc.Streams[0], NoteService_ListNotes_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListNotesRequest, Note]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NoteService_ListNotesClient = grpc.ServerStreamingClient[Note]

// NoteServiceServer is the server API for NoteService service.
// All implementations must embed UnimplementedNoteServiceServer
// for forward compatibility.
type NoteServiceServer interface {
	CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error)
	GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error)
	ListNotes(*ListNotesRequest, grpc.ServerStreamingServer[Note]) error
	mustEmbedUnimplementedNoteServiceServer()
}

// UnimplementedNoteServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNoteServiceServer struct{}

func (UnimplementedNoteServiceServer) CreateNote(context.Context, *CreateNoteRequest) (*CreateNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNote not implemented")
}
func (UnimplementedNoteServiceServer) GetNote(context.Context, *GetNoteRequest) (*GetNoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNote not implemented")
}
func (UnimplementedNoteServiceServer) ListNotes(*ListNotesRequest, grpc.ServerStreamingServer[Note]) error {
	return status.Errorf(codes.Unimplemented, "method ListNotes not implemented")
}
func (UnimplementedNoteServiceServer) mustEmbedUnimplementedNoteServiceServer() {}
func (UnimplementedNoteServiceServer) testEmbeddedByValue()                     {}

// UnsafeNoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NoteServiceServer will
// result in compilation errors.
type UnsafeNoteServiceServer interface {
	mustEmbedUnimplementedNoteServiceServer()
}

func RegisterNoteServiceServer(s grpc.ServiceRegistrar, srv NoteServiceServer) {
	// If the following call pancis, it indicates UnimplementedNoteServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NoteService_ServiceDesc, srv)
}

func _NoteService_CreateNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).CreateNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_CreateNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).CreateNote(ctx, req.(*CreateNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_GetNote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServiceServer).GetNote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NoteService_GetNote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServiceServer).GetNote(ctx, req.(*GetNoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NoteService_ListNotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListNotesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NoteServiceServer).ListNotes(m, &grpc.GenericServerStream[ListNotesRequest, Note]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type NoteService_ListNotesServer = grpc.ServerStreamingServer[Note]

// NoteService_ServiceDesc is the grpc.ServiceDesc for NoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "note.NoteService",
	HandlerType: (*NoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNote",
			Handler:    _NoteService_CreateNote_Handler,
		},
		{
			MethodName: "GetNote",
			Handler:    _NoteService_GetNote_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListNotes",
			Handler:       _NoteService_ListNotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "note/note.proto",
}
