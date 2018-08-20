package im

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
}

// RegisterUserRsp 注册返回的结构
type RegisterUserRsp struct {
	UserName string `json:"username"`
	Error    Error  `json:"error"`
}

// Error error
type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
