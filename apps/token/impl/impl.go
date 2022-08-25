package impl

import (
	"context"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"google.golang.org/grpc"

	"gitee.com/dongdong-0421/keyauth/conf"

	"gitee.com/dongdong-0421/keyauth/apps/token"
	"gitee.com/dongdong-0421/keyauth/apps/user"
)

var (
	// Service 服务实例
	svr = &service{}
)

type service struct {
	col *mongo.Collection
	log logger.Logger
	token.UnimplementedServiceServer
	user user.ServiceServer
}

func (s *service) Config() error {
	// 依赖MongoDB的DB对象
	db, err := conf.C().Mongo.GetDB()
	if err != nil {
		return err
	}
	// 获取一个Collection对象
	s.col = db.Collection(s.Name())
	s.log = zap.L().Named(s.Name())
	s.user = app.GetGrpcApp(user.AppName).(user.ServiceServer)
	// 创建索引
	indexs := []mongo.IndexModel{
		{
			Keys: bsonx.Doc{
				{Key: "data.domain", Value: bsonx.Int32(-1)},
				{Key: "data.name", Value: bsonx.Int32(-1)},
			},
			Options: options.Index().SetUnique(true),
		},
	}

	_, err = s.col.Indexes().CreateMany(context.Background(), indexs)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Name() string {
	return token.AppName
}

func (s *service) Registry(server *grpc.Server) {
	token.RegisterServiceServer(server, svr)
}

func init() {
	app.RegistryGrpcApp(svr)
}
