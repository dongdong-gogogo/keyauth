package rpc_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitee.com/dongdong-0421/keyauth/apps/book"
	"gitee.com/dongdong-0421/keyauth/client/rpc"
	mcenter "github.com/infraboard/mcenter/client/rpc"
)

func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")
	// 设置GRPC服务地址
	// conf.SetAddress("127.0.0.1:8050")
	// 携带认证信息
	// conf.SetClientCredentials("secret_id", "secret_key")
	// 传递Mcenter配置，客户端通过Mcenter进行搜索
	c, err := rpc.NewClient(conf)
	if should.NoError(err) {
		resp, err := c.Book().QueryBook(
			context.Background(),
			book.NewQueryBookRequest(),
		)
		should.NoError(err)
		fmt.Println(resp.Items)
	}
}

func init() {
	// 提前加载好 rpc客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
