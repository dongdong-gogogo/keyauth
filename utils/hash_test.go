package utils_test

import (
	"testing"

	"github.com/dongdong-gogogo/keyauth/utils"
)

func TestHash(t *testing.T) {
	str := "123456"
	hash := utils.Hash(str)
	t.Log(hash)
}

func TestHashPassword(t *testing.T) {
	v := utils.HashPassword("123456")
	t.Log(v)
	ok := utils.CheckPasswordHash("123456", "$2a$14$QFW4O5kn3gIuLmLPUxL8b.b5i3.INinKevfB1OUaeUxfnHf5qvlja")
	t.Log(ok)
}
