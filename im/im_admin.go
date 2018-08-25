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
func (c *client) RegisterUsers(users []User) (ret []RegisterUserRsp, errN *Error) {

	// 参数构造
	data, _ := json.Marshal(users)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.im.jpush.cn/v1/users", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {
		errN = &Error{
			Message: fmt.Errorf("[RegisterUsers] 创建 注册新人员 请求失败, err: %s", err).Error(),
			Code:    ErrCreateReqFail,
		}

		return nil, errN
	}
	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterUsers] 发送 注册新人员 发送请求失败, err: %s", err).Error(),
			Code:    ErrSendReqFail,
		}
		return nil, errN
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterUsers] 发送 注册新人员 请求返回的body无法解析, err: %s", err).Error(),
			Code:    ErrReadRspFail,
		}
		return nil, errN
	}

	// 解析-JSON
	ret = make([]RegisterUserRsp, 0)
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterUsers] 发送 注册新人员 请求返回的JSON无法解析, err: %s", err).Error(),
			Code:    ErrJSONUnmarshalFail,
		}
		return nil, errN

	}

	return
}

// RegisterAdmin 注册管理员
// 极光注册管理员成功返回的消息是空的，需要通过判断状态为201即可
func (c *client) RegisterAdmin(admin User) (ret *RegisterUserRsp, errN *Error) {

	// 参数构造
	data, _ := json.Marshal(admin)

	// 创建请求
	req, err := http.NewRequest("POST", "https://api.im.jpush.cn/v1/admins/", ioutil.NopCloser(bytes.NewReader(data)))
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterAdmin] 创建 注册管理员 请求失败, err: %s", err).Error(),
			Code:    ErrCreateReqFail,
		}
		return nil, errN
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterAdmin] 创建 注册管理员 发送请求失败, err: %s", err).Error(),
			Code:    ErrSendReqFail,
		}
		return nil, errN
	}
	defer rsp.Body.Close()

	fmt.Printf(rsp.Status)

	if rsp.Status == "201 Created" {
		return nil, nil
	}

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回的body无法解析, err: %s", err).Error(),
			Code:    ErrReadRspFail,
		}

		return nil, errN
	}

	if stringutils.StartsWith(rsp.Status, "2") {
		// 解析-JSON
		err = json.Unmarshal(rspBody, &ret)
		if err != nil {

			errN = &Error{
				Message: fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回的JSON无法解析, err: %s", err).Error(),
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
				Message: fmt.Errorf("[RegisterAdmin] 发送 注册管理员 请求返回错误, err: %s", err).Error(),
				Code:    ErrJSONUnmarshalFail,
			}
			return nil, errN
		}

		return nil, &errorRsp.Error
	}

	return
}

// GetAdminsListByAppKey 获取管理员列表
func (c *client) GetAdminsListByAppKey(start, count int) (ret *PageUserRsp, errN *Error) {

	// 创建请求
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.im.jpush.cn/v1/admins?start=%d&count=%d", start, count), nil)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetAdminsListByAppKey] 创建 获取管理员列表 请求失败, err: %s", err).Error(),
			Code:    ErrCreateReqFail,
		}
		return nil, errN
	}

	c.addAuthToHeader(&req.Header)

	// 发送请求
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetAdminsListByAppKey] 创建 获取管理员列表 发送请求失败, err: %s", err).Error(),
			Code:    ErrSendReqFail,
		}
		return nil, errN
	}
	defer rsp.Body.Close()

	// 解析-body
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetAdminsListByAppKey] 创建 获取管理员列表 请求返回的body无法解析, err: %s", err).Error(),
			Code:    ErrReadRspFail,
		}
		return nil, errN
	}

	// 解析-JSON
	err = json.Unmarshal(rspBody, &ret)
	if err != nil {

		errN = &Error{
			Message: fmt.Errorf("[GetAdminsListByAppKey] 创建 获取管理员列表 请求返回的JSON无法解析, err: %s", err).Error(),
			Code:    ErrJSONUnmarshalFail,
		}
		return nil, errN
	}

	if ret.Error.Code != 0 {
		errN = &Error{
			Message: fmt.Errorf("[GetAdminsListByAppKey] 发送 获取管理员列表 请求返回错误, err: %s", err).Error(),
			Code:    ret.Error.Code,
		}
		return nil, errN
	}

	return

}
