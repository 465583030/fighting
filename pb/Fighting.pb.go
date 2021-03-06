// Code generated by protoc-gen-go.
// source: Fighting.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UserInput struct {
	Data int32 `protobuf:"varint,1,opt,name=data" json:"data,omitempty"`
}

func (m *UserInput) Reset()                    { *m = UserInput{} }
func (m *UserInput) String() string            { return proto.CompactTextString(m) }
func (*UserInput) ProtoMessage()               {}
func (*UserInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type UserInputData struct {
	ID   int32 `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Data int32 `protobuf:"varint,2,opt,name=data" json:"data,omitempty"`
}

func (m *UserInputData) Reset()                    { *m = UserInputData{} }
func (m *UserInputData) String() string            { return proto.CompactTextString(m) }
func (*UserInputData) ProtoMessage()               {}
func (*UserInputData) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type CollectUsersInput struct {
	Step           int32            `protobuf:"varint,1,opt,name=step" json:"step,omitempty"`
	UsersInputData []*UserInputData `protobuf:"bytes,2,rep,name=usersInputData" json:"usersInputData,omitempty"`
}

func (m *CollectUsersInput) Reset()                    { *m = CollectUsersInput{} }
func (m *CollectUsersInput) String() string            { return proto.CompactTextString(m) }
func (*CollectUsersInput) ProtoMessage()               {}
func (*CollectUsersInput) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *CollectUsersInput) GetUsersInputData() []*UserInputData {
	if m != nil {
		return m.UsersInputData
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInput)(nil), "pb.UserInput")
	proto.RegisterType((*UserInputData)(nil), "pb.UserInputData")
	proto.RegisterType((*CollectUsersInput)(nil), "pb.CollectUsersInput")
}

func init() { proto.RegisterFile("Fighting.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x73, 0xcb, 0x4c, 0xcf,
	0x28, 0xc9, 0xcc, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92,
	0xe7, 0xe2, 0x0c, 0x2d, 0x4e, 0x2d, 0xf2, 0xcc, 0x2b, 0x28, 0x2d, 0x11, 0x12, 0xe2, 0x62, 0x49,
	0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0xb3, 0x95, 0x8c, 0xb9, 0x78,
	0xe1, 0x0a, 0x5c, 0x80, 0x02, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x2e, 0x50, 0x25, 0x4c, 0x99, 0x2e,
	0x70, 0x4d, 0x4c, 0x48, 0x9a, 0x92, 0xb8, 0x04, 0x9d, 0xf3, 0x73, 0x72, 0x52, 0x93, 0x4b, 0x40,
	0x7a, 0x8b, 0xe1, 0xa6, 0x17, 0x97, 0xa4, 0x16, 0xc0, 0x4c, 0x07, 0xb1, 0x85, 0x2c, 0xb9, 0xf8,
	0x4a, 0xe1, 0x2a, 0x5c, 0x20, 0xc6, 0x30, 0x6b, 0x70, 0x1b, 0x09, 0xea, 0x15, 0x24, 0xe9, 0xa1,
	0xd8, 0x1b, 0x84, 0xa6, 0xd0, 0x89, 0x63, 0x15, 0x13, 0x6b, 0x00, 0xc8, 0x1f, 0x49, 0x6c, 0x60,
	0xef, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x0c, 0xa7, 0x1e, 0xe0, 0x00, 0x00, 0x00,
}
