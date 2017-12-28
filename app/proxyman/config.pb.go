package proxyman

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_serial "v2ray.com/core/common/serial"
import v2ray_core_common_net "v2ray.com/core/common/net"
import v2ray_core_common_net1 "v2ray.com/core/common/net"
import v2ray_core_transport_internet "v2ray.com/core/transport/internet"
import v2ray_core_internet_domainsocket "v2ray.com/core/transport/internet/domainsocket"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type KnownProtocols int32

const (
	KnownProtocols_HTTP KnownProtocols = 0
	KnownProtocols_TLS  KnownProtocols = 1
)

var KnownProtocols_name = map[int32]string{
	0: "HTTP",
	1: "TLS",
}
var KnownProtocols_value = map[string]int32{
	"HTTP": 0,
	"TLS":  1,
}

func (x KnownProtocols) String() string {
	return proto.EnumName(KnownProtocols_name, int32(x))
}
func (KnownProtocols) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategy_Type int32

const (
	// Always allocate all connection handlers.
	AllocationStrategy_Always AllocationStrategy_Type = 0
	// Randomly allocate specific range of handlers.
	AllocationStrategy_Random AllocationStrategy_Type = 1
	// External. Not supported yet.
	AllocationStrategy_External AllocationStrategy_Type = 2
)

var AllocationStrategy_Type_name = map[int32]string{
	0: "Always",
	1: "Random",
	2: "External",
}
var AllocationStrategy_Type_value = map[string]int32{
	"Always":   0,
	"Random":   1,
	"External": 2,
}

func (x AllocationStrategy_Type) String() string {
	return proto.EnumName(AllocationStrategy_Type_name, int32(x))
}
func (AllocationStrategy_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type InboundConfig struct {
}

func (m *InboundConfig) Reset()                    { *m = InboundConfig{} }
func (m *InboundConfig) String() string            { return proto.CompactTextString(m) }
func (*InboundConfig) ProtoMessage()               {}
func (*InboundConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AllocationStrategy struct {
	Type AllocationStrategy_Type `protobuf:"varint,1,opt,name=type,enum=v2ray.core.app.proxyman.AllocationStrategy_Type" json:"type,omitempty"`
	// Number of handlers (ports) running in parallel.
	// Default value is 3 if unset.
	Concurrency *AllocationStrategy_AllocationStrategyConcurrency `protobuf:"bytes,2,opt,name=concurrency" json:"concurrency,omitempty"`
	// Number of minutes before a handler is regenerated.
	// Default value is 5 if unset.
	Refresh *AllocationStrategy_AllocationStrategyRefresh `protobuf:"bytes,3,opt,name=refresh" json:"refresh,omitempty"`
}

func (m *AllocationStrategy) Reset()                    { *m = AllocationStrategy{} }
func (m *AllocationStrategy) String() string            { return proto.CompactTextString(m) }
func (*AllocationStrategy) ProtoMessage()               {}
func (*AllocationStrategy) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllocationStrategy) GetType() AllocationStrategy_Type {
	if m != nil {
		return m.Type
	}
	return AllocationStrategy_Always
}

func (m *AllocationStrategy) GetConcurrency() *AllocationStrategy_AllocationStrategyConcurrency {
	if m != nil {
		return m.Concurrency
	}
	return nil
}

func (m *AllocationStrategy) GetRefresh() *AllocationStrategy_AllocationStrategyRefresh {
	if m != nil {
		return m.Refresh
	}
	return nil
}

type AllocationStrategy_AllocationStrategyConcurrency struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategy_AllocationStrategyConcurrency) Reset() {
	*m = AllocationStrategy_AllocationStrategyConcurrency{}
}
func (m *AllocationStrategy_AllocationStrategyConcurrency) String() string {
	return proto.CompactTextString(m)
}
func (*AllocationStrategy_AllocationStrategyConcurrency) ProtoMessage() {}
func (*AllocationStrategy_AllocationStrategyConcurrency) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *AllocationStrategy_AllocationStrategyConcurrency) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type AllocationStrategy_AllocationStrategyRefresh struct {
	Value uint32 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
}

func (m *AllocationStrategy_AllocationStrategyRefresh) Reset() {
	*m = AllocationStrategy_AllocationStrategyRefresh{}
}
func (m *AllocationStrategy_AllocationStrategyRefresh) String() string {
	return proto.CompactTextString(m)
}
func (*AllocationStrategy_AllocationStrategyRefresh) ProtoMessage() {}
func (*AllocationStrategy_AllocationStrategyRefresh) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 1}
}

