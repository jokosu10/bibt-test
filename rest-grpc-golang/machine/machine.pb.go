// Code generated by protoc-gen-go. DO NOT EDIT.
// source: machine/machine.proto

package machine

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Instruction struct {
	Operator             string   `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Operand              int32    `protobuf:"varint,2,opt,name=operand,proto3" json:"operand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instruction) Reset()         { *m = Instruction{} }
func (m *Instruction) String() string { return proto.CompactTextString(m) }
func (*Instruction) ProtoMessage()    {}
func (*Instruction) Descriptor() ([]byte, []int) {
	return fileDescriptor_84b4f59d98cc997c, []int{0}
}

func (m *Instruction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instruction.Unmarshal(m, b)
}
func (m *Instruction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instruction.Marshal(b, m, deterministic)
}
func (m *Instruction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instruction.Merge(m, src)
}
func (m *Instruction) XXX_Size() int {
	return xxx_messageInfo_Instruction.Size(m)
}
func (m *Instruction) XXX_DiscardUnknown() {
	xxx_messageInfo_Instruction.DiscardUnknown(m)
}

var xxx_messageInfo_Instruction proto.InternalMessageInfo

func (m *Instruction) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *Instruction) GetOperand() int32 {
	if m != nil {
		return m.Operand
	}
	return 0
}

type InstructionSet struct {
	Instructions         []*Instruction `protobuf:"bytes,1,rep,name=instructions,proto3" json:"instructions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *InstructionSet) Reset()         { *m = InstructionSet{} }
func (m *InstructionSet) String() string { return proto.CompactTextString(m) }
func (*InstructionSet) ProtoMessage()    {}
func (*InstructionSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_84b4f59d98cc997c, []int{1}
}

func (m *InstructionSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstructionSet.Unmarshal(m, b)
}
func (m *InstructionSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstructionSet.Marshal(b, m, deterministic)
}
func (m *InstructionSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstructionSet.Merge(m, src)
}
func (m *InstructionSet) XXX_Size() int {
	return xxx_messageInfo_InstructionSet.Size(m)
}
func (m *InstructionSet) XXX_DiscardUnknown() {
	xxx_messageInfo_InstructionSet.DiscardUnknown(m)
}

var xxx_messageInfo_InstructionSet proto.InternalMessageInfo

func (m *InstructionSet) GetInstructions() []*Instruction {
	if m != nil {
		return m.Instructions
	}
	return nil
}

type Result struct {
	Output               float32  `protobuf:"fixed32,1,opt,name=output,proto3" json:"output,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_84b4f59d98cc997c, []int{2}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetOutput() float32 {
	if m != nil {
		return m.Output
	}
	return 0
}

func init() {
	proto.RegisterType((*Instruction)(nil), "machine.Instruction")
	proto.RegisterType((*InstructionSet)(nil), "machine.InstructionSet")
	proto.RegisterType((*Result)(nil), "machine.Result")
}

func init() { proto.RegisterFile("machine/machine.proto", fileDescriptor_84b4f59d98cc997c) }

var fileDescriptor_84b4f59d98cc997c = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x4d, 0x4c, 0xce,
	0xc8, 0xcc, 0x4b, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x33, 0x17, 0xb7, 0x67, 0x5e, 0x71, 0x49, 0x51, 0x69, 0x72, 0x49, 0x66, 0x7e, 0x9e, 0x90,
	0x14, 0x17, 0x47, 0x7e, 0x41, 0x6a, 0x51, 0x62, 0x49, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x9c, 0x2f, 0x24, 0xc1, 0xc5, 0x0e, 0x66, 0xe7, 0xa5, 0x48, 0x30, 0x29, 0x30, 0x6a,
	0xb0, 0x06, 0xc1, 0xb8, 0x4a, 0x5e, 0x5c, 0x7c, 0x48, 0x86, 0x04, 0xa7, 0x96, 0x08, 0x59, 0x70,
	0xf1, 0x64, 0x22, 0x44, 0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x60, 0xae,
	0x40, 0x52, 0x1e, 0x84, 0xa2, 0x52, 0x49, 0x81, 0x8b, 0x2d, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x44,
	0x48, 0x8c, 0x8b, 0x2d, 0xbf, 0xb4, 0xa4, 0xa0, 0xb4, 0x04, 0xec, 0x12, 0xa6, 0x20, 0x28, 0xcf,
	0xc8, 0x81, 0x8b, 0xdd, 0x17, 0x62, 0x8c, 0x90, 0x29, 0x17, 0xbb, 0x6b, 0x45, 0x6a, 0x72, 0x69,
	0x49, 0xaa, 0x90, 0x38, 0x36, 0xb3, 0x83, 0x53, 0x4b, 0xa4, 0xf8, 0xe1, 0x12, 0x10, 0x73, 0x95,
	0x18, 0x92, 0xd8, 0xc0, 0x81, 0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x19, 0x7f, 0x99, 0xca,
	0x1d, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MachineClient is the client API for Machine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MachineClient interface {
	Execute(ctx context.Context, in *InstructionSet, opts ...grpc.CallOption) (*Result, error)
}

type machineClient struct {
	cc *grpc.ClientConn
}

func NewMachineClient(cc *grpc.ClientConn) MachineClient {
	return &machineClient{cc}
}

func (c *machineClient) Execute(ctx context.Context, in *InstructionSet, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/machine.Machine/Execute", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServer is the server API for Machine service.
type MachineServer interface {
	Execute(context.Context, *InstructionSet) (*Result, error)
}

func RegisterMachineServer(s *grpc.Server, srv MachineServer) {
	s.RegisterService(&_Machine_serviceDesc, srv)
}

func _Machine_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstructionSet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/machine.Machine/Execute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServer).Execute(ctx, req.(*InstructionSet))
	}
	return interceptor(ctx, in, info, handler)
}

var _Machine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "machine.Machine",
	HandlerType: (*MachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Execute",
			Handler:    _Machine_Execute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "machine/machine.proto",
}
