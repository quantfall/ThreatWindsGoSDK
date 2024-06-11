// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: plugins.proto

package plugins

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Plugins_Input_FullMethodName        = "/plugins.Plugins/Input"
	Plugins_Parsing_FullMethodName      = "/plugins.Plugins/Parsing"
	Plugins_Analysis_FullMethodName     = "/plugins.Plugins/Analysis"
	Plugins_Notification_FullMethodName = "/plugins.Plugins/Notification"
	Plugins_Anomalies_FullMethodName    = "/plugins.Plugins/Anomalies"
	Plugins_Correlate_FullMethodName    = "/plugins.Plugins/Correlate"
)

// PluginsClient is the client API for Plugins service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginsClient interface {
	Input(ctx context.Context, opts ...grpc.CallOption) (Plugins_InputClient, error)
	Parsing(ctx context.Context, opts ...grpc.CallOption) (Plugins_ParsingClient, error)
	Analysis(ctx context.Context, opts ...grpc.CallOption) (Plugins_AnalysisClient, error)
	Notification(ctx context.Context, opts ...grpc.CallOption) (Plugins_NotificationClient, error)
	Anomalies(ctx context.Context, opts ...grpc.CallOption) (Plugins_AnomaliesClient, error)
	Correlate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Plugins_CorrelateClient, error)
}

type pluginsClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginsClient(cc grpc.ClientConnInterface) PluginsClient {
	return &pluginsClient{cc}
}

func (c *pluginsClient) Input(ctx context.Context, opts ...grpc.CallOption) (Plugins_InputClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[0], Plugins_Input_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsInputClient{stream}
	return x, nil
}

type Plugins_InputClient interface {
	Send(*Log) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type pluginsInputClient struct {
	grpc.ClientStream
}

func (x *pluginsInputClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginsInputClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginsClient) Parsing(ctx context.Context, opts ...grpc.CallOption) (Plugins_ParsingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[1], Plugins_Parsing_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsParsingClient{stream}
	return x, nil
}

type Plugins_ParsingClient interface {
	Send(*Log) error
	Recv() (*Log, error)
	grpc.ClientStream
}

type pluginsParsingClient struct {
	grpc.ClientStream
}

func (x *pluginsParsingClient) Send(m *Log) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginsParsingClient) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginsClient) Analysis(ctx context.Context, opts ...grpc.CallOption) (Plugins_AnalysisClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[2], Plugins_Analysis_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsAnalysisClient{stream}
	return x, nil
}

type Plugins_AnalysisClient interface {
	Send(*Alert) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type pluginsAnalysisClient struct {
	grpc.ClientStream
}

func (x *pluginsAnalysisClient) Send(m *Alert) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginsAnalysisClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginsClient) Notification(ctx context.Context, opts ...grpc.CallOption) (Plugins_NotificationClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[3], Plugins_Notification_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsNotificationClient{stream}
	return x, nil
}

