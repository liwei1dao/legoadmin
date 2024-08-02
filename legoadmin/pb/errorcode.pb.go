// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: errorcode.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrorCode int32

const (
	ErrorCode_Success                 ErrorCode = 0  //成功
	ErrorCode_GatewayException        ErrorCode = 1  //网关执行异常
	ErrorCode_NoFindService           ErrorCode = 10 //没有找到远程服务器
	ErrorCode_NoFindServiceHandleFunc ErrorCode = 11 //远程服务器未找到执行方法
	ErrorCode_RpcFuncExecutionError   ErrorCode = 12 //Rpc方法执行错误
	ErrorCode_CacheReadError          ErrorCode = 13 //缓存读取失败
	ErrorCode_SqlExecutionError       ErrorCode = 14 //数据库执行错误
	ErrorCode_ReqParameterError       ErrorCode = 15 //请求参数错误
	ErrorCode_SignError               ErrorCode = 16 //签名串错误
	ErrorCode_InsufficientPermissions ErrorCode = 17 //权限不足
	ErrorCode_NoLogin                 ErrorCode = 18 //未登录
	ErrorCode_UserSessionNobeing      ErrorCode = 19 //用户不存在
	ErrorCode_StateInvalid            ErrorCode = 20 //无效状态
	ErrorCode_DBError                 ErrorCode = 21 //数据库操作失败
	ErrorCode_SystemError             ErrorCode = 22 //通用错误
	ErrorCode_TokenInvalid            ErrorCode = 23 //token 是空
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:  "Success",
		1:  "GatewayException",
		10: "NoFindService",
		11: "NoFindServiceHandleFunc",
		12: "RpcFuncExecutionError",
		13: "CacheReadError",
		14: "SqlExecutionError",
		15: "ReqParameterError",
		16: "SignError",
		17: "InsufficientPermissions",
		18: "NoLogin",
		19: "UserSessionNobeing",
		20: "StateInvalid",
		21: "DBError",
		22: "SystemError",
		23: "TokenInvalid",
	}
	ErrorCode_value = map[string]int32{
		"Success":                 0,
		"GatewayException":        1,
		"NoFindService":           10,
		"NoFindServiceHandleFunc": 11,
		"RpcFuncExecutionError":   12,
		"CacheReadError":          13,
		"SqlExecutionError":       14,
		"ReqParameterError":       15,
		"SignError":               16,
		"InsufficientPermissions": 17,
		"NoLogin":                 18,
		"UserSessionNobeing":      19,
		"StateInvalid":            20,
		"DBError":                 21,
		"SystemError":             22,
		"TokenInvalid":            23,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errorcode_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_errorcode_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_errorcode_proto_rawDescGZIP(), []int{0}
}

var File_errorcode_proto protoreflect.FileDescriptor

var file_errorcode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xce, 0x02, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4e, 0x6f, 0x46, 0x69, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x4e, 0x6f, 0x46, 0x69, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x46, 0x75, 0x6e, 0x63,
	0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x70, 0x63, 0x46, 0x75, 0x6e, 0x63, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0c, 0x12, 0x12, 0x0a,
	0x0e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x61, 0x64, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10,
	0x0d, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x71, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0e, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x0f, 0x12,
	0x0d, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x10, 0x12, 0x1b,
	0x0a, 0x17, 0x49, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x10, 0x11, 0x12, 0x0b, 0x0a, 0x07, 0x4e,
	0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x12, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x6f, 0x62, 0x65, 0x69, 0x6e, 0x67, 0x10, 0x13,
	0x12, 0x10, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x42, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x15, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x16,
	0x12, 0x10, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x17, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_errorcode_proto_rawDescOnce sync.Once
	file_errorcode_proto_rawDescData = file_errorcode_proto_rawDesc
)

func file_errorcode_proto_rawDescGZIP() []byte {
	file_errorcode_proto_rawDescOnce.Do(func() {
		file_errorcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorcode_proto_rawDescData)
	})
	return file_errorcode_proto_rawDescData
}

var file_errorcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errorcode_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: ErrorCode
}
var file_errorcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errorcode_proto_init() }
func file_errorcode_proto_init() {
	if File_errorcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errorcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errorcode_proto_goTypes,
		DependencyIndexes: file_errorcode_proto_depIdxs,
		EnumInfos:         file_errorcode_proto_enumTypes,
	}.Build()
	File_errorcode_proto = out.File
	file_errorcode_proto_rawDesc = nil
	file_errorcode_proto_goTypes = nil
	file_errorcode_proto_depIdxs = nil
}
