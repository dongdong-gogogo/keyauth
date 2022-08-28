package utils_test

import (
	"testing"

	"github.com/dongdong-gogogo/keyauth/utils"
)

func TestToken(t *testing.T) {
	v := utils.MakeBearer(24)
	t.Log(v)
}
