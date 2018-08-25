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
	c, err := im.Init(ini)
	assert.Nil(t, err)

	var users = []im.User{{UserName: "asdfw3dfas23sdf", Password: "asdfw3dfas23sdf2"}}

	rsp, err := c.RegisterUsers(users)
	assert.Nil(t, err)

	assert.Nil(t, rsp[0].Error)
}

// Test_RegisterAdmin 测试注册管理员
func Test_RegisterAdmin(t *testing.T) {
	ini := im.InitParams{
		AppKey:       "3ac491b5b80577bb48503e12",
		MasterSecret: "388c2a202c6368a9186e26db",
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	var user = im.User{UserName: "asdfw3dfa98ad12", Password: "asdfw3dfas23sdf2"}

	_, err = c.RegisterAdmin(user)
	assert.Nil(t, err)

}
