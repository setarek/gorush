// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gorush.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	gorush.proto

It has these top-level messages:
	Alert
	NotificationRequest
	NotificationReply
	HealthCheckRequest
	HealthCheckResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto1.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

type Alert struct {
	Title        string   `protobuf:"bytes,1,opt,name=title" json:"title,omitempty"`
	Body         string   `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	Subtitle     string   `protobuf:"bytes,3,opt,name=subtitle" json:"subtitle,omitempty"`
	Action       string   `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	ActionLocKey string   `protobuf:"bytes,5,opt,name=actionLocKey" json:"actionLocKey,omitempty"`
	LaunchImage  string   `protobuf:"bytes,6,opt,name=launchImage" json:"launchImage,omitempty"`
	LocKey       string   `protobuf:"bytes,7,opt,name=locKey" json:"locKey,omitempty"`
	TitleLocKey  string   `protobuf:"bytes,8,opt,name=titleLocKey" json:"titleLocKey,omitempty"`
	LocArgs      []string `protobuf:"bytes,9,rep,name=locArgs" json:"locArgs,omitempty"`
	TitleLocArgs []string `protobuf:"bytes,10,rep,name=titleLocArgs" json:"titleLocArgs,omitempty"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto1.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Alert) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Alert) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Alert) GetSubtitle() string {
	if m != nil {
		return m.Subtitle
	}
	return ""
}

func (m *Alert) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Alert) GetActionLocKey() string {
	if m != nil {
		return m.ActionLocKey
	}
	return ""
}

func (m *Alert) GetLaunchImage() string {
	if m != nil {
		return m.LaunchImage
	}
	return ""
}

func (m *Alert) GetLocKey() string {
	if m != nil {
		return m.LocKey
	}
	return ""
}

func (m *Alert) GetTitleLocKey() string {
	if m != nil {
		return m.TitleLocKey
	}
	return ""
}

func (m *Alert) GetLocArgs() []string {
	if m != nil {
		return m.LocArgs
	}
	return nil
}

func (m *Alert) GetTitleLocArgs() []string {
	if m != nil {
		return m.TitleLocArgs
	}
	return nil
}

type NotificationRequest struct {
	Tokens           []string `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
	Platform         int32    `protobuf:"varint,2,opt,name=platform" json:"platform,omitempty"`
	Body             string   `protobuf:"bytes,3,opt,name=body" json:"body,omitempty"`
	Title            string   `protobuf:"bytes,4,opt,name=title" json:"title,omitempty"`
	Topic            string   `protobuf:"bytes,5,opt,name=topic" json:"topic,omitempty"`
	Key              string   `protobuf:"bytes,6,opt,name=key" json:"key,omitempty"`
	Badge            int32    `protobuf:"varint,7,opt,name=badge" json:"badge,omitempty"`
	Category         string   `protobuf:"bytes,8,opt,name=category" json:"category,omitempty"`
	Alert            *Alert   `protobuf:"bytes,9,opt,name=alert" json:"alert,omitempty"`
	Sound            string   `protobuf:"bytes,10,opt,name=sound" json:"sound,omitempty"`
	ContentAvailable bool     `protobuf:"varint,11,opt,name=contentAvailable" json:"contentAvailable,omitempty"`
	ThreadID         string   `protobuf:"bytes,12,opt,name=threadID" json:"threadID,omitempty"`
	MutableContent   bool     `protobuf:"varint,13,opt,name=mutableContent" json:"mutableContent,omitempty"`
}

func (m *NotificationRequest) Reset()                    { *m = NotificationRequest{} }
func (m *NotificationRequest) String() string            { return proto1.CompactTextString(m) }
func (*NotificationRequest) ProtoMessage()               {}
func (*NotificationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NotificationRequest) GetTokens() []string {
	if m != nil {
		return m.Tokens
	}
	return nil
}

func (m *NotificationRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *NotificationRequest) GetMessage() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *NotificationRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *NotificationRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *NotificationRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *NotificationRequest) GetBadge() int32 {
	if m != nil {
		return m.Badge
	}
	return 0
}

func (m *NotificationRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *NotificationRequest) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

func (m *NotificationRequest) GetSound() string {
	if m != nil {
		return m.Sound
	}
	return ""
}

func (m *NotificationRequest) GetContentAvailable() bool {
	if m != nil {
		return m.ContentAvailable
	}
	return false
}

func (m *NotificationRequest) GetThreadID() string {
	if m != nil {
		return m.ThreadID
	}
	return ""
}

func (m *NotificationRequest) GetMutableContent() bool {
	if m != nil {
		return m.MutableContent
	}
	return false
}

type NotificationReply struct {
	Success bool  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Counts  int32 `protobuf:"varint,2,opt,name=counts" json:"counts,omitempty"`
}

