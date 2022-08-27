package rpc

import (
	"fmt"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitee.com/dongdong-0421/keyauth/apps/token"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/auth"
	"github.com/infraboard/mcenter/client/rpc/resolver"
)

var (
	client *ClientSet
)

// SetGlobal todo
func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	return client
}

// NewClient todo
// 传递注册中心的地址
func NewClient(conf *rpc.Config) (*ClientSet, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	// resolver 进行解析的时候 需要mcenter客户端实例已经初始化
	conn, err := grpc.Dial(
		// 127.0.0.1:18010 GRPC server端的地址
		// 基于服务发现  Dial to "passthrough://  dns://keyauth.org "mcenter://keyauth",
		fmt.Sprintf("%s://%s", resolver.Scheme, "keyauth"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(auth.NewAuthentication(conf.ClientID, conf.ClientSecret)),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Token服务的SDK
func (c *ClientSet) Token() token.ServiceClient {
	return token.NewServiceClient(c.conn)
}
