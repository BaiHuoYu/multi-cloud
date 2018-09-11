// Code generated by protoc-gen-go. DO NOT EDIT.
// source: backend.proto

package backend

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

type CreateBackendRequest struct {
	Backend              *BackendDetail `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateBackendRequest) Reset()         { *m = CreateBackendRequest{} }
func (m *CreateBackendRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBackendRequest) ProtoMessage()    {}
func (*CreateBackendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{0}
}
func (m *CreateBackendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBackendRequest.Unmarshal(m, b)
}
func (m *CreateBackendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBackendRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBackendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBackendRequest.Merge(dst, src)
}
func (m *CreateBackendRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBackendRequest.Size(m)
}
func (m *CreateBackendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBackendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBackendRequest proto.InternalMessageInfo

func (m *CreateBackendRequest) GetBackend() *BackendDetail {
	if m != nil {
		return m.Backend
	}
	return nil
}

type CreateBackendResponse struct {
	Backend              *BackendDetail `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateBackendResponse) Reset()         { *m = CreateBackendResponse{} }
func (m *CreateBackendResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBackendResponse) ProtoMessage()    {}
func (*CreateBackendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{1}
}
func (m *CreateBackendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBackendResponse.Unmarshal(m, b)
}
func (m *CreateBackendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBackendResponse.Marshal(b, m, deterministic)
}
func (dst *CreateBackendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBackendResponse.Merge(dst, src)
}
func (m *CreateBackendResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBackendResponse.Size(m)
}
func (m *CreateBackendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBackendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBackendResponse proto.InternalMessageInfo

func (m *CreateBackendResponse) GetBackend() *BackendDetail {
	if m != nil {
		return m.Backend
	}
	return nil
}

type GetBackendRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBackendRequest) Reset()         { *m = GetBackendRequest{} }
func (m *GetBackendRequest) String() string { return proto.CompactTextString(m) }
func (*GetBackendRequest) ProtoMessage()    {}
func (*GetBackendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{2}
}
func (m *GetBackendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBackendRequest.Unmarshal(m, b)
}
func (m *GetBackendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBackendRequest.Marshal(b, m, deterministic)
}
func (dst *GetBackendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBackendRequest.Merge(dst, src)
}
func (m *GetBackendRequest) XXX_Size() int {
	return xxx_messageInfo_GetBackendRequest.Size(m)
}
func (m *GetBackendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBackendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBackendRequest proto.InternalMessageInfo

func (m *GetBackendRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetBackendResponse struct {
	Backend              *BackendDetail `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetBackendResponse) Reset()         { *m = GetBackendResponse{} }
func (m *GetBackendResponse) String() string { return proto.CompactTextString(m) }
func (*GetBackendResponse) ProtoMessage()    {}
func (*GetBackendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{3}
}
func (m *GetBackendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBackendResponse.Unmarshal(m, b)
}
func (m *GetBackendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBackendResponse.Marshal(b, m, deterministic)
}
func (dst *GetBackendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBackendResponse.Merge(dst, src)
}
func (m *GetBackendResponse) XXX_Size() int {
	return xxx_messageInfo_GetBackendResponse.Size(m)
}
func (m *GetBackendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBackendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBackendResponse proto.InternalMessageInfo

func (m *GetBackendResponse) GetBackend() *BackendDetail {
	if m != nil {
		return m.Backend
	}
	return nil
}

type ListBackendRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               string   `protobuf:"bytes,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackendRequest) Reset()         { *m = ListBackendRequest{} }
func (m *ListBackendRequest) String() string { return proto.CompactTextString(m) }
func (*ListBackendRequest) ProtoMessage()    {}
func (*ListBackendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{4}
}
func (m *ListBackendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackendRequest.Unmarshal(m, b)
}
func (m *ListBackendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackendRequest.Marshal(b, m, deterministic)
}
func (dst *ListBackendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackendRequest.Merge(dst, src)
}
func (m *ListBackendRequest) XXX_Size() int {
	return xxx_messageInfo_ListBackendRequest.Size(m)
}
func (m *ListBackendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackendRequest proto.InternalMessageInfo

func (m *ListBackendRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListBackendRequest) GetOffset() string {
	if m != nil {
		return m.Offset
	}
	return ""
}

type ListBackendResponse struct {
	Backends             []*BackendDetail `protobuf:"bytes,1,rep,name=backends,proto3" json:"backends,omitempty"`
	Next                 string           `protobuf:"bytes,2,opt,name=next,proto3" json:"next,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListBackendResponse) Reset()         { *m = ListBackendResponse{} }
func (m *ListBackendResponse) String() string { return proto.CompactTextString(m) }
func (*ListBackendResponse) ProtoMessage()    {}
func (*ListBackendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{5}
}
func (m *ListBackendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackendResponse.Unmarshal(m, b)
}
func (m *ListBackendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackendResponse.Marshal(b, m, deterministic)
}
func (dst *ListBackendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackendResponse.Merge(dst, src)
}
func (m *ListBackendResponse) XXX_Size() int {
	return xxx_messageInfo_ListBackendResponse.Size(m)
}
func (m *ListBackendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackendResponse proto.InternalMessageInfo

func (m *ListBackendResponse) GetBackends() []*BackendDetail {
	if m != nil {
		return m.Backends
	}
	return nil
}

func (m *ListBackendResponse) GetNext() string {
	if m != nil {
		return m.Next
	}
	return ""
}

type UpdateBackendRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Access               string   `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	Security             string   `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBackendRequest) Reset()         { *m = UpdateBackendRequest{} }
func (m *UpdateBackendRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBackendRequest) ProtoMessage()    {}
func (*UpdateBackendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{6}
}
func (m *UpdateBackendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBackendRequest.Unmarshal(m, b)
}
func (m *UpdateBackendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBackendRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBackendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBackendRequest.Merge(dst, src)
}
func (m *UpdateBackendRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBackendRequest.Size(m)
}
func (m *UpdateBackendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBackendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBackendRequest proto.InternalMessageInfo

func (m *UpdateBackendRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateBackendRequest) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *UpdateBackendRequest) GetSecurity() string {
	if m != nil {
		return m.Security
	}
	return ""
}

type UpdateBackendResponse struct {
	Backend              *BackendDetail `protobuf:"bytes,1,opt,name=backend,proto3" json:"backend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateBackendResponse) Reset()         { *m = UpdateBackendResponse{} }
func (m *UpdateBackendResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBackendResponse) ProtoMessage()    {}
func (*UpdateBackendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{7}
}
func (m *UpdateBackendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBackendResponse.Unmarshal(m, b)
}
func (m *UpdateBackendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBackendResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateBackendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBackendResponse.Merge(dst, src)
}
func (m *UpdateBackendResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBackendResponse.Size(m)
}
func (m *UpdateBackendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBackendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBackendResponse proto.InternalMessageInfo

func (m *UpdateBackendResponse) GetBackend() *BackendDetail {
	if m != nil {
		return m.Backend
	}
	return nil
}

type DeleteBackendRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBackendRequest) Reset()         { *m = DeleteBackendRequest{} }
func (m *DeleteBackendRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBackendRequest) ProtoMessage()    {}
func (*DeleteBackendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{8}
}
func (m *DeleteBackendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBackendRequest.Unmarshal(m, b)
}
func (m *DeleteBackendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBackendRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBackendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBackendRequest.Merge(dst, src)
}
func (m *DeleteBackendRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBackendRequest.Size(m)
}
func (m *DeleteBackendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBackendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBackendRequest proto.InternalMessageInfo

func (m *DeleteBackendRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBackendResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBackendResponse) Reset()         { *m = DeleteBackendResponse{} }
func (m *DeleteBackendResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBackendResponse) ProtoMessage()    {}
func (*DeleteBackendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{9}
}
func (m *DeleteBackendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBackendResponse.Unmarshal(m, b)
}
func (m *DeleteBackendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBackendResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteBackendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBackendResponse.Merge(dst, src)
}
func (m *DeleteBackendResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBackendResponse.Size(m)
}
func (m *DeleteBackendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBackendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBackendResponse proto.InternalMessageInfo

type BackendDetail struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TenantId             string   `protobuf:"bytes,2,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Type                 string   `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Endpoint             string   `protobuf:"bytes,6,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	BucketName           string   `protobuf:"bytes,7,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	Access               string   `protobuf:"bytes,8,opt,name=access,proto3" json:"access,omitempty"`
	Security             string   `protobuf:"bytes,9,opt,name=security,proto3" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackendDetail) Reset()         { *m = BackendDetail{} }
func (m *BackendDetail) String() string { return proto.CompactTextString(m) }
func (*BackendDetail) ProtoMessage()    {}
func (*BackendDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_backend_eef016c15b2330d2, []int{10}
}
func (m *BackendDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackendDetail.Unmarshal(m, b)
}
func (m *BackendDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackendDetail.Marshal(b, m, deterministic)
}
func (dst *BackendDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackendDetail.Merge(dst, src)
}
func (m *BackendDetail) XXX_Size() int {
	return xxx_messageInfo_BackendDetail.Size(m)
}
func (m *BackendDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_BackendDetail.DiscardUnknown(m)
}

var xxx_messageInfo_BackendDetail proto.InternalMessageInfo

func (m *BackendDetail) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BackendDetail) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *BackendDetail) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BackendDetail) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BackendDetail) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BackendDetail) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *BackendDetail) GetBucketName() string {
	if m != nil {
		return m.BucketName
	}
	return ""
}

func (m *BackendDetail) GetAccess() string {
	if m != nil {
		return m.Access
	}
	return ""
}

func (m *BackendDetail) GetSecurity() string {
	if m != nil {
		return m.Security
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateBackendRequest)(nil), "CreateBackendRequest")
	proto.RegisterType((*CreateBackendResponse)(nil), "CreateBackendResponse")
	proto.RegisterType((*GetBackendRequest)(nil), "GetBackendRequest")
	proto.RegisterType((*GetBackendResponse)(nil), "GetBackendResponse")
	proto.RegisterType((*ListBackendRequest)(nil), "ListBackendRequest")
	proto.RegisterType((*ListBackendResponse)(nil), "ListBackendResponse")
	proto.RegisterType((*UpdateBackendRequest)(nil), "UpdateBackendRequest")
	proto.RegisterType((*UpdateBackendResponse)(nil), "UpdateBackendResponse")
	proto.RegisterType((*DeleteBackendRequest)(nil), "DeleteBackendRequest")
	proto.RegisterType((*DeleteBackendResponse)(nil), "DeleteBackendResponse")
	proto.RegisterType((*BackendDetail)(nil), "BackendDetail")
}

func init() { proto.RegisterFile("backend.proto", fileDescriptor_backend_eef016c15b2330d2) }

var fileDescriptor_backend_eef016c15b2330d2 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0xb5, 0x9d, 0xf8, 0x23, 0x13, 0x1c, 0xe8, 0x58, 0x52, 0x85, 0x0e, 0x25, 0x6c, 0xa1, 0x84,
	0x1e, 0xf6, 0x90, 0x1e, 0x0a, 0x3d, 0x94, 0x34, 0x0d, 0x94, 0x40, 0xe9, 0x41, 0x90, 0x4b, 0x6f,
	0xb2, 0x34, 0x81, 0x25, 0xb6, 0xa4, 0x6a, 0x57, 0xd0, 0xfc, 0xbf, 0xfe, 0x9b, 0xfe, 0x89, 0xb2,
	0xda, 0x95, 0xab, 0x8f, 0x2d, 0xc5, 0x37, 0xcd, 0x87, 0xde, 0xbc, 0x79, 0x7a, 0x23, 0x58, 0x6f,
	0x93, 0xf4, 0x89, 0xf2, 0x8c, 0x97, 0x55, 0xa1, 0x0a, 0x76, 0x03, 0xde, 0xe7, 0x8a, 0x12, 0x45,
	0xb7, 0x26, 0x1d, 0xd3, 0x8f, 0x9a, 0xa4, 0xc2, 0x2b, 0x58, 0xda, 0xc6, 0x70, 0x7a, 0x39, 0xbd,
	0x3a, 0xbf, 0xbe, 0xe0, 0xb6, 0xe3, 0x8e, 0x54, 0x22, 0x76, 0x71, 0x5b, 0x66, 0x9f, 0xc0, 0x1f,
	0x20, 0xc8, 0xb2, 0xc8, 0x25, 0x1d, 0x01, 0xf1, 0x1a, 0x5e, 0x7c, 0x21, 0x35, 0x60, 0x70, 0x01,
	0x33, 0x61, 0xde, 0x3c, 0x8b, 0x67, 0x22, 0x63, 0x1f, 0x01, 0xbb, 0x4d, 0x47, 0x0f, 0xb9, 0x05,
	0xfc, 0x2a, 0xe4, 0x70, 0x8a, 0x07, 0xf3, 0x9d, 0xd8, 0x0b, 0xd5, 0xbc, 0x3d, 0x8f, 0x4d, 0x80,
	0x01, 0x2c, 0x8a, 0xc7, 0x47, 0x49, 0x2a, 0x9c, 0x35, 0xf3, 0x6d, 0xc4, 0x1e, 0x60, 0xd3, 0xc3,
	0xb0, 0x24, 0xde, 0xc2, 0xca, 0x4e, 0x91, 0xe1, 0xf4, 0xf2, 0xc4, 0xc1, 0xe2, 0x50, 0x47, 0x84,
	0xd3, 0x9c, 0x7e, 0xb6, 0xc0, 0xcd, 0x33, 0xfb, 0x0e, 0xde, 0x43, 0x99, 0x8d, 0x3f, 0xc2, 0x40,
	0x02, 0x4d, 0x2b, 0x49, 0x53, 0x92, 0xb2, 0xa5, 0x65, 0x22, 0x8c, 0x60, 0x25, 0x29, 0xad, 0x2b,
	0xa1, 0x9e, 0xc3, 0x93, 0xa6, 0x72, 0x88, 0xf5, 0xe7, 0x19, 0x60, 0x1f, 0xad, 0xdc, 0x1b, 0xf0,
	0xee, 0x68, 0x47, 0xff, 0xa3, 0xc7, 0x5e, 0x82, 0x3f, 0xe8, 0x33, 0xa3, 0xd8, 0xef, 0x29, 0xac,
	0x7b, 0xd8, 0xa3, 0xcd, 0x22, 0x58, 0x29, 0xca, 0x93, 0x5c, 0xdd, 0x67, 0x76, 0xb7, 0x43, 0xac,
	0xb7, 0xae, 0x25, 0x55, 0xf7, 0x99, 0xdd, 0xcd, 0x46, 0x8d, 0x92, 0xc9, 0x9e, 0xc2, 0x53, 0xab,
	0x64, 0xb2, 0x27, 0x9d, 0x53, 0xcf, 0x25, 0x85, 0x73, 0x93, 0xd3, 0xcf, 0x1a, 0x9b, 0xf2, 0xac,
	0x2c, 0x44, 0xae, 0xc2, 0x85, 0xc1, 0x6e, 0x63, 0x7c, 0x05, 0xb0, 0xad, 0xd3, 0x27, 0x52, 0xdf,
	0x34, 0xd2, 0xb2, 0xa9, 0x76, 0x32, 0x1d, 0xc5, 0x57, 0xff, 0x54, 0xfc, 0xac, 0xaf, 0xf8, 0xf5,
	0xaf, 0x19, 0x2c, 0xed, 0xb6, 0x78, 0x03, 0xeb, 0xde, 0x71, 0xa0, 0xcf, 0x5d, 0xe7, 0x16, 0x05,
	0xdc, 0x79, 0x43, 0x6c, 0x82, 0xef, 0x01, 0xfe, 0xda, 0x1e, 0x91, 0x8f, 0x0e, 0x25, 0xda, 0xf0,
	0xf1, 0x5d, 0xb0, 0x09, 0x7e, 0x80, 0xf3, 0x8e, 0x57, 0x71, 0xc3, 0xc7, 0xee, 0x8f, 0x3c, 0xee,
	0xb0, 0x33, 0x9b, 0x68, 0xda, 0x3d, 0xd3, 0xa0, 0xcf, 0x5d, 0x06, 0x8d, 0x02, 0xee, 0xf4, 0x96,
	0x41, 0xe8, 0x79, 0x01, 0x7d, 0xee, 0xf2, 0x50, 0x14, 0x70, 0xb7, 0x65, 0x26, 0xdb, 0x45, 0xf3,
	0x83, 0x7a, 0xf7, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xc0, 0xe1, 0x6e, 0xb1, 0x04, 0x00, 0x00,
}