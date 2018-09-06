package push

import "github.com/printfcoder/jpush/common"

// MsgObj 消息结构 必填
type MsgObj struct {
	CID          string        `json:"cid,omitempty"`
	Platform     interface{}   `json:"platform"`
	Audience     interface{}   `json:"audience"`
	Notification *Notification `json:"notification,omitempty"`
	Message      *Message      `json:"message,omitempty"`
	SMSMessage   *SMSMessage   `json:"sms_message,omitempty"`
	Options      *Options      `json:"options,omitempty"`
}

// Audience 推送设备指定 必填
type Audience struct {
	Tag            []string `json:"tag,omitempty"`
	TagAnd         []string `json:"tag_and,omitempty"`
	TagNot         []string `json:"tag_not,omitempty"`
	Alias          []string `json:"alias,omitempty"`
	RegistrationID []string `json:"registration_id,omitempty"`
	Segment        string   `json:"segment,omitempty"`
	ABtTest        string   `json:"abtest,omitempty"`
}

// Notification  通知内容体 可选 与message二选一
type Notification struct {
	Android NotificationAndroid `json:"android,omitempty"`
	IOS     NotificationIOS     `json:"ios,omitempty"`
}

// NotificationAndroid 内容结构平台 ANDROID
type NotificationAndroid struct {
	Alert      string      `json:"alert"`
	Title      string      `json:"title,omitempty"`
	BuilderID  int         `json:"builder_id,omitempty"`
	Priority   int         `json:"priority,omitempty"`
	Category   string      `json:"category,omitempty"`
	Style      int         `json:"style,omitempty"`
	AlertType  int         `json:"alert_type,omitempty"`
	BigText    string      `json:"big_text,omitempty"`
	Inbox      interface{} `json:"inbox,omitempty"`
	BigPicPath string      `json:"big_pic_path,omitempty"`
	Extras     interface{} `json:"extras,omitempty"`
}

// NotificationIOS 内容结构平台 IOS
type NotificationIOS struct {
	Alert            string      `json:"alert"`
	Sound            string      `json:"sound,omitempty"`
	Badge            int         `json:"badge,omitempty"`
	ContentAvailable bool        `json:"content-available,omitempty"`
	MutableContent   bool        `json:"mutable-content,omitempty"`
	Category         string      `json:"style,omitempty"`
	Extras           interface{} `json:"extras,omitempty"`
}

// Message 消息结构
type Message struct {
	MsgContent  string      `json:"msg_content"`
	ContentType string      `json:"content_type"`
	Title       string      `json:"title"`
	Extras      interface{} `json:"extras"`
}

// SMSMessage struct
type SMSMessage struct {
	TempID    string      `json:"temp_id"`
	TempPara  interface{} `json:"temp_para"`
	DelayTime int         `json:"delay_time"`
}

// Options 选项
type Options struct {
	SendNo          int    `json:"sendno"`
	TimeToLive      int    `json:"time_to_live"`
	OverrideMsgID   int64  `json:"override_msg_id"`
	APNSProduction  bool   `json:"apns_production"`
	APNSCollapseID  string `json:"apns_collapse_id"`
	BigPushDuration int    `json:"big_push_duration"`
}

// ErrorRsp JPush返回的错误结构
type ErrorRsp struct {
	Error common.Error `json:"error,omitempty"`
}

// CIDList cid 列表
type CIDList struct {
	CIDList []string `json:"cidlist,omitempty"`
}

// Rsp JPush返回的结构
type Rsp struct {
	SendNo string `json:"sendno"`
	MsgID  string `json:"msg_id"`
}
