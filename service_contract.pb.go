// Code generated by protoc-gen-gogo.
// source: service_contract.proto
// DO NOT EDIT!

package gmetric

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// The ResponseInfo message containing generic performance response detail
type QueryRequest struct {
	Query   string            `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Summary bool              `protobuf:"varint,2,opt,name=summary,proto3" json:"summary,omitempty"`
	Units   map[string]string `protobuf:"bytes,3,rep,name=units" json:"units,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptorServiceContract, []int{0} }

func (m *QueryRequest) GetUnits() map[string]string {
	if m != nil {
		return m.Units
	}
	return nil
}

// The ChangeDataset message containing the data difference between upstream node and current sync node details supplied with the sync request.
type QueryResponse struct {
	Error   string                             `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Metrics map[string]*OperationMetricPackage `protobuf:"bytes,2,rep,name=metrics" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptorServiceContract, []int{1} }

func (m *QueryResponse) GetMetrics() map[string]*OperationMetricPackage {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRequest)(nil), "gmetric.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "gmetric.QueryResponse")
}
func (m *QueryRequest) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *QueryRequest) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Query) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintServiceContract(data, i, uint64(len(m.Query)))
		i += copy(data[i:], m.Query)
	}
	if m.Summary {
		data[i] = 0x10
		i++
		if m.Summary {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if len(m.Units) > 0 {
		for k, _ := range m.Units {
			data[i] = 0x1a
			i++
			v := m.Units[k]
			mapSize := 1 + len(k) + sovServiceContract(uint64(len(k))) + 1 + len(v) + sovServiceContract(uint64(len(v)))
			i = encodeVarintServiceContract(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintServiceContract(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintServiceContract(data, i, uint64(len(v)))
			i += copy(data[i:], v)
		}
	}
	return i, nil
}

func (m *QueryResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *QueryResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Error) > 0 {
		data[i] = 0xa
		i++
		i = encodeVarintServiceContract(data, i, uint64(len(m.Error)))
		i += copy(data[i:], m.Error)
	}
	if len(m.Metrics) > 0 {
		for k, _ := range m.Metrics {
			data[i] = 0x12
			i++
			v := m.Metrics[k]
			msgSize := 0
			if v != nil {
				msgSize = v.Size()
				msgSize += 1 + sovServiceContract(uint64(msgSize))
			}
			mapSize := 1 + len(k) + sovServiceContract(uint64(len(k))) + msgSize
			i = encodeVarintServiceContract(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintServiceContract(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			if v != nil {
				data[i] = 0x12
				i++
				i = encodeVarintServiceContract(data, i, uint64(v.Size()))
				n1, err := v.MarshalTo(data[i:])
				if err != nil {
					return 0, err
				}
				i += n1
			}
		}
	}
	return i, nil
}

func encodeFixed64ServiceContract(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32ServiceContract(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintServiceContract(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *QueryRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.Query)
	if l > 0 {
		n += 1 + l + sovServiceContract(uint64(l))
	}
	if m.Summary {
		n += 2
	}
	if len(m.Units) > 0 {
		for k, v := range m.Units {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovServiceContract(uint64(len(k))) + 1 + len(v) + sovServiceContract(uint64(len(v)))
			n += mapEntrySize + 1 + sovServiceContract(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *QueryResponse) Size() (n int) {
	var l int
	_ = l
	l = len(m.Error)
	if l > 0 {
		n += 1 + l + sovServiceContract(uint64(l))
	}
	if len(m.Metrics) > 0 {
		for k, v := range m.Metrics {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovServiceContract(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovServiceContract(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovServiceContract(uint64(mapEntrySize))
		}
	}
	return n
}

func sovServiceContract(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozServiceContract(x uint64) (n int) {
	return sovServiceContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRequest) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Query = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Summary", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Summary = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Units", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthServiceContract
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Units == nil {
				m.Units = make(map[string]string)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceContract
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var stringLenmapvalue uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceContract
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					stringLenmapvalue |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLenmapvalue := int(stringLenmapvalue)
				if intStringLenmapvalue < 0 {
					return ErrInvalidLengthServiceContract
				}
				postStringIndexmapvalue := iNdEx + intStringLenmapvalue
				if postStringIndexmapvalue > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := string(data[iNdEx:postStringIndexmapvalue])
				iNdEx = postStringIndexmapvalue
				m.Units[mapkey] = mapvalue
			} else {
				var mapvalue string
				m.Units[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceContract(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceContract
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
func (m *QueryResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowServiceContract
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthServiceContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Error = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metrics", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthServiceContract
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLenmapkey := int(stringLenmapkey)
			if intStringLenmapkey < 0 {
				return ErrInvalidLengthServiceContract
			}
			postStringIndexmapkey := iNdEx + intStringLenmapkey
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			if m.Metrics == nil {
				m.Metrics = make(map[string]*OperationMetricPackage)
			}
			if iNdEx < postIndex {
				var valuekey uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceContract
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					valuekey |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				var mapmsglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowServiceContract
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
					iNdEx++
					mapmsglen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if mapmsglen < 0 {
					return ErrInvalidLengthServiceContract
				}
				postmsgIndex := iNdEx + mapmsglen
				if mapmsglen < 0 {
					return ErrInvalidLengthServiceContract
				}
				if postmsgIndex > l {
					return io.ErrUnexpectedEOF
				}
				mapvalue := &OperationMetricPackage{}
				if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
					return err
				}
				iNdEx = postmsgIndex
				m.Metrics[mapkey] = mapvalue
			} else {
				var mapvalue *OperationMetricPackage
				m.Metrics[mapkey] = mapvalue
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipServiceContract(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthServiceContract
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
func skipServiceContract(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowServiceContract
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowServiceContract
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthServiceContract
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowServiceContract
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipServiceContract(data[start:])
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
	ErrInvalidLengthServiceContract = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowServiceContract   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("service_contract.proto", fileDescriptorServiceContract) }

var fileDescriptorServiceContract = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0xff, 0x69, 0x69, 0xfb, 0xf7, 0xb6, 0x05, 0x19, 0x44, 0x42, 0x17, 0x31, 0xd4, 0x4d,
	0xdd, 0xcc, 0xa2, 0xa2, 0x14, 0xc1, 0x4d, 0xc1, 0xa5, 0x58, 0x03, 0xae, 0x5c, 0xc8, 0x38, 0x5e,
	0xca, 0x50, 0x33, 0xd3, 0xce, 0x4c, 0x02, 0x79, 0x13, 0x1f, 0xc3, 0x57, 0x70, 0xe7, 0xd2, 0x47,
	0x90, 0xf8, 0x22, 0x92, 0x4c, 0xa3, 0x41, 0xdc, 0xe5, 0xdc, 0x9c, 0x7b, 0xce, 0x37, 0x5c, 0x38,
	0xb0, 0x68, 0x32, 0x29, 0xf0, 0x5e, 0x68, 0xe5, 0x0c, 0x17, 0x8e, 0x6d, 0x8c, 0x76, 0x9a, 0xf6,
	0x56, 0x09, 0x3a, 0x23, 0xc5, 0x78, 0xf8, 0xa8, 0x13, 0x2e, 0x95, 0x1f, 0x4f, 0x5e, 0x08, 0x0c,
	0x6f, 0x52, 0x34, 0x79, 0x8c, 0xdb, 0x14, 0xad, 0xa3, 0xfb, 0xd0, 0xd9, 0x96, 0x3a, 0x20, 0x11,
	0x99, 0xf6, 0x63, 0x2f, 0x68, 0x00, 0x3d, 0x9b, 0x26, 0x09, 0x37, 0x79, 0xd0, 0x8a, 0xc8, 0xf4,
	0x7f, 0x5c, 0x4b, 0x7a, 0x06, 0x9d, 0x54, 0x49, 0x67, 0x83, 0x76, 0xd4, 0x9e, 0x0e, 0x66, 0x11,
	0xdb, 0xf5, 0xb0, 0x66, 0x2a, 0xbb, 0x2d, 0x2d, 0x97, 0xca, 0x99, 0x3c, 0xf6, 0xf6, 0xf1, 0x1c,
	0xe0, 0x67, 0x48, 0xf7, 0xa0, 0xbd, 0xc6, 0xba, 0xb3, 0xfc, 0x2c, 0x39, 0x32, 0xfe, 0x94, 0x62,
	0xd5, 0xd7, 0x8f, 0xbd, 0x38, 0x6f, 0xcd, 0xc9, 0xe4, 0x95, 0xc0, 0x68, 0x17, 0x6e, 0x37, 0x5a,
	0x59, 0x2c, 0xbd, 0x68, 0x8c, 0x36, 0x35, 0x73, 0x25, 0xe8, 0x05, 0xf4, 0x3c, 0x8a, 0x0d, 0x5a,
	0x15, 0xdb, 0xd1, 0x6f, 0x36, 0xbf, 0xce, 0xae, 0xbc, 0xcb, 0xe3, 0xd5, 0x3b, 0xe3, 0x3b, 0x18,
	0x36, 0x7f, 0xfc, 0x81, 0x78, 0xda, 0x44, 0x1c, 0xcc, 0x0e, 0xbf, 0xe3, 0xaf, 0x37, 0x68, 0xb8,
	0x93, 0x5a, 0xf9, 0x80, 0x25, 0x17, 0x6b, 0xbe, 0xc2, 0xc6, 0x1b, 0x16, 0xc7, 0x6f, 0x45, 0x48,
	0xde, 0x8b, 0x90, 0x7c, 0x14, 0x21, 0x79, 0xfe, 0x0c, 0xff, 0xc1, 0x48, 0xe8, 0x84, 0x65, 0x92,
	0x2b, 0xc7, 0x84, 0xc0, 0x45, 0xd7, 0x5f, 0x69, 0x49, 0x1e, 0xba, 0xd5, 0xa1, 0x4e, 0xbe, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x5a, 0x71, 0x9d, 0x4f, 0xd9, 0x01, 0x00, 0x00,
}
