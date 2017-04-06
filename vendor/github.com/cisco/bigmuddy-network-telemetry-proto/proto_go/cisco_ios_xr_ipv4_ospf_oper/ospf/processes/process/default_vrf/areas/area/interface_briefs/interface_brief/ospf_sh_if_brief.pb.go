// Code generated by protoc-gen-go.
// source: ospf_sh_if_brief.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_areas_area_interface_briefs_interface_brief is a generated protocol buffer package.

It is generated from these files:
	ospf_sh_if_brief.proto

It has these top-level messages:
	OspfShIfBrief_KEYS
	OspfShIfBrief
	OspfShInterfaceMadj
*/
package cisco_ios_xr_ipv4_ospf_oper_ospf_processes_process_default_vrf_areas_area_interface_briefs_interface_brief

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

// OSPF Interface Brief Information
type OspfShIfBrief_KEYS struct {
	ProcessName   string `protobuf:"bytes,1,opt,name=process_name,json=processName" json:"process_name,omitempty"`
	AreaId        uint32 `protobuf:"varint,2,opt,name=area_id,json=areaId" json:"area_id,omitempty"`
	InterfaceName string `protobuf:"bytes,3,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
}

func (m *OspfShIfBrief_KEYS) Reset()                    { *m = OspfShIfBrief_KEYS{} }
func (m *OspfShIfBrief_KEYS) String() string            { return proto.CompactTextString(m) }
func (*OspfShIfBrief_KEYS) ProtoMessage()               {}
func (*OspfShIfBrief_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OspfShIfBrief_KEYS) GetProcessName() string {
	if m != nil {
		return m.ProcessName
	}
	return ""
}

func (m *OspfShIfBrief_KEYS) GetAreaId() uint32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *OspfShIfBrief_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

type OspfShIfBrief struct {
	// Interface name
	InterfaceName string `protobuf:"bytes,50,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,51,opt,name=interface_area,json=interfaceArea" json:"interface_area,omitempty"`
	// Interface IP Address
	InterfaceAddress string `protobuf:"bytes,52,opt,name=interface_address,json=interfaceAddress" json:"interface_address,omitempty"`
	// Interface IP Mask
	InterfaceMask uint32 `protobuf:"varint,53,opt,name=interface_mask,json=interfaceMask" json:"interface_mask,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,54,opt,name=interface_link_cost,json=interfaceLinkCost" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,55,opt,name=ospf_interface_state,json=ospfInterfaceState" json:"ospf_interface_state,omitempty"`
	// Interface in fast detect hold down state
	InterfaceFastDetectHoldDown bool `protobuf:"varint,56,opt,name=interface_fast_detect_hold_down,json=interfaceFastDetectHoldDown" json:"interface_fast_detect_hold_down,omitempty"`
	// Total number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,57,opt,name=interface_neighbor_count,json=interfaceNeighborCount" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,58,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount" json:"interface_adj_neighbor_count,omitempty"`
	// If true, interface is multi-area
	InterfaceisMadj bool `protobuf:"varint,59,opt,name=interfaceis_madj,json=interfaceisMadj" json:"interfaceis_madj,omitempty"`
	// Total number of multi-area
	InterfaceMadjCount uint32 `protobuf:"varint,60,opt,name=interface_madj_count,json=interfaceMadjCount" json:"interface_madj_count,omitempty"`
	// Information for multi-area on the interface
	InterfaceMadjList []*OspfShInterfaceMadj `protobuf:"bytes,61,rep,name=interface_madj_list,json=interfaceMadjList" json:"interface_madj_list,omitempty"`
}

