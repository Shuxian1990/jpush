# 极光推送

## IM 即时消息

官方 [RestAPI][RestAPI]

### 初始化

将**AppKey**与**MasterKey**放到初始化参数中：

```golang

    ini := im.InitParams{
    		AppKey:       "你的appKey",
    		MasterSecret: "你的MasterSecret",
    	}
    c, err := im.Init(ini)

```

### 注册普通用户

```golang

    // ...

    // 构造用户账密
    var users = []im.User{{UserName: "asdfw3dfas23sdf", Password: "asdfw3dfas23sdf2"}}

    rsp, err := c.RegisterUsers(users)
    if err != nil {
      // ......
    }

```

### 注册管理员

```golang

    // 构造用户账密
    var user = im.User{UserName: "asdfw3dfa98ad12", Password: "asdfw3dfas23sdf2"}

    rsp, err := c.RegisterAdmin(user)
    if err != nil {
      // ......
    }

```

### 管理员列表

```golang

    // ...

    // 构造用户账密
    rsp, err := c.GetAdminsListByAppKey(0, 5)
    if err != nil {
      // ......
    }

    for i, v := range rsp.Users {
    	t.Logf("[Test_GetAdminsListByAppKey] 第%d个：%s", i, v.UserName)
    }

```

### 获取指定用户

```golang

    // ...

    // 构造用户账密
    rsp, err := c.GetUser("asdfw3dfas8ad12")
    if err != nil {
      // ......
    }

    t.Logf("[Test_GetUser] 用户(asdfw3dfas8ad12)，创建时间：%s", rsp.CTime)

```

### 更新指定用户

```golang

    // ...

   	err = c.UpdateUser(im.User{UserName: "asdfw3dfa98ad12", Nickname: "小三"})
   	assert.Nil(t, err)

   	rsp, err := c.GetUser("asdfw3dfa98ad12")
   	assert.Nil(t, err)

   	t.Logf("[Test_GetUser] 用户(asdfw3dfas8ad12)，更新时间：%s，更新nickName: %s", rsp.CTime, rsp.Nickname)

```


[RestAPI]: https://docs.jiguang.cn/jmessage/server/rest_api_im/
