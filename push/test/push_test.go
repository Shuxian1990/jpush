package test

import (
	"github.com/printfcoder/jpush/common"
	"github.com/printfcoder/jpush/push"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	AppKey       = "3ac491b5b80577bb48503e12"
	MasterSecret = "388c2a202c6368a9186e26db"
)

// Test_GetCID 测试获取CID
func Test_GetCID(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}

	c, err := push.Init(ini)
	assert.Nil(t, err)
	ret, err := c.GetCID(4)
	assert.Nil(t, err)

	assert.Equal(t, len(ret.CIDList), 4)

	for i, s := range ret.CIDList {
		t.Logf("[Test_GetCID] 第%d个CID:%s", i, s)
	}
}
