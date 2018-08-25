package im

var errMap = map[string]string{
	"1":      "请求失败",
	"899003": "参数错误，Request Body参数不符合要求",
	"99001":  "用户已存在",
}

var (
	// ErrCreateReqFail 创建请求失败
	ErrCreateReqFail = 1

	// ErrSendReqFail 发送请求失败
	ErrSendReqFail = 2

	// ErrReadRspFail 读取响应失败
	ErrReadRspFail = 3

	// ErrJSONUnmarshalFail JSON解析失败
	ErrJSONUnmarshalFail = 4

	// ErrErrorJSONUnmarshalFail errorJSON解析失败
	ErrErrorJSONUnmarshalFail = 5
)