func (m *AllocationStrategy_AllocationStrategyRefresh) GetValue() uint32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ReceiverConfig struct {
	PortRange                  *v2ray_core_common_net1.PortRange           `protobuf:"bytes,1,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
	Listen                     *v2ray_core_common_net.IPOrDomain           `protobuf:"bytes,2,opt,name=listen" json:"listen,omitempty"`
	AllocationStrategy         *AllocationStrategy                         `protobuf:"bytes,3,opt,name=allocation_strategy,json=allocationStrategy" json:"allocation_strategy,omitempty"`
	StreamSettings             *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,4,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	ReceiveOriginalDestination bool                                        `protobuf:"varint,5,opt,name=receive_original_destination,json=receiveOriginalDestination" json:"receive_original_destination,omitempty"`
	DomainOverride             []KnownProtocols                            `protobuf:"varint,7,rep,packed,name=domain_override,json=domainOverride,enum=v2ray.core.app.proxyman.KnownProtocols" json:"domain_override,omitempty"`
}

func (m *ReceiverConfig) Reset()                    { *m = ReceiverConfig{} }
func (m *ReceiverConfig) String() string            { return proto.CompactTextString(m) }
func (*ReceiverConfig) ProtoMessage()               {}
func (*ReceiverConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReceiverConfig) GetPortRange() *v2ray_core_common_net1.PortRange {
	if m != nil {
		return m.PortRange
	}
	return nil
}

func (m *ReceiverConfig) GetListen() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Listen
	}
	return nil
}

func (m *ReceiverConfig) GetAllocationStrategy() *AllocationStrategy {
	if m != nil {
		return m.AllocationStrategy
	}
	return nil
}

func (m *ReceiverConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

func (m *ReceiverConfig) GetReceiveOriginalDestination() bool {
	if m != nil {
		return m.ReceiveOriginalDestination
	}
	return false
}

func (m *ReceiverConfig) GetDomainOverride() []KnownProtocols {
	if m != nil {
		return m.DomainOverride
	}
	return nil
}

type UnixReceiverConfig struct {
	DomainSockSettings *v2ray_core_internet_domainsocket.DomainSocketSettings `protobuf:"bytes,2,opt,name=domainSockSettings" json:"domainSockSettings,omitempty"`
	StreamSettings     *v2ray_core_transport_internet.StreamConfig            `protobuf:"bytes,4,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	DomainOverride     []KnownProtocols                                       `protobuf:"varint,7,rep,packed,name=domain_override,json=domainOverride,enum=v2ray.core.app.proxyman.KnownProtocols" json:"domain_override,omitempty"`
}

func (m *UnixReceiverConfig) Reset()                    { *m = UnixReceiverConfig{} }
func (m *UnixReceiverConfig) String() string            { return proto.CompactTextString(m) }
func (*UnixReceiverConfig) ProtoMessage()               {}
func (*UnixReceiverConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UnixReceiverConfig) GetDomainSockSettings() *v2ray_core_internet_domainsocket.DomainSocketSettings {
	if m != nil {
		return m.DomainSockSettings
	}
	return nil
}

func (m *UnixReceiverConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

func (m *UnixReceiverConfig) GetDomainOverride() []KnownProtocols {
	if m != nil {
		return m.DomainOverride
	}
	return nil
}

type InboundHandlerConfig struct {
	Tag              string                                 `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	ReceiverSettings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=receiver_settings,json=receiverSettings" json:"receiver_settings,omitempty"`
	ProxySettings    *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,3,opt,name=proxy_settings,json=proxySettings" json:"proxy_settings,omitempty"`
}

func (m *InboundHandlerConfig) Reset()                    { *m = InboundHandlerConfig{} }
func (m *InboundHandlerConfig) String() string            { return proto.CompactTextString(m) }
func (*InboundHandlerConfig) ProtoMessage()               {}
func (*InboundHandlerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InboundHandlerConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *InboundHandlerConfig) GetReceiverSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.ReceiverSettings
	}
	return nil
}

func (m *InboundHandlerConfig) GetProxySettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.ProxySettings
	}
	return nil
}

type OutboundConfig struct {
}

