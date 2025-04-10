// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: proto/auth/v1/auth.proto

package v1

import (
	v1 "github.com/amanc1361/bilan-rekar-proto/gen/proto/user/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Token message represents an authentication token
type Token struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	TokenType     string                 `protobuf:"bytes,4,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"` // Usually "Bearer"
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Token) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Token) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Token) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

// UserClaims represents the claims embedded in a token
type UserClaims struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        uint64                 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Roles         []string               `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	Permissions   []string               `protobuf:"bytes,4,rep,name=permissions,proto3" json:"permissions,omitempty"`
	IssuedAt      *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserClaims) Reset() {
	*x = UserClaims{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserClaims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserClaims) ProtoMessage() {}

func (x *UserClaims) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserClaims.ProtoReflect.Descriptor instead.
func (*UserClaims) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserClaims) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserClaims) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserClaims) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserClaims) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UserClaims) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *UserClaims) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

// Login request from another microservice
type LoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Token response for successful login
type LoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token         *Token                 `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	User          *v1.User               `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"` // Basic user info
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginResponse) GetToken() *Token {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *LoginResponse) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

// Refresh token request
type RefreshTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

// Validate token request
type ValidateTokenRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenRequest) Reset() {
	*x = ValidateTokenRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenRequest) ProtoMessage() {}

func (x *ValidateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateTokenRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

// Validate token response
type ValidateTokenResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Valid         bool                   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Claims        *UserClaims            `protobuf:"bytes,3,opt,name=claims,proto3" json:"claims,omitempty"` // Only populated if valid is true
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidateTokenResponse) Reset() {
	*x = ValidateTokenResponse{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenResponse) ProtoMessage() {}

func (x *ValidateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateTokenResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{6}
}

func (x *ValidateTokenResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateTokenResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ValidateTokenResponse) GetClaims() *UserClaims {
	if x != nil {
		return x.Claims
	}
	return nil
}

// Check permission request
type CheckPermissionRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Permission    string                 `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckPermissionRequest) Reset() {
	*x = CheckPermissionRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckPermissionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRequest) ProtoMessage() {}

func (x *CheckPermissionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRequest.ProtoReflect.Descriptor instead.
func (*CheckPermissionRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *CheckPermissionRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CheckPermissionRequest) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

// Check role request
type CheckRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	Role          string                 `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CheckRoleRequest) Reset() {
	*x = CheckRoleRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CheckRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRoleRequest) ProtoMessage() {}

