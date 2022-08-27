package token

import (
	"fmt"
	"time"

	"github.com/infraboard/mcube/exception"

	"gitee.com/dongdong-0421/keyauth/utils"

	"gitee.com/dongdong-0421/keyauth/apps/user"
)

const (
	AppName = "token"
)

func NewIssueTokenRequest() *IssueTokenRequest {

	return &IssueTokenRequest{
		UserDomain: user.DefaultDomain,
	}
}

func (req *IssueTokenRequest) Validate() error {
	switch req.GranteType {
	case GranteType_PASSWORD:
		if req.UserName == "" || req.Password == "" {
			return fmt.Errorf("password grant required username and password")
		}
	}
	return nil
}

func NewDefaultToken() *Token {
	return &Token{
		Data: &IssueTokenRequest{},
		Meta: map[string]string{},
	}
}

func NewToken(req *IssueTokenRequest, expiredDuration time.Duration) *Token {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 4)
	return &Token{
		AccessToken:           utils.MakeBearer(24),
		IssueAt:               now.UnixMilli(),
		Data:                  req,
		AccessTokenExpiredAt:  expired.UnixMilli(),
		RefreshToken:          utils.MakeBearer(32),
		RefreshTokenExpiredAt: refresh.UnixMilli(),
	}
}

func (t *Token) Validate() error {
	//判断Token是否过期
	// 判断时间戳
	// 说明已经过期了
	if time.Now().UnixMilli() > t.AccessTokenExpiredAt {
		return exception.NewAccessTokenExpired("access token expired")
	}
	return nil
}

func (t *Token) IsRefreshTokenExpired() bool {
	//判断Refresh Token是否过期
	// 判断时间戳
	// refreshToken 过期
	if time.Now().UnixMilli() > t.RefreshTokenExpiredAt {
		return true
	}
	return false
}

// 延长Token，延迟一个生命周期
func (t *Token) Extend(expiredDuration time.Duration) {
	now := time.Now()
	// Token 10
	expired := now.Add(expiredDuration)
	refresh := now.Add(expiredDuration * 4)

	t.AccessTokenExpiredAt = expired.UnixMilli()
	t.RefreshTokenExpiredAt = refresh.UnixMilli()

}

func NewDescribeTokenRequest(as string) *DescribeTokenRequest {
	return &DescribeTokenRequest{
		AccessToken: as,
	}
}

func NewValidateTokenRequest(as string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: as,
	}
}
func NewRevolkTokenRequest() *RevolkTokenRequest {
	return &RevolkTokenRequest{}
}