func (m *OutboundConfig) Reset()                    { *m = OutboundConfig{} }
func (m *OutboundConfig) String() string            { return proto.CompactTextString(m) }
func (*OutboundConfig) ProtoMessage()               {}
func (*OutboundConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SenderConfig struct {
	// Send traffic through the given IP. Only IP is allowed.
	Via               *v2ray_core_common_net.IPOrDomain           `protobuf:"bytes,1,opt,name=via" json:"via,omitempty"`
	StreamSettings    *v2ray_core_transport_internet.StreamConfig `protobuf:"bytes,2,opt,name=stream_settings,json=streamSettings" json:"stream_settings,omitempty"`
	ProxySettings     *v2ray_core_transport_internet.ProxyConfig  `protobuf:"bytes,3,opt,name=proxy_settings,json=proxySettings" json:"proxy_settings,omitempty"`
	MultiplexSettings *MultiplexingConfig                         `protobuf:"bytes,4,opt,name=multiplex_settings,json=multiplexSettings" json:"multiplex_settings,omitempty"`
}

func (m *SenderConfig) Reset()                    { *m = SenderConfig{} }
func (m *SenderConfig) String() string            { return proto.CompactTextString(m) }
func (*SenderConfig) ProtoMessage()               {}
func (*SenderConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SenderConfig) GetVia() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Via
	}
	return nil
}

func (m *SenderConfig) GetStreamSettings() *v2ray_core_transport_internet.StreamConfig {
	if m != nil {
		return m.StreamSettings
	}
	return nil
}

func (m *SenderConfig) GetProxySettings() *v2ray_core_transport_internet.ProxyConfig {
	if m != nil {
		return m.ProxySettings
	}
	return nil
}

func (m *SenderConfig) GetMultiplexSettings() *MultiplexingConfig {
	if m != nil {
		return m.MultiplexSettings
	}
	return nil
}

type OutboundHandlerConfig struct {
	Tag            string                                 `protobuf:"bytes,1,opt,name=tag" json:"tag,omitempty"`
	SenderSettings *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,2,opt,name=sender_settings,json=senderSettings" json:"sender_settings,omitempty"`
	ProxySettings  *v2ray_core_common_serial.TypedMessage `protobuf:"bytes,3,opt,name=proxy_settings,json=proxySettings" json:"proxy_settings,omitempty"`
	Expire         int64                                  `protobuf:"varint,4,opt,name=expire" json:"expire,omitempty"`
	Comment        string                                 `protobuf:"bytes,5,opt,name=comment" json:"comment,omitempty"`
}

func (m *OutboundHandlerConfig) Reset()                    { *m = OutboundHandlerConfig{} }
func (m *OutboundHandlerConfig) String() string            { return proto.CompactTextString(m) }
func (*OutboundHandlerConfig) ProtoMessage()               {}
func (*OutboundHandlerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *OutboundHandlerConfig) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *OutboundHandlerConfig) GetSenderSettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.SenderSettings
	}
	return nil
}

func (m *OutboundHandlerConfig) GetProxySettings() *v2ray_core_common_serial.TypedMessage {
	if m != nil {
		return m.ProxySettings
	}
	return nil
}

func (m *OutboundHandlerConfig) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *OutboundHandlerConfig) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

type MultiplexingConfig struct {
	// Whether or not Mux is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
	// Max number of concurrent connections that one Mux connection can handle.
	Concurrency uint32 `protobuf:"varint,2,opt,name=concurrency" json:"concurrency,omitempty"`
}

func (m *MultiplexingConfig) Reset()                    { *m = MultiplexingConfig{} }
func (m *MultiplexingConfig) String() string            { return proto.CompactTextString(m) }
func (*MultiplexingConfig) ProtoMessage()               {}
func (*MultiplexingConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *MultiplexingConfig) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *MultiplexingConfig) GetConcurrency() uint32 {
	if m != nil {
		return m.Concurrency
	}
	return 0
}

