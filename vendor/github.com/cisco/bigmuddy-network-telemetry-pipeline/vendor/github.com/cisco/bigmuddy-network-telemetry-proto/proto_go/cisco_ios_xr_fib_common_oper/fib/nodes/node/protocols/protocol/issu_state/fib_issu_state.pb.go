// Code generated by protoc-gen-go.
// source: fib_issu_state.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_issu_state is a generated protocol buffer package.

It is generated from these files:
	fib_issu_state.proto

It has these top-level messages:
	FibIssuState_KEYS
	FibIssuState
	FibIssuProtoState
*/
package cisco_ios_xr_fib_common_oper_fib_nodes_node_protocols_protocol_issu_state

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// FIB ISSU state
type FibIssuState_KEYS struct {
	NodeName     string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	ProtocolName string `protobuf:"bytes,2,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
}

func (m *FibIssuState_KEYS) Reset()                    { *m = FibIssuState_KEYS{} }
func (m *FibIssuState_KEYS) String() string            { return proto.CompactTextString(m) }
func (*FibIssuState_KEYS) ProtoMessage()               {}
func (*FibIssuState_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FibIssuState_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *FibIssuState_KEYS) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

type FibIssuState struct {
	// IMDR supported
	ImdrSupport bool `protobuf:"varint,50,opt,name=imdr_support,json=imdrSupport" json:"imdr_support,omitempty"`
	// SLC supported
	SlcSupport bool `protobuf:"varint,51,opt,name=slc_support,json=slcSupport" json:"slc_support,omitempty"`
	// ISSU restart
	FisIssuRestart bool `protobuf:"varint,52,opt,name=fis_issu_restart,json=fisIssuRestart" json:"fis_issu_restart,omitempty"`
	// IMDR End-of-config implicit
	ImdrEocImplicit bool `protobuf:"varint,53,opt,name=imdr_eoc_implicit,json=imdrEocImplicit" json:"imdr_eoc_implicit,omitempty"`
	// SLC End-of-config implicit
	SlcEocImplicit bool `protobuf:"varint,54,opt,name=slc_eoc_implicit,json=slcEocImplicit" json:"slc_eoc_implicit,omitempty"`
	// End-of-config received from IMDR timestamp
	EocReceivedImdrTimeStamp string `protobuf:"bytes,55,opt,name=eoc_received_imdr_time_stamp,json=eocReceivedImdrTimeStamp" json:"eoc_received_imdr_time_stamp,omitempty"`
	// End-of-config received from SLC timestamp
	EocReceivedSlcTimeStamp string `protobuf:"bytes,56,opt,name=eoc_received_slc_time_stamp,json=eocReceivedSlcTimeStamp" json:"eoc_received_slc_time_stamp,omitempty"`
	// End-of-download received from IM timestamp
	EodReceivedImTimeStamp string `protobuf:"bytes,57,opt,name=eod_received_im_time_stamp,json=eodReceivedImTimeStamp" json:"eod_received_im_time_stamp,omitempty"`
	// End-of-download send to IMDR timestamp
	EodSentImdrTimeStamp string `protobuf:"bytes,58,opt,name=eod_sent_imdr_time_stamp,json=eodSentImdrTimeStamp" json:"eod_sent_imdr_time_stamp,omitempty"`
	// End-of-download send to SLC timestamp
	EodSentSlcTimeStamp string `protobuf:"bytes,59,opt,name=eod_sent_slc_time_stamp,json=eodSentSlcTimeStamp" json:"eod_sent_slc_time_stamp,omitempty"`
	// ISSU error sent to ISSUMGR timetstamp
	FisIssuErrorTs uint64 `protobuf:"varint,60,opt,name=fis_issu_error_ts,json=fisIssuErrorTs" json:"fis_issu_error_ts,omitempty"`
	// IMDR state for the protocols
	FisProtoState []*FibIssuProtoState `protobuf:"bytes,61,rep,name=fis_proto_state,json=fisProtoState" json:"fis_proto_state,omitempty"`
}

func (m *FibIssuState) Reset()                    { *m = FibIssuState{} }
func (m *FibIssuState) String() string            { return proto.CompactTextString(m) }
func (*FibIssuState) ProtoMessage()               {}
func (*FibIssuState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FibIssuState) GetImdrSupport() bool {
	if m != nil {
		return m.ImdrSupport
	}
	return false
}

func (m *FibIssuState) GetSlcSupport() bool {
	if m != nil {
		return m.SlcSupport
	}
	return false
}

func (m *FibIssuState) GetFisIssuRestart() bool {
	if m != nil {
		return m.FisIssuRestart
	}
	return false
}

func (m *FibIssuState) GetImdrEocImplicit() bool {
	if m != nil {
		return m.ImdrEocImplicit
	}
	return false
}

func (m *FibIssuState) GetSlcEocImplicit() bool {
	if m != nil {
		return m.SlcEocImplicit
	}
	return false
}

func (m *FibIssuState) GetEocReceivedImdrTimeStamp() string {
	if m != nil {
		return m.EocReceivedImdrTimeStamp
	}
	return ""
}

func (m *FibIssuState) GetEocReceivedSlcTimeStamp() string {
	if m != nil {
		return m.EocReceivedSlcTimeStamp
	}
	return ""
}

func (m *FibIssuState) GetEodReceivedImTimeStamp() string {
	if m != nil {
		return m.EodReceivedImTimeStamp
	}
	return ""
}

func (m *FibIssuState) GetEodSentImdrTimeStamp() string {
	if m != nil {
		return m.EodSentImdrTimeStamp
	}
	return ""
}

func (m *FibIssuState) GetEodSentSlcTimeStamp() string {
	if m != nil {
		return m.EodSentSlcTimeStamp
	}
	return ""
}

func (m *FibIssuState) GetFisIssuErrorTs() uint64 {
	if m != nil {
		return m.FisIssuErrorTs
	}
	return 0
}

func (m *FibIssuState) GetFisProtoState() []*FibIssuProtoState {
	if m != nil {
		return m.FisProtoState
	}
	return nil
}

// FIB ISSU protocol state
type FibIssuProtoState struct {
	// Protocol name
	ProtocolName    string `protobuf:"bytes,1,opt,name=protocol_name,json=protocolName" json:"protocol_name,omitempty"`
	AibEodTimeStamp string `protobuf:"bytes,2,opt,name=aib_eod_time_stamp,json=aibEodTimeStamp" json:"aib_eod_time_stamp,omitempty"`
	// RSI EOD expected/valid
	RsiEodValid bool `protobuf:"varint,3,opt,name=rsi_eod_valid,json=rsiEodValid" json:"rsi_eod_valid,omitempty"`
	// RSI EOD received timestamp
	RsiEodTimeStamp string `protobuf:"bytes,4,opt,name=rsi_eod_time_stamp,json=rsiEodTimeStamp" json:"rsi_eod_time_stamp,omitempty"`
	// LSD EOD expected/valid
	LsdEodValid bool `protobuf:"varint,5,opt,name=lsd_eod_valid,json=lsdEodValid" json:"lsd_eod_valid,omitempty"`
	// LSD EOD received timestamp
	LsdEodTimeStamp string `protobuf:"bytes,6,opt,name=lsd_eod_time_stamp,json=lsdEodTimeStamp" json:"lsd_eod_time_stamp,omitempty"`
	// LMRIB EOD expected/valid
	LmribEodValid bool `protobuf:"varint,7,opt,name=lmrib_eod_valid,json=lmribEodValid" json:"lmrib_eod_valid,omitempty"`
	// LMRIB EOD received timestamp
	LmribEodTimeStamp string `protobuf:"bytes,8,opt,name=lmrib_eod_time_stamp,json=lmribEodTimeStamp" json:"lmrib_eod_time_stamp,omitempty"`
	// RIB table info valid
	RibInfoValid bool `protobuf:"varint,9,opt,name=rib_info_valid,json=ribInfoValid" json:"rib_info_valid,omitempty"`
	// Number of BCDL tables
	BcdlTables uint32 `protobuf:"varint,10,opt,name=bcdl_tables,json=bcdlTables" json:"bcdl_tables,omitempty"`
	// Number of tables converged
	ConvergedTables uint32 `protobuf:"varint,11,opt,name=converged_tables,json=convergedTables" json:"converged_tables,omitempty"`
	// All RIB tables converged timestamp
	RibTablesConvergedTimeStamp string `protobuf:"bytes,12,opt,name=rib_tables_converged_time_stamp,json=ribTablesConvergedTimeStamp" json:"rib_tables_converged_time_stamp,omitempty"`
	// Protocol EOD expected/valid
	ProtocolEodValid bool `protobuf:"varint,13,opt,name=protocol_eod_valid,json=protocolEodValid" json:"protocol_eod_valid,omitempty"`
	// Protocol EOD sent timestamp
	ProtocolEodTimeStamp string `protobuf:"bytes,14,opt,name=protocol_eod_time_stamp,json=protocolEodTimeStamp" json:"protocol_eod_time_stamp,omitempty"`
}

func (m *FibIssuProtoState) Reset()                    { *m = FibIssuProtoState{} }
func (m *FibIssuProtoState) String() string            { return proto.CompactTextString(m) }
func (*FibIssuProtoState) ProtoMessage()               {}
func (*FibIssuProtoState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FibIssuProtoState) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *FibIssuProtoState) GetAibEodTimeStamp() string {
	if m != nil {
		return m.AibEodTimeStamp
	}
	return ""
}

func (m *FibIssuProtoState) GetRsiEodValid() bool {
	if m != nil {
		return m.RsiEodValid
	}
	return false
}

func (m *FibIssuProtoState) GetRsiEodTimeStamp() string {
	if m != nil {
		return m.RsiEodTimeStamp
	}
	return ""
}

func (m *FibIssuProtoState) GetLsdEodValid() bool {
	if m != nil {
		return m.LsdEodValid
	}
	return false
}

func (m *FibIssuProtoState) GetLsdEodTimeStamp() string {
	if m != nil {
		return m.LsdEodTimeStamp
	}
	return ""
}

func (m *FibIssuProtoState) GetLmribEodValid() bool {
	if m != nil {
		return m.LmribEodValid
	}
	return false
}

func (m *FibIssuProtoState) GetLmribEodTimeStamp() string {
	if m != nil {
		return m.LmribEodTimeStamp
	}
	return ""
}

func (m *FibIssuProtoState) GetRibInfoValid() bool {
	if m != nil {
		return m.RibInfoValid
	}
	return false
}

func (m *FibIssuProtoState) GetBcdlTables() uint32 {
	if m != nil {
		return m.BcdlTables
	}
	return 0
}

func (m *FibIssuProtoState) GetConvergedTables() uint32 {
	if m != nil {
		return m.ConvergedTables
	}
	return 0
}

func (m *FibIssuProtoState) GetRibTablesConvergedTimeStamp() string {
	if m != nil {
		return m.RibTablesConvergedTimeStamp
	}
	return ""
}

func (m *FibIssuProtoState) GetProtocolEodValid() bool {
	if m != nil {
		return m.ProtocolEodValid
	}
	return false
}

func (m *FibIssuProtoState) GetProtocolEodTimeStamp() string {
	if m != nil {
		return m.ProtocolEodTimeStamp
	}
	return ""
}

func init() {
	proto.RegisterType((*FibIssuState_KEYS)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.issu_state.fib_issu_state_KEYS")
	proto.RegisterType((*FibIssuState)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.issu_state.fib_issu_state")
	proto.RegisterType((*FibIssuProtoState)(nil), "cisco_ios_xr_fib_common_oper.fib.nodes.node.protocols.protocol.issu_state.fib_issu_proto_state")
}

func init() { proto.RegisterFile("fib_issu_state.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 660 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4d, 0x6f, 0xd4, 0x3c,
	0x10, 0xc7, 0x95, 0xa7, 0x2f, 0x4f, 0x3b, 0xfb, 0xd6, 0xa6, 0x15, 0x8d, 0x28, 0x52, 0x97, 0x82,
	0xd0, 0x16, 0xd0, 0x22, 0xf5, 0x0d, 0x28, 0x85, 0x0b, 0xec, 0x61, 0x85, 0x84, 0xd0, 0x6e, 0x05,
	0xe2, 0x64, 0x25, 0xb6, 0x83, 0x2c, 0x25, 0xf1, 0xca, 0xe3, 0x56, 0x1c, 0xb9, 0xf1, 0x79, 0xf8,
	0x86, 0xc8, 0xf6, 0xc6, 0x71, 0xd8, 0x2b, 0x97, 0xaa, 0x3b, 0xf9, 0xfd, 0x7f, 0x33, 0x71, 0xc6,
	0xb0, 0x9f, 0x8b, 0x8c, 0x08, 0xc4, 0x5b, 0x82, 0x3a, 0xd5, 0x7c, 0xbc, 0x50, 0x52, 0xcb, 0x78,
	0x4a, 0x05, 0x52, 0x49, 0x84, 0x44, 0xf2, 0x43, 0x11, 0x83, 0x50, 0x59, 0x96, 0xb2, 0x22, 0x72,
	0xc1, 0xd5, 0x38, 0x17, 0xd9, 0xb8, 0x92, 0x8c, 0xa3, 0xfd, 0xeb, 0x22, 0x54, 0x16, 0xe8, 0xff,
	0x1b, 0x37, 0xc2, 0xe3, 0xaf, 0xb0, 0xd7, 0x6e, 0x41, 0x3e, 0x4e, 0xbe, 0xcd, 0xe3, 0x43, 0xd8,
	0x36, 0x61, 0x52, 0xa5, 0x25, 0x4f, 0xa2, 0x61, 0x34, 0xda, 0x9e, 0x6d, 0x99, 0xc2, 0xa7, 0xb4,
	0xe4, 0xf1, 0x23, 0xe8, 0xd5, 0x2a, 0x07, 0xfc, 0x67, 0x81, 0x6e, 0x5d, 0x34, 0xd0, 0xf1, 0xef,
	0x0d, 0xe8, 0xb7, 0xcd, 0xf1, 0x43, 0xe8, 0x8a, 0x92, 0x29, 0x82, 0xb7, 0x8b, 0x85, 0x54, 0x3a,
	0x39, 0x1d, 0x46, 0xa3, 0xad, 0x59, 0xc7, 0xd4, 0xe6, 0xae, 0x14, 0x1f, 0x41, 0x07, 0x0b, 0xea,
	0x89, 0x33, 0x4b, 0x00, 0x16, 0xb4, 0x06, 0x46, 0xb0, 0x93, 0x0b, 0x74, 0x56, 0xc5, 0x51, 0xa7,
	0x4a, 0x27, 0xe7, 0x96, 0xea, 0xe7, 0x02, 0xa7, 0x88, 0xb7, 0x33, 0x57, 0x8d, 0x9f, 0xc2, 0xae,
	0xed, 0xc6, 0x25, 0x25, 0xa2, 0x5c, 0x14, 0x82, 0x0a, 0x9d, 0x5c, 0x58, 0x74, 0x60, 0x1e, 0x4c,
	0x24, 0x9d, 0x2e, 0xcb, 0xc6, 0x6a, 0xda, 0xb6, 0xd0, 0x4b, 0x67, 0xc5, 0x82, 0x86, 0xe4, 0x3b,
	0x78, 0x60, 0x28, 0xc5, 0x29, 0x17, 0x77, 0x9c, 0x11, 0xdb, 0x42, 0x8b, 0x92, 0x9b, 0x77, 0x2c,
	0x17, 0xc9, 0x4b, 0x7b, 0x14, 0x09, 0x97, 0x74, 0xb6, 0x44, 0xa6, 0x25, 0x53, 0x37, 0xa2, 0xe4,
	0x73, 0xf3, 0x3c, 0xbe, 0x86, 0xc3, 0x56, 0xde, 0xb4, 0x0d, 0xe2, 0xaf, 0x6c, 0xfc, 0x20, 0x88,
	0xcf, 0x0b, 0xda, 0xa4, 0xaf, 0xe0, 0x3e, 0x97, 0x2c, 0xec, 0x1e, 0x86, 0x5f, 0xdb, 0xf0, 0x3d,
	0x2e, 0x59, 0xd3, 0xbb, 0xc9, 0x5e, 0x42, 0x62, 0xb2, 0xc8, 0x2b, 0xbd, 0x32, 0xf5, 0x95, 0x4d,
	0xee, 0x73, 0xc9, 0xe6, 0xbc, 0xd2, 0xed, 0x89, 0xcf, 0xe1, 0xc0, 0xe7, 0xfe, 0x9a, 0xf6, 0x8d,
	0x8d, 0xed, 0x2d, 0x63, 0xad, 0x49, 0x4f, 0x60, 0xd7, 0x7f, 0x27, 0xae, 0x94, 0x54, 0x44, 0x63,
	0x72, 0x3d, 0x8c, 0x46, 0xeb, 0xfe, 0x43, 0x4d, 0x4c, 0xf9, 0x06, 0xe3, 0x5f, 0x11, 0x0c, 0x0c,
	0x6b, 0xd7, 0xc7, 0xad, 0x4a, 0xf2, 0x76, 0xb8, 0x36, 0xea, 0x9c, 0x92, 0xf1, 0x3f, 0x5b, 0xf4,
	0xb1, 0xdf, 0xc5, 0xa0, 0xcd, 0xac, 0x97, 0x0b, 0xfc, 0x6c, 0x7e, 0xcf, 0xed, 0x65, 0xf8, 0xb9,
	0x11, 0x5c, 0xb8, 0x80, 0x5b, 0xdd, 0xf8, 0x68, 0x75, 0xe3, 0xe3, 0x67, 0x10, 0xa7, 0x22, 0x23,
	0xe6, 0xb0, 0x82, 0x33, 0x72, 0x77, 0x63, 0x90, 0x8a, 0x6c, 0x22, 0x59, 0x73, 0x3e, 0xc7, 0xd0,
	0x53, 0x28, 0x2c, 0x7c, 0x97, 0x16, 0x82, 0x25, 0x6b, 0xee, 0x32, 0x28, 0x14, 0x13, 0xc9, 0xbe,
	0x98, 0x92, 0x11, 0xd6, 0x4c, 0x20, 0x5c, 0x77, 0x42, 0x07, 0xb6, 0x84, 0x05, 0xb2, 0x40, 0xb8,
	0xe1, 0x84, 0x05, 0xb2, 0x50, 0x58, 0x33, 0x81, 0x70, 0xd3, 0x09, 0x1d, 0xd8, 0x08, 0x9f, 0xc0,
	0xa0, 0x28, 0xd5, 0xf2, 0x85, 0x9c, 0xf2, 0x7f, 0xab, 0xec, 0xd9, 0xb2, 0x97, 0xbe, 0x80, 0xfd,
	0x86, 0x0b, 0xb4, 0x5b, 0x56, 0xbb, 0x5b, 0xc3, 0x8d, 0xf8, 0x31, 0xf4, 0x0d, 0x2e, 0xaa, 0x5c,
	0x2e, 0xbd, 0xdb, 0xd6, 0xdb, 0x55, 0x22, 0x9b, 0x56, 0xb9, 0x74, 0xda, 0x23, 0xe8, 0x64, 0x94,
	0x15, 0x44, 0xa7, 0x59, 0xc1, 0x31, 0x81, 0x61, 0x34, 0xea, 0xcd, 0xc0, 0x94, 0x6e, 0x6c, 0x25,
	0x3e, 0x81, 0x1d, 0x2a, 0xab, 0x3b, 0xae, 0xbe, 0x73, 0x56, 0x53, 0x1d, 0x4b, 0x0d, 0x7c, 0x7d,
	0x89, 0x7e, 0x80, 0x23, 0xd3, 0xd1, 0x41, 0x24, 0x48, 0x35, 0xd3, 0x76, 0xed, 0xb4, 0x87, 0x4a,
	0x64, 0x2e, 0xf3, 0xde, 0x2b, 0xfc, 0xdc, 0xcf, 0x21, 0xf6, 0x4b, 0xd0, 0x9c, 0x49, 0xcf, 0xce,
	0xbe, 0x53, 0x3f, 0xf1, 0xc7, 0x72, 0x01, 0x07, 0x2d, 0x3a, 0xe8, 0xd5, 0x77, 0xb7, 0x2d, 0x88,
	0xf8, 0x26, 0xd9, 0xa6, 0xad, 0x9e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x11, 0x82, 0xfd, 0x22,
	0xf9, 0x05, 0x00, 0x00,
}
