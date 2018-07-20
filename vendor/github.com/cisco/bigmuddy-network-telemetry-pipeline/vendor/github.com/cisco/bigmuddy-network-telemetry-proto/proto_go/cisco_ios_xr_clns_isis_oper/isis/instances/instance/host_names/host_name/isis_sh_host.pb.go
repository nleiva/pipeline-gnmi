// Code generated by protoc-gen-go.
// source: isis_sh_host.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_clns_isis_oper_isis_instances_instance_host_names_host_name is a generated protocol buffer package.

It is generated from these files:
	isis_sh_host.proto

It has these top-level messages:
	IsisShHost_KEYS
	IsisShHost
*/
package cisco_ios_xr_clns_isis_oper_isis_instances_instance_host_names_host_name

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

// Bag representing a host (IS)
type IsisShHost_KEYS struct {
	InstanceName string `protobuf:"bytes,1,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
	SystemId     string `protobuf:"bytes,2,opt,name=system_id,json=systemId" json:"system_id,omitempty"`
}

func (m *IsisShHost_KEYS) Reset()                    { *m = IsisShHost_KEYS{} }
func (m *IsisShHost_KEYS) String() string            { return proto.CompactTextString(m) }
func (*IsisShHost_KEYS) ProtoMessage()               {}
func (*IsisShHost_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsisShHost_KEYS) GetInstanceName() string {
	if m != nil {
		return m.InstanceName
	}
	return ""
}

func (m *IsisShHost_KEYS) GetSystemId() string {
	if m != nil {
		return m.SystemId
	}
	return ""
}

type IsisShHost struct {
	// TRUE if this is the local IS
	LocalIsFlag bool `protobuf:"varint,50,opt,name=local_is_flag,json=localIsFlag" json:"local_is_flag,omitempty"`
	// Host levels
	HostLevels string `protobuf:"bytes,51,opt,name=host_levels,json=hostLevels" json:"host_levels,omitempty"`
	// Host name
	HostName string `protobuf:"bytes,52,opt,name=host_name,json=hostName" json:"host_name,omitempty"`
}

func (m *IsisShHost) Reset()                    { *m = IsisShHost{} }
func (m *IsisShHost) String() string            { return proto.CompactTextString(m) }
func (*IsisShHost) ProtoMessage()               {}
func (*IsisShHost) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsisShHost) GetLocalIsFlag() bool {
	if m != nil {
		return m.LocalIsFlag
	}
	return false
}

func (m *IsisShHost) GetHostLevels() string {
	if m != nil {
		return m.HostLevels
	}
	return ""
}

func (m *IsisShHost) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func init() {
	proto.RegisterType((*IsisShHost_KEYS)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.host_names.host_name.isis_sh_host_KEYS")
	proto.RegisterType((*IsisShHost)(nil), "cisco_ios_xr_clns_isis_oper.isis.instances.instance.host_names.host_name.isis_sh_host")
}

func init() { proto.RegisterFile("isis_sh_host.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xa9, 0x07, 0xd9, 0x9d, 0xdd, 0x3d, 0x98, 0x53, 0xc0, 0x83, 0x4b, 0xbd, 0xec, 0x29,
	0x07, 0xd7, 0xbf, 0xa0, 0x58, 0x14, 0x0f, 0x15, 0x0f, 0x9e, 0x86, 0x98, 0xc6, 0x36, 0x90, 0x26,
	0xa5, 0x13, 0x44, 0xff, 0xbd, 0x64, 0x4a, 0x6d, 0x6f, 0x8f, 0x8f, 0x37, 0xef, 0x3d, 0x06, 0x84,
	0x23, 0x47, 0x48, 0x1d, 0x76, 0x91, 0x92, 0x1a, 0xc6, 0x98, 0xa2, 0x78, 0x32, 0x8e, 0x4c, 0x44,
	0x17, 0x09, 0x7f, 0x46, 0x34, 0x3e, 0x10, 0xb2, 0x2b, 0x0e, 0x76, 0x54, 0x59, 0x29, 0x17, 0x28,
	0xe9, 0x60, 0xec, 0xa2, 0x54, 0xbe, 0xc7, 0xa0, 0x7b, 0x4b, 0x8b, 0x2c, 0xdf, 0xe1, 0x6a, 0x9d,
	0x8f, 0xcf, 0x0f, 0x1f, 0x6f, 0xe2, 0x16, 0x0e, 0xf3, 0x15, 0xbb, 0x64, 0x71, 0x2c, 0x4e, 0xdb,
	0x7a, 0x3f, 0xc3, 0x57, 0xdd, 0x5b, 0x71, 0x0d, 0x5b, 0xfa, 0xa5, 0x64, 0x7b, 0x74, 0x8d, 0xbc,
	0x60, 0xc3, 0x66, 0x02, 0x55, 0x53, 0x0e, 0xb0, 0x5f, 0xc7, 0x8a, 0x12, 0x0e, 0x3e, 0x1a, 0xed,
	0xd1, 0x11, 0x7e, 0x79, 0xdd, 0xca, 0xbb, 0x63, 0x71, 0xda, 0xd4, 0x3b, 0x86, 0x15, 0x3d, 0x7a,
	0xdd, 0x8a, 0x1b, 0xd8, 0xf1, 0x04, 0x6f, 0xbf, 0xad, 0x27, 0x79, 0xe6, 0x48, 0xc8, 0xe8, 0x85,
	0x49, 0x6e, 0xfc, 0x1f, 0x2e, 0xef, 0xa7, 0xc6, 0x0c, 0xf2, 0x9c, 0xcf, 0x4b, 0xfe, 0xcc, 0xf9,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xed, 0x68, 0xd1, 0x2f, 0x01, 0x00, 0x00,
}
