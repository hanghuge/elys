// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/tier/userdata.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type UserData struct {
	User  string  `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Pools []*Pool `protobuf:"bytes,2,rep,name=pools,proto3" json:"pools,omitempty"`
}

func (m *UserData) Reset()         { *m = UserData{} }
func (m *UserData) String() string { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()    {}
func (*UserData) Descriptor() ([]byte, []int) {
	return fileDescriptor_604a2627f200671e, []int{0}
}
func (m *UserData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserData.Merge(m, src)
}
func (m *UserData) XXX_Size() int {
	return m.Size()
}
func (m *UserData) XXX_DiscardUnknown() {
	xxx_messageInfo_UserData.DiscardUnknown(m)
}

var xxx_messageInfo_UserData proto.InternalMessageInfo

func (m *UserData) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *UserData) GetPools() []*Pool {
	if m != nil {
		return m.Pools
	}
	return nil
}

type Pool struct {
	PoolId    string                `protobuf:"bytes,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	Pool      string                `protobuf:"bytes,2,opt,name=pool,proto3" json:"pool,omitempty"`
	FiatValue string                `protobuf:"bytes,3,opt,name=fiat_value,json=fiatValue,proto3" json:"fiat_value,omitempty"`
	Amount    cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_604a2627f200671e, []int{1}
}
func (m *Pool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pool.Merge(m, src)
}
func (m *Pool) XXX_Size() int {
	return m.Size()
}
func (m *Pool) XXX_DiscardUnknown() {
	xxx_messageInfo_Pool.DiscardUnknown(m)
}

var xxx_messageInfo_Pool proto.InternalMessageInfo

func (m *Pool) GetPoolId() string {
	if m != nil {
		return m.PoolId
	}
	return ""
}

func (m *Pool) GetPool() string {
	if m != nil {
		return m.Pool
	}
	return ""
}

func (m *Pool) GetFiatValue() string {
	if m != nil {
		return m.FiatValue
	}
	return ""
}

func init() {
	proto.RegisterType((*UserData)(nil), "elys.tier.UserData")
	proto.RegisterType((*Pool)(nil), "elys.tier.Pool")
}

func init() { proto.RegisterFile("elys/tier/userdata.proto", fileDescriptor_604a2627f200671e) }

var fileDescriptor_604a2627f200671e = []byte{
	// 292 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x33, 0x6d, 0xff, 0xfe, 0x76, 0x5c, 0x08, 0x83, 0xe2, 0x20, 0x74, 0x5a, 0x0a, 0x42,
	0x5d, 0x38, 0x03, 0x8a, 0x2f, 0x50, 0x75, 0xd1, 0x9d, 0x04, 0x74, 0xe1, 0xa6, 0x4c, 0x9b, 0x31,
	0x0d, 0x4d, 0x72, 0x43, 0x66, 0xa2, 0xf6, 0x11, 0xdc, 0xf9, 0x58, 0x5d, 0x76, 0x29, 0x2e, 0x8a,
	0x24, 0x2f, 0x22, 0x37, 0x29, 0xee, 0xce, 0x39, 0xdf, 0x9d, 0x03, 0x73, 0x28, 0x37, 0xf1, 0xda,
	0x2a, 0x17, 0x99, 0x5c, 0x15, 0xd6, 0xe4, 0x81, 0x76, 0x5a, 0x66, 0x39, 0x38, 0x60, 0x3d, 0x24,
	0x12, 0xc9, 0xd9, 0x71, 0x08, 0x21, 0xd4, 0xa9, 0x42, 0xd5, 0x1c, 0x8c, 0xee, 0xe9, 0xc1, 0xa3,
	0x35, 0xf9, 0x9d, 0x76, 0x9a, 0x31, 0xda, 0xc1, 0xe7, 0x9c, 0x0c, 0xc9, 0xb8, 0xe7, 0xd7, 0x9a,
	0x9d, 0xd3, 0x7f, 0x19, 0x40, 0x6c, 0x79, 0x6b, 0xd8, 0x1e, 0x1f, 0x5e, 0x1d, 0xc9, 0xbf, 0x42,
	0xf9, 0x00, 0x10, 0xfb, 0x0d, 0x1d, 0x7d, 0x10, 0xda, 0x41, 0xcf, 0x4e, 0xe9, 0x7f, 0x4c, 0x66,
	0x51, 0xb0, 0xaf, 0xe9, 0xa2, 0x9d, 0x06, 0x58, 0x8e, 0x8a, 0xb7, 0x9a, 0x72, 0xd4, 0xac, 0x4f,
	0xe9, 0x4b, 0xa4, 0xdd, 0xec, 0x55, 0xc7, 0x85, 0xe1, 0xed, 0x9a, 0xf4, 0x30, 0x79, 0xc2, 0x80,
	0xdd, 0xd0, 0xae, 0x4e, 0xa0, 0x48, 0x1d, 0xef, 0x20, 0x9a, 0xf4, 0x37, 0xbb, 0x81, 0xf7, 0xbd,
	0x1b, 0x9c, 0x2c, 0xc0, 0x26, 0x60, 0x6d, 0xb0, 0x92, 0x11, 0xa8, 0x44, 0xbb, 0xa5, 0x9c, 0xa6,
	0xce, 0xdf, 0x1f, 0x4f, 0x6e, 0x37, 0xa5, 0x20, 0xdb, 0x52, 0x90, 0x9f, 0x52, 0x90, 0xcf, 0x4a,
	0x78, 0xdb, 0x4a, 0x78, 0x5f, 0x95, 0xf0, 0x9e, 0x2f, 0xc2, 0xc8, 0x2d, 0x8b, 0xb9, 0x5c, 0x40,
	0xa2, 0xf0, 0x1f, 0x97, 0xa9, 0x71, 0x6f, 0x90, 0xaf, 0x6a, 0xa3, 0xde, 0x9b, 0x05, 0xdd, 0x3a,
	0x33, 0x76, 0xde, 0xad, 0xe7, 0xb9, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x47, 0xf8, 0xba,
	0x5b, 0x01, 0x00, 0x00,
}

