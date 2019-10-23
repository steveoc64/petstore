// Code generated by protoc-gen-go. DO NOT EDIT.
// source: petstore.proto

package petstore

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_available Status = 0
	Status_pending   Status = 1
	Status_sold      Status = 2
)

var Status_name = map[int32]string{
	0: "available",
	1: "pending",
	2: "sold",
}

var Status_value = map[string]int32{
	"available": 0,
	"pending":   1,
	"sold":      2,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{0}
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Category struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Category) Reset()         { *m = Category{} }
func (m *Category) String() string { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()    {}
func (*Category) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{1}
}

func (m *Category) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Category.Unmarshal(m, b)
}
func (m *Category) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Category.Marshal(b, m, deterministic)
}
func (m *Category) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Category.Merge(m, src)
}
func (m *Category) XXX_Size() int {
	return xxx_messageInfo_Category.Size(m)
}
func (m *Category) XXX_DiscardUnknown() {
	xxx_messageInfo_Category.DiscardUnknown(m)
}

var xxx_messageInfo_Category proto.InternalMessageInfo

func (m *Category) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Tag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{2}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PetID struct {
	PetId                int64    `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetID) Reset()         { *m = PetID{} }
func (m *PetID) String() string { return proto.CompactTextString(m) }
func (*PetID) ProtoMessage()    {}
func (*PetID) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{3}
}

func (m *PetID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetID.Unmarshal(m, b)
}
func (m *PetID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetID.Marshal(b, m, deterministic)
}
func (m *PetID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetID.Merge(m, src)
}
func (m *PetID) XXX_Size() int {
	return xxx_messageInfo_PetID.Size(m)
}
func (m *PetID) XXX_DiscardUnknown() {
	xxx_messageInfo_PetID.DiscardUnknown(m)
}

var xxx_messageInfo_PetID proto.InternalMessageInfo

func (m *PetID) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

type Pet struct {
	PetId                int64     `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Category             *Category `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name                 string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	PhotoUrls            []string  `protobuf:"bytes,4,rep,name=photo_urls,json=photoUrls,proto3" json:"photo_urls,omitempty"`
	Tags                 []*Tag    `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	Status               string    `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}
func (*Pet) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{4}
}

func (m *Pet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pet.Unmarshal(m, b)
}
func (m *Pet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pet.Marshal(b, m, deterministic)
}
func (m *Pet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pet.Merge(m, src)
}
func (m *Pet) XXX_Size() int {
	return xxx_messageInfo_Pet.Size(m)
}
func (m *Pet) XXX_DiscardUnknown() {
	xxx_messageInfo_Pet.DiscardUnknown(m)
}

var xxx_messageInfo_Pet proto.InternalMessageInfo

func (m *Pet) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

func (m *Pet) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *Pet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Pet) GetPhotoUrls() []string {
	if m != nil {
		return m.PhotoUrls
	}
	return nil
}

func (m *Pet) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Pet) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type UpdatePetWithFormReq struct {
	PetId                int64    `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string   `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePetWithFormReq) Reset()         { *m = UpdatePetWithFormReq{} }
func (m *UpdatePetWithFormReq) String() string { return proto.CompactTextString(m) }
func (*UpdatePetWithFormReq) ProtoMessage()    {}
func (*UpdatePetWithFormReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{5}
}

func (m *UpdatePetWithFormReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePetWithFormReq.Unmarshal(m, b)
}
func (m *UpdatePetWithFormReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePetWithFormReq.Marshal(b, m, deterministic)
}
func (m *UpdatePetWithFormReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePetWithFormReq.Merge(m, src)
}
func (m *UpdatePetWithFormReq) XXX_Size() int {
	return xxx_messageInfo_UpdatePetWithFormReq.Size(m)
}
func (m *UpdatePetWithFormReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePetWithFormReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePetWithFormReq proto.InternalMessageInfo

func (m *UpdatePetWithFormReq) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

func (m *UpdatePetWithFormReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePetWithFormReq) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type DeletePetReq struct {
	// Note - api_key is passed in via a header, not a param
	PetId                int64    `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePetReq) Reset()         { *m = DeletePetReq{} }
func (m *DeletePetReq) String() string { return proto.CompactTextString(m) }
func (*DeletePetReq) ProtoMessage()    {}
func (*DeletePetReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{6}
}

