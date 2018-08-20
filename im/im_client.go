package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/printfcoder/goutils/stringutils"
)

// Client 客户端接口
type Client interface {
	RegisterUsers(users []User) ([]RegisterUserRsp, error)
}

// InitParams 初始化参数
type InitParams struct {
	AppID        string
	AppSecret    string
	AppKey       string
	MasterSecret string
	// AuthHeartbeat Auth刷新时间 单位小时 默认20小时
	AuthHeartbeat time.Duration
}

type client struct {
	InitParams
}

var single *client

// initClient 客户端-单例
func initClient(parms InitParams) Client {
	if single == nil {
		single = new(client)
		single.AppID = parms.AppID
		single.AppSecret = parms.AppSecret
		single.AppKey = parms.AppKey
		single.MasterSecret = parms.MasterSecret
		single.AuthHeartbeat = parms.AuthHeartbeat
	}

	return single
}

// RegisterUsers 注册新人员
func (c *client) RegisterUsers(users []User) (ret []RegisterUserRsp, err error) {

	// 参数构造
	data, _ := json.Marshal(users)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.im.jpush.cn/v1/users", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 创建auth请求失败, err: %s", err)
	}
	req.Header.Add("Content-Type", "application/json")

	sign := stringutils.ToBase64(c.AppKey + ":" + c.MasterSecret)

	req.Header.Add("Authorization", "Basic "+sign)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送auth请求失败, err: %s", err)
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送auth请求返回的body无法解析, err: %s", err)
	}

	// 解析-JSON
	ret = make([]RegisterUserRsp, 0)
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		return nil, fmt.Errorf("[RegisterUsers] 发送auth请求返回的JSON无法解析, err: %s", err)
	}

	return
}
