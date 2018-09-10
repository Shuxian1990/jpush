package im

import (
	"encoding/json"
	"fmt"
	"github.com/printfcoder/jpush/common"
)

// GetUser 获取用户
func (c *client) GetUser(userName string) (ret *User, errN *common.Error) {

	ret = &User{}
	errN = c.get(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", userName), "获取用户", ret)

	return
}

// UpdateUser 更新用户
func (c *client) UpdateUser(user User) (errN *common.Error) {

	data, _ := json.Marshal(user)
	errN = c.putOrPost(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", user.UserName), "put", "更新用户", data, nil)

	return
}

// GetUserStat 获取用户状态
func (c *client) GetUserStat(userName string) (ret *UserStat, errN *common.Error) {

	ret = &UserStat{}
	errN = c.get(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s/userstat", userName), "获取用户状态", ret)

	return
}

// GetUsersStat 获取用户状态 批量
func (c *client) GetUsersStat(userNames []string) (ret []*UsersState, errN *common.Error) {

	data, _ := json.Marshal(userNames)
	errN = c.putOrPost("https://api.im.jpush.cn/v1/users/userstat", "POST", "获取用户状态 批量", data, &ret)

	return
}