type Plugins_NotificationClient interface {
	Send(*Message) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type pluginsNotificationClient struct {
	grpc.ClientStream
}

func (x *pluginsNotificationClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginsNotificationClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginsClient) Anomalies(ctx context.Context, opts ...grpc.CallOption) (Plugins_AnomaliesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[4], Plugins_Anomalies_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsAnomaliesClient{stream}
	return x, nil
}

type Plugins_AnomaliesClient interface {
	Send(*Alert) error
	Recv() (*Ack, error)
	grpc.ClientStream
}

type pluginsAnomaliesClient struct {
	grpc.ClientStream
}

func (x *pluginsAnomaliesClient) Send(m *Alert) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pluginsAnomaliesClient) Recv() (*Ack, error) {
	m := new(Ack)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pluginsClient) Correlate(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Plugins_CorrelateClient, error) {
	stream, err := c.cc.NewStream(ctx, &Plugins_ServiceDesc.Streams[5], Plugins_Correlate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &pluginsCorrelateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Plugins_CorrelateClient interface {
	Recv() (*Alert, error)
	grpc.ClientStream
}

type pluginsCorrelateClient struct {
	grpc.ClientStream
}

func (x *pluginsCorrelateClient) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PluginsServer is the server API for Plugins service.
// All implementations must embed UnimplementedPluginsServer
// for forward compatibility
type PluginsServer interface {
	Input(Plugins_InputServer) error
	Parsing(Plugins_ParsingServer) error
	Analysis(Plugins_AnalysisServer) error
	Notification(Plugins_NotificationServer) error
	Anomalies(Plugins_AnomaliesServer) error
	Correlate(*emptypb.Empty, Plugins_CorrelateServer) error
	mustEmbedUnimplementedPluginsServer()
}

// UnimplementedPluginsServer must be embedded to have forward compatible implementations.
type UnimplementedPluginsServer struct {
}

func (UnimplementedPluginsServer) Input(Plugins_InputServer) error {
	return status.Errorf(codes.Unimplemented, "method Input not implemented")
}
func (UnimplementedPluginsServer) Parsing(Plugins_ParsingServer) error {
	return status.Errorf(codes.Unimplemented, "method Parsing not implemented")
}
func (UnimplementedPluginsServer) Analysis(Plugins_AnalysisServer) error {
	return status.Errorf(codes.Unimplemented, "method Analysis not implemented")
}
func (UnimplementedPluginsServer) Notification(Plugins_NotificationServer) error {
	return status.Errorf(codes.Unimplemented, "method Notification not implemented")
}
func (UnimplementedPluginsServer) Anomalies(Plugins_AnomaliesServer) error {
	return status.Errorf(codes.Unimplemented, "method Anomalies not implemented")
}
func (UnimplementedPluginsServer) Correlate(*emptypb.Empty, Plugins_CorrelateServer) error {
	return status.Errorf(codes.Unimplemented, "method Correlate not implemented")
}
func (UnimplementedPluginsServer) mustEmbedUnimplementedPluginsServer() {}

// UnsafePluginsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginsServer will
// result in compilation errors.
type UnsafePluginsServer interface {
	mustEmbedUnimplementedPluginsServer()
}

func RegisterPluginsServer(s grpc.ServiceRegistrar, srv PluginsServer) {
	s.RegisterService(&Plugins_ServiceDesc, srv)
}

func _Plugins_Input_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginsServer).Input(&pluginsInputServer{stream})
}

type Plugins_InputServer interface {
	Send(*Ack) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type pluginsInputServer struct {
	grpc.ServerStream
}

func (x *pluginsInputServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginsInputServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugins_Parsing_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginsServer).Parsing(&pluginsParsingServer{stream})
}

type Plugins_ParsingServer interface {
	Send(*Log) error
	Recv() (*Log, error)
	grpc.ServerStream
}

type pluginsParsingServer struct {
	grpc.ServerStream
}

func (x *pluginsParsingServer) Send(m *Log) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginsParsingServer) Recv() (*Log, error) {
	m := new(Log)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugins_Analysis_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginsServer).Analysis(&pluginsAnalysisServer{stream})
}

type Plugins_AnalysisServer interface {
	Send(*Event) error
	Recv() (*Alert, error)
	grpc.ServerStream
}

type pluginsAnalysisServer struct {
	grpc.ServerStream
}

func (x *pluginsAnalysisServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginsAnalysisServer) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugins_Notification_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginsServer).Notification(&pluginsNotificationServer{stream})
}

type Plugins_NotificationServer interface {
	Send(*Ack) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type pluginsNotificationServer struct {
	grpc.ServerStream
}

func (x *pluginsNotificationServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginsNotificationServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugins_Anomalies_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PluginsServer).Anomalies(&pluginsAnomaliesServer{stream})
}

type Plugins_AnomaliesServer interface {
	Send(*Ack) error
	Recv() (*Alert, error)
	grpc.ServerStream
}

type pluginsAnomaliesServer struct {
	grpc.ServerStream
}

func (x *pluginsAnomaliesServer) Send(m *Ack) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pluginsAnomaliesServer) Recv() (*Alert, error) {
	m := new(Alert)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Plugins_Correlate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PluginsServer).Correlate(m, &pluginsCorrelateServer{stream})
}

type Plugins_CorrelateServer interface {
	Send(*Alert) error
	grpc.ServerStream
}

type pluginsCorrelateServer struct {
	grpc.ServerStream
}

func (x *pluginsCorrelateServer) Send(m *Alert) error {
	return x.ServerStream.SendMsg(m)
}

// Plugins_ServiceDesc is the grpc.ServiceDesc for Plugins service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Plugins_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.Plugins",
	HandlerType: (*PluginsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Input",
			Handler:       _Plugins_Input_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Parsing",
			Handler:       _Plugins_Parsing_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Analysis",
			Handler:       _Plugins_Analysis_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Notification",
			Handler:       _Plugins_Notification_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Anomalies",
			Handler:       _Plugins_Anomalies_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Correlate",
			Handler:       _Plugins_Correlate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plugins.proto",
}