func (m *UserData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Pools) > 0 {
		for iNdEx := len(m.Pools) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Pools[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintUserdata(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintUserdata(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Pool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintUserdata(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.FiatValue) > 0 {
		i -= len(m.FiatValue)
		copy(dAtA[i:], m.FiatValue)
		i = encodeVarintUserdata(dAtA, i, uint64(len(m.FiatValue)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Pool) > 0 {
		i -= len(m.Pool)
		copy(dAtA[i:], m.Pool)
		i = encodeVarintUserdata(dAtA, i, uint64(len(m.Pool)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PoolId) > 0 {
		i -= len(m.PoolId)
		copy(dAtA[i:], m.PoolId)
		i = encodeVarintUserdata(dAtA, i, uint64(len(m.PoolId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUserdata(dAtA []byte, offset int, v uint64) int {
	offset -= sovUserdata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovUserdata(uint64(l))
	}
	if len(m.Pools) > 0 {
		for _, e := range m.Pools {
			l = e.Size()
			n += 1 + l + sovUserdata(uint64(l))
		}
	}
	return n
}

func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PoolId)
	if l > 0 {
		n += 1 + l + sovUserdata(uint64(l))
	}
	l = len(m.Pool)
	if l > 0 {
		n += 1 + l + sovUserdata(uint64(l))
	}
	l = len(m.FiatValue)
	if l > 0 {
		n += 1 + l + sovUserdata(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovUserdata(uint64(l))
	return n
}

func sovUserdata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUserdata(x uint64) (n int) {
	return sovUserdata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserdata
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
			return fmt.Errorf("proto: UserData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
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
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pools", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pools = append(m.Pools, &Pool{})
			if err := m.Pools[len(m.Pools)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserdata
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
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUserdata
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
			return fmt.Errorf("proto: Pool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
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
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pool", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
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
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Pool = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FiatValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
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
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FiatValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUserdata
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
				return ErrInvalidLengthUserdata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUserdata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUserdata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUserdata
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
func skipUserdata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUserdata
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
					return 0, ErrIntOverflowUserdata
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
					return 0, ErrIntOverflowUserdata
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
				return 0, ErrInvalidLengthUserdata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUserdata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUserdata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUserdata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUserdata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUserdata = fmt.Errorf("proto: unexpected end of group")
)