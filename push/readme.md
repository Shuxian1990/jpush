# 极光推送

## PUSH

官方 [PushAPI][PushAPI]

### 初始化

将**AppKey**与**MasterKey**放到初始化参数中：

```golang

    ini := im.InitParams{
    		AppKey:       "你的appKey",
    		MasterSecret: "你的MasterSecret",
    	}
    c, err := im.Init(ini)

```

### 获取CID 推送唯一标识符

```golang

    // ...

    ret, err := c.GetCID(4)
    if err != nil {
      // ......
    }

   for i, s := range ret.CIDList {
   		t.Logf("[Test_GetCID] 第%d个CID:%s", i, s)
   }

```

[PushAPI]: https://docs.jiguang.cn/jpush/server/push/rest_api_v3_push/