func (x *CheckRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRoleRequest.ProtoReflect.Descriptor instead.
func (*CheckRoleRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{8}
}

func (x *CheckRoleRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CheckRoleRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

// Authorization check response
type AuthorizationResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authorized    bool                   `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthorizationResponse) Reset() {
	*x = AuthorizationResponse{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthorizationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizationResponse) ProtoMessage() {}

func (x *AuthorizationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizationResponse.ProtoReflect.Descriptor instead.
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *AuthorizationResponse) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

func (x *AuthorizationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Logout request
type LogoutRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessToken   string                 `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	UserId        uint64                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *LogoutRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LogoutRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// Status response
type StatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	mi := &file_proto_auth_v1_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_auth_v1_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_auth_v1_auth_proto_rawDescGZIP(), []int{11}
}

func (x *StatusResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *StatusResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_auth_v1_auth_proto protoreflect.FileDescriptor

const file_proto_auth_v1_auth_proto_rawDesc = "" +
	"\n" +
	"\x18proto/auth/v1/auth.proto\x12\x04auth\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18proto/user/v1/user.proto\"\xa9\x01\n" +
	"\x05Token\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12#\n" +
	"\rrefresh_token\x18\x02 \x01(\tR\frefreshToken\x129\n" +
	"\n" +
	"expires_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\x12\x1d\n" +
	"\n" +
	"token_type\x18\x04 \x01(\tR\ttokenType\"\xed\x01\n" +
	"\n" +
	"UserClaims\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\x04R\x06userId\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05roles\x18\x03 \x03(\tR\x05roles\x12 \n" +
	"\vpermissions\x18\x04 \x03(\tR\vpermissions\x127\n" +
	"\tissued_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\bissuedAt\x129\n" +
	"\n" +
	"expires_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\"F\n" +
	"\fLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1a\n" +
	"\bpassword\x18\x02 \x01(\tR\bpassword\"\x86\x01\n" +
	"\rLoginResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12!\n" +
	"\x05token\x18\x03 \x01(\v2\v.auth.TokenR\x05token\x12\x1e\n" +
	"\x04user\x18\x04 \x01(\v2\n" +
	".user.UserR\x04user\":\n" +
	"\x13RefreshTokenRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"9\n" +
	"\x14ValidateTokenRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\"q\n" +
	"\x15ValidateTokenResponse\x12\x14\n" +
	"\x05valid\x18\x01 \x01(\bR\x05valid\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12(\n" +
	"\x06claims\x18\x03 \x01(\v2\x10.auth.UserClaimsR\x06claims\"[\n" +
	"\x16CheckPermissionRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12\x1e\n" +
	"\n" +
	"permission\x18\x02 \x01(\tR\n" +
	"permission\"I\n" +
	"\x10CheckRoleRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12\x12\n" +
	"\x04role\x18\x02 \x01(\tR\x04role\"Q\n" +
	"\x15AuthorizationResponse\x12\x1e\n" +
	"\n" +
	"authorized\x18\x01 \x01(\bR\n" +
	"authorized\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"K\n" +
	"\rLogoutRequest\x12!\n" +
	"\faccess_token\x18\x01 \x01(\tR\vaccessToken\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\x04R\x06userId\"D\n" +
	"\x0eStatusResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2\xe8\x04\n" +
	"\vAuthService\x120\n" +
	"\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x12>\n" +
	"\fRefreshToken\x12\x19.auth.RefreshTokenRequest\x1a\x13.auth.LoginResponse\x12H\n" +
	"\rValidateToken\x12\x1a.auth.ValidateTokenRequest\x1a\x1b.auth.ValidateTokenResponse\x12L\n" +
	"\x0fCheckPermission\x12\x1c.auth.CheckPermissionRequest\x1a\x1b.auth.AuthorizationResponse\x12@\n" +
	"\tCheckRole\x12\x16.auth.CheckRoleRequest\x1a\x1b.auth.AuthorizationResponse\x123\n" +
	"\x06Logout\x12\x13.auth.LogoutRequest\x1a\x14.auth.StatusResponse\x12A\n" +
	"\fLoginWithOTP\x12\x1c.user.VerifyMobileOTPRequest\x1a\x13.auth.LoginResponse\x12I\n" +
	"\x14RequestPasswordReset\x12\x1b.user.GetUserByEmailRequest\x1a\x14.auth.StatusResponse\x12J\n" +
	"\x16ResetPasswordWithToken\x12\x1a.user.ResetPasswordRequest\x1a\x14.auth.StatusResponseBy\n" +
	"\bcom.authB\tAuthProtoP\x01Z2github.com/amanc1361/bilan-rekar/gen/proto/auth/v1\xa2\x02\x03AXX\xaa\x02\x04Auth\xca\x02\x04Auth\xe2\x02\x10Auth\\GPBMetadata\xea\x02\x04Authb\x06proto3"

var (
	file_proto_auth_v1_auth_proto_rawDescOnce sync.Once
	file_proto_auth_v1_auth_proto_rawDescData []byte
)

func file_proto_auth_v1_auth_proto_rawDescGZIP() []byte {
	file_proto_auth_v1_auth_proto_rawDescOnce.Do(func() {
		file_proto_auth_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_auth_v1_auth_proto_rawDesc), len(file_proto_auth_v1_auth_proto_rawDesc)))
	})
	return file_proto_auth_v1_auth_proto_rawDescData
}

