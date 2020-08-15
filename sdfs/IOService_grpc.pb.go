// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sdfs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FileIOServiceClient is the client API for FileIOService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileIOServiceClient interface {
	// Define a RPC operation
	MkDir(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error)
	RmDir(ctx context.Context, in *RmDirRequest, opts ...grpc.CallOption) (*RmDirResponse, error)
	Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*UnlinkResponse, error)
	Write(ctx context.Context, in *DataWriteRequest, opts ...grpc.CallOption) (*DataWriteResponse, error)
	Read(ctx context.Context, in *DataReadRequest, opts ...grpc.CallOption) (*DataReadResponse, error)
	Release(ctx context.Context, in *FileCloseRequest, opts ...grpc.CallOption) (*FileCloseResponse, error)
	Mknod(ctx context.Context, in *MkNodRequest, opts ...grpc.CallOption) (*MkNodResponse, error)
	Open(ctx context.Context, in *FileOpenRequest, opts ...grpc.CallOption) (*FileOpenResponse, error)
	GetFileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileMessageResponse, error)
	CreateCopy(ctx context.Context, in *FileSnapshotRequest, opts ...grpc.CallOption) (*FileSnapshotResponse, error)
	FileExists(ctx context.Context, in *FileExistsRequest, opts ...grpc.CallOption) (*FileExistsResponse, error)
	MkDirAll(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error)
	Stat(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileMessageResponse, error)
	Rename(ctx context.Context, in *FileRenameRequest, opts ...grpc.CallOption) (*FileRenameResponse, error)
	CopyExtent(ctx context.Context, in *CopyExtentRequest, opts ...grpc.CallOption) (*CopyExtentResponse, error)
	SetUserMetaData(ctx context.Context, in *SetUserMetaDataRequest, opts ...grpc.CallOption) (*SetUserMetaDataResponse, error)
}

type fileIOServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFileIOServiceClient(cc grpc.ClientConnInterface) FileIOServiceClient {
	return &fileIOServiceClient{cc}
}

