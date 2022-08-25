package impl

import (
	"context"

	"gitee.com/dongdong-0421/keyauth/utils"

	"github.com/infraboard/mcube/exception"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitee.com/dongdong-0421/keyauth/apps/user"
)

func (s *service) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate create user error,%s", err)
	}

	ins := user.NewUser(req)
	ins.Data.Password = utils.HashPassword(ins.Data.Password)
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return nil, exception.NewInternalServerError("inserted user(%s) document error, %s", ins.Data.Name, err)
	}
	return ins, nil
}
func (s *service) QueryUser(ctx context.Context, req *user.QueryUserRequest) (*user.UserSet, error) {
	query := newQueryUserRequest(req)
	return s.query(ctx, query)
}
func (s *service) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	return s.get(ctx, req.UserId)
}
func (s *service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (s *service) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
