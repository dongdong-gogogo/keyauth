package rpc_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"gitee.com/dongdong-0421/keyauth/apps/token"

	"github.com/stretchr/testify/assert"

	"gitee.com/dongdong-0421/keyauth/client/rpc"
	mcenter "github.com/infraboard/mcenter/client/rpc"
)

func TestBookQuery(t *testing.T) {
	should := assert.New(t)

	conf := mcenter.NewDefaultConfig()
	conf.Address = os.Getenv("MCENTER_ADDRESS")
	conf.ClientID = os.Getenv("MCENTER_CDMB_CLINET_ID")
	conf.ClientSecret = os.Getenv("MCENTER_CMDB_CLIENT_SECRET")

	// 把Mcenter的配置传递给Keyauth客户端
	c, err := rpc.NewClient(conf)
	// 使用SDK，调用keyauth进行 凭证的校验
	//c.Token().ValidateToken()
	if should.NoError(err) {
		resp, err := c.Token().ValidateToken(
			context.Background(),
			token.NewValidateTokenRequest(""),
		)
		should.NoError(err)
		fmt.Println(resp)
	}
}

func init() {
	// 提前加载好 rpc客户端, resolver需要使用
	err := mcenter.LoadClientFromEnv()
	if err != nil {
		panic(err)
	}
}
