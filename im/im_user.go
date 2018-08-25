package im

import (
	"encoding/json"
	"fmt"
)

// GetUser 获取用户
func (c *client) GetUser(userName string) (ret *User, errN *Error) {

	ret = &User{}
	errN = c.get(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", userName), "获取用户", ret)

	return
}

// UpdateUser 更新用户
func (c *client) UpdateUser(user User) (errN *Error) {

	data, _ := json.Marshal(user)
	errN = c.putOrPost(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", user.UserName), "更新用户", data, nil)

	return
}

// GetUserStat 获取用户状态
func (c *client) GetUserStat(userName string) (ret *UserStat, errN *Error) {

	ret = &UserStat{}
	errN = c.get(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s/userstat", userName), "获取用户状态", ret)

	return
}