func init() {
	proto.RegisterType((*InboundConfig)(nil), "v2ray.core.app.proxyman.InboundConfig")
	proto.RegisterType((*AllocationStrategy)(nil), "v2ray.core.app.proxyman.AllocationStrategy")
	proto.RegisterType((*AllocationStrategy_AllocationStrategyConcurrency)(nil), "v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyConcurrency")
	proto.RegisterType((*AllocationStrategy_AllocationStrategyRefresh)(nil), "v2ray.core.app.proxyman.AllocationStrategy.AllocationStrategyRefresh")
	proto.RegisterType((*ReceiverConfig)(nil), "v2ray.core.app.proxyman.ReceiverConfig")
	proto.RegisterType((*UnixReceiverConfig)(nil), "v2ray.core.app.proxyman.UnixReceiverConfig")
	proto.RegisterType((*InboundHandlerConfig)(nil), "v2ray.core.app.proxyman.InboundHandlerConfig")
	proto.RegisterType((*OutboundConfig)(nil), "v2ray.core.app.proxyman.OutboundConfig")
	proto.RegisterType((*SenderConfig)(nil), "v2ray.core.app.proxyman.SenderConfig")
	proto.RegisterType((*OutboundHandlerConfig)(nil), "v2ray.core.app.proxyman.OutboundHandlerConfig")
	proto.RegisterType((*MultiplexingConfig)(nil), "v2ray.core.app.proxyman.MultiplexingConfig")
	proto.RegisterEnum("v2ray.core.app.proxyman.KnownProtocols", KnownProtocols_name, KnownProtocols_value)
	proto.RegisterEnum("v2ray.core.app.proxyman.AllocationStrategy_Type", AllocationStrategy_Type_name, AllocationStrategy_Type_value)
}

