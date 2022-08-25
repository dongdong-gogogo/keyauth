package utils_test

import (
	"testing"

	"gitee.com/dongdong-0421/keyauth/utils"
)

func TestToken(t *testing.T) {
	v := utils.MakeBearer(24)
	t.Log(v)
}
