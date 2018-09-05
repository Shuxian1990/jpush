package test

import (
	"github.com/printfcoder/jpush/common"
	"github.com/printfcoder/jpush/push"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	AppKey       = ""
	MasterSecret = ""
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

// Test_Push_All 测试推送 所有人
func Test_Push_All(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}

	c, err := push.Init(ini)
	assert.Nil(t, err)

	msgObj := push.MsgObj{
		Platform: []string{"android", "ios"},
		Audience: "all",
		Notification: &push.Notification{
			Android: push.NotificationAndroid{
				Alert: "I am the King",
				Title: "You can so you do",
			},
		},
	}

	ret, err := c.Push(msgObj)
	assert.Nil(t, err)

	t.Logf("[Test_Push_Single] msgID: %s, sendNo: %s", ret.MsgID, ret.SendNo)

}

// Test_Push_Single 测试推送 所有人
func Test_Push_Single(t *testing.T) {
	ini := common.InitParams{
		AppKey:       AppKey,
		MasterSecret: MasterSecret,
	}

	c, err := push.Init(ini)
	assert.Nil(t, err)

	msgObj := push.MsgObj{
		Platform: []string{"android", "ios"},
		Audience: push.Audience{
			RegistrationID: []string{"160a3797c853e44cd30"},
		},
		Notification: &push.Notification{
			Android: push.NotificationAndroid{
				Alert: "I am the King2",
				Title: "You can so you do2",
			},
		},
	}

	ret, err := c.Push(msgObj)
	assert.Nil(t, err)

	t.Logf("[Test_Push_Single] msgID: %s, sendNo: %s", ret.MsgID, ret.SendNo)

}