func (m *NotificationReply) Reset()                    { *m = NotificationReply{} }
func (m *NotificationReply) String() string            { return proto1.CompactTextString(m) }
func (*NotificationReply) ProtoMessage()               {}
func (*NotificationReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NotificationReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *NotificationReply) GetCounts() int32 {
	if m != nil {
		return m.Counts
	}
	return 0
}

type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=proto.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto1.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto1.RegisterType((*Alert)(nil), "proto.Alert")
	proto1.RegisterType((*NotificationRequest)(nil), "proto.NotificationRequest")
	proto1.RegisterType((*NotificationReply)(nil), "proto.NotificationReply")
	proto1.RegisterType((*HealthCheckRequest)(nil), "proto.HealthCheckRequest")
	proto1.RegisterType((*HealthCheckResponse)(nil), "proto.HealthCheckResponse")
	proto1.RegisterEnum("proto.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Gorush service

type GorushClient interface {
	Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error)
}

type gorushClient struct {
	cc *grpc.ClientConn
}

func NewGorushClient(cc *grpc.ClientConn) GorushClient {
	return &gorushClient{cc}
}

func (c *gorushClient) Send(ctx context.Context, in *NotificationRequest, opts ...grpc.CallOption) (*NotificationReply, error) {
	out := new(NotificationReply)
	err := grpc.Invoke(ctx, "/proto.Gorush/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Gorush service

type GorushServer interface {
	Send(context.Context, *NotificationRequest) (*NotificationReply, error)
}

func RegisterGorushServer(s *grpc.Server, srv GorushServer) {
	s.RegisterService(&_Gorush_serviceDesc, srv)
}

func _Gorush_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GorushServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Gorush/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GorushServer).Send(ctx, req.(*NotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gorush_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Gorush",
	HandlerType: (*GorushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Gorush_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/proto.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gorush.proto",
}

func init() { proto1.RegisterFile("gorush.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 562 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcb, 0x8e, 0xda, 0x30,
	0x14, 0x9d, 0x00, 0xe1, 0x71, 0x61, 0xa6, 0xd4, 0x53, 0x55, 0x2e, 0x2b, 0x94, 0x45, 0x85, 0xba,
	0x60, 0x41, 0x77, 0x5d, 0x8c, 0x8a, 0xa6, 0xa3, 0x19, 0x34, 0x55, 0x46, 0x0a, 0x7d, 0x2c, 0x2b,
	0x63, 0x3c, 0x21, 0xc2, 0xc4, 0x69, 0xec, 0x8c, 0xc4, 0x77, 0xf4, 0x13, 0xfa, 0x9d, 0x95, 0x2a,
	0x5f, 0x27, 0x94, 0xb4, 0x74, 0x15, 0x9f, 0x73, 0x1f, 0x3e, 0x3e, 0xf7, 0x06, 0x06, 0xb1, 0xca,
	0x0b, 0xbd, 0x99, 0x66, 0xb9, 0x32, 0x8a, 0xf8, 0xf8, 0x09, 0x7e, 0x36, 0xc0, 0x9f, 0x4b, 0x91,
	0x1b, 0xf2, 0x02, 0x7c, 0x93, 0x18, 0x29, 0xa8, 0x37, 0xf6, 0x26, 0xbd, 0xc8, 0x01, 0x42, 0xa0,
	0xb5, 0x52, 0xeb, 0x3d, 0x6d, 0x20, 0x89, 0x67, 0x32, 0x82, 0xae, 0x2e, 0x56, 0x2e, 0xb9, 0x89,
	0xfc, 0x01, 0x93, 0x97, 0xd0, 0x66, 0xdc, 0x24, 0x2a, 0xa5, 0x2d, 0x8c, 0x94, 0x88, 0x04, 0x30,
	0x70, 0xa7, 0x8f, 0x8a, 0xdf, 0x8b, 0x3d, 0xf5, 0x31, 0x5a, 0xe3, 0xc8, 0x18, 0xfa, 0x92, 0x15,
	0x29, 0xdf, 0x2c, 0x76, 0x2c, 0x16, 0xb4, 0x8d, 0x29, 0xc7, 0x94, 0xed, 0x2e, 0x5d, 0x7d, 0xc7,
	0x75, 0x97, 0x87, 0x4a, 0xbc, 0xbe, 0x6c, 0xde, 0x75, 0x95, 0x47, 0x14, 0xa1, 0xd0, 0x91, 0x8a,
	0xcf, 0xf3, 0x58, 0xd3, 0xde, 0xb8, 0x39, 0xe9, 0x45, 0x15, 0xb4, 0xca, 0xaa, 0x44, 0x0c, 0x03,
	0x86, 0x6b, 0x5c, 0xf0, 0xab, 0x01, 0x97, 0xa1, 0x32, 0xc9, 0x63, 0xc2, 0x99, 0x15, 0x1c, 0x89,
	0xef, 0x85, 0xd0, 0xc6, 0xea, 0x31, 0x6a, 0x2b, 0x52, 0x4d, 0x3d, 0xac, 0x2a, 0x91, 0x75, 0x28,
	0x93, 0xcc, 0x3c, 0xaa, 0x7c, 0x87, 0xce, 0xf9, 0xd1, 0x01, 0x5b, 0x25, 0x3b, 0xa1, 0xb5, 0x7d,
	0xa1, 0x33, 0xaf, 0x82, 0x7f, 0x26, 0xd0, 0x3a, 0x9e, 0x80, 0x65, 0x55, 0x96, 0xf0, 0xd2, 0x32,
	0x07, 0xc8, 0x10, 0x9a, 0x5b, 0xb1, 0x2f, 0x3d, 0xb2, 0x47, 0x9b, 0xb7, 0x62, 0xeb, 0x58, 0xa0,
	0x35, 0x7e, 0xe4, 0x80, 0x55, 0xc2, 0x99, 0x11, 0xb1, 0xca, 0x2b, 0x5b, 0x0e, 0x98, 0x04, 0xe0,
	0x33, 0x3b, 0x7a, 0xda, 0x1b, 0x7b, 0x93, 0xfe, 0x6c, 0xe0, 0x36, 0x63, 0x8a, 0xeb, 0x10, 0xb9,
	0x90, 0xed, 0xaa, 0x55, 0x91, 0xae, 0x29, 0xb8, 0xdb, 0x11, 0x90, 0x37, 0x30, 0xe4, 0x2a, 0x35,
	0x22, 0x35, 0xf3, 0x27, 0x96, 0x48, 0xb6, 0x92, 0x82, 0xf6, 0xc7, 0xde, 0xa4, 0x1b, 0xfd, 0xc3,
	0x5b, 0x05, 0x66, 0x93, 0x0b, 0xb6, 0x5e, 0x7c, 0xa0, 0x03, 0xa7, 0xa0, 0xc2, 0xe4, 0x35, 0x5c,
	0xec, 0x0a, 0x63, 0xd3, 0xae, 0x5d, 0x19, 0x3d, 0xc7, 0x2e, 0x7f, 0xb1, 0xc1, 0x0d, 0x3c, 0xaf,
	0xdb, 0x9f, 0x49, 0x1c, 0xa9, 0x2e, 0x38, 0x17, 0x5a, 0xe3, 0xca, 0x76, 0xa3, 0x0a, 0xda, 0xb1,
	0x70, 0x55, 0xa4, 0x46, 0x97, 0xe6, 0x97, 0x28, 0x98, 0x02, 0xb9, 0x13, 0x4c, 0x9a, 0xcd, 0xf5,
	0x46, 0xf0, 0x6d, 0x35, 0x44, 0xdb, 0x47, 0xe4, 0x4f, 0x09, 0xaf, 0x56, 0xbf, 0x82, 0xc1, 0x0f,
	0x0f, 0x2e, 0x6b, 0x05, 0x3a, 0x53, 0xa9, 0x16, 0xe4, 0x3d, 0xb4, 0xb5, 0x61, 0xa6, 0x70, 0x17,
	0x5f, 0xcc, 0x26, 0xa5, 0x73, 0x27, 0x72, 0xa7, 0x4b, 0xdb, 0x2b, 0x8d, 0x97, 0x98, 0x1f, 0x95,
	0x75, 0xc1, 0x3b, 0x38, 0xaf, 0x05, 0x48, 0x1f, 0x3a, 0x9f, 0xc3, 0xfb, 0xf0, 0xe1, 0x6b, 0x38,
	0x3c, 0xb3, 0x60, 0x79, 0x13, 0x7d, 0x59, 0x84, 0xb7, 0x43, 0x8f, 0x3c, 0x83, 0x7e, 0xf8, 0xf0,
	0xe9, 0x5b, 0x45, 0x34, 0x66, 0x77, 0xd0, 0xbe, 0xc5, 0x3f, 0x99, 0x5c, 0x41, 0x6b, 0x29, 0xd2,
	0x35, 0x19, 0x95, 0xf7, 0x9f, 0x58, 0xd1, 0x11, 0x3d, 0x19, 0xcb, 0xe4, 0x3e, 0x38, 0xb3, 0x9d,
	0x9c, 0x64, 0x72, 0x05, 0x3e, 0xca, 0x26, 0xaf, 0x4e, 0x3d, 0xc5, 0x75, 0x1a, 0xfd, 0xff, 0x95,
	0xab, 0x36, 0x86, 0xde, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xb1, 0x58, 0xde, 0x64, 0x04,
	0x00, 0x00,
}
