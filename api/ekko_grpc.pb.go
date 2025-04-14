// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: ekko/api/ekko.proto

package ekko

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Ekko_CreateField_FullMethodName       = "/ekko.Ekko/CreateField"
	Ekko_UpdateField_FullMethodName       = "/ekko.Ekko/UpdateField"
	Ekko_DeleteField_FullMethodName       = "/ekko.Ekko/DeleteField"
	Ekko_CreateScenario_FullMethodName    = "/ekko.Ekko/CreateScenario"
	Ekko_UpdateScenario_FullMethodName    = "/ekko.Ekko/UpdateScenario"
	Ekko_DeleteScenario_FullMethodName    = "/ekko.Ekko/DeleteScenario"
	Ekko_ListScenario_FullMethodName      = "/ekko.Ekko/ListScenario"
	Ekko_FavoriteScenario_FullMethodName  = "/ekko.Ekko/FavoriteScenario"
	Ekko_RatingScenario_FullMethodName    = "/ekko.Ekko/RatingScenario"
	Ekko_ListAttempt_FullMethodName       = "/ekko.Ekko/ListAttempt"
	Ekko_GetAttempt_FullMethodName        = "/ekko.Ekko/GetAttempt"
	Ekko_SubmitAnswer_FullMethodName      = "/ekko.Ekko/SubmitAnswer"
	Ekko_ListAllSubmission_FullMethodName = "/ekko.Ekko/ListAllSubmission"
)

