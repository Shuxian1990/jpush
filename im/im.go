package im

// Init 初始化
func Init(parms InitParams) (client Client, err error) {
	return initClient(parms)
}