func (m *DeletePetReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePetReq.Unmarshal(m, b)
}
func (m *DeletePetReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePetReq.Marshal(b, m, deterministic)
}
func (m *DeletePetReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePetReq.Merge(m, src)
}
func (m *DeletePetReq) XXX_Size() int {
	return xxx_messageInfo_DeletePetReq.Size(m)
}
func (m *DeletePetReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePetReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePetReq proto.InternalMessageInfo

func (m *DeletePetReq) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

type UploadFileReq struct {
	PetId                int64    `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	AdditionalMetadata   string   `protobuf:"bytes,2,opt,name=additional_metadata,json=additionalMetadata,proto3" json:"additional_metadata,omitempty"`
	File                 string   `protobuf:"bytes,3,opt,name=file,proto3" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileReq) Reset()         { *m = UploadFileReq{} }
func (m *UploadFileReq) String() string { return proto.CompactTextString(m) }
func (*UploadFileReq) ProtoMessage()    {}
func (*UploadFileReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{7}
}

func (m *UploadFileReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileReq.Unmarshal(m, b)
}
func (m *UploadFileReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileReq.Marshal(b, m, deterministic)
}
func (m *UploadFileReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileReq.Merge(m, src)
}
func (m *UploadFileReq) XXX_Size() int {
	return xxx_messageInfo_UploadFileReq.Size(m)
}
func (m *UploadFileReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileReq.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileReq proto.InternalMessageInfo

func (m *UploadFileReq) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

func (m *UploadFileReq) GetAdditionalMetadata() string {
	if m != nil {
		return m.AdditionalMetadata
	}
	return ""
}

func (m *UploadFileReq) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

type PetBody struct {
	PetId                int64    `protobuf:"varint,1,opt,name=pet_id,json=petId,proto3" json:"pet_id,omitempty"`
	Body                 string   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PetBody) Reset()         { *m = PetBody{} }
func (m *PetBody) String() string { return proto.CompactTextString(m) }
func (*PetBody) ProtoMessage()    {}
func (*PetBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{8}
}

func (m *PetBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PetBody.Unmarshal(m, b)
}
func (m *PetBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PetBody.Marshal(b, m, deterministic)
}
func (m *PetBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PetBody.Merge(m, src)
}
func (m *PetBody) XXX_Size() int {
	return xxx_messageInfo_PetBody.Size(m)
}
func (m *PetBody) XXX_DiscardUnknown() {
	xxx_messageInfo_PetBody.DiscardUnknown(m)
}

var xxx_messageInfo_PetBody proto.InternalMessageInfo

func (m *PetBody) GetPetId() int64 {
	if m != nil {
		return m.PetId
	}
	return 0
}

func (m *PetBody) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type StatusReq struct {
	Status               []string `protobuf:"bytes,1,rep,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusReq) Reset()         { *m = StatusReq{} }
func (m *StatusReq) String() string { return proto.CompactTextString(m) }
func (*StatusReq) ProtoMessage()    {}
func (*StatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{9}
}

func (m *StatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusReq.Unmarshal(m, b)
}
func (m *StatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusReq.Marshal(b, m, deterministic)
}
func (m *StatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusReq.Merge(m, src)
}
func (m *StatusReq) XXX_Size() int {
	return xxx_messageInfo_StatusReq.Size(m)
}
func (m *StatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_StatusReq proto.InternalMessageInfo

func (m *StatusReq) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type Pets struct {
	Pets                 []*Pet   `protobuf:"bytes,1,rep,name=pets,proto3" json:"pets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pets) Reset()         { *m = Pets{} }
func (m *Pets) String() string { return proto.CompactTextString(m) }
func (*Pets) ProtoMessage()    {}
func (*Pets) Descriptor() ([]byte, []int) {
	return fileDescriptor_749e5a3d28fcc1b1, []int{10}
}

func (m *Pets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pets.Unmarshal(m, b)
}
func (m *Pets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pets.Marshal(b, m, deterministic)
}
func (m *Pets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pets.Merge(m, src)
}
func (m *Pets) XXX_Size() int {
	return xxx_messageInfo_Pets.Size(m)
}
func (m *Pets) XXX_DiscardUnknown() {
	xxx_messageInfo_Pets.DiscardUnknown(m)
}

var xxx_messageInfo_Pets proto.InternalMessageInfo

func (m *Pets) GetPets() []*Pet {
	if m != nil {
		return m.Pets
	}
	return nil
}

func init() {
	proto.RegisterEnum("petstore.Status", Status_name, Status_value)
	proto.RegisterType((*Empty)(nil), "petstore.Empty")
	proto.RegisterType((*Category)(nil), "petstore.Category")
	proto.RegisterType((*Tag)(nil), "petstore.Tag")
	proto.RegisterType((*PetID)(nil), "petstore.PetID")
	proto.RegisterType((*Pet)(nil), "petstore.Pet")
	proto.RegisterType((*UpdatePetWithFormReq)(nil), "petstore.UpdatePetWithFormReq")
	proto.RegisterType((*DeletePetReq)(nil), "petstore.DeletePetReq")
	proto.RegisterType((*UploadFileReq)(nil), "petstore.UploadFileReq")
	proto.RegisterType((*PetBody)(nil), "petstore.PetBody")
	proto.RegisterType((*StatusReq)(nil), "petstore.StatusReq")
	proto.RegisterType((*Pets)(nil), "petstore.Pets")
}

func init() { proto.RegisterFile("petstore.proto", fileDescriptor_749e5a3d28fcc1b1) }

var fileDescriptor_749e5a3d28fcc1b1 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xd1, 0x6e, 0xd3, 0x40,
	0x10, 0x8c, 0xe3, 0xc4, 0x8d, 0xb7, 0xb4, 0x0d, 0x57, 0x28, 0x51, 0x04, 0x55, 0x30, 0x42, 0x4a,
	0x91, 0x70, 0xa5, 0x82, 0x10, 0x3c, 0x52, 0x42, 0x50, 0x1e, 0x90, 0x2c, 0xb7, 0x11, 0xe2, 0x29,
	0xba, 0xf4, 0x16, 0xf7, 0x84, 0xe3, 0x33, 0xf6, 0xb6, 0x92, 0xbf, 0x82, 0xff, 0xe1, 0xeb, 0x90,
	0x2f, 0x8e, 0x1d, 0xd3, 0xa4, 0xe2, 0xed, 0xce, 0xb3, 0x3b, 0x3b, 0xb3, 0x63, 0x1b, 0xf6, 0x63,
	0xa4, 0x94, 0x54, 0x82, 0x6e, 0x9c, 0x28, 0x52, 0xac, 0xb3, 0xba, 0xf7, 0x9f, 0x06, 0x4a, 0x05,
	0x21, 0x9e, 0xf2, 0x58, 0x9e, 0xf2, 0x28, 0x52, 0xc4, 0x49, 0xaa, 0x28, 0x5d, 0xd6, 0x39, 0x3b,
	0xd0, 0xfe, 0xbc, 0x88, 0x29, 0x73, 0x5c, 0xe8, 0x7c, 0xe2, 0x84, 0x81, 0x4a, 0x32, 0xb6, 0x0f,
	0x4d, 0x29, 0x7a, 0xc6, 0xc0, 0x18, 0x9a, 0x7e, 0x53, 0x0a, 0xc6, 0xa0, 0x15, 0xf1, 0x05, 0xf6,
	0x9a, 0x03, 0x63, 0x68, 0xfb, 0xfa, 0xec, 0x9c, 0x80, 0x79, 0xc9, 0x83, 0xff, 0x2a, 0x3d, 0x86,
	0xb6, 0x87, 0x34, 0x19, 0xb1, 0xc7, 0x60, 0xc5, 0x48, 0xb3, 0xb2, 0xa1, 0x1d, 0x23, 0x4d, 0x84,
	0xf3, 0xc7, 0x00, 0xd3, 0x43, 0xda, 0x02, 0x33, 0x17, 0x3a, 0x57, 0x85, 0x32, 0x4d, 0xbb, 0x7b,
	0xc6, 0xdc, 0xd2, 0xed, 0x4a, 0xb3, 0x5f, 0xd6, 0x94, 0x12, 0xcc, 0x4a, 0x02, 0x7b, 0x06, 0x10,
	0x5f, 0x2b, 0x52, 0xb3, 0x9b, 0x24, 0x4c, 0x7b, 0xad, 0x81, 0x39, 0xb4, 0x7d, 0x5b, 0x3f, 0x99,
	0x26, 0x61, 0xca, 0x9e, 0x43, 0x8b, 0x78, 0x90, 0xf6, 0xda, 0x03, 0x73, 0xb8, 0x7b, 0xb6, 0x57,
	0xd1, 0x5f, 0xf2, 0xc0, 0xd7, 0x10, 0x3b, 0x02, 0x2b, 0x25, 0x4e, 0x37, 0x69, 0xcf, 0xd2, 0xbc,
	0xc5, 0xcd, 0xf9, 0x0e, 0x8f, 0xa6, 0xb1, 0xe0, 0x84, 0x1e, 0xd2, 0x37, 0x49, 0xd7, 0x63, 0x95,
	0x2c, 0x7c, 0xfc, 0xb5, 0xcd, 0xcc, 0x86, 0xfd, 0xac, 0x51, 0x9b, 0x35, 0xea, 0x97, 0xf0, 0x60,
	0x84, 0x21, 0x6a, 0xea, 0xed, 0x94, 0xce, 0x4f, 0xd8, 0x9b, 0xc6, 0xa1, 0xe2, 0x62, 0x2c, 0x43,
	0xbc, 0x67, 0xf4, 0x29, 0x1c, 0x72, 0x21, 0x64, 0x9e, 0x3e, 0x0f, 0x67, 0x0b, 0x24, 0x2e, 0x38,
	0xf1, 0x42, 0x09, 0xab, 0xa0, 0xaf, 0x05, 0x92, 0x6b, 0xfd, 0x21, 0xc3, 0x72, 0x91, 0xf9, 0xd9,
	0x79, 0x0b, 0x3b, 0x1e, 0xd2, 0xb9, 0x12, 0xd9, 0x3d, 0x0e, 0xe7, 0x4a, 0x64, 0x2b, 0x87, 0xf9,
	0xd9, 0x79, 0x01, 0xf6, 0x85, 0xf6, 0x94, 0xcb, 0xab, 0xec, 0x1a, 0x3a, 0x87, 0x95, 0xdd, 0x13,
	0x68, 0x79, 0x48, 0x3a, 0x8c, 0x7c, 0xff, 0x1a, 0xad, 0x85, 0x91, 0xaf, 0x41, 0x43, 0xaf, 0x5c,
	0xb0, 0x96, 0x7c, 0x6c, 0x0f, 0x6c, 0x7e, 0xcb, 0x65, 0xc8, 0xe7, 0x21, 0x76, 0x1b, 0x6c, 0x17,
	0x76, 0x62, 0x8c, 0x84, 0x8c, 0x82, 0xae, 0xc1, 0x3a, 0xd0, 0x4a, 0x55, 0x28, 0xba, 0xcd, 0xb3,
	0xdf, 0x26, 0x1c, 0x78, 0x05, 0xcd, 0x05, 0x26, 0xb7, 0xf2, 0x0a, 0x99, 0x0b, 0xf0, 0x05, 0x29,
	0x37, 0x93, 0x4d, 0x46, 0xec, 0xa0, 0x36, 0x66, 0x32, 0xea, 0xd7, 0xe7, 0x3a, 0x0d, 0x36, 0x86,
	0x87, 0x77, 0x82, 0x66, 0xc7, 0x55, 0xd5, 0xa6, 0xb7, 0xa0, 0xbf, 0x46, 0xbb, 0xfc, 0xcc, 0x1a,
	0xec, 0x1d, 0xd8, 0x65, 0xaa, 0xec, 0xa8, 0xc2, 0xd7, 0xa3, 0xde, 0xd4, 0xf7, 0x1e, 0xa0, 0x8a,
	0x99, 0x3d, 0x59, 0x1f, 0xbc, 0x16, 0xfe, 0xa6, 0xce, 0x21, 0x58, 0x1f, 0x85, 0xc8, 0xc7, 0xd5,
	0x4d, 0xdd, 0xf5, 0xf8, 0x1a, 0xec, 0xd2, 0xc6, 0xbf, 0xc5, 0x1b, 0x88, 0x3f, 0x40, 0x77, 0x2c,
	0xa3, 0x9c, 0x39, 0x3d, 0xcf, 0x8a, 0x40, 0x0e, 0xab, 0xb2, 0x32, 0xf2, 0xfe, 0x7e, 0x8d, 0x2a,
	0x75, 0x1a, 0x73, 0x4b, 0xff, 0x7e, 0xde, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xff, 0x60,
	0x83, 0xb8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PetstoreServiceClient is the client API for PetstoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PetstoreServiceClient interface {
	GetPetByID(ctx context.Context, in *PetID, opts ...grpc.CallOption) (*Pet, error)
	UpdatePetWithForm(ctx context.Context, in *UpdatePetWithFormReq, opts ...grpc.CallOption) (*Empty, error)
	DeletePet(ctx context.Context, in *DeletePetReq, opts ...grpc.CallOption) (*Empty, error)
	UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*Empty, error)
	AddPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error)
	UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error)
	FindPetsByStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*Pets, error)
}

