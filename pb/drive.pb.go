// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/drive.proto

package drive_pb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Contract struct {
	Drive                []byte   `protobuf:"bytes,1,opt,name=drive,proto3" json:"drive,omitempty"`
	Owner                []byte   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Members              [][]byte `protobuf:"bytes,3,rep,name=members,proto3" json:"members,omitempty"`
	Duration             uint64   `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Created              uint64   `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	Space                uint64   `protobuf:"varint,6,opt,name=space,proto3" json:"space,omitempty"`
	Root                 []byte   `protobuf:"bytes,7,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Contract) Reset()         { *m = Contract{} }
func (m *Contract) String() string { return proto.CompactTextString(m) }
func (*Contract) ProtoMessage()    {}
func (*Contract) Descriptor() ([]byte, []int) {
	return fileDescriptor_412ba19c731df5fd, []int{0}
}
func (m *Contract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract.Merge(m, src)
}
func (m *Contract) XXX_Size() int {
	return m.Size()
}
func (m *Contract) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract.DiscardUnknown(m)
}

var xxx_messageInfo_Contract proto.InternalMessageInfo

func (m *Contract) GetDrive() []byte {
	if m != nil {
		return m.Drive
	}
	return nil
}

func (m *Contract) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Contract) GetMembers() [][]byte {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Contract) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Contract) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *Contract) GetSpace() uint64 {
	if m != nil {
		return m.Space
	}
	return 0
}

func (m *Contract) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

type Invite struct {
	Drive                []byte   `protobuf:"bytes,1,opt,name=drive,proto3" json:"drive,omitempty"`
	Owner                []byte   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Duration             uint64   `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Space                uint64   `protobuf:"varint,4,opt,name=space,proto3" json:"space,omitempty"`
	Created              uint64   `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invite) Reset()         { *m = Invite{} }
func (m *Invite) String() string { return proto.CompactTextString(m) }
func (*Invite) ProtoMessage()    {}
func (*Invite) Descriptor() ([]byte, []int) {
	return fileDescriptor_412ba19c731df5fd, []int{1}
}
func (m *Invite) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Invite) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Invite.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Invite) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invite.Merge(m, src)
}
func (m *Invite) XXX_Size() int {
	return m.Size()
}
func (m *Invite) XXX_DiscardUnknown() {
	xxx_messageInfo_Invite.DiscardUnknown(m)
}

var xxx_messageInfo_Invite proto.InternalMessageInfo

func (m *Invite) GetDrive() []byte {
	if m != nil {
		return m.Drive
	}
	return nil
}

func (m *Invite) GetOwner() []byte {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *Invite) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Invite) GetSpace() uint64 {
	if m != nil {
		return m.Space
	}
	return 0
}

func (m *Invite) GetCreated() uint64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func init() {
	proto.RegisterType((*Contract)(nil), "drive_pb.Contract")
	proto.RegisterType((*Invite)(nil), "drive_pb.Invite")
}

func init() { proto.RegisterFile("pb/drive.proto", fileDescriptor_412ba19c731df5fd) }

var fileDescriptor_412ba19c731df5fd = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x48, 0xd2, 0x4f,
	0x29, 0xca, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x73, 0xe2, 0x0b,
	0x92, 0xa4, 0x74, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3,
	0xd3, 0xf3, 0xf5, 0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x54,
	0xda, 0xc0, 0xc8, 0xc5, 0xe1, 0x9c, 0x9f, 0x57, 0x52, 0x94, 0x98, 0x5c, 0x22, 0x24, 0xc2, 0xc5,
	0x0a, 0x36, 0x47, 0x82, 0x51, 0x81, 0x51, 0x83, 0x27, 0x08, 0xc2, 0x01, 0x89, 0xe6, 0x97, 0xe7,
	0xa5, 0x16, 0x49, 0x30, 0x41, 0x44, 0xc1, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xdc, 0xa4,
	0xd4, 0xa2, 0x62, 0x09, 0x66, 0x05, 0x66, 0x0d, 0x9e, 0x20, 0x18, 0x57, 0x48, 0x8a, 0x8b, 0x23,
	0xa5, 0xb4, 0x28, 0xb1, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x45, 0x81, 0x51, 0x83, 0x25, 0x08, 0xce,
	0x07, 0xe9, 0x4a, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x4d, 0x91, 0x60, 0x05, 0x4b, 0xc1, 0xb8, 0x20,
	0x5b, 0x8a, 0x0b, 0x12, 0x93, 0x53, 0x25, 0xd8, 0xc0, 0xe2, 0x10, 0x8e, 0x90, 0x10, 0x17, 0x4b,
	0x51, 0x7e, 0x7e, 0x89, 0x04, 0x3b, 0xd8, 0x6a, 0x30, 0x5b, 0xa9, 0x81, 0x91, 0x8b, 0xcd, 0x33,
	0xaf, 0x2c, 0xb3, 0x24, 0x95, 0x24, 0x07, 0x23, 0x3b, 0x8b, 0x19, 0xcd, 0x59, 0x70, 0xcb, 0x59,
	0x90, 0x2d, 0xc7, 0xe9, 0x58, 0x27, 0x81, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc6, 0x63, 0x39, 0x86, 0x24, 0x36, 0x70, 0x70, 0x1a, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xcf, 0x83, 0x89, 0xbb, 0x99, 0x01, 0x00, 0x00,
}

func (m *Contract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Drive) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDrive(dAtA, i, uint64(len(m.Drive)))
		i += copy(dAtA[i:], m.Drive)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDrive(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if len(m.Members) > 0 {
		for _, b := range m.Members {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintDrive(dAtA, i, uint64(len(b)))
			i += copy(dAtA[i:], b)
		}
	}
	if m.Duration != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Duration))
	}
	if m.Created != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Created))
	}
	if m.Space != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Space))
	}
	if len(m.Root) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintDrive(dAtA, i, uint64(len(m.Root)))
		i += copy(dAtA[i:], m.Root)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Invite) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Invite) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Drive) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintDrive(dAtA, i, uint64(len(m.Drive)))
		i += copy(dAtA[i:], m.Drive)
	}
	if len(m.Owner) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintDrive(dAtA, i, uint64(len(m.Owner)))
		i += copy(dAtA[i:], m.Owner)
	}
	if m.Duration != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Duration))
	}
	if m.Space != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Space))
	}
	if m.Created != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintDrive(dAtA, i, uint64(m.Created))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintDrive(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Contract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Drive)
	if l > 0 {
		n += 1 + l + sovDrive(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDrive(uint64(l))
	}
	if len(m.Members) > 0 {
		for _, b := range m.Members {
			l = len(b)
			n += 1 + l + sovDrive(uint64(l))
		}
	}
	if m.Duration != 0 {
		n += 1 + sovDrive(uint64(m.Duration))
	}
	if m.Created != 0 {
		n += 1 + sovDrive(uint64(m.Created))
	}
	if m.Space != 0 {
		n += 1 + sovDrive(uint64(m.Space))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovDrive(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Invite) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Drive)
	if l > 0 {
		n += 1 + l + sovDrive(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovDrive(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovDrive(uint64(m.Duration))
	}
	if m.Space != 0 {
		n += 1 + sovDrive(uint64(m.Space))
	}
	if m.Created != 0 {
		n += 1 + sovDrive(uint64(m.Created))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDrive(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDrive(x uint64) (n int) {
	return sovDrive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrive
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
			return fmt.Errorf("proto: Contract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drive", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Drive = append(m.Drive[:0], dAtA[iNdEx:postIndex]...)
			if m.Drive == nil {
				m.Drive = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, make([]byte, postIndex-iNdEx))
			copy(m.Members[len(m.Members)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			m.Created = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Created |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Space", wireType)
			}
			m.Space = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Space |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDrive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDrive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDrive
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
func (m *Invite) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDrive
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
			return fmt.Errorf("proto: Invite: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Invite: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Drive", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Drive = append(m.Drive[:0], dAtA[iNdEx:postIndex]...)
			if m.Drive == nil {
				m.Drive = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthDrive
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthDrive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = append(m.Owner[:0], dAtA[iNdEx:postIndex]...)
			if m.Owner == nil {
				m.Owner = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Space", wireType)
			}
			m.Space = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Space |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			m.Created = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDrive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Created |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDrive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDrive
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDrive
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
func skipDrive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDrive
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
					return 0, ErrIntOverflowDrive
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
					return 0, ErrIntOverflowDrive
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
				return 0, ErrInvalidLengthDrive
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDrive
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDrive
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
				next, err := skipDrive(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDrive
				}
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
	ErrInvalidLengthDrive = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDrive   = fmt.Errorf("proto: integer overflow")
)
