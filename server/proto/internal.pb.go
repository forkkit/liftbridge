// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/proto/internal.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		server/proto/internal.proto

	It has these top-level messages:
		RaftLog
		CreateStreamOp
		Stream
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type RaftLog_Op int32

const (
	RaftLog_CREATE_STREAM RaftLog_Op = 0
)

var RaftLog_Op_name = map[int32]string{
	0: "CREATE_STREAM",
}
var RaftLog_Op_value = map[string]int32{
	"CREATE_STREAM": 0,
}

func (x RaftLog_Op) String() string {
	return proto1.EnumName(RaftLog_Op_name, int32(x))
}
func (RaftLog_Op) EnumDescriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0, 0} }

type RaftLog struct {
	Op             RaftLog_Op      `protobuf:"varint,1,opt,name=op,proto3,enum=proto.RaftLog_Op" json:"op,omitempty"`
	CreateStreamOp *CreateStreamOp `protobuf:"bytes,2,opt,name=createStreamOp" json:"createStreamOp,omitempty"`
}

func (m *RaftLog) Reset()                    { *m = RaftLog{} }
func (m *RaftLog) String() string            { return proto1.CompactTextString(m) }
func (*RaftLog) ProtoMessage()               {}
func (*RaftLog) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *RaftLog) GetOp() RaftLog_Op {
	if m != nil {
		return m.Op
	}
	return RaftLog_CREATE_STREAM
}

func (m *RaftLog) GetCreateStreamOp() *CreateStreamOp {
	if m != nil {
		return m.CreateStreamOp
	}
	return nil
}

type CreateStreamOp struct {
	Stream *Stream `protobuf:"bytes,1,opt,name=stream" json:"stream,omitempty"`
}

func (m *CreateStreamOp) Reset()                    { *m = CreateStreamOp{} }
func (m *CreateStreamOp) String() string            { return proto1.CompactTextString(m) }
func (*CreateStreamOp) ProtoMessage()               {}
func (*CreateStreamOp) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *CreateStreamOp) GetStream() *Stream {
	if m != nil {
		return m.Stream
	}
	return nil
}

type Stream struct {
	Subject           string `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ReplicationFactor int32  `protobuf:"varint,3,opt,name=replicationFactor,proto3" json:"replicationFactor,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto1.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *Stream) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Stream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Stream) GetReplicationFactor() int32 {
	if m != nil {
		return m.ReplicationFactor
	}
	return 0
}

func init() {
	proto1.RegisterType((*RaftLog)(nil), "proto.RaftLog")
	proto1.RegisterType((*CreateStreamOp)(nil), "proto.CreateStreamOp")
	proto1.RegisterType((*Stream)(nil), "proto.Stream")
	proto1.RegisterEnum("proto.RaftLog_Op", RaftLog_Op_name, RaftLog_Op_value)
}
func (m *RaftLog) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftLog) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Op != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.Op))
	}
	if m.CreateStreamOp != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.CreateStreamOp.Size()))
		n1, err := m.CreateStreamOp.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *CreateStreamOp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateStreamOp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Stream != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.Stream.Size()))
		n2, err := m.Stream.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Stream) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stream) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Subject) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if len(m.Name) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintInternal(dAtA, i, uint64(len(m.Name)))
		i += copy(dAtA[i:], m.Name)
	}
	if m.ReplicationFactor != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintInternal(dAtA, i, uint64(m.ReplicationFactor))
	}
	return i, nil
}

func encodeVarintInternal(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RaftLog) Size() (n int) {
	var l int
	_ = l
	if m.Op != 0 {
		n += 1 + sovInternal(uint64(m.Op))
	}
	if m.CreateStreamOp != nil {
		l = m.CreateStreamOp.Size()
		n += 1 + l + sovInternal(uint64(l))
	}
	return n
}

func (m *CreateStreamOp) Size() (n int) {
	var l int
	_ = l
	if m.Stream != nil {
		l = m.Stream.Size()
		n += 1 + l + sovInternal(uint64(l))
	}
	return n
}

func (m *Stream) Size() (n int) {
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovInternal(uint64(l))
	}
	if m.ReplicationFactor != 0 {
		n += 1 + sovInternal(uint64(m.ReplicationFactor))
	}
	return n
}

func sovInternal(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozInternal(x uint64) (n int) {
	return sovInternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftLog) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: RaftLog: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftLog: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= (RaftLog_Op(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreateStreamOp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreateStreamOp == nil {
				m.CreateStreamOp = &CreateStreamOp{}
			}
			if err := m.CreateStreamOp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateStreamOp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: CreateStreamOp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateStreamOp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stream", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stream == nil {
				m.Stream = &Stream{}
			}
			if err := m.Stream.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Stream) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInternal
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
			return fmt.Errorf("proto: Stream: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stream: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInternal
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicationFactor", wireType)
			}
			m.ReplicationFactor = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReplicationFactor |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthInternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
					return 0, ErrIntOverflowInternal
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
				return 0, ErrInvalidLengthInternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowInternal
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
				next, err := skipInternal(dAtA[start:])
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
	ErrInvalidLengthInternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInternal   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("server/proto/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b,
	0xcc, 0xd1, 0x03, 0x73, 0x85, 0x58, 0xc1, 0x94, 0x52, 0x0b, 0x23, 0x17, 0x7b, 0x50, 0x62, 0x5a,
	0x89, 0x4f, 0x7e, 0xba, 0x90, 0x22, 0x17, 0x53, 0x7e, 0x81, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x9f,
	0x91, 0x20, 0x44, 0x99, 0x1e, 0x54, 0x4e, 0xcf, 0xbf, 0x20, 0x88, 0x29, 0xbf, 0x40, 0xc8, 0x96,
	0x8b, 0x2f, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x35, 0xb8, 0xa4, 0x28, 0x35, 0x31, 0xd7, 0xbf, 0x40,
	0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x14, 0xaa, 0xdc, 0x19, 0x45, 0x32, 0x08, 0x4d, 0xb1,
	0x92, 0x38, 0x17, 0x93, 0x7f, 0x81, 0x90, 0x20, 0x17, 0xaf, 0x73, 0x90, 0xab, 0x63, 0x88, 0x6b,
	0x7c, 0x70, 0x48, 0x90, 0xab, 0xa3, 0xaf, 0x00, 0x83, 0x92, 0x39, 0x17, 0x1f, 0xaa, 0x56, 0x21,
	0x55, 0x2e, 0xb6, 0x62, 0x30, 0x1b, 0xec, 0x20, 0x6e, 0x23, 0x5e, 0xa8, 0x0d, 0x10, 0x05, 0x41,
	0x50, 0x49, 0xa5, 0x14, 0x2e, 0x36, 0x88, 0x88, 0x90, 0x04, 0x17, 0x7b, 0x71, 0x69, 0x52, 0x56,
	0x6a, 0x72, 0x09, 0x58, 0x07, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b,
	0x0a, 0x76, 0x2a, 0x67, 0x10, 0x98, 0x2d, 0xa4, 0xc3, 0x25, 0x58, 0x94, 0x5a, 0x90, 0x93, 0x99,
	0x9c, 0x58, 0x92, 0x99, 0x9f, 0xe7, 0x96, 0x98, 0x5c, 0x92, 0x5f, 0x24, 0xc1, 0xac, 0xc0, 0xa8,
	0xc1, 0x1a, 0x84, 0x29, 0xe1, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f,
	0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x76, 0x8d, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xc8, 0x8d, 0xe5, 0x47, 0x64, 0x01, 0x00, 0x00,
}