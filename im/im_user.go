package im

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/printfcoder/goutils/stringutils"
	"io/ioutil"
	"net/http"
)

// GetUser 获取用户
func (c *client) GetUser(userName string) (ret *User, errN *Error) {

	// 创建请求
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", userName), nil)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetUser] 创建 获取用户 请求失败, err: %s", err).Error(),
			Code:    ErrCreateReqFail,
		}
		return nil, errN
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetUser] 发送 获取用户 发送请求失败, err: %s", err).Error(),
			Code:    ErrSendReqFail,
		}
		return nil, errN
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetUser] 发送 获取用户 请求返回的body无法解析, err: %s", err).Error(),
			Code:    ErrReadRspFail,
		}
		return nil, errN
	}

	if stringutils.StartsWith(rsp.Status, "2") {
		// 解析-JSON
		err = json.Unmarshal(rspBody, &ret)
		if err != nil {

			errN = &Error{
				Message: fmt.Errorf("[GetUser] 发送 获取用户 请求返回的JSON无法解析, err: %s", err).Error(),
				Code:    ErrJSONUnmarshalFail,
			}
			return nil, errN
		}
	} else {

		// 解析-JSON
		var errorRsp ErrorRsp
		err = json.Unmarshal(rspBody, &errorRsp)
		if err != nil {
			errN = &Error{
				Message: fmt.Errorf("[GetUser] 发送 获取用户 请求返回的Error无法解析, err: %s", err).Error(),
				Code:    ErrErrorJSONUnmarshalFail,
			}
			return nil, errN
		}

		return nil, &errorRsp.Error
	}

	return

}

// UpdateUser 更新用户
func (c *client) UpdateUser(user User) (errN *Error) {

	// 参数构造
	data, _ := json.Marshal(user)

	// 创建请求
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.im.jpush.cn/v1/users/%s", user.UserName), ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[UpdateUser] 创建 更新用户 请求失败, err: %s", err).Error(),
			Code:    ErrCreateReqFail,
		}
		return
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[UpdateUser] 发送 更新用户 发送请求失败, err: %s", err).Error(),
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
				Message: fmt.Errorf("[UpdateUser] 发送 更新用户 请求返回的body无法解析, err: %s", err).Error(),
				Code:    ErrReadRspFail,
			}
			return
		}

		// 解析-JSON
		var errorRsp ErrorRsp
		err = json.Unmarshal(rspBody, &errorRsp)
		if err != nil {
			errN = &Error{
				Message: fmt.Errorf("[UpdateUser] 发送 更新用户 请求返回的Error无法解析, err: %s", err).Error(),
				Code:    ErrErrorJSONUnmarshalFail,
			}
			return
		}

		return &errorRsp.Error
	}

	return

}
