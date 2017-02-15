// Code generated by protoc-gen-gogo.
// source: issue261.proto
// DO NOT EDIT!

/*
	Package issue261 is a generated protocol buffer package.

	It is generated from these files:
		issue261.proto

	It has these top-level messages:
		RepStdTypes
*/
package issue261

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/gogo/protobuf/types"

import strings "strings"
import reflect "reflect"

import github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"

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

type RepStdTypes struct {
	Durations []*time.Duration `protobuf:"bytes,4,rep,name=durations,stdduration" json:"durations,omitempty"`
}

func (m *RepStdTypes) Reset()                    { *m = RepStdTypes{} }
func (*RepStdTypes) ProtoMessage()               {}
func (*RepStdTypes) Descriptor() ([]byte, []int) { return fileDescriptorIssue261, []int{0} }

func (m *RepStdTypes) GetDurations() []*time.Duration {
	if m != nil {
		return m.Durations
	}
	return nil
}

func init() {
	proto.RegisterType((*RepStdTypes)(nil), "issue261.RepStdTypes")
}
func (this *RepStdTypes) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*RepStdTypes)
	if !ok {
		that2, ok := that.(RepStdTypes)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Durations) != len(that1.Durations) {
		return false
	}
	for i := range this.Durations {
		if dthis, dthat := this.Durations[i], that1.Durations[i]; (dthis != nil && dthat != nil && *dthis != *dthat) || (dthis != nil && dthat == nil) || (dthis == nil && dthat != nil) {
			return false
		}
	}
	return true
}
func (this *RepStdTypes) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&issue261.RepStdTypes{")
	s = append(s, "Durations: "+fmt.Sprintf("%#v", this.Durations)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringIssue261(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RepStdTypes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RepStdTypes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Durations) > 0 {
		for _, msg := range m.Durations {
			dAtA[i] = 0x22
			i++
			i = encodeVarintIssue261(dAtA, i, uint64(github_com_gogo_protobuf_types.SizeOfStdDuration(*msg)))
			n, err := github_com_gogo_protobuf_types.StdDurationMarshalTo(*msg, dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Issue261(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Issue261(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintIssue261(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RepStdTypes) Size() (n int) {
	var l int
	_ = l
	if len(m.Durations) > 0 {
		for _, e := range m.Durations {
			l = github_com_gogo_protobuf_types.SizeOfStdDuration(*e)
			n += 1 + l + sovIssue261(uint64(l))
		}
	}
	return n
}

func sovIssue261(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozIssue261(x uint64) (n int) {
	return sovIssue261(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RepStdTypes) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RepStdTypes{`,
		`Durations:` + strings.Replace(fmt.Sprintf("%v", this.Durations), "Duration", "google_protobuf1.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringIssue261(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RepStdTypes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssue261
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
			return fmt.Errorf("proto: RepStdTypes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RepStdTypes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Durations", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssue261
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
				return ErrInvalidLengthIssue261
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Durations = append(m.Durations, new(time.Duration))
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(m.Durations[len(m.Durations)-1], dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssue261(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssue261
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
func skipIssue261(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssue261
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
					return 0, ErrIntOverflowIssue261
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
					return 0, ErrIntOverflowIssue261
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
				return 0, ErrInvalidLengthIssue261
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowIssue261
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
				next, err := skipIssue261(dAtA[start:])
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
	ErrInvalidLengthIssue261 = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssue261   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("issue261.proto", fileDescriptorIssue261) }

var fileDescriptorIssue261 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x2c, 0x2e, 0x2e,
	0x4d, 0x35, 0x32, 0x33, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0xa5, 0x74,
	0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xd3, 0xf3, 0xf5,
	0xc1, 0x0a, 0x92, 0x4a, 0xd3, 0xc0, 0x3c, 0x30, 0x07, 0xcc, 0x82, 0x68, 0x94, 0x92, 0x4b, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0x45, 0xa8, 0x4a, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0xc8,
	0x2b, 0xf9, 0x70, 0x71, 0x07, 0xa5, 0x16, 0x04, 0x97, 0xa4, 0x84, 0x54, 0x16, 0xa4, 0x16, 0x0b,
	0xd9, 0x72, 0x71, 0xc2, 0x14, 0x14, 0x4b, 0xb0, 0x28, 0x30, 0x6b, 0x70, 0x1b, 0x49, 0xea, 0x41,
	0x8c, 0xd0, 0x83, 0x19, 0xa1, 0xe7, 0x02, 0x55, 0xe1, 0xc4, 0x32, 0xe3, 0xbe, 0x3c, 0x63, 0x10,
	0x42, 0x87, 0x93, 0xce, 0x85, 0x87, 0x72, 0x0c, 0x37, 0x1e, 0xca, 0x31, 0x7c, 0x78, 0x28, 0xc7,
	0xd8, 0xf0, 0x48, 0x8e, 0x71, 0xc5, 0x23, 0x39, 0xc6, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92,
	0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c,
	0x96, 0x63, 0x48, 0x62, 0x03, 0x9b, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x98, 0xb5, 0xf1,
	0xf7, 0xed, 0x00, 0x00, 0x00,
}
