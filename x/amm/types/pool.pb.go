// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type Pool struct {
	PoolId      uint64                                 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
	Address     string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PoolParams  *PoolParams                            `protobuf:"bytes,3,opt,name=poolParams,proto3" json:"poolParams,omitempty"`
	TotalShares types.Coin                             `protobuf:"bytes,4,opt,name=totalShares,proto3" json:"totalShares"`
	PoolAssets  []*PoolAsset                           `protobuf:"bytes,5,rep,name=poolAssets,proto3" json:"poolAssets,omitempty"`
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalWeight" yaml:"total_weight"`
}

func (m *Pool) Reset()         { *m = Pool{} }
func (m *Pool) String() string { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()    {}
func (*Pool) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ac3be9a215271f9, []int{0}
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

func (m *Pool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *Pool) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Pool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *Pool) GetTotalShares() types.Coin {
	if m != nil {
		return m.TotalShares
	}
	return types.Coin{}
}

func (m *Pool) GetPoolAssets() []*PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func init() {
	proto.RegisterType((*Pool)(nil), "elysnetwork.elys.amm.Pool")
}

func init() { proto.RegisterFile("elys/amm/pool.proto", fileDescriptor_3ac3be9a215271f9) }

var fileDescriptor_3ac3be9a215271f9 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x3f, 0x6f, 0xe2, 0x30,
	0x14, 0x4f, 0x20, 0xc7, 0xe9, 0xc2, 0x16, 0xd0, 0x29, 0x30, 0x24, 0x11, 0xc3, 0x29, 0x0b, 0xb6,
	0xe0, 0xb6, 0x2e, 0x2d, 0xa9, 0x3a, 0xb0, 0xa1, 0x74, 0xa8, 0xd4, 0x05, 0x39, 0xc4, 0x0a, 0x11,
	0x71, 0x1c, 0xc5, 0x6e, 0x29, 0xdf, 0xa2, 0x1f, 0x8b, 0x91, 0xb1, 0xea, 0x80, 0x2a, 0x18, 0xba,
	0xf7, 0x13, 0x54, 0x76, 0x4c, 0x15, 0xa4, 0x76, 0xca, 0x7b, 0x7e, 0xef, 0xf7, 0xcf, 0xb1, 0xd9,
	0xc1, 0xd9, 0x86, 0x41, 0x44, 0x08, 0x2c, 0x28, 0xcd, 0x40, 0x51, 0x52, 0x4e, 0xad, 0xae, 0x38,
	0xcc, 0x31, 0x5f, 0xd3, 0x72, 0x05, 0x44, 0x0d, 0x10, 0x21, 0xfd, 0xfe, 0xd9, 0xea, 0xbc, 0x40,
	0x25, 0x22, 0xac, 0x42, 0xf4, 0x7b, 0xe7, 0x33, 0xc4, 0x18, 0xe6, 0x6a, 0xd4, 0x4d, 0x68, 0x42,
	0x65, 0x09, 0x45, 0xa5, 0x4e, 0x9d, 0x05, 0x65, 0x84, 0x32, 0x18, 0x21, 0x86, 0xe1, 0xe3, 0x28,
	0xc2, 0x1c, 0x8d, 0xe0, 0x82, 0xa6, 0xf9, 0x89, 0xb0, 0x9a, 0xcf, 0x2b, 0x60, 0xd5, 0x54, 0xa3,
	0xc1, 0x7b, 0xc3, 0x34, 0x66, 0x94, 0x66, 0xd6, 0x5f, 0xb3, 0x25, 0xd4, 0xa6, 0xb1, 0xad, 0x7b,
	0xba, 0x6f, 0x84, 0xaa, 0xb3, 0x6c, 0xf3, 0x37, 0x8a, 0xe3, 0x12, 0x33, 0x66, 0x37, 0x3c, 0xdd,
	0xff, 0x13, 0x9e, 0x5a, 0xeb, 0xca, 0x34, 0xc5, 0xce, 0x4c, 0x5a, 0xb7, 0x9b, 0x9e, 0xee, 0xb7,
	0xc7, 0x1e, 0xf8, 0x2e, 0x2d, 0x98, 0x7d, 0xed, 0x85, 0x35, 0x8c, 0x35, 0x31, 0xdb, 0x9c, 0x72,
	0x94, 0xdd, 0x2e, 0x51, 0x89, 0x99, 0x6d, 0x48, 0x8a, 0x1e, 0x50, 0x06, 0x45, 0x1a, 0xa0, 0xd2,
	0x80, 0x6b, 0x9a, 0xe6, 0x81, 0xb1, 0xdd, 0xbb, 0x5a, 0x58, 0xc7, 0x58, 0x97, 0x95, 0x89, 0x89,
	0xb8, 0x23, 0x66, 0xff, 0xf2, 0x9a, 0x7e, 0x7b, 0xec, 0xfe, 0x6c, 0x42, 0xee, 0x85, 0x35, 0x88,
	0x95, 0x28, 0x0f, 0x77, 0x38, 0x4d, 0x96, 0xdc, 0x6e, 0x89, 0x8c, 0xc1, 0x8d, 0x10, 0x7a, 0xdd,
	0xbb, 0xff, 0x92, 0x94, 0x2f, 0x1f, 0x22, 0xb0, 0xa0, 0x44, 0x5d, 0x9b, 0xfa, 0x0c, 0x59, 0xbc,
	0x82, 0x7c, 0x53, 0x60, 0x06, 0xa6, 0x39, 0xff, 0xd8, 0xbb, 0x9d, 0x0d, 0x22, 0xd9, 0xc5, 0x40,
	0x52, 0xcd, 0xd7, 0x92, 0x6b, 0x10, 0xd6, 0x99, 0x83, 0x60, 0x7b, 0x70, 0xf4, 0xdd, 0xc1, 0xd1,
	0xdf, 0x0e, 0x8e, 0xfe, 0x7c, 0x74, 0xb4, 0xdd, 0xd1, 0xd1, 0x5e, 0x8e, 0x8e, 0x76, 0xef, 0xd7,
	0x54, 0x84, 0xdb, 0xa1, 0xb2, 0x2e, 0x1b, 0xf8, 0x24, 0x5f, 0x82, 0xd4, 0x8a, 0x5a, 0xf2, 0xa7,
	0xfd, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x17, 0xed, 0x73, 0x69, 0x02, 0x00, 0x00,
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
		size := m.TotalWeight.Size()
		i -= size
		if _, err := m.TotalWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	{
		size, err := m.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPool(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPool(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovPool(uint64(m.PoolId))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovPool(uint64(l))
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovPool(uint64(l))
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovPool(uint64(l))
		}
	}
	l = m.TotalWeight.Size()
	n += 1 + l + sovPool(uint64(l))
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, &PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)