// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: benderchain/add.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Add struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Post    string `protobuf:"bytes,3,opt,name=post,proto3" json:"post,omitempty"`
	Title   string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Body    string `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	Ipfs    string `protobuf:"bytes,6,opt,name=ipfs,proto3" json:"ipfs,omitempty"`
}

func (m *Add) Reset()         { *m = Add{} }
func (m *Add) String() string { return proto.CompactTextString(m) }
func (*Add) ProtoMessage()    {}
func (*Add) Descriptor() ([]byte, []int) {
	return fileDescriptor_baa186102f7f032b, []int{0}
}
func (m *Add) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Add) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Add.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Add) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Add.Merge(m, src)
}
func (m *Add) XXX_Size() int {
	return m.Size()
}
func (m *Add) XXX_DiscardUnknown() {
	xxx_messageInfo_Add.DiscardUnknown(m)
}

var xxx_messageInfo_Add proto.InternalMessageInfo

func (m *Add) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Add) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Add) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *Add) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Add) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Add) GetIpfs() string {
	if m != nil {
		return m.Ipfs
	}
	return ""
}

func init() {
	proto.RegisterType((*Add)(nil), "faddat.benderchain.benderchain.Add")
}

func init() { proto.RegisterFile("benderchain/add.proto", fileDescriptor_baa186102f7f032b) }

var fileDescriptor_baa186102f7f032b = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0xe3, 0x34, 0x2d, 0xc2, 0x03, 0x83, 0x55, 0x24, 0x8b, 0xc1, 0xaa, 0x98, 0xba, 0x10,
	0x4b, 0xf0, 0x05, 0x30, 0x33, 0x75, 0x64, 0xb3, 0xf3, 0xdc, 0xd4, 0x12, 0xf4, 0x59, 0xce, 0x43,
	0xa2, 0x23, 0x7f, 0xc0, 0x67, 0x31, 0x76, 0x64, 0x44, 0xc9, 0x8f, 0x20, 0xdb, 0x42, 0xca, 0x76,
	0xee, 0xf5, 0xf5, 0x70, 0x1e, 0xbf, 0xb6, 0xee, 0x08, 0x2e, 0x76, 0x07, 0xe3, 0x8f, 0xda, 0x00,
	0xb4, 0x21, 0x22, 0xa1, 0x50, 0x7b, 0x03, 0x60, 0xa8, 0x9d, 0xbd, 0xce, 0xf9, 0x66, 0xdd, 0x63,
	0x8f, 0x79, 0xaa, 0x13, 0x95, 0x5f, 0xb7, 0x9f, 0x8c, 0x2f, 0x1e, 0x01, 0x84, 0xe4, 0x17, 0x5d,
	0x74, 0x86, 0x30, 0x4a, 0xb6, 0x61, 0xdb, 0xcb, 0xdd, 0x7f, 0x14, 0x57, 0xbc, 0xf6, 0x20, 0xeb,
	0x0d, 0xdb, 0x36, 0xbb, 0xda, 0x83, 0x10, 0xbc, 0x09, 0x38, 0x90, 0x5c, 0xe4, 0x59, 0x66, 0xb1,
	0xe6, 0x4b, 0xf2, 0xf4, 0xea, 0x64, 0x93, 0xcb, 0x12, 0xd2, 0xd2, 0x22, 0x9c, 0xe4, 0xb2, 0x2c,
	0x13, 0xa7, 0xce, 0x87, 0xfd, 0x20, 0x57, 0xa5, 0x4b, 0xfc, 0xf4, 0xfc, 0x3d, 0x2a, 0x76, 0x1e,
	0x15, 0xfb, 0x1d, 0x15, 0xfb, 0x9a, 0x54, 0x75, 0x9e, 0x54, 0xf5, 0x33, 0xa9, 0xea, 0xe5, 0xbe,
	0xf7, 0x74, 0x78, 0xb7, 0x6d, 0x87, 0x6f, 0xba, 0xe8, 0xe9, 0xa2, 0x74, 0x57, 0xec, 0x3f, 0xf4,
	0xfc, 0x16, 0x74, 0x0a, 0x6e, 0xb0, 0xab, 0x2c, 0xf6, 0xf0, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0xe5, 0x81, 0xc8, 0x27, 0x01, 0x00, 0x00,
}

func (m *Add) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Add) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Add) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Ipfs) > 0 {
		i -= len(m.Ipfs)
		copy(dAtA[i:], m.Ipfs)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Ipfs)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Post) > 0 {
		i -= len(m.Post)
		copy(dAtA[i:], m.Post)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Post)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintAdd(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintAdd(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdd(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdd(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Add) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovAdd(uint64(m.Id))
	}
	l = len(m.Post)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	l = len(m.Ipfs)
	if l > 0 {
		n += 1 + l + sovAdd(uint64(l))
	}
	return n
}

func sovAdd(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdd(x uint64) (n int) {
	return sovAdd(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Add) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdd
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
			return fmt.Errorf("proto: Add: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Add: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Post", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Post = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ipfs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAdd
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdd
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ipfs = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdd(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdd
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
func skipAdd(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdd
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
					return 0, ErrIntOverflowAdd
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAdd
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
			if length < 0 {
				return 0, ErrInvalidLengthAdd
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdd
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdd
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdd        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdd          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdd = fmt.Errorf("proto: unexpected end of group")
)