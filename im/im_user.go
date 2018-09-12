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
	errN = c.do(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", user.UserName), "PUT", "更新用户", data, nil, nil)

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
	errN = c.do("https://api.im.jpush.cn/v1/users/userstat", "POST", "获取用户状态 批量", data, &ret, nil)

	return
}

// ToBlackList 添加黑名单
func (c *client) ToBlackList(userName string, blacklist []string) (errN *JustError) {

	data, _ := json.Marshal(blacklist)

	rspError := make([]*BlacklistErrRsp, 0)
	err := c.do(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s/blacklist", userName), "PUT", "添加黑名单", data, nil, &rspError)
	if err != nil {
		errN = &JustError{ErrorRsp: &ErrorRsp{}}
		errN.ErrorRsp.Error = err
	}

	if len(rspError) != 0 {
		errN = &JustError{}
		errN.BlacklistErrRsp = rspError
	}
	return
}

// DeleteBlackList 移除黑名单
func (c *client) DeleteBlackList(userName string, blacklist []string) (errN *JustError) {

	data, _ := json.Marshal(blacklist)
	rspError := make([]*BlacklistErrRsp, 0)
	err := c.do(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s/blacklist", userName), "DELETE", "移除黑名单", data, nil, &rspError)
	if err != nil {
		errN = &JustError{ErrorRsp: &ErrorRsp{}}
		errN.ErrorRsp.Error = err
	}

	if len(rspError) != 0 {
		errN = &JustError{}
		errN.BlacklistErrRsp = rspError
	}

	return
}

// GetBlackList 黑名单列表
func (c *client) GetBlackList(userName string, blacklist []string) (errN *common.Error) {

	data, _ := json.Marshal(blacklist)
	rspError := make([]*BlacklistErrRsp, 0)
	errN = c.do(fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s/blacklist", userName), "GET", "黑名单列表", data, nil, &rspError)

	return
}
