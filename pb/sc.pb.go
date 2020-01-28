// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sc.proto

/*
	Package sc_pb is a generated protocol buffer package.

	It is generated from these files:
		sc.proto

	It has these top-level messages:
		SContract
*/
package sc_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import drive_pb "."

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SContract struct {
	Id        []byte             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Drive     *drive_pb.Contract `protobuf:"bytes,2,opt,name=drive" json:"drive,omitempty"`
	File      []byte             `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	Vmversion uint64             `protobuf:"varint,4,opt,name=vmversion,proto3" json:"vmversion,omitempty"`
	Functions []string           `protobuf:"bytes,5,rep,name=functions" json:"functions,omitempty"`
}

func (m *SContract) Reset()                    { *m = SContract{} }
func (m *SContract) String() string            { return proto.CompactTextString(m) }
func (*SContract) ProtoMessage()               {}
func (*SContract) Descriptor() ([]byte, []int) { return fileDescriptorSc, []int{0} }

func (m *SContract) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *SContract) GetDrive() *drive_pb.Contract {
	if m != nil {
		return m.Drive
	}
	return nil
}

func (m *SContract) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *SContract) GetVmversion() uint64 {
	if m != nil {
		return m.Vmversion
	}
	return 0
}

func (m *SContract) GetFunctions() []string {
	if m != nil {
		return m.Functions
	}
	return nil
}

func init() {
	proto.RegisterType((*SContract)(nil), "sc_pb.SContract")
}
func (m *SContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SContract) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSc(dAtA, i, uint64(len(m.Id)))
		i += copy(dAtA[i:], m.Id)
	}
	if m.Drive != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSc(dAtA, i, uint64(m.Drive.Size()))
		n1, err := m.Drive.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.File) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSc(dAtA, i, uint64(len(m.File)))
		i += copy(dAtA[i:], m.File)
	}
	if m.Vmversion != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintSc(dAtA, i, uint64(m.Vmversion))
	}
	if len(m.Functions) > 0 {
		for _, s := range m.Functions {
			dAtA[i] = 0x2a
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func encodeFixed64Sc(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Sc(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSc(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SContract) Size() (n int) {
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSc(uint64(l))
	}
	if m.Drive != nil {
		l = m.Drive.Size()
		n += 1 + l + sovSc(uint64(l))
	}
	l = len(m.File)
	if l > 0 {
		n += 1 + l + sovSc(uint64(l))
	}
	if m.Vmversion != 0 {
		n += 1 + sovSc(uint64(m.Vmversion))
	}
	if len(m.Functions) > 0 {
		for _, s := range m.Functions {
			l = len(s)
			n += 1 + l + sovSc(uint64(l))
		}
	}
	return n
}

func sovSc(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSc(x uint64) (n int) {
	return sovSc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSc
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
			return fmt.Errorf("proto: SContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drive", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSc
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
				return ErrInvalidLengthSc
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Drive == nil {
				m.Drive = &drive_pb.Contract{}
			}
			if err := m.Drive.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthSc
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.File = append(m.File[:0], dAtA[iNdEx:postIndex]...)
			if m.File == nil {
				m.File = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vmversion", wireType)
			}
			m.Vmversion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vmversion |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Functions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSc
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
				return ErrInvalidLengthSc
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Functions = append(m.Functions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSc
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
func skipSc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSc
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
					return 0, ErrIntOverflowSc
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
					return 0, ErrIntOverflowSc
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
				return 0, ErrInvalidLengthSc
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSc
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
				next, err := skipSc(dAtA[start:])
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
	ErrInvalidLengthSc = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSc   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sc.proto", fileDescriptorSc) }

var fileDescriptorSc = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x4e, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x4e, 0x8e, 0x2f, 0x48, 0x92, 0xd2, 0x4d, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xcb, 0x26,
	0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x25, 0xc5, 0x9d, 0x52, 0x94, 0x59, 0x96,
	0x0a, 0xe1, 0x28, 0x4d, 0x65, 0xe4, 0xe2, 0x0c, 0x76, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e,
	0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x62, 0xca, 0x4c,
	0x11, 0xd2, 0xe0, 0x62, 0x05, 0x2b, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x36, 0x12, 0xd2, 0x03,
	0xf3, 0xe2, 0x0b, 0x92, 0xf4, 0x60, 0x5a, 0x82, 0x20, 0x0a, 0x84, 0x84, 0xb8, 0x58, 0xd2, 0x32,
	0x73, 0x52, 0x25, 0x98, 0xc1, 0x7a, 0xc1, 0x6c, 0x21, 0x19, 0x2e, 0xce, 0xb2, 0xdc, 0xb2, 0xd4,
	0xa2, 0xe2, 0xcc, 0xfc, 0x3c, 0x09, 0x16, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x84, 0x00, 0x48, 0x36,
	0xad, 0x34, 0x2f, 0xb9, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0x82, 0x55, 0x81, 0x59, 0x83, 0x33, 0x08,
	0x21, 0xe0, 0x24, 0x70, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31,
	0xce, 0x78, 0x2c, 0xc7, 0x90, 0xc4, 0x06, 0x76, 0xb0, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0xb8, 0x3a, 0xbc, 0xff, 0x00, 0x00, 0x00,
}
