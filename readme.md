# 极光推送Golang SDK

## 安装

### jpush

```bash
$ go get -u github.com/printfcoder/jpush
```

### goutils

```bash
$ go get -u https://github.com/printfcoder/goutils/...
```

## 即时消息

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

    // 构造用户账密
    rsp, err := c.GetAdminsListByAppKey(0, 5)
    if err != nil {
      // ......
    }

    for i, v := range rsp.Users {
    	t.Logf("[Test_GetAdminsListByAppKey] 第%d个：%s", i, v.UserName)
    }

```


[RestAPI]: https://docs.jiguang.cn/jmessage/server/rest_api_im/