type petstoreServiceClient struct {
	cc *grpc.ClientConn
}

func NewPetstoreServiceClient(cc *grpc.ClientConn) PetstoreServiceClient {
	return &petstoreServiceClient{cc}
}

func (c *petstoreServiceClient) GetPetByID(ctx context.Context, in *PetID, opts ...grpc.CallOption) (*Pet, error) {
	out := new(Pet)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/GetPetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) UpdatePetWithForm(ctx context.Context, in *UpdatePetWithFormReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/UpdatePetWithForm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) DeletePet(ctx context.Context, in *DeletePetReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/DeletePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/UploadFile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) AddPet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Pet, error) {
	out := new(Pet)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/AddPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/UpdatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petstoreServiceClient) FindPetsByStatus(ctx context.Context, in *StatusReq, opts ...grpc.CallOption) (*Pets, error) {
	out := new(Pets)
	err := c.cc.Invoke(ctx, "/petstore.PetstoreService/FindPetsByStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetstoreServiceServer is the server API for PetstoreService service.
type PetstoreServiceServer interface {
	GetPetByID(context.Context, *PetID) (*Pet, error)
	UpdatePetWithForm(context.Context, *UpdatePetWithFormReq) (*Empty, error)
	DeletePet(context.Context, *DeletePetReq) (*Empty, error)
	UploadFile(context.Context, *UploadFileReq) (*Empty, error)
	AddPet(context.Context, *Pet) (*Pet, error)
	UpdatePet(context.Context, *Pet) (*Empty, error)
	FindPetsByStatus(context.Context, *StatusReq) (*Pets, error)
}

// UnimplementedPetstoreServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPetstoreServiceServer struct {
}

func (*UnimplementedPetstoreServiceServer) GetPetByID(ctx context.Context, req *PetID) (*Pet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPetByID not implemented")
}
func (*UnimplementedPetstoreServiceServer) UpdatePetWithForm(ctx context.Context, req *UpdatePetWithFormReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePetWithForm not implemented")
}
func (*UnimplementedPetstoreServiceServer) DeletePet(ctx context.Context, req *DeletePetReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}
func (*UnimplementedPetstoreServiceServer) UploadFile(ctx context.Context, req *UploadFileReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (*UnimplementedPetstoreServiceServer) AddPet(ctx context.Context, req *Pet) (*Pet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPet not implemented")
}
func (*UnimplementedPetstoreServiceServer) UpdatePet(ctx context.Context, req *Pet) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePet not implemented")
}
func (*UnimplementedPetstoreServiceServer) FindPetsByStatus(ctx context.Context, req *StatusReq) (*Pets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPetsByStatus not implemented")
}

func RegisterPetstoreServiceServer(s *grpc.Server, srv PetstoreServiceServer) {
	s.RegisterService(&_PetstoreService_serviceDesc, srv)
}

func _PetstoreService_GetPetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).GetPetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/GetPetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).GetPetByID(ctx, req.(*PetID))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_UpdatePetWithForm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePetWithFormReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).UpdatePetWithForm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/UpdatePetWithForm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).UpdatePetWithForm(ctx, req.(*UpdatePetWithFormReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/DeletePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).DeletePet(ctx, req.(*DeletePetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/UploadFile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).UploadFile(ctx, req.(*UploadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_AddPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).AddPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/AddPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).AddPet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_UpdatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).UpdatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/UpdatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).UpdatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetstoreService_FindPetsByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetstoreServiceServer).FindPetsByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.PetstoreService/FindPetsByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetstoreServiceServer).FindPetsByStatus(ctx, req.(*StatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _PetstoreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "petstore.PetstoreService",
	HandlerType: (*PetstoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPetByID",
			Handler:    _PetstoreService_GetPetByID_Handler,
		},
		{
			MethodName: "UpdatePetWithForm",
			Handler:    _PetstoreService_UpdatePetWithForm_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetstoreService_DeletePet_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _PetstoreService_UploadFile_Handler,
		},
		{
			MethodName: "AddPet",
			Handler:    _PetstoreService_AddPet_Handler,
		},
		{
			MethodName: "UpdatePet",
			Handler:    _PetstoreService_UpdatePet_Handler,
		},
		{
			MethodName: "FindPetsByStatus",
			Handler:    _PetstoreService_FindPetsByStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petstore.proto",
}
