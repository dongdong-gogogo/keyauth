package auth

import (
	"github.com/infraboard/mcube/logger"

	"github.com/dongdong-gogogo/keyauth/apps/token"
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
