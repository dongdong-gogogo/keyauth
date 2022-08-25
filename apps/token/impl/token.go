package impl

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitee.com/dongdong-0421/keyauth/apps/token"
	"github.com/infraboard/mcube/exception"

	"gitee.com/dongdong-0421/keyauth/apps/user"
)

var (
	AUTH_ERROR = "user or password not correct"
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
		tk := token.NewToken(req, 10*time.Minute)
		return tk, err

	default:
		return nil, fmt.Errorf("grant type %s not implemented", req.GranteType)
	}
	return nil, nil
}

func (i *service) RevolkToken(ctx context.Context, req *token.RevolkTokenRequest) (*token.Token, error) {

	return nil, nil
}

func (i *service) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {

	return nil, nil
}

func (i *service) QueryToken(ctx context.Context, req *token.QueryTokenRequest) (*token.TokenSet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryToken not implemented")
}