func (m *OspfShIfBrief) Reset()                    { *m = OspfShIfBrief{} }
func (m *OspfShIfBrief) String() string            { return proto.CompactTextString(m) }
func (*OspfShIfBrief) ProtoMessage()               {}
func (*OspfShIfBrief) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OspfShIfBrief) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceAddress() string {
	if m != nil {
		return m.InterfaceAddress
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceMask() uint32 {
	if m != nil {
		return m.InterfaceMask
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShIfBrief) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func (m *OspfShIfBrief) GetInterfaceFastDetectHoldDown() bool {
	if m != nil {
		return m.InterfaceFastDetectHoldDown
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceisMadj() bool {
	if m != nil {
		return m.InterfaceisMadj
	}
	return false
}

func (m *OspfShIfBrief) GetInterfaceMadjCount() uint32 {
	if m != nil {
		return m.InterfaceMadjCount
	}
	return 0
}

func (m *OspfShIfBrief) GetInterfaceMadjList() []*OspfShInterfaceMadj {
	if m != nil {
		return m.InterfaceMadjList
	}
	return nil
}

// OSPF Interface Multi-Area Information
type OspfShInterfaceMadj struct {
	// Area ID string in decimal or dotted-decimal format
	InterfaceArea string `protobuf:"bytes,1,opt,name=interface_area,json=interfaceArea" json:"interface_area,omitempty"`
	// Area ID
	MadjAreaId uint32 `protobuf:"varint,2,opt,name=madj_area_id,json=madjAreaId" json:"madj_area_id,omitempty"`
	// Number of Neighbors
	InterfaceNeighborCount uint32 `protobuf:"varint,3,opt,name=interface_neighbor_count,json=interfaceNeighborCount" json:"interface_neighbor_count,omitempty"`
	// Total number of Adjacent Neighbors
	InterfaceAdjNeighborCount uint32 `protobuf:"varint,4,opt,name=interface_adj_neighbor_count,json=interfaceAdjNeighborCount" json:"interface_adj_neighbor_count,omitempty"`
	// Interface link cost
	InterfaceLinkCost uint32 `protobuf:"varint,5,opt,name=interface_link_cost,json=interfaceLinkCost" json:"interface_link_cost,omitempty"`
	// Interface OSPF state
	OspfInterfaceState string `protobuf:"bytes,6,opt,name=ospf_interface_state,json=ospfInterfaceState" json:"ospf_interface_state,omitempty"`
}

func (m *OspfShInterfaceMadj) Reset()                    { *m = OspfShInterfaceMadj{} }
func (m *OspfShInterfaceMadj) String() string            { return proto.CompactTextString(m) }
func (*OspfShInterfaceMadj) ProtoMessage()               {}
func (*OspfShInterfaceMadj) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *OspfShInterfaceMadj) GetInterfaceArea() string {
	if m != nil {
		return m.InterfaceArea
	}
	return ""
}

func (m *OspfShInterfaceMadj) GetMadjAreaId() uint32 {
	if m != nil {
		return m.MadjAreaId
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceAdjNeighborCount() uint32 {
	if m != nil {
		return m.InterfaceAdjNeighborCount
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetInterfaceLinkCost() uint32 {
	if m != nil {
		return m.InterfaceLinkCost
	}
	return 0
}

func (m *OspfShInterfaceMadj) GetOspfInterfaceState() string {
	if m != nil {
		return m.OspfInterfaceState
	}
	return ""
}

func init() {
	proto.RegisterType((*OspfShIfBrief_KEYS)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.interface_briefs.interface_brief.ospf_sh_if_brief_KEYS")
	proto.RegisterType((*OspfShIfBrief)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.interface_briefs.interface_brief.ospf_sh_if_brief")
	proto.RegisterType((*OspfShInterfaceMadj)(nil), "cisco_ios_xr_ipv4_ospf_oper.ospf.processes.process.default_vrf.areas.area.interface_briefs.interface_brief.ospf_sh_interface_madj")
}

func init() { proto.RegisterFile("ospf_sh_if_brief.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xd5, 0x36, 0x34, 0x80, 0xdb, 0x42, 0x31, 0x50, 0x8c, 0x40, 0x62, 0x89, 0x84, 0x14, 0x84,
	0xb4, 0x42, 0x6d, 0x81, 0xf2, 0x25, 0x54, 0x35, 0x20, 0x2a, 0x5a, 0x0e, 0xe9, 0x89, 0x93, 0xe5,
	0xac, 0xbd, 0xc4, 0x9b, 0xcd, 0x3a, 0xf2, 0xb8, 0x2d, 0xe2, 0xc8, 0x5f, 0xe2, 0xce, 0x0f, 0xe2,
	0x57, 0x20, 0x4f, 0xd2, 0xac, 0x77, 0x95, 0x82, 0x10, 0x12, 0x97, 0xc8, 0x99, 0x79, 0x6f, 0xde,
	0xf8, 0x79, 0x66, 0xc9, 0x86, 0x81, 0x49, 0xc6, 0x61, 0xc8, 0x75, 0xc6, 0x07, 0x56, 0xab, 0x2c,
	0x99, 0x58, 0xe3, 0x0c, 0xcd, 0x53, 0x0d, 0xa9, 0xe1, 0xda, 0x00, 0xff, 0x62, 0xb9, 0x9e, 0x9c,
	0x6c, 0x73, 0x44, 0x9a, 0x89, 0xb2, 0x89, 0x3f, 0x79, 0x5c, 0xaa, 0x00, 0x14, 0x9c, 0x9d, 0x12,
	0xa9, 0x32, 0x71, 0x5c, 0x38, 0x7e, 0x62, 0xb3, 0x44, 0x58, 0x25, 0x00, 0x7f, 0x13, 0x5d, 0x3a,
	0x65, 0x33, 0x91, 0xaa, 0xa9, 0x00, 0x34, 0x03, 0x9d, 0xaf, 0xe4, 0x66, 0xb3, 0x0b, 0xfe, 0xe1,
	0xed, 0xa7, 0x23, 0x7a, 0x9f, 0xac, 0xce, 0x6a, 0xf3, 0x52, 0x8c, 0x15, 0x8b, 0xe2, 0xa8, 0x7b,
	0xb9, 0xbf, 0x32, 0x8b, 0x7d, 0x14, 0x63, 0x45, 0x6f, 0x91, 0x8b, 0x5e, 0x84, 0x6b, 0xc9, 0x96,
	0xe2, 0xa8, 0xbb, 0xd6, 0x6f, 0xfb, 0xbf, 0xfb, 0x92, 0x3e, 0x20, 0x57, 0x2a, 0x1d, 0x64, 0xb7,
	0x90, 0xbd, 0x36, 0x8f, 0x7a, 0x7e, 0xe7, 0xe7, 0x32, 0x59, 0x6f, 0x8a, 0x2f, 0xe0, 0x6e, 0x2e,
	0xe0, 0xd6, 0x61, 0x5e, 0x96, 0x6d, 0x35, 0x60, 0xbb, 0x56, 0x09, 0xfa, 0x88, 0x5c, 0x0b, 0x60,
	0x52, 0x5a, 0x05, 0xc0, 0xb6, 0x11, 0xb9, 0x5e, 0x21, 0xa7, 0xf1, 0x7a, 0xcd, 0xb1, 0x80, 0x11,
	0x7b, 0x82, 0xd7, 0xaa, 0x6a, 0x1e, 0x0a, 0x18, 0xd1, 0x84, 0x5c, 0xaf, 0x60, 0x85, 0x2e, 0x47,
	0x3c, 0x35, 0xe0, 0xd8, 0x53, 0xc4, 0x56, 0x72, 0x07, 0xba, 0x1c, 0xed, 0x19, 0x70, 0xf4, 0x31,
	0xb9, 0x81, 0xb7, 0xac, 0x48, 0xe0, 0x84, 0x53, 0xec, 0x19, 0xb6, 0x41, 0x7d, 0x6e, 0xff, 0x2c,
	0x75, 0xe4, 0x33, 0xb4, 0x47, 0xee, 0x55, 0xe0, 0x4c, 0x80, 0xe3, 0x52, 0x39, 0x95, 0x3a, 0x3e,
	0x34, 0x85, 0xe4, 0xd2, 0x9c, 0x96, 0x6c, 0x27, 0x8e, 0xba, 0x97, 0xfa, 0x77, 0xe6, 0xb0, 0x77,
	0x02, 0x5c, 0x0f, 0x41, 0xef, 0x4d, 0x21, 0x7b, 0xe6, 0xb4, 0xa4, 0x3b, 0x84, 0x05, 0x4e, 0x2a,
	0xfd, 0x79, 0x38, 0x30, 0x96, 0xa7, 0xe6, 0xb8, 0x74, 0xec, 0x39, 0x36, 0xbb, 0x51, 0x79, 0x3a,
	0x4b, 0xef, 0xf9, 0x2c, 0x7d, 0x43, 0xee, 0x86, 0xae, 0xe5, 0x4d, 0xf6, 0x0b, 0x64, 0xdf, 0x0e,
	0x0c, 0xcc, 0xeb, 0x05, 0x1e, 0x92, 0xca, 0x5d, 0x0d, 0x7c, 0x2c, 0x64, 0xce, 0x5e, 0x62, 0xc7,
	0x57, 0x83, 0xf8, 0xa1, 0x90, 0xb9, 0x77, 0x27, 0x34, 0x5d, 0xe6, 0x33, 0x8d, 0x57, 0xa8, 0x41,
	0x03, 0xeb, 0x65, 0x3e, 0x2d, 0xfe, 0x3d, 0x0a, 0x1f, 0x00, 0x29, 0x85, 0x06, 0xc7, 0x5e, 0xc7,
	0xad, 0xee, 0xca, 0xe6, 0xb7, 0x28, 0xf9, 0x7f, 0xeb, 0x93, 0xcc, 0xc7, 0xb7, 0xd6, 0x4f, 0x30,
	0x05, 0xbe, 0xed, 0x03, 0x0d, 0xae, 0xf3, 0x63, 0x29, 0xd8, 0xf7, 0x1a, 0x7a, 0xc1, 0x2c, 0x47,
	0x8b, 0x66, 0x39, 0x26, 0xab, 0x78, 0xd9, 0xfa, 0xce, 0x11, 0x1f, 0xdb, 0x9d, 0xee, 0xdd, 0xef,
	0x5e, 0xbc, 0xf5, 0x4f, 0x2f, 0x7e, 0xe1, 0x4f, 0x2f, 0x7e, 0xce, 0x52, 0x2c, 0xff, 0xed, 0x52,
	0xb4, 0xcf, 0x5b, 0x8a, 0x41, 0x1b, 0x3f, 0x8e, 0x5b, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe2,
	0x39, 0xf3, 0x71, 0x36, 0x05, 0x00, 0x00,
}