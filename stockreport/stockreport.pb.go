// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stockreport.proto

package stockreport

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StockReportRequest struct {
	Resolution           int32    `protobuf:"varint,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	Start                int64    `protobuf:"varint,2,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,3,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockReportRequest) Reset()         { *m = StockReportRequest{} }
func (m *StockReportRequest) String() string { return proto.CompactTextString(m) }
func (*StockReportRequest) ProtoMessage()    {}
func (*StockReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_stockreport_49be123f76f1a0c8, []int{0}
}
func (m *StockReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockReportRequest.Unmarshal(m, b)
}
func (m *StockReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockReportRequest.Marshal(b, m, deterministic)
}
func (dst *StockReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockReportRequest.Merge(dst, src)
}
func (m *StockReportRequest) XXX_Size() int {
	return xxx_messageInfo_StockReportRequest.Size(m)
}
func (m *StockReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StockReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StockReportRequest proto.InternalMessageInfo

func (m *StockReportRequest) GetResolution() int32 {
	if m != nil {
		return m.Resolution
	}
	return 0
}

func (m *StockReportRequest) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *StockReportRequest) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

type StockReportItem struct {
	Stockname            string   `protobuf:"bytes,1,opt,name=stockname,proto3" json:"stockname,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Rating               int32    `protobuf:"varint,3,opt,name=rating,proto3" json:"rating,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StockReportItem) Reset()         { *m = StockReportItem{} }
func (m *StockReportItem) String() string { return proto.CompactTextString(m) }
func (*StockReportItem) ProtoMessage()    {}
func (*StockReportItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_stockreport_49be123f76f1a0c8, []int{1}
}
func (m *StockReportItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockReportItem.Unmarshal(m, b)
}
func (m *StockReportItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockReportItem.Marshal(b, m, deterministic)
}
func (dst *StockReportItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockReportItem.Merge(dst, src)
}
func (m *StockReportItem) XXX_Size() int {
	return xxx_messageInfo_StockReportItem.Size(m)
}
func (m *StockReportItem) XXX_DiscardUnknown() {
	xxx_messageInfo_StockReportItem.DiscardUnknown(m)
}

var xxx_messageInfo_StockReportItem proto.InternalMessageInfo

func (m *StockReportItem) GetStockname() string {
	if m != nil {
		return m.Stockname
	}
	return ""
}

func (m *StockReportItem) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *StockReportItem) GetRating() int32 {
	if m != nil {
		return m.Rating
	}
	return 0
}

type StockReportResponse struct {
	StockData            []*StockReportItem `protobuf:"bytes,1,rep,name=StockData,proto3" json:"StockData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StockReportResponse) Reset()         { *m = StockReportResponse{} }
func (m *StockReportResponse) String() string { return proto.CompactTextString(m) }
func (*StockReportResponse) ProtoMessage()    {}
func (*StockReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_stockreport_49be123f76f1a0c8, []int{2}
}
func (m *StockReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StockReportResponse.Unmarshal(m, b)
}
func (m *StockReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StockReportResponse.Marshal(b, m, deterministic)
}
func (dst *StockReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StockReportResponse.Merge(dst, src)
}
func (m *StockReportResponse) XXX_Size() int {
	return xxx_messageInfo_StockReportResponse.Size(m)
}
func (m *StockReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StockReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StockReportResponse proto.InternalMessageInfo

func (m *StockReportResponse) GetStockData() []*StockReportItem {
	if m != nil {
		return m.StockData
	}
	return nil
}

func init() {
	proto.RegisterType((*StockReportRequest)(nil), "StockReportRequest")
	proto.RegisterType((*StockReportItem)(nil), "StockReportItem")
	proto.RegisterType((*StockReportResponse)(nil), "StockReportResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockReportDataServiceClient is the client API for StockReportDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockReportDataServiceClient interface {
	GetStockReportData(ctx context.Context, in *StockReportRequest, opts ...grpc.CallOption) (*StockReportResponse, error)
}

type stockReportDataServiceClient struct {
	cc *grpc.ClientConn
}

func NewStockReportDataServiceClient(cc *grpc.ClientConn) StockReportDataServiceClient {
	return &stockReportDataServiceClient{cc}
}

func (c *stockReportDataServiceClient) GetStockReportData(ctx context.Context, in *StockReportRequest, opts ...grpc.CallOption) (*StockReportResponse, error) {
	out := new(StockReportResponse)
	err := c.cc.Invoke(ctx, "/StockReportDataService/GetStockReportData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockReportDataServiceServer is the server API for StockReportDataService service.
type StockReportDataServiceServer interface {
	GetStockReportData(context.Context, *StockReportRequest) (*StockReportResponse, error)
}

func RegisterStockReportDataServiceServer(s *grpc.Server, srv StockReportDataServiceServer) {
	s.RegisterService(&_StockReportDataService_serviceDesc, srv)
}

func _StockReportDataService_GetStockReportData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockReportDataServiceServer).GetStockReportData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/StockReportDataService/GetStockReportData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockReportDataServiceServer).GetStockReportData(ctx, req.(*StockReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StockReportDataService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "StockReportDataService",
	HandlerType: (*StockReportDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStockReportData",
			Handler:    _StockReportDataService_GetStockReportData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stockreport.proto",
}

func init() { proto.RegisterFile("stockreport.proto", fileDescriptor_stockreport_49be123f76f1a0c8) }

var fileDescriptor_stockreport_49be123f76f1a0c8 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xad, 0x61, 0x0b, 0x3b, 0x1e, 0xac, 0xd3, 0x52, 0x16, 0x11, 0x91, 0x9c, 0x7a, 0xca,
	0xa1, 0x7e, 0x02, 0x41, 0x11, 0xaf, 0xe9, 0x51, 0x2f, 0xb1, 0x0e, 0x12, 0x34, 0x7f, 0x4c, 0xa6,
	0x7e, 0x7e, 0x37, 0xb1, 0xb0, 0x6b, 0xed, 0x2d, 0xf3, 0x7b, 0xe1, 0xcd, 0x9b, 0x07, 0x17, 0x99,
	0xc3, 0xf6, 0x23, 0x51, 0x0c, 0x89, 0x55, 0x4c, 0x81, 0x83, 0x7c, 0x01, 0xdc, 0x14, 0xa8, 0x2b,
	0xd4, 0xf4, 0xb5, 0xa3, 0xcc, 0x78, 0x0d, 0x90, 0x28, 0x87, 0xcf, 0x1d, 0xdb, 0xe0, 0xbb, 0xc9,
	0xcd, 0x64, 0xd5, 0xe8, 0x11, 0xc1, 0x05, 0x34, 0x99, 0x4d, 0xe2, 0xee, 0xb4, 0x97, 0x84, 0xfe,
	0x1d, 0x70, 0x06, 0x82, 0xfc, 0x5b, 0x27, 0x2a, 0x2b, 0x4f, 0x49, 0x70, 0x3e, 0x72, 0x7f, 0x62,
	0x72, 0x78, 0x05, 0x6d, 0x4d, 0xe1, 0x8d, 0xa3, 0xea, 0xdc, 0xea, 0x01, 0x14, 0x95, 0xad, 0xeb,
	0x23, 0x18, 0x17, 0xf7, 0xe6, 0x03, 0xc0, 0x25, 0x4c, 0x93, 0x61, 0xeb, 0xdf, 0xeb, 0x8e, 0x46,
	0xef, 0x27, 0xf9, 0x00, 0xf3, 0x3f, 0x47, 0xe4, 0x18, 0x7c, 0x26, 0x54, 0xd0, 0x56, 0x7c, 0x6f,
	0xd8, 0xf4, 0xab, 0xc4, 0xea, 0x6c, 0x3d, 0x53, 0x07, 0x79, 0xf4, 0xf0, 0x65, 0xfd, 0x0c, 0xcb,
	0x91, 0x5a, 0xd0, 0x86, 0xd2, 0xb7, 0xdd, 0x12, 0xde, 0x01, 0x3e, 0x12, 0x1f, 0x88, 0x38, 0x57,
	0xff, 0xab, 0xbb, 0x5c, 0xa8, 0x23, 0x51, 0xe4, 0xc9, 0xeb, 0xb4, 0xf6, 0x7d, 0xfb, 0x13, 0x00,
	0x00, 0xff, 0xff, 0x2c, 0x80, 0x93, 0xdb, 0x84, 0x01, 0x00, 0x00,
}
