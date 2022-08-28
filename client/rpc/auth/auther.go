package auth

import (
	"github.com/infraboard/mcube/logger"

	"gitee.com/dongdong-0421/keyauth/apps/token"
)

func NewKeyauthAuther(auth token.ServiceClient) *KeyauthAuther {
	return &KeyauthAuther{
		auth: auth,
	}
}

// 有keyauth提供的 HTTP认证中间件
type KeyauthAuther struct {
	log  logger.Logger
	auth token.ServiceClient
}