func init() { proto.RegisterFile("v2ray.com/core/app/proxyman/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xd1, 0x6e, 0x23, 0x35,
	0x14, 0xdd, 0xc9, 0xa4, 0x69, 0x7a, 0xbb, 0x9d, 0xce, 0x9a, 0x85, 0x0d, 0x01, 0xa4, 0x10, 0x10,
	0x1b, 0x2d, 0x68, 0xb2, 0x64, 0x05, 0x12, 0xe2, 0x01, 0x4a, 0xbb, 0xd2, 0xb6, 0x50, 0x25, 0x38,
	0x81, 0x87, 0x15, 0x52, 0xe4, 0xce, 0xb8, 0xc1, 0xea, 0x8c, 0x3d, 0xb2, 0x9d, 0x6c, 0xf2, 0x4b,
	0x7c, 0x03, 0x0f, 0x3c, 0xf2, 0xc0, 0x17, 0xf0, 0x2b, 0xbc, 0xa0, 0x19, 0x7b, 0xa6, 0x69, 0x93,
	0x6c, 0x29, 0xab, 0xbe, 0xd9, 0xd3, 0x73, 0x8e, 0x7b, 0xce, 0xbd, 0xd7, 0x0e, 0x74, 0x66, 0x3d,
	0x49, 0x16, 0x41, 0x28, 0x92, 0x6e, 0x28, 0x24, 0xed, 0x92, 0x34, 0xed, 0xa6, 0x52, 0xcc, 0x17,
	0x09, 0xe1, 0xdd, 0x50, 0xf0, 0x73, 0x36, 0x09, 0x52, 0x29, 0xb4, 0x40, 0x8f, 0x0a, 0xa4, 0xa4,
	0x01, 0x49, 0xd3, 0xa0, 0x40, 0x35, 0x9f, 0x5e, 0x93, 0x08, 0x45, 0x92, 0x08, 0xde, 0x55, 0x54,
	0x32, 0x12, 0x77, 0xf5, 0x22, 0xa5, 0xd1, 0x38, 0xa1, 0x4a, 0x91, 0x09, 0x35, 0x52, 0xcd, 0xc7,
	0xeb, 0x19, 0x9c, 0xea, 0x2e, 0x89, 0x22, 0x49, 0x95, 0xb2, 0xc0, 0x8f, 0x37, 0x03, 0x53, 0x21,
	0xb5, 0x45, 0x05, 0xd7, 0x50, 0x5a, 0x12, 0xae, 0xb2, 0xbf, 0x77, 0x19, 0xd7, 0x54, 0x66, 0xe8,
	0x65, 0x27, 0xcd, 0xaf, 0x6f, 0xc6, 0x47, 0x22, 0x21, 0x8c, 0x2b, 0x11, 0x5e, 0x5c, 0x23, 0xb7,
	0xf7, 0x61, 0xef, 0x98, 0x9f, 0x89, 0x29, 0x8f, 0x0e, 0xf3, 0xcf, 0xed, 0x3f, 0x5c, 0x40, 0x07,
	0x71, 0x2c, 0x42, 0xa2, 0x99, 0xe0, 0x43, 0x2d, 0x89, 0xa6, 0x93, 0x05, 0x3a, 0x82, 0x6a, 0x66,
	0xbd, 0xe1, 0xb4, 0x9c, 0x8e, 0xd7, 0x7b, 0x1a, 0x6c, 0x48, 0x2f, 0x58, 0xa5, 0x06, 0xa3, 0x45,
	0x4a, 0x71, 0xce, 0x46, 0x17, 0xb0, 0x1b, 0x0a, 0x1e, 0x4e, 0xa5, 0xa4, 0x3c, 0x5c, 0x34, 0x2a,
	0x2d, 0xa7, 0xb3, 0xdb, 0x3b, 0xbe, 0x8d, 0xd8, 0xea, 0xa7, 0xc3, 0x4b, 0x41, 0xbc, 0xac, 0x8e,
	0xc6, 0xb0, 0x2d, 0xe9, 0xb9, 0xa4, 0xea, 0xd7, 0x86, 0x9b, 0x1f, 0xf4, 0xfc, 0xcd, 0x0e, 0xc2,
	0x46, 0x0c, 0x17, 0xaa, 0xcd, 0x2f, 0xe0, 0x83, 0xd7, 0xfe, 0x3b, 0xe8, 0x21, 0x6c, 0xcd, 0x48,
	0x3c, 0x35, 0xa9, 0xed, 0x61, 0xb3, 0x69, 0x7e, 0x0e, 0xef, 0x6e, 0x14, 0x5f, 0x4f, 0x69, 0x7f,
	0x06, 0xd5, 0x2c, 0x45, 0x04, 0x50, 0x3b, 0x88, 0x5f, 0x91, 0x85, 0xf2, 0xef, 0x65, 0x6b, 0x4c,
	0x78, 0x24, 0x12, 0xdf, 0x41, 0xf7, 0xa1, 0xfe, 0x7c, 0x9e, 0xd5, 0x9a, 0xc4, 0x7e, 0xa5, 0xfd,
	0xb7, 0x0b, 0x1e, 0xa6, 0x21, 0x65, 0x33, 0x2a, 0x4d, 0x55, 0xd1, 0x37, 0x00, 0x59, 0x47, 0x8c,
	0x25, 0xe1, 0x13, 0xa3, 0xbd, 0xdb, 0x6b, 0x2d, 0xc7, 0x61, 0x5a, 0x31, 0xe0, 0x54, 0x07, 0x03,
	0x21, 0x35, 0xce, 0x70, 0x78, 0x27, 0x2d, 0x96, 0xe8, 0x2b, 0xa8, 0xc5, 0x4c, 0x69, 0xca, 0x6d,
	0xd1, 0x3e, 0xdc, 0x40, 0x3e, 0x1e, 0xf4, 0xe5, 0x51, 0xde, 0x6d, 0xd8, 0x12, 0xd0, 0x2f, 0xf0,
	0x16, 0x29, 0xfd, 0x8e, 0x95, 0x35, 0x6c, 0x6b, 0xf2, 0xe9, 0x2d, 0x6a, 0x82, 0x11, 0x59, 0x6d,
	0xcc, 0x11, 0xec, 0x2b, 0x2d, 0x29, 0x49, 0xc6, 0x8a, 0x6a, 0xcd, 0xf8, 0x44, 0x35, 0xaa, 0xab,
	0xca, 0xe5, 0x4c, 0x04, 0xc5, 0x4c, 0x04, 0xc3, 0x9c, 0x65, 0xf2, 0xc1, 0x9e, 0xd1, 0x18, 0x5a,
	0x09, 0xf4, 0x2d, 0xbc, 0x2f, 0x4d, 0x82, 0x63, 0x21, 0xd9, 0x84, 0x71, 0x12, 0x8f, 0x23, 0xaa,
	0x34, 0xe3, 0xf9, 0xe9, 0x8d, 0xad, 0x96, 0xd3, 0xa9, 0xe3, 0xa6, 0xc5, 0xf4, 0x2d, 0xe4, 0xe8,
	0x12, 0x81, 0x06, 0xb0, 0x6f, 0xa6, 0x6e, 0x2c, 0x66, 0x54, 0x4a, 0x16, 0xd1, 0xc6, 0x76, 0xcb,
	0xed, 0x78, 0xbd, 0xc7, 0x1b, 0x1d, 0x7f, 0xcf, 0xc5, 0x2b, 0x3e, 0xc8, 0xc6, 0x32, 0x14, 0xb1,
	0xc2, 0x9e, 0xe1, 0xf7, 0x2d, 0xfd, 0xa4, 0x5a, 0xaf, 0xf9, 0xdb, 0xed, 0xdf, 0x2b, 0x80, 0x7e,
	0xe2, 0x6c, 0x7e, 0xad, 0xc0, 0xe7, 0x80, 0x0c, 0x7c, 0x28, 0xc2, 0x8b, 0xc2, 0x86, 0xad, 0xd5,
	0x97, 0xcb, 0x27, 0x96, 0xfe, 0x97, 0xef, 0x84, 0xe0, 0xa8, 0xe4, 0x52, 0x5d, 0xb0, 0xf1, 0x1a,
	0xc5, 0x3b, 0x8a, 0xfb, 0x2e, 0xc2, 0x72, 0xfc, 0xca, 0x49, 0xb5, 0xee, 0xfa, 0xd5, 0x93, 0x6a,
	0x7d, 0xcb, 0xaf, 0xd9, 0xf8, 0xfe, 0x72, 0xe0, 0xa1, 0xbd, 0xf0, 0x5e, 0x10, 0x1e, 0xc5, 0x65,
	0x80, 0x3e, 0xb8, 0x9a, 0x4c, 0xf2, 0xd1, 0xd8, 0xc1, 0xd9, 0x12, 0x0d, 0xe1, 0x81, 0xad, 0xaf,
	0xbc, 0x34, 0x6b, 0x12, 0xfd, 0x64, 0x4d, 0xf7, 0x9b, 0x07, 0x22, 0xbf, 0xed, 0xa2, 0x53, 0xf3,
	0x3e, 0x60, 0xbf, 0x10, 0x28, 0x9d, 0x9e, 0x82, 0x97, 0x5b, 0xb8, 0x54, 0x74, 0x6f, 0xa5, 0xb8,
	0x97, 0xb3, 0x0b, 0xb9, 0xb6, 0x0f, 0x5e, 0x7f, 0xaa, 0x97, 0xef, 0xef, 0x3f, 0x2b, 0x70, 0x7f,
	0x48, 0x79, 0x54, 0x1a, 0x7b, 0x06, 0xee, 0x8c, 0x11, 0x3b, 0xf3, 0xff, 0x61, 0x6c, 0x33, 0xf4,
	0xba, 0x32, 0x57, 0xde, 0xbc, 0xcc, 0x3f, 0x6e, 0x30, 0xff, 0xe4, 0x06, 0xd1, 0x41, 0x46, 0xb2,
	0x9a, 0x57, 0x03, 0x40, 0x2f, 0x01, 0x25, 0xd3, 0x58, 0xb3, 0x34, 0xa6, 0xf3, 0xd7, 0xb6, 0xe4,
	0x95, 0xe6, 0x39, 0x2d, 0x28, 0x8c, 0x4f, 0xac, 0xee, 0x83, 0x52, 0xa6, 0x0c, 0xf7, 0x1f, 0x07,
	0xde, 0x2e, 0xd2, 0xbd, 0xa9, 0x59, 0xfa, 0xb0, 0xaf, 0xf2, 0xd4, 0xff, 0x6f, 0xab, 0x78, 0x86,
	0x7e, 0x47, 0x8d, 0x82, 0xde, 0x81, 0x1a, 0x9d, 0xa7, 0x4c, 0xd2, 0x3c, 0x1b, 0x17, 0xdb, 0x1d,
	0x6a, 0xc0, 0x76, 0x26, 0x42, 0xb9, 0xce, 0xef, 0xb4, 0x1d, 0x5c, 0x6c, 0xdb, 0x03, 0x40, 0xab,
	0x31, 0x65, 0x78, 0xca, 0xc9, 0x59, 0x4c, 0xa3, 0xdc, 0x7d, 0x1d, 0x17, 0x5b, 0xd4, 0x5a, 0x7d,
	0xdb, 0xf7, 0xae, 0x3c, 0xc8, 0x4f, 0x3e, 0x02, 0xef, 0xea, 0xd4, 0xa2, 0x3a, 0x54, 0x5f, 0x8c,
	0x46, 0x03, 0xff, 0x1e, 0xda, 0x06, 0x77, 0xf4, 0xc3, 0xd0, 0x77, 0xbe, 0x3b, 0x84, 0xf7, 0x42,
	0x91, 0x6c, 0xaa, 0xdc, 0xc0, 0x79, 0x59, 0x2f, 0xd6, 0xbf, 0x55, 0x1e, 0xfd, 0xdc, 0xc3, 0x64,
	0x11, 0x1c, 0x66, 0xa8, 0x83, 0x34, 0x35, 0x7d, 0x92, 0x10, 0x7e, 0x56, 0xcb, 0x7f, 0xdc, 0x3c,
	0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x85, 0xa0, 0x69, 0xec, 0x0f, 0x0a, 0x00, 0x00,
}
