// Code generated by protoc-gen-go-vtproto. DO NOT EDIT.
// protoc-gen-go-vtproto version: (devel)
// source: testproto/grpc/grpc.proto

package grpc

import (
	context "context"
	fmt "fmt"
	protohelpers "github.com/maheeshap-canopus/vtprotobuf/protohelpers"
	inner "github.com/maheeshap-canopus/vtprotobuf/testproto/grpc/inner"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	proto "google.golang.org/protobuf/proto"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

func (m *LocalTestMessageRequest) CloneVT() *LocalTestMessageRequest {
	if m == nil {
		return (*LocalTestMessageRequest)(nil)
	}
	r := LocalTestMessageRequestFromVTPool()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *LocalTestMessageRequest) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (m *LocalTestMessageResponse) CloneVT() *LocalTestMessageResponse {
	if m == nil {
		return (*LocalTestMessageResponse)(nil)
	}
	r := LocalTestMessageResponseFromVTPool()
	if len(m.unknownFields) > 0 {
		r.unknownFields = make([]byte, len(m.unknownFields))
		copy(r.unknownFields, m.unknownFields)
	}
	return r
}

func (m *LocalTestMessageResponse) CloneMessageVT() proto.Message {
	return m.CloneVT()
}

func (this *LocalTestMessageRequest) EqualVT(that *LocalTestMessageRequest) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *LocalTestMessageRequest) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*LocalTestMessageRequest)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}
func (this *LocalTestMessageResponse) EqualVT(that *LocalTestMessageResponse) bool {
	if this == that {
		return true
	} else if this == nil || that == nil {
		return false
	}
	return string(this.unknownFields) == string(that.unknownFields)
}

func (this *LocalTestMessageResponse) EqualMessageVT(thatMsg proto.Message) bool {
	that, ok := thatMsg.(*LocalTestMessageResponse)
	if !ok {
		return false
	}
	return this.EqualVT(that)
}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	TestRPC(ctx context.Context, in *inner.TestMessageRequest, opts ...grpc.CallOption) (*inner.TestMessageResponse, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) TestRPC(ctx context.Context, in *inner.TestMessageRequest, opts ...grpc.CallOption) (*inner.TestMessageResponse, error) {
	out := inner.TestMessageResponseFromVTPool()
	err := c.cc.Invoke(ctx, "/TestService/TestRPC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations must embed UnimplementedTestServiceServer
// for forward compatibility
type TestServiceServer interface {
	TestRPC(context.Context, *inner.TestMessageRequest) (*inner.TestMessageResponse, error)
	mustEmbedUnimplementedTestServiceServer()
}

// UnimplementedTestServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceServer struct {
}

func (UnimplementedTestServiceServer) TestRPC(context.Context, *inner.TestMessageRequest) (*inner.TestMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRPC not implemented")
}
func (UnimplementedTestServiceServer) mustEmbedUnimplementedTestServiceServer() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_TestRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := inner.TestMessageRequestFromVTPool()
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).TestRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestService/TestRPC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).TestRPC(ctx, req.(*inner.TestMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRPC",
			Handler:    _TestService_TestRPC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testproto/grpc/grpc.proto",
}

// TestServiceLocalClient is the client API for TestServiceLocal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceLocalClient interface {
	TestRPCLocal(ctx context.Context, in *LocalTestMessageRequest, opts ...grpc.CallOption) (*LocalTestMessageResponse, error)
}

type testServiceLocalClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceLocalClient(cc grpc.ClientConnInterface) TestServiceLocalClient {
	return &testServiceLocalClient{cc}
}