var file_proto_auth_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_proto_auth_v1_auth_proto_goTypes = []any{
	(*Token)(nil),                     // 0: auth.Token
	(*UserClaims)(nil),                // 1: auth.UserClaims
	(*LoginRequest)(nil),              // 2: auth.LoginRequest
	(*LoginResponse)(nil),             // 3: auth.LoginResponse
	(*RefreshTokenRequest)(nil),       // 4: auth.RefreshTokenRequest
	(*ValidateTokenRequest)(nil),      // 5: auth.ValidateTokenRequest
	(*ValidateTokenResponse)(nil),     // 6: auth.ValidateTokenResponse
	(*CheckPermissionRequest)(nil),    // 7: auth.CheckPermissionRequest
	(*CheckRoleRequest)(nil),          // 8: auth.CheckRoleRequest
	(*AuthorizationResponse)(nil),     // 9: auth.AuthorizationResponse
	(*LogoutRequest)(nil),             // 10: auth.LogoutRequest
	(*StatusResponse)(nil),            // 11: auth.StatusResponse
	(*timestamppb.Timestamp)(nil),     // 12: google.protobuf.Timestamp
	(*v1.User)(nil),                   // 13: user.User
	(*v1.VerifyMobileOTPRequest)(nil), // 14: user.VerifyMobileOTPRequest
	(*v1.GetUserByEmailRequest)(nil),  // 15: user.GetUserByEmailRequest
	(*v1.ResetPasswordRequest)(nil),   // 16: user.ResetPasswordRequest
}
var file_proto_auth_v1_auth_proto_depIdxs = []int32{
	12, // 0: auth.Token.expires_at:type_name -> google.protobuf.Timestamp
	12, // 1: auth.UserClaims.issued_at:type_name -> google.protobuf.Timestamp
	12, // 2: auth.UserClaims.expires_at:type_name -> google.protobuf.Timestamp
	0,  // 3: auth.LoginResponse.token:type_name -> auth.Token
	13, // 4: auth.LoginResponse.user:type_name -> user.User
	1,  // 5: auth.ValidateTokenResponse.claims:type_name -> auth.UserClaims
	2,  // 6: auth.AuthService.Login:input_type -> auth.LoginRequest
	4,  // 7: auth.AuthService.RefreshToken:input_type -> auth.RefreshTokenRequest
	5,  // 8: auth.AuthService.ValidateToken:input_type -> auth.ValidateTokenRequest
	7,  // 9: auth.AuthService.CheckPermission:input_type -> auth.CheckPermissionRequest
	8,  // 10: auth.AuthService.CheckRole:input_type -> auth.CheckRoleRequest
	10, // 11: auth.AuthService.Logout:input_type -> auth.LogoutRequest
	14, // 12: auth.AuthService.LoginWithOTP:input_type -> user.VerifyMobileOTPRequest
	15, // 13: auth.AuthService.RequestPasswordReset:input_type -> user.GetUserByEmailRequest
	16, // 14: auth.AuthService.ResetPasswordWithToken:input_type -> user.ResetPasswordRequest
	3,  // 15: auth.AuthService.Login:output_type -> auth.LoginResponse
	3,  // 16: auth.AuthService.RefreshToken:output_type -> auth.LoginResponse
	6,  // 17: auth.AuthService.ValidateToken:output_type -> auth.ValidateTokenResponse
	9,  // 18: auth.AuthService.CheckPermission:output_type -> auth.AuthorizationResponse
	9,  // 19: auth.AuthService.CheckRole:output_type -> auth.AuthorizationResponse
	11, // 20: auth.AuthService.Logout:output_type -> auth.StatusResponse
	3,  // 21: auth.AuthService.LoginWithOTP:output_type -> auth.LoginResponse
	11, // 22: auth.AuthService.RequestPasswordReset:output_type -> auth.StatusResponse
	11, // 23: auth.AuthService.ResetPasswordWithToken:output_type -> auth.StatusResponse
	15, // [15:24] is the sub-list for method output_type
	6,  // [6:15] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_auth_v1_auth_proto_init() }
func file_proto_auth_v1_auth_proto_init() {
	if File_proto_auth_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_auth_v1_auth_proto_rawDesc), len(file_proto_auth_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_auth_v1_auth_proto_goTypes,
		DependencyIndexes: file_proto_auth_v1_auth_proto_depIdxs,
		MessageInfos:      file_proto_auth_v1_auth_proto_msgTypes,
	}.Build()
	File_proto_auth_v1_auth_proto = out.File
	file_proto_auth_v1_auth_proto_goTypes = nil
	file_proto_auth_v1_auth_proto_depIdxs = nil
}
