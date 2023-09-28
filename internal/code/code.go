package code

import (
	_ "embed"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	CacheSetError      = 10106
	CacheGetError      = 10107
	CacheDelError      = 10108
	CacheNotExist      = 10109
	ResubmitError      = 10110
	HashIdsEncodeError = 10111
	HashIdsDecodeError = 10112
	RBACError          = 10113
	RedisConnectError  = 10114
	MySQLConnectError  = 10115
	WriteConfigError   = 10116
	SendEmailError     = 10117
	MySQLExecError     = 10118
	GoVersionError     = 10119
	SocketConnectError = 10120
	SocketSendError    = 10121

	AuthorizedCreateError    = 20101
	AuthorizedListError      = 20102
	AuthorizedDeleteError    = 20103
	AuthorizedUpdateError    = 20104
	AuthorizedDetailError    = 20105
	AuthorizedCreateAPIError = 20106
	AuthorizedListAPIError   = 20107
	AuthorizedDeleteAPIError = 20108

	ArticleCreateError = 20201
	ArticleListError   = 20202
	ArticleDeleteError = 20203
	ArticleUpdateError = 20204
	ArticleDetailError = 20213
)
