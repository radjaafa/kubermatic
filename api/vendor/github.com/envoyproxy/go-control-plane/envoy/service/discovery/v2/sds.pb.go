// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/service/discovery/v2/sds.proto

package v2

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import v2 "github.com/envoyproxy/go-control-plane/envoy/api/v2"
import _ "github.com/gogo/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#not-implemented-hide:] Not configuration. Workaround c++ protobuf issue with importing
// services: https://github.com/google/protobuf/issues/4221
type SdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SdsDummy) Reset()         { *m = SdsDummy{} }
func (m *SdsDummy) String() string { return proto.CompactTextString(m) }
func (*SdsDummy) ProtoMessage()    {}
func (*SdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_sds_b4c12a2ca9e02474, []int{0}
}
func (m *SdsDummy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SdsDummy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SdsDummy.Merge(dst, src)
}
func (m *SdsDummy) XXX_Size() int {
	return m.Size()
}
func (m *SdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_SdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_SdsDummy proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SdsDummy)(nil), "envoy.service.discovery.v2.SdsDummy")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretDiscoveryServiceClient is the client API for SecretDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretDiscoveryServiceClient interface {
	StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error)
	FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type secretDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretDiscoveryServiceClient(cc *grpc.ClientConn) SecretDiscoveryServiceClient {
	return &secretDiscoveryServiceClient{cc}
}

func (c *secretDiscoveryServiceClient) StreamSecrets(ctx context.Context, opts ...grpc.CallOption) (SecretDiscoveryService_StreamSecretsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SecretDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v2.SecretDiscoveryService/StreamSecrets", opts...)
	if err != nil {
		return nil, err
	}
	x := &secretDiscoveryServiceStreamSecretsClient{stream}
	return x, nil
}

type SecretDiscoveryService_StreamSecretsClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type secretDiscoveryServiceStreamSecretsClient struct {
	grpc.ClientStream
}

func (x *secretDiscoveryServiceStreamSecretsClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *secretDiscoveryServiceClient) FetchSecrets(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretDiscoveryServiceServer is the server API for SecretDiscoveryService service.
type SecretDiscoveryServiceServer interface {
	StreamSecrets(SecretDiscoveryService_StreamSecretsServer) error
	FetchSecrets(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
}

func RegisterSecretDiscoveryServiceServer(s *grpc.Server, srv SecretDiscoveryServiceServer) {
	s.RegisterService(&_SecretDiscoveryService_serviceDesc, srv)
}

func _SecretDiscoveryService_StreamSecrets_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SecretDiscoveryServiceServer).StreamSecrets(&secretDiscoveryServiceStreamSecretsServer{stream})
}

type SecretDiscoveryService_StreamSecretsServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type secretDiscoveryServiceStreamSecretsServer struct {
	grpc.ServerStream
}

func (x *secretDiscoveryServiceStreamSecretsServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *secretDiscoveryServiceStreamSecretsServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SecretDiscoveryService_FetchSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v2.SecretDiscoveryService/FetchSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretDiscoveryServiceServer).FetchSecrets(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v2.SecretDiscoveryService",
	HandlerType: (*SecretDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchSecrets",
			Handler:    _SecretDiscoveryService_FetchSecrets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamSecrets",
			Handler:       _SecretDiscoveryService_StreamSecrets_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v2/sds.proto",
}

func (m *SdsDummy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SdsDummy) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SdsDummy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSds(x uint64) (n int) {
	return sovSds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SdsDummy) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSds
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SdsDummy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SdsDummy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipSds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSds
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSds
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSds
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSds
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSds(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSds   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/service/discovery/v2/sds.proto", fileDescriptor_sds_b4c12a2ca9e02474)
}

var fileDescriptor_sds_b4c12a2ca9e02474 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x90, 0xb1, 0x4a, 0x04, 0x31,
	0x10, 0x86, 0x8d, 0x85, 0x48, 0xd0, 0x66, 0x41, 0x8b, 0x70, 0xac, 0xb2, 0x58, 0x1c, 0x16, 0x59,
	0x59, 0x1b, 0xb9, 0x52, 0x0e, 0x6b, 0x71, 0xc1, 0x3e, 0xee, 0x0e, 0x67, 0xc0, 0xcd, 0xe4, 0x32,
	0xb9, 0xe0, 0xb6, 0xbe, 0x82, 0x2f, 0x65, 0x29, 0xf8, 0x02, 0xb2, 0xfa, 0x20, 0x62, 0x72, 0x9e,
	0x58, 0x6c, 0x67, 0xfd, 0xfd, 0xff, 0x37, 0xc3, 0xcf, 0x4f, 0xc0, 0x04, 0xec, 0x4b, 0x02, 0x17,
	0x74, 0x03, 0x65, 0xab, 0xa9, 0xc1, 0x00, 0xae, 0x2f, 0x43, 0x55, 0x52, 0x4b, 0xd2, 0x3a, 0xf4,
	0x98, 0x89, 0x98, 0x92, 0xeb, 0x94, 0xdc, 0xa4, 0x64, 0xa8, 0xc4, 0x24, 0x19, 0x94, 0xd5, 0xdf,
	0x9d, 0x5f, 0x14, 0x9b, 0x62, 0xb2, 0x40, 0x5c, 0x3c, 0x40, 0xc4, 0xca, 0x18, 0xf4, 0xca, 0x6b,
	0x34, 0x6b, 0x6f, 0xc1, 0xf9, 0x6e, 0xdd, 0xd2, 0x7c, 0xd5, 0x75, 0x7d, 0xf5, 0xc1, 0xf8, 0x61,
	0x0d, 0x8d, 0x03, 0x3f, 0xff, 0x71, 0xd4, 0xe9, 0x5e, 0x76, 0xcb, 0xf7, 0x6b, 0xef, 0x40, 0x75,
	0x89, 0x53, 0x96, 0xcb, 0xf4, 0x90, 0xb2, 0x5a, 0x86, 0x4a, 0x6e, 0x0a, 0x37, 0xb0, 0x5c, 0x01,
	0x79, 0x71, 0x34, 0xca, 0xc9, 0xa2, 0x21, 0x28, 0xb6, 0xa6, 0xec, 0x8c, 0x65, 0x4b, 0xbe, 0x77,
	0x05, 0xbe, 0xb9, 0xff, 0x37, 0xed, 0xf1, 0xd3, 0xdb, 0xe7, 0xf3, 0xb6, 0x28, 0x0e, 0xfe, 0x4c,
	0x31, 0xa3, 0xe4, 0x9f, 0xb1, 0xd3, 0xcb, 0x8b, 0x97, 0x21, 0x67, 0xaf, 0x43, 0xce, 0xde, 0x87,
	0x9c, 0xf1, 0xa9, 0xc6, 0xa4, 0xb4, 0x0e, 0x1f, 0x7b, 0x39, 0xbe, 0xf2, 0x35, 0xbb, 0xdb, 0x89,
	0x93, 0x9d, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xd7, 0x7c, 0x41, 0xb2, 0x01, 0x00, 0x00,
}