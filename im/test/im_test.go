package test

import (
	"github.com/printfcoder/jpush/common"
	"testing"

	"github.com/printfcoder/jpush/im"
	"github.com/stretchr/testify/assert"
)

var (
	AppKey       = "3ac491b5b80577bb48503e12"
	MasterSecret = "388c2a202c6368a9186e26db"
)

// Test_RegisterUser 测试注册用户
func Test_RegisterUser(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	var users = []im.User{{UserName: "asdfw3dfas23sdf", Password: "asdfw3dfas23sdf2", Nickname: "asdfw3dfas23sdf"}}

	_, err = c.RegisterUsers(users)
	assert.Nil(t, err)

}

// Test_RegisterAdmin 测试注册管理员
func Test_RegisterAdmin(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	var user = im.User{UserName: "asdfw3dfa98ad92", Password: "asdfw3dfas23sdf2"}

	_, err = c.RegisterAdmin(user)
	assert.Nil(t, err)

}

// Test_GetAdminsListByAppKey 测试管理员列表
func Test_GetAdminsListByAppKey(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	rsp, err := c.GetAdminsListByAppKey(0, 5)
	assert.Nil(t, err)

	for i, v := range rsp.Users {
		t.Logf("[Test_GetAdminsListByAppKey] 第%d个：%s", i, v.UserName)
	}

}

// Test_GetUser 测试获取人员
func Test_GetUser(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	rsp, err := c.GetUser("asdfw3dfa98ad12")
	assert.Nil(t, err)

	t.Logf("[Test_GetUser] 用户(asdfw3dfas8ad12)，创建时间：%s", rsp.CTime)
}

// Test_Update 测试更新人员
func Test_UpdateUser(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	err = c.UpdateUser(im.User{UserName: "asdfw3dfa98ad12", Nickname: "小三三"})
	assert.Nil(t, err)

	rsp, err := c.GetUser("asdfw3dfa98ad12")
	assert.Nil(t, err)

	t.Logf("[Test_GetUser] 用户(asdfw3dfas8ad12)，更新时间：%s，更新nickName: %s", rsp.CTime, rsp.Nickname)
}

// Test_GetUserStat 测试获取用户状态
func Test_GetUserStat(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}
	c, err := im.Init(ini)
	assert.Nil(t, err)

	userStat, err := c.GetUserStat("asdfw3dfa98ad12")
	assert.Nil(t, err)

	t.Logf("[Test_GetUser] 用户(asdfw3dfas8ad12)，已经登录：%t，在线: %t", userStat.Login, userStat.Online)
}