func (c *testServiceLocalClient) TestRPCLocal(ctx context.Context, in *LocalTestMessageRequest, opts ...grpc.CallOption) (*LocalTestMessageResponse, error) {
	out := LocalTestMessageResponseFromVTPool()
	err := c.cc.Invoke(ctx, "/TestServiceLocal/TestRPCLocal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceLocalServer is the server API for TestServiceLocal service.
// All implementations must embed UnimplementedTestServiceLocalServer
// for forward compatibility
type TestServiceLocalServer interface {
	TestRPCLocal(context.Context, *LocalTestMessageRequest) (*LocalTestMessageResponse, error)
	mustEmbedUnimplementedTestServiceLocalServer()
}

// UnimplementedTestServiceLocalServer must be embedded to have forward compatible implementations.
type UnimplementedTestServiceLocalServer struct {
}

func (UnimplementedTestServiceLocalServer) TestRPCLocal(context.Context, *LocalTestMessageRequest) (*LocalTestMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestRPCLocal not implemented")
}
func (UnimplementedTestServiceLocalServer) mustEmbedUnimplementedTestServiceLocalServer() {}

// UnsafeTestServiceLocalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceLocalServer will
// result in compilation errors.
type UnsafeTestServiceLocalServer interface {
	mustEmbedUnimplementedTestServiceLocalServer()
}

func RegisterTestServiceLocalServer(s grpc.ServiceRegistrar, srv TestServiceLocalServer) {
	s.RegisterService(&TestServiceLocal_ServiceDesc, srv)
}

func _TestServiceLocal_TestRPCLocal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := LocalTestMessageRequestFromVTPool()
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceLocalServer).TestRPCLocal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TestServiceLocal/TestRPCLocal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceLocalServer).TestRPCLocal(ctx, req.(*LocalTestMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TestServiceLocal_ServiceDesc is the grpc.ServiceDesc for TestServiceLocal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestServiceLocal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TestServiceLocal",
	HandlerType: (*TestServiceLocalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestRPCLocal",
			Handler:    _TestServiceLocal_TestRPCLocal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "testproto/grpc/grpc.proto",
}

func (m *LocalTestMessageRequest) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalTestMessageRequest) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *LocalTestMessageRequest) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *LocalTestMessageResponse) MarshalVT() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVT(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalTestMessageResponse) MarshalToVT(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVT(dAtA[:size])
}

func (m *LocalTestMessageResponse) MarshalToSizedBufferVT(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *LocalTestMessageRequest) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalTestMessageRequest) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *LocalTestMessageRequest) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

func (m *LocalTestMessageResponse) MarshalVTStrict() (dAtA []byte, err error) {
	if m == nil {
		return nil, nil
	}
	size := m.SizeVT()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBufferVTStrict(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LocalTestMessageResponse) MarshalToVTStrict(dAtA []byte) (int, error) {
	size := m.SizeVT()
	return m.MarshalToSizedBufferVTStrict(dAtA[:size])
}

func (m *LocalTestMessageResponse) MarshalToSizedBufferVTStrict(dAtA []byte) (int, error) {
	if m == nil {
		return 0, nil
	}
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.unknownFields != nil {
		i -= len(m.unknownFields)
		copy(dAtA[i:], m.unknownFields)
	}
	return len(dAtA) - i, nil
}

var vtprotoPool_LocalTestMessageRequest = sync.Pool{
	New: func() interface{} {
		return &LocalTestMessageRequest{}
	},
}

func (m *LocalTestMessageRequest) ResetVT() {
	if m != nil {
		m.Reset()
	}
}
func (m *LocalTestMessageRequest) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_LocalTestMessageRequest.Put(m)
	}
}
func LocalTestMessageRequestFromVTPool() *LocalTestMessageRequest {
	return vtprotoPool_LocalTestMessageRequest.Get().(*LocalTestMessageRequest)
}

var vtprotoPool_LocalTestMessageResponse = sync.Pool{
	New: func() interface{} {
		return &LocalTestMessageResponse{}
	},
}

func (m *LocalTestMessageResponse) ResetVT() {
	if m != nil {
		m.Reset()
	}
}
func (m *LocalTestMessageResponse) ReturnToVTPool() {
	if m != nil {
		m.ResetVT()
		vtprotoPool_LocalTestMessageResponse.Put(m)
	}
}
func LocalTestMessageResponseFromVTPool() *LocalTestMessageResponse {
	return vtprotoPool_LocalTestMessageResponse.Get().(*LocalTestMessageResponse)
}
func (m *LocalTestMessageRequest) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *LocalTestMessageResponse) SizeVT() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += len(m.unknownFields)
	return n
}

func (m *LocalTestMessageRequest) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalTestMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalTestMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalTestMessageResponse) UnmarshalVT(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalTestMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalTestMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalTestMessageRequest) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalTestMessageRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalTestMessageRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *LocalTestMessageResponse) UnmarshalVTUnsafe(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return protohelpers.ErrIntOverflow
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: LocalTestMessageResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LocalTestMessageResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := protohelpers.Skip(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return protohelpers.ErrInvalidLength
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.unknownFields = append(m.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
