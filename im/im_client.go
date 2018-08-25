package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/printfcoder/goutils/stringutils"
	"io/ioutil"
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
	GetUserStat(userName string) (*UserStat, *Error)
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

func (c *client) get(url, funcMsg string, ret interface{}) (errN *Error) {
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[get] 创建 %s 请求失败, err: %s", funcMsg, err).Error(),
			Code:    ErrCreateReqFail,
		}
		return
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[get] 发送 %s 发送请求失败, err: %s", funcMsg, err).Error(),
			Code:    ErrSendReqFail,
		}
		return
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[get] 发送 %s 请求返回的body无法解析, err: %s", funcMsg, err).Error(),
			Code:    ErrReadRspFail,
		}
		return
	}

	if stringutils.StartsWith(rsp.Status, "2") {
		// 解析-JSON
		err = json.Unmarshal(rspBody, &ret)
		if err != nil {

			errN = &Error{
				Message: fmt.Errorf("[get] 发送 %s 请求返回的JSON无法解析, err: %s", funcMsg, err).Error(),
				Code:    ErrJSONUnmarshalFail,
			}
			return errN
		}
	} else {

		// 解析-JSON
		var errorRsp ErrorRsp
		err = json.Unmarshal(rspBody, &errorRsp)
		if err != nil {
			errN = &Error{
				Message: fmt.Errorf("[get] 发送 %s 请求返回的Error无法解析, err: %s", funcMsg, err).Error(),
				Code:    ErrErrorJSONUnmarshalFail,
			}
			return errN
		}

		return &errorRsp.Error
	}

	return
}

func (c *client) putOrPost(url, funcMsg string, body []byte, ret interface{}) (errN *Error) {

	// 创建请求
	req, err := http.NewRequest("PUT", url, ioutil.NopCloser(bytes.NewReader(body)))
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[putOrPost] 创建 %s 请求失败, err: %s", funcMsg, err).Error(),
			Code:    ErrCreateReqFail,
		}
		return
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[putOrPost] 发送 %s 发送请求失败, err: %s", funcMsg, err).Error(),
			Code:    ErrSendReqFail,
		}
		return
	}
	defer rsp.Body.Close()

	if stringutils.StartsWith(rsp.Status, "2") {
		return nil
	} else {

		// 解析-body
		rspBody, err := ioutil.ReadAll(rsp.Body)
		if err != nil {

			errN = &Error{
				Message: fmt.Errorf("[putOrPost] 发送 %s 请求返回的body无法解析, err: %s", funcMsg, err).Error(),
				Code:    ErrReadRspFail,
			}
			return
		}

		// 解析-JSON
		var errorRsp ErrorRsp
		err = json.Unmarshal(rspBody, &errorRsp)
		if err != nil {
			errN = &Error{
				Message: fmt.Errorf("[putOrPost] 发送 %s 请求返回的Error无法解析, err: %s", funcMsg, err).Error(),
				Code:    ErrErrorJSONUnmarshalFail,
			}
			return
		}

		return &errorRsp.Error
	}

	return

}
