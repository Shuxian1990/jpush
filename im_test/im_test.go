package im_test

import (
	"testing"

	"github.com/printfcoder/jpush/im"
	"github.com/stretchr/testify/assert"
)

// Test_RegisterUser 测试注册用户
func Test_RegisterUser(t *testing.T) {
	ini := im.InitParams{
		AppID:        "你的appID",
		AppSecret:    "你的AppSecret",
		AppKey:       "你的appKey",
		MasterSecret: "你的MasterSecret",
	}
	c := im.Init(ini)

	var users = []im.User{{UserName: "asdfw3dfas23sdf", Password: "asdfw3dfas23sdf2"}}

	rsp, err := c.RegisterUsers(users)
	assert.Nil(t, err)

	assert.Nil(t, rsp[0].Error)
}
