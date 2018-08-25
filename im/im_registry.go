package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/printfcoder/goutils/stringutils"
	"io/ioutil"
	"net/http"
)

// RegisterUsers 注册新人员
func (c *client) RegisterUsers(users []User) (ret []RegisterUserRsp, err error) {

	// 参数构造
	data, _ := json.Marshal(users)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.im.jpush.cn/v1/users", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 创建 注册新人员 请求失败, err: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")

	sign := stringutils.ToBase64(c.AppKey + ":" + c.MasterSecret)

	req.Header.Add("Authorization", "Basic "+sign)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送 注册新人员 请求失败, err: %s", err)
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送 注册新人员 请求返回的body无法解析, err: %s", err)
	}

	// 解析-JSON
	ret = make([]RegisterUserRsp, 0)
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送 注册新人员 请求返回的JSON无法解析, err: %s", err)
	}

	return
}

// RegisterAdmin 注册管理员
// 极光注册管理员成功返回的消息是空的，需要通过判断状态为201即可
func (c *client) RegisterAdmin(admin User) (ret *RegisterUserRsp, err error) {

	// 参数构造
	data, _ := json.Marshal(admin)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.im.jpush.cn/v1/admins/", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, fmt.Errorf("[RegisterAdmin] 创建 注册管理员 请求失败, err: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")

	sign := stringutils.ToBase64(c.AppKey + ":" + c.MasterSecret)

	req.Header.Add("Authorization", "Basic "+sign)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求失败, err: %s", err)
	}
	defer rsp.Body.Close()

	fmt.Printf(rsp.Status)

	if rsp.Status == "201 Created" {
		return nil, nil
	}

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回的body无法解析, err: %s", err)
	}

	// 解析-JSON
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		return nil, fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回的JSON无法解析, err: %s", err)
	}

	if ret.Error.Code != 0 {
		return nil, fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回错误, err: %s", ret.Error.Message)
	}

	return
}
