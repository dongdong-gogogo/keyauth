package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/dongdong-gogogo/keyauth/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dongdong-gogogo/keyauth/apps/token"
	"github.com/infraboard/mcube/exception"

	"github.com/dongdong-gogogo/keyauth/apps/user"
)

var (
	AUTH_ERROR = "user or password not correct"
)

var (
	DefaultTokenDuration = 10 * time.Minute
)

func (s *service) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate issue token error,%s", err)
	}
	// 根据不同授权模型来做不同的验证
	switch req.GranteType {
	case token.GranteType_PASSWORD:
		// 1 、获取用户对象（User Object）
		descReq := user.NewDescribeUserRequestByName(req.UserDomain, req.UserName)
		u, err := s.user.DescribeUser(ctx, descReq)
		if err != nil {
			s.log.Debug("describe user error, %s", err)
			if exception.IsNotFoundError(err) {
				// 401
				return nil, exception.NewUnauthorized(AUTH_ERROR)
			}
			return nil, err
		}
		s.log.Debug(u)
		// 2、校验用户密码是否正确
		if ok := u.CheckPassword(req.Password); !ok {
			return nil, exception.NewUnauthorized(AUTH_ERROR)
		}

		// 3、颁发一个token
		tk := token.NewToken(req, DefaultTokenDuration)

		// 4、脱敏
		tk.Data.Password = ""
		// 5、入库
		if err := s.save(ctx, tk); err != nil {
			return nil, err
		}
		return tk, err

	default:
		return nil, fmt.Errorf("grant type %s not implemented", req.GranteType)
	}
	return nil, nil
}

func (s *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {
	// 1、获取AccessToken

	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}
	// 2、检查RefreshToken是否匹配
	if tk.RefreshToken != req.RefreshToken {
		return nil, exception.NewBadRequest("refresh token not conrrect")
	}
	// 3、删除
	if err := s.delete(ctx, tk); err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	// 1、获取AccessToken
	tk, err := s.get(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}
	// 2、校验 Token的合法性
	if tk.Validate(); err != nil {
		// 2.1 如果Access Token过期
		if utils.IsAccessTokenExpiredError(err) {
			if tk.IsRefreshTokenExpired() {
				return nil, exception.NewRefreshTokenExpired("refresh token expired")
			}
			// 2.2 如果Refresh 没过期，延长时间
			//类似执行一个Update，Update Expired 时间
			tk.Extend(DefaultTokenDuration)
			if err := s.update(ctx, tk); err != nil {
				return nil, err
			}
			// 返回续约后的token
			return tk, err
		}
		return nil, err
	}
	return tk, nil
}

func (s *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
