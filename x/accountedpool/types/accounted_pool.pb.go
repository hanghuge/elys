// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/accountedpool/accounted_pool.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/elys-network/elys/x/amm/types"
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

type AccountedPool struct {
	PoolId      uint64                                 `protobuf:"varint,1,opt,name=poolId,proto3" json:"poolId,omitempty"`
	TotalShares types.Coin                             `protobuf:"bytes,2,opt,name=totalShares,proto3" json:"totalShares"`
	PoolAssets  []types1.PoolAsset                     `protobuf:"bytes,3,rep,name=poolAssets,proto3" json:"poolAssets"`
	TotalWeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=totalWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"totalWeight"`
}

func (m *AccountedPool) Reset()         { *m = AccountedPool{} }
func (m *AccountedPool) String() string { return proto.CompactTextString(m) }
func (*AccountedPool) ProtoMessage()    {}
func (*AccountedPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_237592a8964da530, []int{0}
}
func (m *AccountedPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountedPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountedPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountedPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountedPool.Merge(m, src)
}
func (m *AccountedPool) XXX_Size() int {
	return m.Size()
}
func (m *AccountedPool) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountedPool.DiscardUnknown(m)
}

var xxx_messageInfo_AccountedPool proto.InternalMessageInfo

func (m *AccountedPool) GetPoolId() uint64 {
	if m != nil {
		return m.PoolId
	}
	return 0
}

func (m *AccountedPool) GetTotalShares() types.Coin {
	if m != nil {
		return m.TotalShares
	}
	return types.Coin{}
}

func (m *AccountedPool) GetPoolAssets() []types1.PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountedPool)(nil), "elys.accountedpool.AccountedPool")
}

func init() {
	proto.RegisterFile("elys/accountedpool/accounted_pool.proto", fileDescriptor_237592a8964da530)
}

var fileDescriptor_237592a8964da530 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x51, 0x3f, 0x4f, 0xc2, 0x40,
	0x14, 0xef, 0x09, 0x21, 0xb1, 0xc4, 0xa5, 0x1a, 0x53, 0x18, 0x8e, 0xc6, 0x41, 0xbb, 0x70, 0x17,
	0x70, 0x72, 0x04, 0x27, 0x12, 0x07, 0x82, 0x83, 0x89, 0x0b, 0xb9, 0x96, 0x4b, 0x69, 0xe8, 0xf5,
	0x11, 0xee, 0x50, 0xf9, 0x16, 0x7e, 0x2c, 0x46, 0x46, 0xe3, 0x40, 0x0c, 0xfd, 0x10, 0xae, 0xe6,
	0x7a, 0x47, 0xac, 0x4e, 0x7d, 0xbf, 0xfe, 0xfe, 0xf5, 0xbd, 0xba, 0x37, 0x3c, 0xdb, 0x48, 0xca,
	0xe2, 0x18, 0xd6, 0xb9, 0xe2, 0xb3, 0x25, 0x40, 0xf6, 0x8b, 0xa6, 0x1a, 0x92, 0xe5, 0x0a, 0x14,
	0x78, 0x9e, 0x16, 0x92, 0x3f, 0xc2, 0x76, 0xcb, 0x98, 0x85, 0xa0, 0x1a, 0x4e, 0x99, 0x94, 0x5c,
	0x19, 0x79, 0xfb, 0x22, 0x81, 0x04, 0xca, 0x91, 0xea, 0xc9, 0xbe, 0xc5, 0x31, 0x48, 0x01, 0x92,
	0x46, 0x4c, 0x72, 0xfa, 0xd2, 0x8b, 0xb8, 0x62, 0x3d, 0x1a, 0x43, 0x9a, 0x5b, 0xbe, 0x65, 0xf8,
	0xa9, 0x31, 0x1a, 0x60, 0xa8, 0xab, 0x6f, 0xe4, 0x9e, 0x0d, 0x8e, 0xed, 0x63, 0x80, 0xcc, 0xbb,
	0x74, 0x1b, 0xba, 0x76, 0x34, 0xf3, 0x51, 0x80, 0xc2, 0xfa, 0xc4, 0x22, 0x6f, 0xe0, 0x36, 0x15,
	0x28, 0x96, 0x3d, 0xce, 0xd9, 0x8a, 0x4b, 0xff, 0x24, 0x40, 0x61, 0xb3, 0xdf, 0x22, 0x36, 0x4d,
	0x57, 0x13, 0x5b, 0x4d, 0xee, 0x21, 0xcd, 0x87, 0xf5, 0xed, 0xbe, 0xe3, 0x4c, 0xaa, 0x1e, 0xef,
	0xce, 0x75, 0x75, 0xd8, 0x40, 0x2f, 0x24, 0xfd, 0x5a, 0x50, 0x0b, 0x9b, 0xfd, 0x73, 0x62, 0x2e,
	0x20, 0x04, 0x19, 0x1f, 0x39, 0xeb, 0xad, 0x88, 0xbd, 0xb1, 0x6d, 0x7f, 0xe2, 0x69, 0x32, 0x57,
	0x7e, 0x3d, 0x40, 0xe1, 0xe9, 0x90, 0x68, 0xd9, 0xe7, 0xbe, 0x73, 0x9d, 0xa4, 0x6a, 0xbe, 0x8e,
	0x48, 0x0c, 0xc2, 0x6e, 0x67, 0x1f, 0x5d, 0x39, 0x5b, 0x50, 0xb5, 0x59, 0x72, 0x49, 0x46, 0xb9,
	0x9a, 0x54, 0x23, 0x86, 0x0f, 0xdb, 0x03, 0x46, 0xbb, 0x03, 0x46, 0x5f, 0x07, 0x8c, 0xde, 0x0b,
	0xec, 0xec, 0x0a, 0xec, 0x7c, 0x14, 0xd8, 0x79, 0xee, 0x57, 0xe2, 0xf4, 0xc7, 0x75, 0x73, 0xae,
	0x5e, 0x61, 0xb5, 0x28, 0x01, 0x7d, 0xfb, 0xf7, 0x5b, 0xcb, 0xf8, 0xa8, 0x51, 0x9e, 0xf3, 0xf6,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x51, 0x3f, 0x71, 0x78, 0xf9, 0x01, 0x00, 0x00,
}

func (m *AccountedPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountedPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountedPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintAccountedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccountedPool(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.TotalShares.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintAccountedPool(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PoolId != 0 {
		i = encodeVarintAccountedPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccountedPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccountedPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AccountedPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovAccountedPool(uint64(m.PoolId))
	}
	l = m.TotalShares.Size()
	n += 1 + l + sovAccountedPool(uint64(l))
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovAccountedPool(uint64(l))
		}
	}
	l = m.TotalWeight.Size()
	n += 1 + l + sovAccountedPool(uint64(l))
	return n
}

func sovAccountedPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccountedPool(x uint64) (n int) {
	return sovAccountedPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AccountedPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccountedPool
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
			return fmt.Errorf("proto: AccountedPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountedPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountedPool
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
				return fmt.Errorf("proto: wrong wireType = %d for field TotalShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountedPool
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
				return ErrInvalidLengthAccountedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalShares.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountedPool
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
				return ErrInvalidLengthAccountedPool
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccountedPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolAssets = append(m.PoolAssets, types1.PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccountedPool
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
				return ErrInvalidLengthAccountedPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccountedPool
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
			skippy, err := skipAccountedPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccountedPool
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
func skipAccountedPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccountedPool
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
					return 0, ErrIntOverflowAccountedPool
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
					return 0, ErrIntOverflowAccountedPool
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
				return 0, ErrInvalidLengthAccountedPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccountedPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccountedPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccountedPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccountedPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccountedPool = fmt.Errorf("proto: unexpected end of group")
)