func (c *fileIOServiceClient) MkDir(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error) {
	out := new(MkDirResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/MkDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) RmDir(ctx context.Context, in *RmDirRequest, opts ...grpc.CallOption) (*RmDirResponse, error) {
	out := new(RmDirResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/RmDir", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Unlink(ctx context.Context, in *UnlinkRequest, opts ...grpc.CallOption) (*UnlinkResponse, error) {
	out := new(UnlinkResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Unlink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Write(ctx context.Context, in *DataWriteRequest, opts ...grpc.CallOption) (*DataWriteResponse, error) {
	out := new(DataWriteResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Read(ctx context.Context, in *DataReadRequest, opts ...grpc.CallOption) (*DataReadResponse, error) {
	out := new(DataReadResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Release(ctx context.Context, in *FileCloseRequest, opts ...grpc.CallOption) (*FileCloseResponse, error) {
	out := new(FileCloseResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Release", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Mknod(ctx context.Context, in *MkNodRequest, opts ...grpc.CallOption) (*MkNodResponse, error) {
	out := new(MkNodResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Mknod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Open(ctx context.Context, in *FileOpenRequest, opts ...grpc.CallOption) (*FileOpenResponse, error) {
	out := new(FileOpenResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Open", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) GetFileInfo(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileMessageResponse, error) {
	out := new(FileMessageResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/GetFileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) CreateCopy(ctx context.Context, in *FileSnapshotRequest, opts ...grpc.CallOption) (*FileSnapshotResponse, error) {
	out := new(FileSnapshotResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/CreateCopy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) FileExists(ctx context.Context, in *FileExistsRequest, opts ...grpc.CallOption) (*FileExistsResponse, error) {
	out := new(FileExistsResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/FileExists", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) MkDirAll(ctx context.Context, in *MkDirRequest, opts ...grpc.CallOption) (*MkDirResponse, error) {
	out := new(MkDirResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/MkDirAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Stat(ctx context.Context, in *FileInfoRequest, opts ...grpc.CallOption) (*FileMessageResponse, error) {
	out := new(FileMessageResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Stat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) Rename(ctx context.Context, in *FileRenameRequest, opts ...grpc.CallOption) (*FileRenameResponse, error) {
	out := new(FileRenameResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/Rename", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) CopyExtent(ctx context.Context, in *CopyExtentRequest, opts ...grpc.CallOption) (*CopyExtentResponse, error) {
	out := new(CopyExtentResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/CopyExtent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileIOServiceClient) SetUserMetaData(ctx context.Context, in *SetUserMetaDataRequest, opts ...grpc.CallOption) (*SetUserMetaDataResponse, error) {
	out := new(SetUserMetaDataResponse)
	err := c.cc.Invoke(ctx, "/org.opendedup.grpc.FileIOService/SetUserMetaData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileIOServiceServer is the server API for FileIOService service.
// All implementations must embed UnimplementedFileIOServiceServer
// for forward compatibility
type FileIOServiceServer interface {
	// Define a RPC operation
	MkDir(context.Context, *MkDirRequest) (*MkDirResponse, error)
	RmDir(context.Context, *RmDirRequest) (*RmDirResponse, error)
	Unlink(context.Context, *UnlinkRequest) (*UnlinkResponse, error)
	Write(context.Context, *DataWriteRequest) (*DataWriteResponse, error)
	Read(context.Context, *DataReadRequest) (*DataReadResponse, error)
	Release(context.Context, *FileCloseRequest) (*FileCloseResponse, error)
	Mknod(context.Context, *MkNodRequest) (*MkNodResponse, error)
	Open(context.Context, *FileOpenRequest) (*FileOpenResponse, error)
	GetFileInfo(context.Context, *FileInfoRequest) (*FileMessageResponse, error)
	CreateCopy(context.Context, *FileSnapshotRequest) (*FileSnapshotResponse, error)
	FileExists(context.Context, *FileExistsRequest) (*FileExistsResponse, error)
	MkDirAll(context.Context, *MkDirRequest) (*MkDirResponse, error)
	Stat(context.Context, *FileInfoRequest) (*FileMessageResponse, error)
	Rename(context.Context, *FileRenameRequest) (*FileRenameResponse, error)
	CopyExtent(context.Context, *CopyExtentRequest) (*CopyExtentResponse, error)
	SetUserMetaData(context.Context, *SetUserMetaDataRequest) (*SetUserMetaDataResponse, error)
	mustEmbedUnimplementedFileIOServiceServer()
}

// UnimplementedFileIOServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFileIOServiceServer struct {
}

func (*UnimplementedFileIOServiceServer) MkDir(context.Context, *MkDirRequest) (*MkDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MkDir not implemented")
}
func (*UnimplementedFileIOServiceServer) RmDir(context.Context, *RmDirRequest) (*RmDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RmDir not implemented")
}
func (*UnimplementedFileIOServiceServer) Unlink(context.Context, *UnlinkRequest) (*UnlinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlink not implemented")
}
func (*UnimplementedFileIOServiceServer) Write(context.Context, *DataWriteRequest) (*DataWriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (*UnimplementedFileIOServiceServer) Read(context.Context, *DataReadRequest) (*DataReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedFileIOServiceServer) Release(context.Context, *FileCloseRequest) (*FileCloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Release not implemented")
}
func (*UnimplementedFileIOServiceServer) Mknod(context.Context, *MkNodRequest) (*MkNodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mknod not implemented")
}
func (*UnimplementedFileIOServiceServer) Open(context.Context, *FileOpenRequest) (*FileOpenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (*UnimplementedFileIOServiceServer) GetFileInfo(context.Context, *FileInfoRequest) (*FileMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileInfo not implemented")
}
func (*UnimplementedFileIOServiceServer) CreateCopy(context.Context, *FileSnapshotRequest) (*FileSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCopy not implemented")
}
func (*UnimplementedFileIOServiceServer) FileExists(context.Context, *FileExistsRequest) (*FileExistsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FileExists not implemented")
}
func (*UnimplementedFileIOServiceServer) MkDirAll(context.Context, *MkDirRequest) (*MkDirResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MkDirAll not implemented")
}
func (*UnimplementedFileIOServiceServer) Stat(context.Context, *FileInfoRequest) (*FileMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stat not implemented")
}
func (*UnimplementedFileIOServiceServer) Rename(context.Context, *FileRenameRequest) (*FileRenameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rename not implemented")
}
func (*UnimplementedFileIOServiceServer) CopyExtent(context.Context, *CopyExtentRequest) (*CopyExtentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyExtent not implemented")
}
func (*UnimplementedFileIOServiceServer) SetUserMetaData(context.Context, *SetUserMetaDataRequest) (*SetUserMetaDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserMetaData not implemented")
}
func (*UnimplementedFileIOServiceServer) mustEmbedUnimplementedFileIOServiceServer() {}

func RegisterFileIOServiceServer(s *grpc.Server, srv FileIOServiceServer) {
	s.RegisterService(&_FileIOService_serviceDesc, srv)
}

func _FileIOService_MkDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).MkDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/MkDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).MkDir(ctx, req.(*MkDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_RmDir_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RmDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).RmDir(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/RmDir",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).RmDir(ctx, req.(*RmDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Unlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Unlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Unlink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Unlink(ctx, req.(*UnlinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataWriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Write(ctx, req.(*DataWriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Read(ctx, req.(*DataReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Release_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileCloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Release(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Release",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Release(ctx, req.(*FileCloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Mknod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkNodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Mknod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Mknod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Mknod(ctx, req.(*MkNodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileOpenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Open",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Open(ctx, req.(*FileOpenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_GetFileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).GetFileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/GetFileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).GetFileInfo(ctx, req.(*FileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_CreateCopy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).CreateCopy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/CreateCopy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).CreateCopy(ctx, req.(*FileSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_FileExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).FileExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/FileExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).FileExists(ctx, req.(*FileExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_MkDirAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MkDirRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).MkDirAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/MkDirAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).MkDirAll(ctx, req.(*MkDirRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Stat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Stat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Stat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Stat(ctx, req.(*FileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_Rename_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRenameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).Rename(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/Rename",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).Rename(ctx, req.(*FileRenameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_CopyExtent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyExtentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).CopyExtent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/CopyExtent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).CopyExtent(ctx, req.(*CopyExtentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FileIOService_SetUserMetaData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserMetaDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileIOServiceServer).SetUserMetaData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/org.opendedup.grpc.FileIOService/SetUserMetaData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileIOServiceServer).SetUserMetaData(ctx, req.(*SetUserMetaDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileIOService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "org.opendedup.grpc.FileIOService",
	HandlerType: (*FileIOServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MkDir",
			Handler:    _FileIOService_MkDir_Handler,
		},
		{
			MethodName: "RmDir",
			Handler:    _FileIOService_RmDir_Handler,
		},
		{
			MethodName: "Unlink",
			Handler:    _FileIOService_Unlink_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _FileIOService_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _FileIOService_Read_Handler,
		},
		{
			MethodName: "Release",
			Handler:    _FileIOService_Release_Handler,
		},
		{
			MethodName: "Mknod",
			Handler:    _FileIOService_Mknod_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _FileIOService_Open_Handler,
		},
		{
			MethodName: "GetFileInfo",
			Handler:    _FileIOService_GetFileInfo_Handler,
		},
		{
			MethodName: "CreateCopy",
			Handler:    _FileIOService_CreateCopy_Handler,
		},
		{
			MethodName: "FileExists",
			Handler:    _FileIOService_FileExists_Handler,
		},
		{
			MethodName: "MkDirAll",
			Handler:    _FileIOService_MkDirAll_Handler,
		},
		{
			MethodName: "Stat",
			Handler:    _FileIOService_Stat_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _FileIOService_Rename_Handler,
		},
		{
			MethodName: "CopyExtent",
			Handler:    _FileIOService_CopyExtent_Handler,
		},
		{
			MethodName: "SetUserMetaData",
			Handler:    _FileIOService_SetUserMetaData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "IOService.proto",
}
