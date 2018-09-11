package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/printfcoder/goutils/stringutils"
	"github.com/printfcoder/jpush/common"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Client 客户端接口
type Client interface {
	RegisterUsers(users []User) ([]RegisterUserRsp, *common.Error)
	RegisterAdmin(admin User) (*RegisterUserRsp, *common.Error)
	GetAdminsListByAppKey(start, count int) (*PageUserRsp, *common.Error)
	GetUser(userName string) (*User, *common.Error)
	UpdateUser(user User) *common.Error
	GetUserStat(userName string) (*UserStat, *common.Error)
	GetUsersStat(userNames []string) ([]*UsersState, *common.Error)
	ToBlackList(userName string, blacklist []string) (errN *JustError)
	DeleteBlackList(userName string, blacklist []string) (errN *JustError)
}

type client struct {
	common.Client
}

var single *client

// initClient 客户端-单例
func initClient(params common.InitParams) (c Client, err error) {
	if single == nil {
		single = new(client)
		single.AppID = params.AppID
		single.AppSecret = params.AppSecret
		single.AppKey = params.AppKey
		single.MasterSecret = params.MasterSecret
		single.AuthHeartbeat = params.AuthHeartbeat
	}

	return single, nil
}

func (c *client) get(url, funcMsg string, ret interface{}) (errN *common.Error) {
	// 创建请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[get] 创建 %s 请求失败, err: %s", funcMsg, err).Error(),
			Code:    common.ErrCreateReqFail,
		}
		return
	}

	c.AddAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[get] 发送 %s 发送请求失败, err: %s", funcMsg, err).Error(),
			Code:    common.ErrSendReqFail,
		}
		return
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[get] 发送 %s 请求返回的body无法解析, err: %s", funcMsg, err).Error(),
			Code:    common.ErrReadRspFail,
		}
		return
	}

	if stringutils.StartsWith(rsp.Status, "2") {
		// 解析-JSON
		err = json.Unmarshal(rspBody, &ret)
		if err != nil {

			errN = &common.Error{
				Message: fmt.Errorf("[get] 发送 %s 请求返回的JSON无法解析, err: %s", funcMsg, err).Error(),
				Code:    common.ErrJSONUnmarshalFail,
			}
			return errN
		}
	} else {

		// 解析-JSON
		var errorRsp ErrorRsp
		err = json.Unmarshal(rspBody, &errorRsp)
		if err != nil {
			errN = &common.Error{
				Message: fmt.Errorf("[get] 发送 %s 请求返回的Error无法解析, err: %s", funcMsg, err).Error(),
				Code:    common.ErrErrorJSONUnmarshalFail,
			}
			return errN
		}

		return &errorRsp.Error
	}

	return
}

func (c *client) do(url, method, funcMsg string, body []byte, ret interface{}, rspErr interface{}) (errN *common.Error) {

	// 创建请求
	req, err := http.NewRequest(method, url, ioutil.NopCloser(bytes.NewReader(body)))
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[do] 创建 %s 请求失败, err: %s", funcMsg, err).Error(),
			Code:    common.ErrCreateReqFail,
		}
		return
	}

	c.AddAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[do] 发送 %s 发送请求失败, err: %s", funcMsg, err).Error(),
			Code:    common.ErrSendReqFail,
		}
		return
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &common.Error{
			Message: fmt.Errorf("[do] 发送 %s 请求返回的body无法解析, err: %s", funcMsg, err).Error(),
			Code:    common.ErrReadRspFail,
		}
		return
	}

	if stringutils.StartsWith(rsp.Status, "2") {

		if ret != nil {
			// 解析-JSON
			err = json.Unmarshal(rspBody, &ret)
			if err != nil {

				errN = &common.Error{
					Message: fmt.Errorf("[do] 发送 %s 请求返回的JSON无法解析, err: %s", funcMsg, err).Error(),
					Code:    common.ErrJSONUnmarshalFail,
				}
				return errN
			}
		}

		return nil
	} else {

		// 需要自定义错误类型的执行另外的解析错误逻辑
		if rspErr == nil || (reflect.ValueOf(rspErr).Kind() == reflect.Ptr && reflect.ValueOf(rspErr).IsNil()) {
			// 解析-JSON
			var errorRsp ErrorRsp
			err = json.Unmarshal(rspBody, &errorRsp)
			if err != nil {
				errN = &common.Error{
					Message: fmt.Errorf("[do] 发送 %s 请求返回的Error无法解析, err: %s", funcMsg, err).Error(),
					Code:    common.ErrErrorJSONUnmarshalFail,
				}
				return
			}

			return &errorRsp.Error
		} else {

			err = json.Unmarshal(rspBody, rspErr)
			if err != nil {
				errN = &common.Error{
					Message: fmt.Errorf("[do] 发送 %s 请求返回的自定义Error无法解析, err: %s", funcMsg, err).Error(),
					Code:    common.ErrCustomErrorJSONUnmarshalFail,
				}
				return
			}
		}

	}

	return

}
