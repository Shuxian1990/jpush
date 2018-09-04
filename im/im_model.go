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
	Login  bool `json:"login"`
	Online bool `json:"online"`
}

// RegisterUserRsp 注册返回的结构
type RegisterUserRsp struct {
	UserName string `json:"username"`
}

// ErrorRsp JPush返回的错误结构
type ErrorRsp struct {
	Error common.Error `json:"error,omitempty"`
}

// PageUserRsp 用户分页结构
type PageUserRsp struct {
	Total int          `json:"total"`
	Start int          `json:"start"`
	Count int          `json:"count"`
	Users []User       `json:"users"`
	Error common.Error `json:"error"`
}
