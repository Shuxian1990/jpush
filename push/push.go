package push

import "github.com/printfcoder/jpush/common"

// Init 初始化
func Init(params common.InitParams) (client Client, err error) {
	return initClient(params)
}
