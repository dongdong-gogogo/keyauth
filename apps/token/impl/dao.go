package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dongdong-gogogo/keyauth/apps/token"
)

func (s *service) save(ctx context.Context, ins *token.Token) error {
	//s.col.InsertMany()
	if _, err := s.col.InsertOne(ctx, ins); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			ins.AccessToken, err)
	}
	return nil
}

func (s *service) get(ctx context.Context, accessToken string) (*token.Token, error) {
	filter := bson.M{"_id": accessToken}

	ins := token.NewDefaultToken()
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("as %s not found", accessToken)
		}

		return nil, exception.NewInternalServerError("find as %s error, %s", accessToken, err)
	}

	return ins, nil
}

// 通过主键来更新一个对象
func (s *service) update(ctx context.Context, ins *token.Token) error {
	if _, err := s.col.UpdateByID(ctx, ins.AccessToken, ins); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			ins.AccessToken, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, ins *token.Token) error {
	if ins == nil || ins.AccessToken == "" {
		return fmt.Errorf("access token is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.AccessToken})
	if err != nil {
		return exception.NewInternalServerError("delete access token (%s) error, %s", ins.AccessToken, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("access token %s not found", ins.AccessToken)
	}

	return nil
}