// EkkoClient is the client API for Ekko service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EkkoClient interface {
	CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error)
	UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteField(ctx context.Context, in *DeleteFieldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Scenario api
	CreateScenario(ctx context.Context, in *CreateScenarioRequest, opts ...grpc.CallOption) (*CreateScenarioResponse, error)
	UpdateScenario(ctx context.Context, in *UpdateScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteScenario(ctx context.Context, in *DeleteScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListScenario(ctx context.Context, in *ListScenarioRequest, opts ...grpc.CallOption) (*ListScenarioResponse, error)
	FavoriteScenario(ctx context.Context, in *FavoriteScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RatingScenario(ctx context.Context, in *RatingScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Submission for candidate
	ListAttempt(ctx context.Context, in *ListAttemptRequest, opts ...grpc.CallOption) (*ListAttemptResponse, error)
	GetAttempt(ctx context.Context, in *GetAttemptRequest, opts ...grpc.CallOption) (*GetAttemptResponse, error)
	SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error)
	// Submission for bm
	ListAllSubmission(ctx context.Context, in *ListAllSubmissionRequest, opts ...grpc.CallOption) (*ListAllSubmissionResponse, error)
}

type ekkoClient struct {
	cc grpc.ClientConnInterface
}

func NewEkkoClient(cc grpc.ClientConnInterface) EkkoClient {
	return &ekkoClient{cc}
}

func (c *ekkoClient) CreateField(ctx context.Context, in *CreateFieldRequest, opts ...grpc.CallOption) (*CreateFieldResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateFieldResponse)
	err := c.cc.Invoke(ctx, Ekko_CreateField_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_UpdateField_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) DeleteField(ctx context.Context, in *DeleteFieldRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_DeleteField_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) CreateScenario(ctx context.Context, in *CreateScenarioRequest, opts ...grpc.CallOption) (*CreateScenarioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateScenarioResponse)
	err := c.cc.Invoke(ctx, Ekko_CreateScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) UpdateScenario(ctx context.Context, in *UpdateScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_UpdateScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) DeleteScenario(ctx context.Context, in *DeleteScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_DeleteScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) ListScenario(ctx context.Context, in *ListScenarioRequest, opts ...grpc.CallOption) (*ListScenarioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScenarioResponse)
	err := c.cc.Invoke(ctx, Ekko_ListScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) FavoriteScenario(ctx context.Context, in *FavoriteScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_FavoriteScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) RatingScenario(ctx context.Context, in *RatingScenarioRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Ekko_RatingScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) ListAttempt(ctx context.Context, in *ListAttemptRequest, opts ...grpc.CallOption) (*ListAttemptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAttemptResponse)
	err := c.cc.Invoke(ctx, Ekko_ListAttempt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) GetAttempt(ctx context.Context, in *GetAttemptRequest, opts ...grpc.CallOption) (*GetAttemptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAttemptResponse)
	err := c.cc.Invoke(ctx, Ekko_GetAttempt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) SubmitAnswer(ctx context.Context, in *SubmitAnswerRequest, opts ...grpc.CallOption) (*SubmitAnswerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubmitAnswerResponse)
	err := c.cc.Invoke(ctx, Ekko_SubmitAnswer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ekkoClient) ListAllSubmission(ctx context.Context, in *ListAllSubmissionRequest, opts ...grpc.CallOption) (*ListAllSubmissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllSubmissionResponse)
	err := c.cc.Invoke(ctx, Ekko_ListAllSubmission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EkkoServer is the server API for Ekko service.
// All implementations must embed UnimplementedEkkoServer
// for forward compatibility.
type EkkoServer interface {
	CreateField(context.Context, *CreateFieldRequest) (*CreateFieldResponse, error)
	UpdateField(context.Context, *UpdateFieldRequest) (*emptypb.Empty, error)
	DeleteField(context.Context, *DeleteFieldRequest) (*emptypb.Empty, error)
	// Scenario api
	CreateScenario(context.Context, *CreateScenarioRequest) (*CreateScenarioResponse, error)
	UpdateScenario(context.Context, *UpdateScenarioRequest) (*emptypb.Empty, error)
	DeleteScenario(context.Context, *DeleteScenarioRequest) (*emptypb.Empty, error)
	ListScenario(context.Context, *ListScenarioRequest) (*ListScenarioResponse, error)
	FavoriteScenario(context.Context, *FavoriteScenarioRequest) (*emptypb.Empty, error)
	RatingScenario(context.Context, *RatingScenarioRequest) (*emptypb.Empty, error)
	// Submission for candidate
	ListAttempt(context.Context, *ListAttemptRequest) (*ListAttemptResponse, error)
	GetAttempt(context.Context, *GetAttemptRequest) (*GetAttemptResponse, error)
	SubmitAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error)
	// Submission for bm
	ListAllSubmission(context.Context, *ListAllSubmissionRequest) (*ListAllSubmissionResponse, error)
	mustEmbedUnimplementedEkkoServer()
}

// UnimplementedEkkoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEkkoServer struct{}

func (UnimplementedEkkoServer) CreateField(context.Context, *CreateFieldRequest) (*CreateFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateField not implemented")
}
func (UnimplementedEkkoServer) UpdateField(context.Context, *UpdateFieldRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateField not implemented")
}
func (UnimplementedEkkoServer) DeleteField(context.Context, *DeleteFieldRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteField not implemented")
}
func (UnimplementedEkkoServer) CreateScenario(context.Context, *CreateScenarioRequest) (*CreateScenarioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateScenario not implemented")
}
func (UnimplementedEkkoServer) UpdateScenario(context.Context, *UpdateScenarioRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScenario not implemented")
}
func (UnimplementedEkkoServer) DeleteScenario(context.Context, *DeleteScenarioRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScenario not implemented")
}
func (UnimplementedEkkoServer) ListScenario(context.Context, *ListScenarioRequest) (*ListScenarioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScenario not implemented")
}
func (UnimplementedEkkoServer) FavoriteScenario(context.Context, *FavoriteScenarioRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FavoriteScenario not implemented")
}
func (UnimplementedEkkoServer) RatingScenario(context.Context, *RatingScenarioRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RatingScenario not implemented")
}
func (UnimplementedEkkoServer) ListAttempt(context.Context, *ListAttemptRequest) (*ListAttemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAttempt not implemented")
}
func (UnimplementedEkkoServer) GetAttempt(context.Context, *GetAttemptRequest) (*GetAttemptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAttempt not implemented")
}
func (UnimplementedEkkoServer) SubmitAnswer(context.Context, *SubmitAnswerRequest) (*SubmitAnswerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitAnswer not implemented")
}
func (UnimplementedEkkoServer) ListAllSubmission(context.Context, *ListAllSubmissionRequest) (*ListAllSubmissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllSubmission not implemented")
}
func (UnimplementedEkkoServer) mustEmbedUnimplementedEkkoServer() {}
func (UnimplementedEkkoServer) testEmbeddedByValue()              {}

// UnsafeEkkoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EkkoServer will
// result in compilation errors.
type UnsafeEkkoServer interface {
	mustEmbedUnimplementedEkkoServer()
}

func RegisterEkkoServer(s grpc.ServiceRegistrar, srv EkkoServer) {
	// If the following call pancis, it indicates UnimplementedEkkoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Ekko_ServiceDesc, srv)
}

func _Ekko_CreateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).CreateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_CreateField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).CreateField(ctx, req.(*CreateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_UpdateField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).UpdateField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_UpdateField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).UpdateField(ctx, req.(*UpdateFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_DeleteField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).DeleteField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_DeleteField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).DeleteField(ctx, req.(*DeleteFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_CreateScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).CreateScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_CreateScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).CreateScenario(ctx, req.(*CreateScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_UpdateScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).UpdateScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_UpdateScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).UpdateScenario(ctx, req.(*UpdateScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_DeleteScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).DeleteScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_DeleteScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).DeleteScenario(ctx, req.(*DeleteScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_ListScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).ListScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_ListScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).ListScenario(ctx, req.(*ListScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_FavoriteScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FavoriteScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).FavoriteScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_FavoriteScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).FavoriteScenario(ctx, req.(*FavoriteScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_RatingScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RatingScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).RatingScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_RatingScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).RatingScenario(ctx, req.(*RatingScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_ListAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAttemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).ListAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_ListAttempt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).ListAttempt(ctx, req.(*ListAttemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_GetAttempt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAttemptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).GetAttempt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_GetAttempt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).GetAttempt(ctx, req.(*GetAttemptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_SubmitAnswer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitAnswerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).SubmitAnswer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_SubmitAnswer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).SubmitAnswer(ctx, req.(*SubmitAnswerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ekko_ListAllSubmission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAllSubmissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EkkoServer).ListAllSubmission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Ekko_ListAllSubmission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EkkoServer).ListAllSubmission(ctx, req.(*ListAllSubmissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Ekko_ServiceDesc is the grpc.ServiceDesc for Ekko service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ekko_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ekko.Ekko",
	HandlerType: (*EkkoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateField",
			Handler:    _Ekko_CreateField_Handler,
		},
		{
			MethodName: "UpdateField",
			Handler:    _Ekko_UpdateField_Handler,
		},
		{
			MethodName: "DeleteField",
			Handler:    _Ekko_DeleteField_Handler,
		},
		{
			MethodName: "CreateScenario",
			Handler:    _Ekko_CreateScenario_Handler,
		},
		{
			MethodName: "UpdateScenario",
			Handler:    _Ekko_UpdateScenario_Handler,
		},
		{
			MethodName: "DeleteScenario",
			Handler:    _Ekko_DeleteScenario_Handler,
		},
		{
			MethodName: "ListScenario",
			Handler:    _Ekko_ListScenario_Handler,
		},
		{
			MethodName: "FavoriteScenario",
			Handler:    _Ekko_FavoriteScenario_Handler,
		},
		{
			MethodName: "RatingScenario",
			Handler:    _Ekko_RatingScenario_Handler,
		},
		{
			MethodName: "ListAttempt",
			Handler:    _Ekko_ListAttempt_Handler,
		},
		{
			MethodName: "GetAttempt",
			Handler:    _Ekko_GetAttempt_Handler,
		},
		{
			MethodName: "SubmitAnswer",
			Handler:    _Ekko_SubmitAnswer_Handler,
		},
		{
			MethodName: "ListAllSubmission",
			Handler:    _Ekko_ListAllSubmission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ekko/api/ekko.proto",
}

const (
	Chronobreak_ListField_FullMethodName    = "/ekko.Chronobreak/ListField"
	Chronobreak_ListScenario_FullMethodName = "/ekko.Chronobreak/ListScenario"
	Chronobreak_GetScenario_FullMethodName  = "/ekko.Chronobreak/GetScenario"
)

// ChronobreakClient is the client API for Chronobreak service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChronobreakClient interface {
	ListField(ctx context.Context, in *ListFieldRequest, opts ...grpc.CallOption) (*ListFieldResponse, error)
	ListScenario(ctx context.Context, in *ListScenarioRequest, opts ...grpc.CallOption) (*ListScenarioResponse, error)
	GetScenario(ctx context.Context, in *GetScenarioRequest, opts ...grpc.CallOption) (*GetScenarioResponse, error)
}

type chronobreakClient struct {
	cc grpc.ClientConnInterface
}

func NewChronobreakClient(cc grpc.ClientConnInterface) ChronobreakClient {
	return &chronobreakClient{cc}
}

func (c *chronobreakClient) ListField(ctx context.Context, in *ListFieldRequest, opts ...grpc.CallOption) (*ListFieldResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFieldResponse)
	err := c.cc.Invoke(ctx, Chronobreak_ListField_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chronobreakClient) ListScenario(ctx context.Context, in *ListScenarioRequest, opts ...grpc.CallOption) (*ListScenarioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListScenarioResponse)
	err := c.cc.Invoke(ctx, Chronobreak_ListScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chronobreakClient) GetScenario(ctx context.Context, in *GetScenarioRequest, opts ...grpc.CallOption) (*GetScenarioResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetScenarioResponse)
	err := c.cc.Invoke(ctx, Chronobreak_GetScenario_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChronobreakServer is the server API for Chronobreak service.
// All implementations must embed UnimplementedChronobreakServer
// for forward compatibility.
type ChronobreakServer interface {
	ListField(context.Context, *ListFieldRequest) (*ListFieldResponse, error)
	ListScenario(context.Context, *ListScenarioRequest) (*ListScenarioResponse, error)
	GetScenario(context.Context, *GetScenarioRequest) (*GetScenarioResponse, error)
	mustEmbedUnimplementedChronobreakServer()
}

// UnimplementedChronobreakServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChronobreakServer struct{}

func (UnimplementedChronobreakServer) ListField(context.Context, *ListFieldRequest) (*ListFieldResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListField not implemented")
}
func (UnimplementedChronobreakServer) ListScenario(context.Context, *ListScenarioRequest) (*ListScenarioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScenario not implemented")
}
func (UnimplementedChronobreakServer) GetScenario(context.Context, *GetScenarioRequest) (*GetScenarioResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScenario not implemented")
}
func (UnimplementedChronobreakServer) mustEmbedUnimplementedChronobreakServer() {}
func (UnimplementedChronobreakServer) testEmbeddedByValue()                     {}

// UnsafeChronobreakServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChronobreakServer will
// result in compilation errors.
type UnsafeChronobreakServer interface {
	mustEmbedUnimplementedChronobreakServer()
}

func RegisterChronobreakServer(s grpc.ServiceRegistrar, srv ChronobreakServer) {
	// If the following call pancis, it indicates UnimplementedChronobreakServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Chronobreak_ServiceDesc, srv)
}

func _Chronobreak_ListField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFieldRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChronobreakServer).ListField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chronobreak_ListField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChronobreakServer).ListField(ctx, req.(*ListFieldRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chronobreak_ListScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChronobreakServer).ListScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chronobreak_ListScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChronobreakServer).ListScenario(ctx, req.(*ListScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chronobreak_GetScenario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScenarioRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChronobreakServer).GetScenario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Chronobreak_GetScenario_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChronobreakServer).GetScenario(ctx, req.(*GetScenarioRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chronobreak_ServiceDesc is the grpc.ServiceDesc for Chronobreak service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chronobreak_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ekko.Chronobreak",
	HandlerType: (*ChronobreakServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListField",
			Handler:    _Chronobreak_ListField_Handler,
		},
		{
			MethodName: "ListScenario",
			Handler:    _Chronobreak_ListScenario_Handler,
		},
		{
			MethodName: "GetScenario",
			Handler:    _Chronobreak_GetScenario_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ekko/api/ekko.proto",
}
