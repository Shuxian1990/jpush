package im

import (
	"github.com/printfcoder/goutils/stringutils"
	"net/http"
	"time"
)

// Client 客户端接口
type Client interface {
	RegisterUsers(users []User) ([]RegisterUserRsp, *Error)
	RegisterAdmin(admin User) (*RegisterUserRsp, *Error)
	GetAdminsListByAppKey(start, count int) (*PageUserRsp, *Error)
	GetUser(userName string) (*User, *Error)
	UpdateUser(user User) *Error
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
func initClient(parms InitParams) (c Client, err error) {
	if single == nil {
		single = new(client)
		single.AppID = parms.AppID
		single.AppSecret = parms.AppSecret
		single.AppKey = parms.AppKey
		single.MasterSecret = parms.MasterSecret
		single.AuthHeartbeat = parms.AuthHeartbeat
	}

	return single, nil
}

func (c *client) addAuthToHeader(header *http.Header) {
	header.Add("Content-Type", "application/json")
	sign := stringutils.ToBase64(c.AppKey + ":" + c.MasterSecret)
	header.Add("Authorization", "Basic "+sign)
}
