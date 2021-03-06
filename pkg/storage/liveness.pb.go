// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/storage/liveness.proto
// DO NOT EDIT!

package storage

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Liveness holds information about a node's latest heartbeat and epoch.
type Liveness struct {
	NodeID github_com_cockroachdb_cockroach_pkg_roachpb.NodeID `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	// Epoch is a monotonically-increasing value for node liveness. It
	// may be incremented if the liveness record expires (current time
	// is later than the expiration timestamp).
	Epoch int64 `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// The timestamp at which this liveness record expires.
	Expiration cockroach_util_hlc.Timestamp `protobuf:"bytes,3,opt,name=expiration" json:"expiration"`
	Draining   bool                         `protobuf:"varint,4,opt,name=draining,proto3" json:"draining,omitempty"`
}

func (m *Liveness) Reset()                    { *m = Liveness{} }
func (m *Liveness) String() string            { return proto.CompactTextString(m) }
func (*Liveness) ProtoMessage()               {}
func (*Liveness) Descriptor() ([]byte, []int) { return fileDescriptorLiveness, []int{0} }

func init() {
	proto.RegisterType((*Liveness)(nil), "cockroach.storage.Liveness")
}
func (m *Liveness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Liveness) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NodeID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintLiveness(dAtA, i, uint64(m.NodeID))
	}
	if m.Epoch != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintLiveness(dAtA, i, uint64(m.Epoch))
	}
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLiveness(dAtA, i, uint64(m.Expiration.Size()))
	n1, err := m.Expiration.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if m.Draining {
		dAtA[i] = 0x20
		i++
		if m.Draining {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func encodeFixed64Liveness(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Liveness(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintLiveness(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Liveness) Size() (n int) {
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovLiveness(uint64(m.NodeID))
	}
	if m.Epoch != 0 {
		n += 1 + sovLiveness(uint64(m.Epoch))
	}
	l = m.Expiration.Size()
	n += 1 + l + sovLiveness(uint64(l))
	if m.Draining {
		n += 2
	}
	return n
}

func sovLiveness(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLiveness(x uint64) (n int) {
	return sovLiveness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Liveness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiveness
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
			return fmt.Errorf("proto: Liveness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Liveness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= (github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
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
				return ErrInvalidLengthLiveness
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Expiration.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Draining", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiveness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Draining = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipLiveness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLiveness
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
func skipLiveness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiveness
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
					return 0, ErrIntOverflowLiveness
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
					return 0, ErrIntOverflowLiveness
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
				return 0, ErrInvalidLengthLiveness
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLiveness
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
				next, err := skipLiveness(dAtA[start:])
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
	ErrInvalidLengthLiveness = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiveness   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/storage/liveness.proto", fileDescriptorLiveness) }

var fileDescriptorLiveness = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x6b, 0xfa, 0x2b, 0x33, 0x11, 0x75, 0x88, 0x2a, 0xe1, 0x06, 0x84, 0x50, 0x26, 0x5b,
	0xa2, 0x4f, 0x40, 0x61, 0xa9, 0x84, 0x18, 0x22, 0xa6, 0x2e, 0x28, 0xb1, 0x2d, 0xc7, 0x6a, 0x9a,
	0x6b, 0x25, 0x2e, 0xe2, 0x31, 0x78, 0xac, 0x8e, 0x8c, 0x2c, 0x44, 0x10, 0xde, 0x82, 0x09, 0xe5,
	0x87, 0x40, 0xb7, 0x73, 0xad, 0xfb, 0xf9, 0xdc, 0x73, 0xf0, 0x05, 0x07, 0xbe, 0xc9, 0x20, 0xe4,
	0x31, 0x33, 0x1b, 0xc5, 0x72, 0x0b, 0x59, 0xa8, 0x24, 0x4b, 0xf4, 0x93, 0x4c, 0x65, 0x9e, 0x53,
	0x93, 0x81, 0x05, 0xe7, 0xa4, 0xdb, 0xa2, 0xed, 0xc6, 0xec, 0xf2, 0x10, 0xdc, 0x59, 0x9d, 0xb0,
	0x38, 0xe1, 0xcc, 0xea, 0xad, 0xcc, 0x6d, 0xb8, 0x35, 0x0d, 0x3a, 0x9b, 0x2a, 0x50, 0x50, 0x4b,
	0x56, 0xa9, 0xe6, 0xf5, 0xfc, 0x1d, 0xe1, 0xc9, 0x5d, 0xeb, 0xe1, 0xac, 0xf1, 0x38, 0x05, 0x21,
	0x1f, 0xb5, 0x70, 0x91, 0x87, 0xfc, 0xe1, 0xf2, 0xba, 0x2c, 0xe6, 0xa3, 0x7b, 0x10, 0x72, 0x75,
	0xfb, 0x5d, 0xcc, 0x17, 0x4a, 0xdb, 0x78, 0x17, 0x51, 0x0e, 0x5b, 0xd6, 0x99, 0x8a, 0x88, 0x1d,
	0x1e, 0x50, 0x2b, 0x13, 0xd1, 0x06, 0x0b, 0x46, 0xd5, 0x8f, 0x2b, 0xe1, 0x4c, 0xf1, 0x50, 0x1a,
	0xe0, 0xb1, 0x7b, 0xe4, 0x21, 0xbf, 0x1f, 0x34, 0x83, 0x73, 0x83, 0xb1, 0x7c, 0x36, 0x3a, 0x0b,
	0xad, 0x86, 0xd4, 0xed, 0x7b, 0xc8, 0x3f, 0xbe, 0x3a, 0xa5, 0x7f, 0x21, 0xab, 0x34, 0x34, 0x4e,
	0x38, 0x7d, 0xf8, 0x4d, 0xb3, 0x1c, 0xec, 0x8b, 0x79, 0x2f, 0xf8, 0x87, 0x39, 0x33, 0x3c, 0x11,
	0x59, 0xa8, 0x53, 0x9d, 0x2a, 0x77, 0xe0, 0x21, 0x7f, 0x12, 0x74, 0xf3, 0xf2, 0x6c, 0xff, 0x49,
	0x7a, 0xfb, 0x92, 0xa0, 0xd7, 0x92, 0xa0, 0xb7, 0x92, 0xa0, 0x8f, 0x92, 0xa0, 0x97, 0x2f, 0xd2,
	0x5b, 0x8f, 0xdb, 0x02, 0xa3, 0x51, 0xdd, 0xc4, 0xe2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x6d, 0xd3,
	0x75, 0x7e, 0x82, 0x01, 0x00, 0x00,
}
