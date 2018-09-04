package common

import (
	"time"
)

// InitParams 初始化参数
type InitParams struct {
	AppID        string
	AppSecret    string
	AppKey       string
	MasterSecret string
	// AuthHeartbeat Auth刷新时间 单位小时 默认20小时
	AuthHeartbeat time.Duration
}

// Error error
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}
