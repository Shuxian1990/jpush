package im

import "github.com/printfcoder/jpush/common"

// User im人员
type User struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	Nickname  string `json:"nickname,omitempty"`
	Birthday  string `json:"avatar,omitempty"`
	Signature string `json:"signature,omitempty"`
	Gender    int    `json:"gender,omitempty"` // 0 - 未知， 1 - 男 ，2 - 女
	Region    string `json:"region,omitempty"`
	Address   string `json:"address,omitempty"`
	Extras    string `json:"extras,omitempty"`
	MTime     string `json:"mtime,omitempty"`
	CTime     string `json:"ctime,omitempty"`
}

// UserStat 用户状态
type UserStat struct {
	Login    bool   `json:"login"`
	Online   bool   `json:"online"`
	Platform string `json:"platform"`
}

// UsersState 批量用户状态
type UsersState struct {
	UserName string     `json:"userName"`
	Devices  []UserStat `json:"devices"`
}

// RegisterUserRsp 注册返回的结构
type RegisterUserRsp struct {
	UserName string `json:"username"`
}

// ErrorRsp JPush返回的错误结构
type ErrorRsp struct {
	Error common.Error `json:"error,omitempty"`
}

// BlacklistErrRsp 黑名单返回的错误结构
type BlacklistErrRsp struct {
	UserName string `json:"username,omitempty"`
	ErrorRsp
}

// JustError 就是有错误，可能是请求本身有问题，可能是响应的参数错误，
// 可能是数据查询出有不符合逻辑的（比如重复加入黑名单）
type JustError struct {
	BlacklistErrRsp []*BlacklistErrRsp
	*ErrorRsp
}

// PageUserRsp 用户分页结构
type PageUserRsp struct {
	Total int          `json:"total"`
	Start int          `json:"start"`
	Count int          `json:"count"`
	Users []User       `json:"users"`
	Error common.Error `json:"error"`
}
