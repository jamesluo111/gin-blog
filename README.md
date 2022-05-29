**本篇是根据煎鱼大佬的go-gin-example写的，感谢煎鱼大佬，其项目地址为 [eddycjy/go-gin-example: An example of gin (github.com)](https://github.com/EDDYCJY/go-gin-example)** 

# Gin学习笔记

### 一、Go介绍与环境配置

#### 1.安装Go

**centos**

下载go环境

```
$ wget https://studygolang.com/dl/golang/go1.13.1.linux-amd64.tar.gz

$ tar -zxvf go1.13.1.linux-amd64.tar.gz

$ mv go/ /usr/local/
```

配置环境变量 

```
#打开环境变量配置文件
$ vi /etc/profile

#添加环境变量 GOROOT 和将 GOBIN 添加到 PATH 中
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin 

#配置完毕执行命令使其生效
$ source /etc/profile

#在控制台输入go version，若输出版本号则安装成功，如下：
$ go version
go version go1.13.1 linux/amd64
```



**MacOS**

在 MacOS 上安装 Go 最方便的办法就是使用 brew，安装如下：

```
$ brew install go

#升级
$ brew upgrade go
```



#### 2.初始化

首先你需要有一个你喜欢的目录，例如：`$ mkdir ~/go-application && cd ~/go-application`，然后执行如下命令：

```

$ mkdir go-gin-example && cd go-gin-example

$ go env -w GO111MODULE=on

$ go env -w GOPROXY=https://goproxy.cn,direct

$ go mod init github.com/EDDYCJY/go-gin-example
go: creating new go.mod: module github.com/EDDYCJY/go-gin-example

$ ls
go.mod
```

- `go env -w GO111MODULE=on`：打开 Go modules 开关（目前在 Go1.13 中默认值为 `auto`）。
- `go env -w GOPROXY=...`：设置 GOPROXY 代理，这里主要涉及到两个值，第一个是 `https://goproxy.cn`，它是由七牛云背书的一个强大稳定的 Go 模块代理，可以有效地解决你的外网问题；第二个是 `direct`，它是一个特殊的 fallback 选项，它的作用是用于指示 Go 在拉取模块时遇到错误会回源到模块版本的源地址去抓取（比如 GitHub 等）。
- `go mod init [MODULE_PATH]`：初始化 Go modules，它将会生成 go.mod 文件，需要注意的是 `MODULE_PATH` 填写的是模块引入路径，你可以根据自己的情况修改路径。



在执行了上述步骤后，初始化工作已完成，我们打开 `go.mod` 文件看看，如下：

```
module github.com/jamesluo111/go-gin-example

go 1.17
```



#### 3.基础使用

用 `go get` 拉取新的依赖 

- 拉取最新的版本(优先择取 tag)：`go get golang.org/x/text@latest`
- 拉取 `master` 分支的最新 commit：`go get golang.org/x/text@master`
- 拉取 tag 为 v0.3.2 的 commit：`go get golang.org/x/text@v0.3.2`
- 拉取 hash 为 342b231 的 commit，最终会被转换为 v0.3.2：`go get golang.org/x/text@342b2e`
- 用 `go get -u` 更新现有的依赖
- 用 `go mod download` 下载 go.mod 文件中指明的所有依赖
- 用 `go mod tidy` 整理现有的依赖
- 用 `go mod graph` 查看现有的依赖结构
- 用 `go mod init` 生成 go.mod 文件 (Go 1.17 中唯一一个可以生成 go.mod 文件的子命令)



#### 4.gin是什么

Gin 是用 Go 开发的一个微框架，类似 Martinier 的 API，重点是小巧、易用、性能好很多，也因为 **httprouter**[3] 的性能提高了 40 倍。



#### 5.gin安装

我们回到刚刚创建的 `go-gin-example` 目录下，在命令行下执行如下命令：

```
$ go get -u github.com/gin-gonic/gin
go: downloading github.com/gin-gonic/gin v1.7.7
go: downloading github.com/gin-contrib/sse v0.1.0
go: downloading github.com/mattn/go-isatty v0.0.12
go: downloading github.com/golang/protobuf v1.3.3
go: downloading github.com/ugorji/go/codec v1.1.7
go: downloading gopkg.in/yaml.v2 v2.2.8
go: downloading github.com/json-iterator/go v1.1.9
go: downloading github.com/go-playground/validator/v10 v10.4.1
go: downloading github.com/ugorji/go v1.1.7
...
```



go.sum

这时候你再检查一下该目录下，会发现多了个 `go.sum` 文件，如下：

```
github.com/creack/pty v1.1.9/go.mod h1:oKZEueFk5CKHvIhNR5MUki03XCEU+Q6VDXinZuGJ33E=
github.com/davecgh/go-spew v1.1.0/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/davecgh/go-spew v1.1.1/go.mod h1:J7Y8YcW2NihsgmVo/mv3lAwl/skON4iLHjSsI+c5H38=
github.com/gin-contrib/sse v0.1.0 h1:Y/yl/+YNO8GZSjAhjMsSuLt29uWRFHdHYUb5lYOV9qE=
github.com/gin-contrib/sse v0.1.0/go.mod h1:RHrZQHXnP2xjPF+u1gW/2HnVO7nvIa9PG3Gm+fLHvGI=
github.com/gin-gonic/gin v1.7.7 h1:3DoBmSbJbZAWqXJC3SLjAPfutPJJRN1U5pALB7EeTTs=
github.com/gin-gonic/gin v1.7.7/go.mod h1:axIBovoeJpVj8S3BwE0uPMTeReE4+AfFtqpqaZ1qq1U=
github.com/go-playground/assert/v2 v2.0.1/go.mod h1:VDjEfimB/XKnb+ZQfWdccd7VUvScMdVu0Titje2rxJ4=
github.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=
...
```



go.mod

既然我们下载了依赖包，`go.mod` 文件会不会有所改变呢，我们再去看看，如下：

```
module github.com/jamesluo111/go-gin-example

go 1.17

require (
    github.com/gin-contrib/sse v0.1.0 // indirect
    github.com/gin-gonic/gin v1.7.7 // indirect
    github.com/go-playground/locales v0.14.0 // indirect
    github.com/go-playground/universal-translator v0.18.0 // indirect
    github.com/go-playground/validator/v10 v10.11.0 // indirect
    github.com/golang/protobuf v1.5.2 // indirect
    github.com/json-iterator/go v1.1.12 // indirect
    github.com/leodido/go-urn v1.2.1 // indirect
    github.com/mattn/go-isatty v0.0.14 // indirect
    github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
    github.com/modern-go/reflect2 v1.0.2 // indirect
    github.com/ugorji/go/codec v1.2.7 // indirect
    golang.org/x/crypto v0.0.0-20220511200225-c6db032c6c88 // indirect
    golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
    golang.org/x/text v0.3.7 // indirect
    google.golang.org/protobuf v1.28.0 // indirect
    gopkg.in/yaml.v2 v2.4.0 // indirect
)
```

确确实实发生了改变，那多出来的东西又是什么呢，`go.mod` 文件又保存了什么信息呢，实际上 `go.mod` 文件是启用了 Go modules 的项目所必须的最重要的文件，因为它描述了当前项目（也就是当前模块）的元信息，每一行都以一个动词开头，目前有以下 5 个动词:

- module：用于定义当前项目的模块路径。
- go：用于设置预期的 Go 版本。
- require：用于设置一个特定的模块版本。
- exclude：用于从使用中排除一个特定的模块版本。
- replace：用于将一个模块版本替换为另外一个模块版本。

你可能还会疑惑 `indirect` 是什么东西，`indirect` 的意思是传递依赖，也就是非直接依赖。



#### 6.测试

编写一个`test.go`文件

```
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}
```

执行`test.go`

```
$ go run test.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> main.main.func1 (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Environment variable PORT is undefined. Using port :8080 by default
[GIN-debug] Listening and serving HTTP on :8080
```

访问 `$HOST:8080/ping`，若返回`{"message":"pong"}`则正确

```
curl 127.0.0.1:8080/ping
```



#### 7.整理

刚刚在执行了命令 `$ go get -u github.com/gin-gonic/gin` 后，我们查看了 `go.mod` 文件，如下：

```
...
require (
    github.com/gin-contrib/sse v0.1.0 // indirect
    github.com/gin-gonic/gin v1.7.7 // indirect
    ...
)
```

你会发现 `go.mod` 里的 `github.com/gin-gonic/gin` 是 `indirect` 模式，这显然不对啊，因为我们的应用程序已经实际的编写了 gin server 代码了，我就想把它调对，怎么办呢，在应用根目录下执行如下命令：

```
$ go mod tidy
```

该命令主要的作用是整理现有的依赖，非常的常用，执行后 `go.mod` 文件内容为：

```

require github.com/gin-gonic/gin v1.7.7

require (
    github.com/gin-contrib/sse v0.1.0 // indirect
    github.com/go-playground/locales v0.14.0 // indirect
```

可以看到 `github.com/gin-gonic/gin` 已经变成了直接依赖，调整完毕。



### 二、初始化项目及公共库

首先，在一个初始项目开始前，大家都要思考一下

- 程序的文本配置写在代码中，好吗？
- API 的错误码硬编码在程序中，合适吗？
- db 句柄谁都去`Open`，没有统一管理，好吗？
- 获取分页等公共参数，谁都自己写一套逻辑，好吗？

显然在较正规的项目中，这些问题的答案都是**不可以**，为了解决这些问题，我们挑选一款读写配置文件的库，目前比较火的有 viper，有兴趣你未来可以简单了解一下，没兴趣的话等以后接触到再说。

但是本系列选用 go-ini/ini ，它的 中文文档。大家是必须需要要简单阅读它的文档，再接着完成后面的内容。

#### 1.初始化项目目录

在前一章节中，我们初始化了一个 `go-gin-example` 项目，接下来我们需要继续新增如下目录结构：

```
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime
```

- conf：用于存储配置文件
- middleware：应用中间件
- models：应用数据库模型
- pkg：第三方包
- routers 路由逻辑处理
- runtime：应用运行时数据



#### 2.添加 Go Modules Replace

打开 `go.mod` 文件，新增 `replace` 配置项，如下：

```
replace (
        //本机测试在windows，所以路径为windows下的绝对路径
        github.com/EDDYCJY/go-gin-example/pkg/setting => D:/dnmp/wwwroot/go-gin-example/pkg/setting
        github.com/EDDYCJY/go-gin-example/conf        => D:/dnmp/wwwroot/go-gin-example/pkg/conf
        github.com/EDDYCJY/go-gin-example/middleware  => D:/dnmp/wwwroot/go-gin-example/middleware
        github.com/EDDYCJY/go-gin-example/models      => D:/dnmp/wwwroot/go-gin-example/models
        github.com/EDDYCJY/go-gin-example/routers     => D:/dnmp/wwwroot/go-gin-example/routers
)
```

可能你会不理解为什么要特意跑来加 `replace` 配置项，首先你要看到我们使用的是完整的外部模块引用路径（`github.com/jamesluo111/go-gin-example/xxx`），而这个模块还没推送到远程，是没有办法下载下来的，因此需要用 `replace` 将其指定读取本地的模块路径，这样子就可以解决本地模块读取的问题。

**注：后续每新增一个本地应用目录，你都需要主动去 go.mod 文件里新增一条 replace（我不会提醒你），如果你漏了，那么编译时会出现报错，找不到那个模块。**



#### 3.初始化数据库

新建 `blog` 数据库，编码为`utf8_general_ci`，在 `blog` 数据库下，新建以下表

1.标签表

```
CREATE TABLE `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '标签名称',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT '创建时间',
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用、1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签管理';
```

2.文章表

```
CREATE TABLE `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
  `title` varchar(100) DEFAULT '' COMMENT '文章标题',
  `desc` varchar(255) DEFAULT '' COMMENT '简述',
  `content` text,
  `created_on` int(11) DEFAULT NULL,
  `created_by` varchar(100) DEFAULT '' COMMENT '创建人',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT '修改时间',
  `modified_by` varchar(255) DEFAULT '' COMMENT '修改人',
  `deleted_on` int(10) unsigned DEFAULT '0',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
```

3.认证表

```
CREATE TABLE `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `blog`.`blog_auth` (`id`, `username`, `password`) VALUES (null, 'test', 'test123456');
```



#### 4.编写项目配置包

在 `go-gin-example` 应用目录下，拉取 `go-ini/ini` 的依赖包，如下：

```
$ go get -u github.com/go-ini/ini
go: finding github.com/go-ini/ini v1.48.0
go: downloading github.com/go-ini/ini v1.48.0
go: extracting github.com/go-ini/ini v1.48.0
```



接下来我们需要编写基础的应用配置文件，在 `go-gin-example` 的`conf`目录下新建`app.ini`文件，写入内容：

```
#debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 23347$040412

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = root
PASSWORD = root
#127.0.0.1:3306
HOST = 172.27.108.187:3306
NAME = blog
TABLE_PREFIX = blog_
```

建立调用配置的`setting`模块，在`go-gin-example`的`pkg`目录下新建`setting`目录（注意新增 replace 配置），新建 `setting.go` 文件，写入内容：

```
package setting

import (
    "github.com/go-ini/ini"
    "log"
    "time"
)

var (
    //配置文件
    Cfg *ini.File

    //项目环境
    RunModel string

    //项目端口号
    HttpPort int
    //最大读取时间
    ReadTimeOut time.Duration
    //最大写入时间
    WriteTimeOut time.Duration

    //分页数
    PageSize int

    //jwt密钥
    JwtSecret string
)

func init() {
    var err error
    Cfg, err = ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    LoadBase()
    LoadServer()
    LoadApp()
}

func LoadBase() {
    RunModel = Cfg.Section("").Key("RUN_MODE").String()
}

func LoadServer() {
    sec, err := Cfg.GetSection("server")
    if err != nil {
        log.Fatalf("Fail to get section 'server': %v", err)
    }
    HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
    ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
    WriteTimeOut = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
    sec, err := Cfg.GetSection("app")
    if err != nil {
        log.Fatalf("Fail to get section 'app': %v", err)
    }
    PageSize = sec.Key("PAGE_SIZE").MustInt(10)
    JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
}

```

当前的目录结构：

```

go-gin-example
├── conf
│   └── app.ini
├── go.mod
├── go.sum
├── middleware
├── models
├── pkg
│   └── setting.go
├── routers
└── runtime
```



#### 5.编写api错误码包

建立错误码的`e`模块，在`go-gin-example`的`pkg`目录下新建`e`目录（注意新增 replace 配置），新建`code.go`和`msg.go`文件，写入内容：

1.code.go

```
package e

const (
    SUCCESS        = 200
    ERROR          = 500
    INVALID_PARAMS = 400

    ERROR_EXIST_TAG         = 10001
    ERROR_NOT_EXIST_TAG     = 10002
    ERROR_NOT_EXIST_ARTICLE = 10003

    ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
    ERROR_AUTH_TOKEN               = 20003
    ERROR_AUTH                     = 20004
)
```



2.msg.go

```
package e

var MsgFlags = map[int]string{
    SUCCESS:                        "ok",
    ERROR:                          "fail",
    INVALID_PARAMS:                 "请求参数错误",
    ERROR_EXIST_TAG:                "已存在该标签名称",
    ERROR_NOT_EXIST_TAG:            "该标签不存在",
    ERROR_NOT_EXIST_ARTICLE:        "该文章不存在",
    ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
    ERROR_AUTH_TOKEN:               "Token生成失败",
    ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[code]
}

```



#### 6.编写工具包

在`go-gin-example`的`pkg`目录下新建`util`目录（注意新增 replace 配置），并拉取`com`的依赖包，如下：

```
$ go get -u github.com/unknwon/com
```



#### 7.编写分页页码的获取方法

在`util`目录下新建`pagination.go`，写入内容：

```
package util

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
    "github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
    result := 0
    page, _ := com.StrTo(c.Query("page")).Int()
    if page > 0 {
        result = (page - 1) * setting.PageSize
    }
    return result
}

```



#### 8.编写 models init

拉取`gorm`的依赖包，如下：

```
$ go get -u github.com/jinzhu/gorm
```

拉取`mysql`驱动的依赖包，如下：

```
$ go get -u github.com/go-sql-driver/mysql
```

完成后，在`go-gin-example`的`models`目录下新建`models.go`，用于`models`的初始化使用

```
package models

import (
    "fmt"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
    "github.com/jinzhu/gorm"
    "log"
)

var db *gorm.DB

type Model struct {
    Id         int `gorm:"primary_key" json:"id"`
    CreatedOn  int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init() {
    var (
        err                                               error
        dbType, dbName, user, password, host, tablePrefix string
    )
    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    dbType = sec.Key("TYPE").String()
    dbName = sec.Key("NAME").String()
    user = sec.Key("USER").String()
    password = sec.Key("PASSWORD").String()
    host = sec.Key("HOST").String()
    tablePrefix = sec.Key("TABLE_PREFIX").String()
    db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", user, password, host, dbName))
    if err != nil {
        log.Println(err)
    }

    //设置数据库前缀
    gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
        return tablePrefix + defaultTableName
    }

    db.SingularTable(true)
    db.LogMode(true)
    db.DB().SetMaxIdleConns(10)
    db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
    defer db.Close()
}

```



#### 9.编写项目启动、路由文件

**编写 Demo**

在`go-gin-example`下建立`main.go`作为启动文件（也就是`main`包），我们先写个**Demo**，帮助大家理解，写入文件内容：

```
package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
    "net/http"
)

func main() {
    router := gin.Default()
    router.GET("/test", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "message": "test",
        })
    })

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HttpPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeOut,
        WriteTimeout:   setting.WriteTimeOut,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}

```

执行`go run main.go`，查看命令行是否显示

```
PS D:\dnmp\wwwroot\go-gin-example> go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> main.main.func1 (3 handlers)
```

在本机执行`curl 127.0.0.1:8000/test`，检查是否返回`{"message":"test"}`。



#### 10.知识点

**标准库**

- fmt：实现了类似 C 语言 printf 和 scanf 的格式化 I/O。格式化动作（'verb'）源自 C 语言但更简单
- net/http：提供了 HTTP 客户端和服务端的实现

**Gin**

- gin.Default()：返回 Gin 的`type Engine struct{...}`，里面包含`RouterGroup`，相当于创建一个路由`Handlers`，可以后期绑定各类的路由规则和函数、中间件等
- router.GET(...){...}：创建不同的 HTTP 方法绑定到`Handlers`中，也支持 POST、PUT、DELETE、PATCH、OPTIONS、HEAD 等常用的 Restful 方法
- gin.H{...}：就是一个`map[string]interface{}`
- gin.Context：`Context`是`gin`中的上下文，它允许我们在中间件之间传递变量、管理流、验证 JSON 请求、响应 JSON 请求等，在`gin`中包含大量`Context`的方法，例如我们常用的`DefaultQuery`、`Query`、`DefaultPostForm`、`PostForm`等等

**&http.Server 和 ListenAndServe？**

1、http.Server：

```
type Server struct {
    Addr    string
    Handler Handler
    TLSConfig *tls.Config
    ReadTimeout time.Duration
    ReadHeaderTimeout time.Duration
    WriteTimeout time.Duration
    IdleTimeout time.Duration
    MaxHeaderBytes int
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}
```

- Addr：监听的 TCP 地址，格式为`:8000`
- Handler：http 句柄，实质为`ServeHTTP`，用于处理程序响应 HTTP 请求
- TLSConfig：安全传输层协议（TLS）的配置
- ReadTimeout：允许读取的最大时间
- ReadHeaderTimeout：允许读取请求头的最大时间
- WriteTimeout：允许写入的最大时间
- IdleTimeout：等待的最大时间
- MaxHeaderBytes：请求头的最大字节数
- ConnState：指定一个可选的回调函数，当客户端连接发生变化时调用
- ErrorLog：指定一个可选的日志记录器，用于接收程序的意外行为和底层系统错误；如果未设置或为`nil`则默认以日志包的标准日志记录器完成（也就是在控制台输出）

2、 ListenAndServe：

```
func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}
```

开始监听服务，监听 TCP 网络地址，Addr 和调用应用程序处理连接上的请求。

我们在源码中看到`Addr`是调用我们在`&http.Server`中设置的参数，因此我们在设置时要用`&`，我们要改变参数的值，因为我们`ListenAndServe`和其他一些方法需要用到`&http.Server`中的参数，他们是相互影响的。

3、 `http.ListenAndServe`和 `r.Run()`有区别吗？

我们看看`r.Run`的实现：

```
func (engine *Engine) Run(addr ...string) (err error) {
    defer func() { debugPrintError(err) }()

    address := resolveAddress(addr)
    debugPrint("Listening and serving HTTP on %s\n", address)
    err = http.ListenAndServe(address, engine)
    return
}
```

通过分析源码，得知**本质上没有区别**，同时也得知了启动`gin`时的监听 debug 信息在这里输出

4、 为什么 Demo 里会有`WARNING`？

首先我们可以看下`Default()`的实现

```

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
    debugPrintWARNINGDefault()
    engine := New()
    engine.Use(Logger(), Recovery())
    return engine
}
```

大家可以看到默认情况下，已经附加了日志、恢复中间件的引擎实例。并且在开头调用了`debugPrintWARNINGDefault()`，而它的实现就是输出该行日志

```

func debugPrintWARNINGDefault() {
    debugPrint(`[WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
`)
}
```

而另外一个`Running in "debug" mode. Switch to "release" mode in production.`，是运行模式原因，并不难理解，已在配置文件的管控下 :-)，运维人员随时就可以修改它的配置。

5、 Demo 的`router.GET`等路由规则可以不写在`main`包中吗？

我们发现`router.GET`等路由规则，在 Demo 中被编写在了`main`包中，感觉很奇怪，我们去抽离这部分逻辑！

在`go-gin-example`下`routers`目录新建`router.go`文件，写入内容：

```
package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunModel)

    r.GET("/test", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "message": "success",
        })
    })

    return r
}

```

修改`main.go`的文件内容：

```
package main

import (
    "fmt"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
    "github.com/jamesluo111/go-gin-example/routers"
    "net/http"
)

func main() {
    router := routers.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HttpPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeOut,
        WriteTimeout:   setting.WriteTimeOut,
        MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}

```

当前目录结构：

```
go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   └── models.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   └── router.go
├── runtime
```

重启服务，执行 `curl 127.0.0.1:8000/test`查看是否正确返回。



### 三、开发标签模块

#### 1.定义接口

本节正是编写标签的逻辑，我们想一想，一般接口为增删改查是基础的，那么我们定义一下接口吧！

- 获取标签列表：GET("/tags")
- 新建标签：POST("/tags")
- 更新指定标签：PUT("/tags/:id")
- 删除指定标签：DELETE("/tags/:id")



#### 2.编写路由空壳

 开始编写路由文件逻辑，在`routers`下新建`api`目录，我们当前是第一个 API 大版本，因此在`api`下新建`v1`目录，再新建`tag.go`文件，写入内容： 

```
package v1

import (
    "github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
```



#### 3.注册路由

 我们打开`routers`下的`router.go`文件，修改文件内容为： 

```
package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/go-gin-example/pkg/setting"
    v1 "github.com/jamesluo111/go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunModel)

    r.GET("/test", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "message": "success",
        })
    })

    apivi := r.Group("/api/v1")
    {
        apivi.GET("/tags", v1.GetTag)
        apivi.POST("/tags", v1.AddTag)
        apivi.PUT("/tags/:id", v1.EditTag)
        apivi.DELETE("/tags/:id", v1.DeleteTag)
    }

    return r
}

```

 当前目录结构： 

```
gin-blog/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   └── models.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       └── tag.go
│   └── router.go
├── runtime
```



#### 4.检验路由是否注册成功

 回到命令行，执行`go run main.go`，检查路由规则是否注册成功。 

```
PS D:\dnmp\wwwroot\gin-blog> go run .\main.go
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release                                               
 - using code:  gin.SetMode(gin.ReleaseMode)                                          

[GIN-debug] GET    /test                     --> github.com/jamesluo111/gin-blog/routers.InitRouter.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/jamesluo111/gin-blog/routers/api/v1.GetTag (3 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/jamesluo111/gin-blog/routers/api/v1.AddTag (3 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/jamesluo111/gin-blog/routers/api/v1.EditTag (3 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/jamesluo111/gin-blog/routers/api/v1.DeleteTag (3 handlers)
```



#### 5.下载依赖包

 首先我们要拉取`validation`的依赖包，在后面的接口里会使用到表单验证 

```
$ go get -u github.com/astaxie/beego/validation
```



#### 6.编写标签列表的 models 逻辑

 创建`models`目录下的`tag.go`，写入文件内容： 

```
package models

type Tag struct {
    Model

    Name string `json:"name"`
    CreatedBy string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State int `json:"state"`
}

func GetTags(pageNum int, pageSize int, maps interface {}) (tags []Tag) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

    return
}

func GetTagTotal(maps interface {}) (count int){
    db.Model(&Tag{}).Where(maps).Count(&count)

    return
}
```

1. 我们创建了一个`Tag struct{}`，用于`Gorm`的使用。并给予了附属属性`json`，这样子在`c.JSON`的时候就会自动转换格式，非常的便利
2. 可能会有的初学者看到`return`，而后面没有跟着变量，会不理解；其实你可以看到在函数末端，我们已经显示声明了返回值，这个变量在函数体内也可以直接使用，因为他在一开始就被声明了
3. 有人会疑惑`db`是哪里来的；因为在同个`models`包下，因此`db *gorm.DB`是可以直接使用的



#### 7.编写标签列表的路由逻辑

打开`routers`目录下 v1 版本的`tag.go`，第一我们先编写**获取标签列表的接口**

修改文件内容：

```
package v1

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/models"
    "github.com/jamesluo111/gin-blog/pkg/e"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "github.com/jamesluo111/gin-blog/pkg/util"
    "github.com/unknwon/com"
    "net/http"
)

//获取文章标签
func GetTag(c *gin.Context) {
    //获取标签名称
    name := c.Query("name")

    //请求条件
    maps := make(map[string]interface{})
    //返回数据
    data := make(map[string]interface{})

    if name != "" {
        maps["name"] = name
    }

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        maps["state"] = state
    }

    code := e.SUCCESS

    data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
    data["count"] = models.GetTagTotal(maps)

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

//新增文章标签
func AddTag(c *gin.Context) {

}

//编辑文章标签
func EditTag(c *gin.Context) {

}

//删除文章标签
func DeleteTag(c *gin.Context) {

}

```

1. `c.Query`可用于获取`?name=test&state=1`这类 URL 参数，而`c.DefaultQuery`则支持设置一个默认值
2. `code`变量使用了`e`模块的错误编码，这正是先前规划好的错误码，方便排错和识别记录
3. `util.GetPage`保证了各接口的`page`处理是一致的
4. `c *gin.Context`是`Gin`很重要的组成部分，可以理解为上下文，它允许我们在中间件之间传递变量、管理流、验证请求的 JSON 和呈现 JSON 响应

在本机执行`curl 127.0.0.1:8000/api/v1/tags`，正确的返回值为`{"code":200,"data":{"lists":[],"total":0},"msg":"ok"}`，若存在问题请结合 gin 结果进行拍错。

在获取标签列表接口中，我们可以根据`name`、`state`、`page`来筛选查询条件，分页的步长可通过`app.ini`进行配置，以`lists`、`total`的组合返回达到分页效果。



#### 8.编写新增标签的 models 逻辑

接下来我们编写**新增标签**的接口

打开`models`目录下的`tag.go`，修改文件（增加 2 个方法）：

```
func ExistTagByName(name string) bool {
    var tag Tag
    db.Select("id").Where("name = ?", name).First(&tag)
    if tag.Id > 0 {
        return true
    }
    return false
}

func AddTag(name string, state int, createdBy string) bool {
    db.Create(&Tag{
        Name:       name,
        State:      state,
        CreatedBy:  createdBy,
        ModifiedBy: createdBy,
    })
    return true
}
```



#### 9.编写新增标签的路由逻辑

 打开`routers`目录下的`tag.go`，修改文件（变动 AddTag 方法）： 

```
//新增文章标签
func AddTag(c *gin.Context) {
    name := c.Query("name")
    state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
    createdBy := c.Query("created_by")

    //验证
    valid := validation.Validation{}
    valid.Required(name, "name").Message("名称不能为空")
    valid.MaxSize(name, 100, "name").Message("名称最长为100个字符")
    valid.Required(createdBy, "created_by").Message("创建人不能为空")
    valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100个字符")
    valid.Range(state, 0, 1, "state").Message("状态在0和1之间")

    code := e.INVALID_PARAMS

    if !valid.HasErrors() {
        if !models.ExistTagByName(name) {
            code = e.SUCCESS
            models.AddTag(name, state, createdBy)
        } else {
            code = e.ERROR_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}
```

 用`Postman`用 POST 访问`http://127.0.0.1:8000/api/v1/tags?name=1&state=1&created_by=test`，查看`code`是否返回`200`及`blog_tag`表中是否有值，有值则正确。 



#### 10编写 models callbacks

 但是这个时候大家会发现，我明明新增了标签，但`created_on`居然没有值，那做修改标签的时候`modified_on`会不会也存在这个问题？ 

 为了解决这个问题，我们需要打开`models`目录下的`tag.go`文件，修改文件内容（修改包引用和增加 2 个方法）： 

```
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())
    return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}
```

 重启服务，再在用`Postman`用 POST 访问`http://127.0.0.1:8000/api/v1/tags?name=2&state=1&created_by=test`，发现`created_on`已经有值了！ 

 **在这几段代码中，涉及到知识点：** 

 这属于`gorm`的`Callbacks`，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。 

 `gorm`所支持的回调方法： 

- 创建：BeforeSave、BeforeCreate、AfterCreate、AfterSave
- 更新：BeforeSave、BeforeUpdate、AfterUpdate、AfterSave
- 删除：BeforeDelete、AfterDelete
- 查询：AfterFind



#### 11.编写其余接口的路由逻辑

接下来，我们一口气把剩余的两个接口（EditTag、DeleteTag）完成吧

打开`routers`目录下 v1 版本的`tag.go`文件，修改内容：

```
//编辑文章标签
func EditTag(c *gin.Context) {
    name := c.Query("name")
    id := com.StrTo(c.Param("id")).MustInt()
    modifiedBy := c.Query("modified_by")

    valid := validation.Validation{}

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")
    }

    valid.Required(name, "name").Message("name不允许为空")
    valid.Required(id, "id").Message("id不允许为空")
    valid.Required(modifiedBy, "modifiedBy").Message("modifiedBy不允许为空")
    valid.MaxSize(name, 100, "name").Message("name最大字符小于100")
    valid.MaxSize(modifiedBy, 100, "modifiedBy").Message("modifiedBy最大字符小于100")

    code := e.INVALID_PARAMS

    if !valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistTagById(id) {
            data := make(map[string]interface{})
            data["name"] = name
            data["modified_by"] = modifiedBy
            if state != -1 {
                data["state"] = state
            }
            models.EditTag(id, data)
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}

//删除文章标签
func DeleteTag(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()

    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("id必须大于0")
    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        code = e.SUCCESS
        if models.ExistTagById(id) {
            models.DeleteTag(id)
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}
```



#### 12.编写其余接口的 models 逻辑

 打开`models`下的`tag.go`，修改文件内容： 

```
func ExistTagById(id int) bool {
    var tag Tag
    db.Select("id").Where("id = ?", id).First(&tag)
    if tag.Id > 0 {
        return true
    }
    return false
}

func EditTag(id int, data map[string]interface{}) bool {
    db.Model(&Tag{}).Where("id = ?", id).Updates(data)
    return true
}

func DeleteTag(id int) bool {
    db.Where("id = ?", id).Delete(&Tag{})
    return true
}
```

**验证功能**

重启服务，用 Postman

- PUT 访问 http://127.0.0.1:8000/api/v1/tags/1?name=edit1&state=0&modified_by=edit1 ，查看 code 是否返回 200
- DELETE 访问 http://127.0.0.1:8000/api/v1/tags/1 ，查看 code 是否返回 200

至此，Tag 的 API's 完成，下一节我们将开始 Article 的 API's 编写！



### 四、开发文章模块

#### 1.定义接口

- 获取文章列表：GET("/articles")
- 获取指定文章：POST("/articles/:id")
- 新建文章：POST("/articles")
- 更新指定文章：PUT("/articles/:id")
- 删除指定文章：DELETE("/articles/:id"



#### 2.编写路由逻辑

 在`routers`的 v1 版本下，新建`article.go`文件，写入内容： 

```
package v1

import "github.com/gin-gonic/gin"

//获取单个文章
func GetArticle(c *gin.Context) {

}

//获取文章列表
func GetArticles(c *gin.Context) {

}

//新增文章
func AddArticle(c *gin.Context) {

}

//修改文章
func EditArticle(c *gin.Context) {

}

//删除文章
func DeleteArticle(c *gin.Context) {

}

```

我们打开`routers`下的`router.go`文件，修改文件内容为：

```
        //文章模块
        apivi.GET("/articles", v1.GetArticles)
        apivi.POST("/articles/:id", v1.GetArticle)
        apivi.POST("/articles", v1.AddArticle)
        apivi.PUT("/articles/:id", v1.EditArticle)
        apivi.DELETE("/articles/:id", v1.DeleteArticle)
```

当前目录结构：

```

go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```



#### 3.编写 models 逻辑 

创建`models`目录下的`article.go`，写入文件内容：

```
package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Article struct {
    Model

    TagId int    `json:"tag_id" gorm:"index"`
    Tag   string `json:"tag"`

    Title      string `json:"title"`
    Desc       string `json:"desc"`
    Content    string `json:"content"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())
    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}

```

我们创建了一个`Article struct {}`，与`Tag`不同的是，`Article`多了几项，如下：

1. `gorm:index`，用于声明这个字段为索引，如果你使用了自动迁移功能则会有所影响，在不使用则无影响
2. `Tag`字段，实际是一个嵌套的`struct`，它利用`TagID`与`Tag`模型相互关联，在执行查询的时候，能够达到`Article`、`Tag`关联查询的功能
3. `time.Now().Unix()` 返回当前的时间戳

接下来，请确保已对上一章节的内容通读且了解，由于逻辑偏差不会太远，我们本节直接编写这五个接口

打开`models`目录下的`article.go`，修改文件内容：

```
package models

import (
    "github.com/jinzhu/gorm"
    "time"
)

type Article struct {
    Model

    TagId int `json:"tag_id" gorm:"index"`
    Tag   Tag `json:"tag"`

    Title      string `json:"title"`
    Desc       string `json:"desc"`
    Content    string `json:"content"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State      int    `json:"state"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedOn", time.Now().Unix())
    return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("ModifiedOn", time.Now().Unix())
    return nil
}

func ExistArticleById(id int) bool {
    var article Article
    db.Select("id").Where("id = ?", id).First(&article)
    if article.Id > 0 {
        return true
    }
    return false
}

func GetArticleTotal(maps interface{}) (count int) {
    db.Model(&Article{}).Where(maps).Count(count)
    return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
    return
}

func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).First(&article)
    db.Model(&article).Related(&article.Tag)
    return
}

func EditArticle(id int, data interface{}) bool {
    db.Model(&Article{}).Where("id = ?", id).Update(data)
    return true

}

func AddArticle(data map[string]interface{}) bool {
    db.Create(&Article{
        TagId:     data["tag_id"].(int),
        Title:     data["title"].(string),
        Desc:      data["desc"].(string),
        Content:   data["content"].(string),
        CreatedBy: data["created_by"].(string),
        State:     data["state"].(int),
    })

    return true
}

func DeleteArticle(id int) bool {
    db.Where("id = ?", id).Delete(Article{})

    return true
}

```

在这里，我们拿出三点不同来讲，如下：

**1、 我们的`Article`是如何关联到`Tag`？**

```

func GetArticle(id int) (article Article) {
    db.Where("id = ?", id).First(&article)
    db.Model(&article).Related(&article.Tag)

    return
}
```

能够达到关联，首先是`gorm`本身做了大量的约定俗成

- `Article`有一个结构体成员是`TagID`，就是外键。`gorm`会通过类名+ID 的方式去找到这两个类之间的关联关系
- `Article`有一个结构体成员是`Tag`，就是我们嵌套在`Article`里的`Tag`结构体，我们可以通过`Related`进行关联查询

**2、 `Preload`是什么东西，为什么查询可以得出每一项的关联`Tag`？**

```

func GetArticles(pageNum int, pageSize int, maps interface {}) (articles []Article) {
    db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

    return
}
```

`Preload`就是一个预加载器，它会执行两条 SQL，分别是`SELECT * FROM blog_articles;`和`SELECT * FROM blog_tag WHERE id IN (1,2,3,4);`，那么在查询出结构后，`gorm`内部处理对应的映射逻辑，将其填充到`Article`的`Tag`中，会特别方便，并且避免了循环查询

那么有没有别的办法呢，大致是两种

- `gorm`的`Join`
- 循环`Related`

**3、 `v.(I)` 是什么？**

`v`表示一个接口值，`I`表示接口类型。这个实际就是 Golang 中的**类型断言**，用于判断一个接口值的实际类型是否为某个类型，或一个非接口值的类型是否实现了某个接口类型



#### 4.路由逻辑

打开`routers`目录下 v1 版本的`article.go`文件，修改文件内容：

```
package v1

import (
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/models"
    "github.com/jamesluo111/gin-blog/pkg/e"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "github.com/jamesluo111/gin-blog/pkg/util"
    "github.com/unknwon/com"
    "log"
    "net/http"
)

//获取单个文章
func GetArticle(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()
    valid := validation.Validation{}

    valid.Min(id, 1, "id").Message("id最小值为1")

    code := e.INVALID_PARAMS
    var data interface{}
    if !valid.HasErrors() {
        if models.ExistArticleById(id) {
            data = models.GetArticle(id)
            code = e.SUCCESS
        } else {
            code = e.ERROR
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key:%s, err.Msg:%s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

//获取文章列表
func GetArticles(c *gin.Context) {
    data := make(map[string]interface{})
    maps := make(map[string]interface{})

    valid := validation.Validation{}

    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("state在0到1之间")
        maps["state"] = state
    }

    var tagId int = -1
    if arg := c.Query("tag_id"); arg != "" {
        tagId = com.StrTo(arg).MustInt()
        valid.Min(tagId, 1, "tagId").Message("tagId不能小于1")
        maps["tag_id"] = tagId
    }

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        code = e.SUCCESS
        data["list"] = models.GetArticles(util.GetPage(c), setting.PageSize, maps)
        data["count"] = models.GetArticleTotal(maps)
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key:%s, err.Msg:%s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

//新增文章
func AddArticle(c *gin.Context) {
    tagId := com.StrTo(c.Query("tag_id")).MustInt()
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    createdBy := c.Query("created_by")
    state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

    valid := validation.Validation{}

    valid.Min(tagId, 1, "tagId").Message("标签id最小为1")
    valid.Required(title, "title").Message("标题不能为空")
    valid.Required(desc, "desc").Message("描述不能为空")
    valid.Required(content, "content").Message("内容不能为空")
    valid.Required(createdBy, "createdBy").Message("创建人不能为空")
    valid.Range(state, 0, 1, "state").Message("状态必须在0和1之间")

    code := e.INVALID_PARAMS

    if !valid.HasErrors() {
        code = e.SUCCESS
        maps := make(map[string]interface{})
        maps["tag_id"] = tagId
        maps["title"] = title
        maps["desc"] = desc
        maps["content"] = content
        maps["created_by"] = createdBy
        maps["state"] = state
        models.AddArticle(maps)
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key:%s, err.Msg:%s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}

//修改文章
func EditArticle(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()
    tagId := com.StrTo(c.Query("tag_id")).MustInt()
    title := c.Query("title")
    desc := c.Query("desc")
    content := c.Query("content")
    createdBy := c.Query("created_by")

    valid := validation.Validation{}
    var state int = -1
    if arg := c.Query("state"); arg != "" {
        state = com.StrTo(arg).MustInt()
        valid.Range(state, 0, 1, "state").Message("state在0到1之间")
    }

    valid.Required(id, "id").Message("必须有文章id")
    valid.Required(tagId, "tagId").Message("必须有标签id")
    valid.Min(id, 1, "id").Message("文章id必须大于0")
    valid.MaxSize(title, 100, "title").Message("标题最长为100个字符")
    valid.MaxSize(desc, 100, "desc").Message("简述最长为255个字符")
    valid.MaxSize(content, 65535, "content").Message("文章内容最长为65535个字符")
    valid.Required(createdBy, "createdBy").Message("修改人不能为空")
    valid.MaxSize(createdBy, 100, "createdBy").Message("修改人最长为100个字符")

    code := e.INVALID_PARAMS

    if !valid.HasErrors() {
        if models.ExistTagById(tagId) {
            if models.ExistArticleById(id) {
                maps := make(map[string]interface{})
                if title != "" {
                    maps["title"] = title
                }
                if desc != "" {
                    maps["desc"] = desc
                }
                if content != "" {
                    maps["content"] = content
                }
                maps["createdBy"] = createdBy
                models.EditArticle(id, maps)
                code = e.SUCCESS
            } else {
                code = e.ERROR_NOT_EXIST_ARTICLE
            }
        } else {
            code = e.ERROR_NOT_EXIST_TAG
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key:%s, err.Msg:%s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })

}

//删除文章
func DeleteArticle(c *gin.Context) {
    id := com.StrTo(c.Param("id")).MustInt()

    valid := validation.Validation{}

    valid.Min(id, 1, "id").Message("id最小值大于1")

    code := e.INVALID_PARAMS
    if !valid.HasErrors() {
        if models.ExistArticleById(id) {
            models.DeleteArticle(id)
            code = e.SUCCESS
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        for _, err := range valid.Errors {
            log.Printf("err.key:%s, err.Msg:%s", err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": make(map[string]string),
    })
}

```

当前目录结构：

```

go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
├── models
│   ├── article.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       └── pagination.go
├── routers
│   ├── api
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```



#### 5.验证功能

我们重启服务，执行`go run main.go`，检查控制台输出结果

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in produ
ction.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /test                     --> github.com/jamesluo111/gin-blog/routers.InitRouter.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/jamesluo111/gin-blog/routers/api/v1.GetTag (3 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/jamesluo111/gin-blog/routers/api/v1.AddTag (3 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/jamesluo111/gin-blog/routers/api/v1.EditTag (3 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/jamesluo111/gin-blog/routers/api/v1.DeleteTag (3 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/jamesluo111/gin-blog/routers/api/v1.GetArticles (3 handlers)
[GIN-debug] POST   /api/v1/articles/:id      --> github.com/jamesluo111/gin-blog/routers/api/v1.GetArticle (3 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/jamesluo111/gin-blog/routers/api/v1.AddArticle (3 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/jamesluo111/gin-blog/routers/api/v1.EditArticle (3 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/jamesluo111/gin-blog/routers/api/v1.DeleteArticle (3 handlers)

```

使用`Postman`检验接口是否正常，在这里大家可以选用合适的参数传递方式，此处为了方便展示我选用了 GET/Param 传参的方式，而后期会改为 POST。

- POST：http://127.0.0.1:8000/api/v1/articles?tag_id=1&title=test1&desc=test-desc&content=test-content&created_by=test-created&state=1
- GET：http://127.0.0.1:8000/api/v1/articles
- GET：http://127.0.0.1:8000/api/v1/articles/1
- PUT：http://127.0.0.1:8000/api/v1/articles/1?tag_id=1&title=test-edit1&desc=test-desc-edit&content=test-content-edit&modified_by=test-created-edit&state=0
- DELETE：http://127.0.0.1:8000/api/v1/articles/1

至此，我们的 API's 编写就到这里，下一节我们将介绍另外的一些技巧！



### 五、使用 JWT 进行身份校验

在前面几节中，我们已经基本的完成了 API's 的编写，但是，还存在一些非常严重的问题，例如，我们现在的 API 是可以随意调用的，这显然还不安全全，在本文中我们通过 jwt-go （GoDoc）的方式来简单解决这个问题。

#### 1.下载依赖包

首先，我们下载 jwt-go 的依赖包，如下：

```
go get -u github.com/dgrijalva/jwt-go
```



#### 2.编写 jwt 工具包

我们需要编写一个`jwt`的工具包，我们在`pkg`下的`util`目录新建`jwt.go`，写入文件内容：

```
package util

import (
    "github.com/dgrijalva/jwt-go"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
    Username string `json:"username"`
    Password string `json:"password"`
    jwt.StandardClaims
}

//生成token
func GenerateToken(username, password string) (string, error) {
    nowTime := time.Now()
    //过期时间3小时
    expireTime := nowTime.Add(3 * time.Hour)

    claims := Claims{
        Username: username,
        Password: password,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expireTime.Unix(),
            Issuer:    "gin-blog",
        },
    }

    //生成token
    tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    //对token进行签发
    token, err := tokenClaims.SignedString(jwtSecret)
    return token, err
}

//对token解析
func ParseToken(token string) (*Claims, error) {
    //对签名进行反解析
    tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if tokenClaims != nil {
        //对token进行验证解析,最后生成Claims
        if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
            return claims, nil
        }
    }

    return nil, err
}

```

在这个工具包，我们涉及到

- `NewWithClaims(method SigningMethod, claims Claims)`，`method`对应着`SigningMethodHMAC struct{}`，其包含`SigningMethodHS256`、`SigningMethodHS384`、`SigningMethodHS512`三种`crypto.Hash`方案
- `func (t *Token) SignedString(key interface{})` 该方法内部生成签名字符串，再用于获取完整、已签名的`token`
- `func (p *Parser) ParseWithClaims` 用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回`*Token`
- `func (m MapClaims) Valid()` 验证基于时间的声明`exp, iat, nbf`，注意如果没有任何声明在令牌中，仍然会被认为是有效的。并且对于时区偏差没有计算方法

有了`jwt`工具包，接下来我们要编写要用于`Gin`的中间件，我们在`middleware`下新建`jwt`目录，新建`jwt.go`文件，写入内容：

```
package jwt

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/pkg/e"
    "github.com/jamesluo111/gin-blog/pkg/util"
    "net/http"
    "time"
)

func JWT() gin.HandlerFunc {
    return func(context *gin.Context) {
        var code int
        var data interface{}
        code = e.SUCCESS
        token := context.Query("token")
        if token == "" {
            code = e.INVALID_PARAMS
        } else {
            claims, err := util.ParseToken(token)

            if err != nil {
                code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
            } else if time.Now().Unix() > claims.ExpiresAt {
                code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
            }
        }

        if code != e.SUCCESS {
            context.JSON(http.StatusUnauthorized, gin.H{
                "code": code,
                "msg":  e.GetMsg(code),
                "data": data,
            })
            context.Abort()
            return
        }

        context.Next()
    }
}

```



#### 3.如何获取`Token`

那么我们如何调用它呢，我们还要获取`Token`呢？

1、 我们要新增一个获取`Token`的 API

在`models`下新建`auth.go`文件，写入内容：

```
package models

type Auth struct {
    Id       int    `gorm:"primary_key" json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
}

func CheckoutAuth(username, password string) bool {
    var auth Auth
    db.Select("id").Where("username = ? and password = ?", username, password).First(&auth)

    if auth.Id > 0 {
        return true
    }

    return false
}

```

在`routers`下的`api`目录新建`auth.go`文件，写入内容：

```
package api

import (
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/models"
    "github.com/jamesluo111/gin-blog/pkg/e"
    "github.com/jamesluo111/gin-blog/pkg/util"
    "log"
    "net/http"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    valid := validation.Validation{}
    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS

    if ok {
        isExist := models.CheckoutAuth(username, password)

        if isExist {
            token, err := util.GenerateToken(username, password)

            if err != nil {
                code = e.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token
                code = e.SUCCESS
            }
        } else {
            code = e.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            log.Println(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

```

我们打开`routers`目录下的`router.go`文件，修改文件内容（新增获取 token 的方法）：

```
package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "github.com/jamesluo111/gin-blog/routers/api"
    v1 "github.com/jamesluo111/gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunModel)

    //新增auth路由
    r.GET("/auth", api.GetAuth)

    apivi := r.Group("/api/v1")
    {
        //标签模块
        apivi.GET("/tags", v1.GetTag)
        apivi.POST("/tags", v1.AddTag)
        apivi.PUT("/tags/:id", v1.EditTag)
        apivi.DELETE("/tags/:id", v1.DeleteTag)

        //文章模块
        apivi.GET("/articles", v1.GetArticles)
        apivi.POST("/articles/:id", v1.GetArticle)
        apivi.POST("/articles", v1.AddArticle)
        apivi.PUT("/articles/:id", v1.EditArticle)
        apivi.DELETE("/articles/:id", v1.DeleteArticle)
    }

    return r
}

```



#### 4.验证`Token`

获取`token`的 API 方法就到这里啦，让我们来测试下是否可以正常使用吧！

重启服务后，用`GET`方式访问`http://127.0.0.1:8000/auth?username=test&password=test123456`，查看返回值是否正确

```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2NTI3MDEwNDYsImlzcyI6Imdpbi1ibG9nIn0.fw-4JNxRkdGbKBjFIxbmkFyRyLBkVcrZk7WtHaiEdi8"
    },
    "msg": "ok"
}
```

我们有了`token`的 API，也调用成功了



#### 5.将中间件接入`Gin`

接下来我们将中间件接入到`Gin`的访问流程中

我们打开`routers`目录下的`router.go`文件，修改文件内容（新增引用包和中间件引用）

```
package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/middleware/jwt"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "github.com/jamesluo111/gin-blog/routers/api"
    v1 "github.com/jamesluo111/gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunModel)

    r.GET("/auth", api.GetAuth)

    apivi := r.Group("/api/v1")
    //使用中间件
    apivi.Use(jwt.JWT())
    {
        //标签模块
        apivi.GET("/tags", v1.GetTag)
        apivi.POST("/tags", v1.AddTag)
        apivi.PUT("/tags/:id", v1.EditTag)
        apivi.DELETE("/tags/:id", v1.DeleteTag)

        //文章模块
        apivi.GET("/articles", v1.GetArticles)
        apivi.POST("/articles/:id", v1.GetArticle)
        apivi.POST("/articles", v1.AddArticle)
        apivi.PUT("/articles/:id", v1.EditArticle)
        apivi.DELETE("/articles/:id", v1.DeleteArticle)
    }

    return r
}


```

当前目录结构：

```

go-gin-example/
├── conf
│   └── app.ini
├── main.go
├── middleware
│   └── jwt
│       └── jwt.go
├── models
│   ├── article.go
│   ├── auth.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── routers
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```

到这里，我们的`JWT`编写就完成啦！



#### 6.验证功能

我们来测试一下，再次访问

- http://127.0.0.1:8000/api/v1/articles
- http://127.0.0.1:8000/api/v1/articles?token=23131

正确的反馈应该是

```
{
  "code": 400,
  "data": null,
  "msg": "请求参数错误"
}

{
  "code": 20001,
  "data": null,
  "msg": "Token鉴权失败"
}
```

我们需要访问`http://127.0.0.1:8000/auth?username=test&password=test123456`，得到`token`

```
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2NTI3MDE3NjAsImlzcyI6Imdpbi1ibG9nIn0.hrpUjHTBlaPadWE32fwXBkQwE-zxNBvRSm-8tna-9yU"
    },
    "msg": "ok"
}
```

再用包含`token`的 URL 参数去访问我们的应用 API，

访问`http://127.0.0.1:8000/api/v1/articles?token=eyJhbGci...`，检查接口返回值

```
{
    "code": 200,
    "data": {
        "count": 0,
        "list": [
            {
                "id": 1,
                "created_on": 1652672753,
                "modified_on": 1652673056,
                "tag_id": 1,
                "tag": {
                    "id": 1,
                    "created_on": 1652672347,
                    "modified_on": 0,
                    "name": "golang",
                    "created_by": "luo",
                    "modified_by": "",
                    "state": 1
                },
                "title": "how to study gin",
                "desc": "how to study gin",
                "content": "read and read",
                "created_by": "yang",
                "modified_by": "",
                "state": 1
            }
        ]
    },
    "msg": "ok"
}
```

返回正确，至此我们的`jwt-go`在`Gin`中的验证就完成了！



### 六、编写一个简单的文件日志

在上一节中，我们解决了 API's 可以任意访问的问题，那么我们现在还有一个问题，就是我们的日志，都是输出到控制台上的，这显然对于一个项目来说是不合理的，因此我们这一节简单封装`log`库，使其支持简单的文件日志！

#### 1.新建`logging`包

我们在`pkg`下新建`logging`目录，新建`file.go`和`log.go`文件，写入内容



#### 2.编写`file`文件

```
package logging

import (
    "fmt"
    "log"
    "os"
    "time"
)

var (
    LogSavePath = "runtime/logs/"
    LogSaveName = "log"
    LogFileExt  = "log"
    TimeFormat  = "20060102"
)

func getLogFilePath() string {
    return fmt.Sprintf("%s", LogSavePath)
}

func getFileFullPath() string {
    prefixPath := getLogFilePath()
    suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

    return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filePath string) *os.File {
    _, err := os.Stat(filePath)

    switch {
    case os.IsNotExist(err):
        mkDir()
    case os.IsPermission(err):
        log.Fatalf("Permission :%v", err)
    }

    handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Fail to open:%v", err)
    }

    return handle
}

func mkDir() {
    dir, _ := os.Getwd()
    err := os.MkdirAll(dir+"/"+getFileFullPath(), os.ModePerm)
    if err != nil {
        panic(err)
    }
}

```

- `os.Stat`：返回文件信息结构描述文件。如果出现错误，会返回`*PathError`

```
type PathError struct {
    Op   string
    Path string
    Err  error
}
```

- `os.IsNotExist`：能够接受`ErrNotExist`、`syscall`的一些错误，它会返回一个布尔值，能够得知文件不存在或目录不存在
- `os.IsPermission`：能够接受`ErrPermission`、`syscall`的一些错误，它会返回一个布尔值，能够得知权限是否满足
- `os.OpenFile`：调用文件，支持传入文件名称、指定的模式调用文件、文件权限，返回的文件的方法可以用于 I/O。如果出现错误，则为`*PathError`。

```

const (
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
    O_RDONLY int = syscall.O_RDONLY // 以只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 以只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 以读写模式打开文件
    // The remaining values may be or'ed in to control behavior.
    O_APPEND int = syscall.O_APPEND // 在写入时将数据追加到文件中
    O_CREATE int = syscall.O_CREAT  // 如果不存在，则创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 使用O_CREATE时，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 同步IO
    O_TRUNC  int = syscall.O_TRUNC  // 如果可以，打开时
)
```

- `os.Getwd`：返回与当前目录对应的根路径名
- `os.MkdirAll`：创建对应的目录以及所需的子目录，若成功则返回`nil`，否则返回`error`
- `os.ModePerm`：`const`定义`ModePerm FileMode = 0777`



#### 3.编写`log`文件

```
package logging

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "runtime"
)

type Level int

var (
    F                  *os.File
    DefaultPrefix      = ""
    DefaultCallerDepth = 2

    logger     *log.Logger
    logPrefix  = ""
    LevelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
    DEBUG Level = iota
    INFO
    WARN
    ERROR
    FATAL
)

func init() {
    filePath := getFileFullPath()
    F = openLogFile(filePath)
    logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
    setPrefix(DEBUG)
    logger.Println(v)
}

func Info(v ...interface{}) {
    setPrefix(INFO)
    logger.Println(v)
}

func Warn(v ...interface{}) {
    setPrefix(WARN)
    logger.Println(v)
}

func Error(v ...interface{}) {
    setPrefix(ERROR)
    logger.Println(v)
}

func Fatal(v ...interface{}) {
    setPrefix(FATAL)
    logger.Println(v)
}

func setPrefix(level Level) {
    _, file, line, ok := runtime.Caller(DefaultCallerDepth)
    if ok {
        logPrefix = fmt.Sprintf("[%s][%s:%d]", LevelFlags[level], filepath.Base(file), line)
    } else {
        logPrefix = fmt.Sprintf("[%s]", LevelFlags[level])
    }

    log.SetPrefix(logPrefix)
}

```

- `log.New`：创建一个新的日志记录器。`out`定义要写入日志数据的`IO`句柄。`prefix`定义每个生成的日志行的开头。`flag`定义了日志记录属性

```
func New(out io.Writer, prefix string, flag int) *Logger {
    return &Logger{out: out, prefix: prefix, flag: flag}
}
```

- `log.LstdFlags`：日志记录的格式属性之一，其余的选项如下

```

const (
    Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
    Ltime                         // the time in the local time zone: 01:23:23
    Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    Llongfile                     // full file name and line number: /a/b/c/d.go:23
    Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
    LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
    LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

当前目录结构：

```

gin-blog/
├── conf
│   └── app.ini
├── main.go
├── middleware
│   └── jwt
│       └── jwt.go
├── models
│   ├── article.go
│   ├── auth.go
│   ├── models.go
│   └── tag.go
├── pkg
│   ├── e
│   │   ├── code.go
│   │   └── msg.go
│   ├── logging
│   │   ├── file.go
│   │   └── log.go
│   ├── setting
│   │   └── setting.go
│   └── util
│       ├── jwt.go
│       └── pagination.go
├── routers
│   ├── api
│   │   ├── auth.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime
```

我们自定义的`logging`包，已经基本完成了，接下来让它接入到我们的项目之中吧。我们打开先前包含`log`包的代码，如下：

1. 打开`routers`目录下的`article.go`、`tag.go`、`auth.go`。
2. 将`log`包的引用删除，修改引用我们自己的日志包为`github.com/jamesluo111/gin-blog/pkg/logging`。
3. 将原本的`log.Println(...)`改为`logging.Info(...)`。

例如`auth.go`文件的修改内容：

```
package api

import (
    "github.com/astaxie/beego/validation"
    "github.com/gin-gonic/gin"
    "github.com/jamesluo111/gin-blog/models"
    "github.com/jamesluo111/gin-blog/pkg/e"
    "github.com/jamesluo111/gin-blog/pkg/logging"
    "github.com/jamesluo111/gin-blog/pkg/util"
    "net/http"
)

type auth struct {
    Username string `valid:"Required; MaxSize(50)"`
    Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
    username := c.Query("username")
    password := c.Query("password")

    valid := validation.Validation{}
    a := auth{Username: username, Password: password}
    ok, _ := valid.Valid(&a)

    data := make(map[string]interface{})
    code := e.INVALID_PARAMS

    if ok {
        isExist := models.CheckoutAuth(username, password)

        if isExist {
            token, err := util.GenerateToken(username, password)

            if err != nil {
                code = e.ERROR_AUTH_TOKEN
            } else {
                data["token"] = token
                code = e.SUCCESS
            }
        } else {
            code = e.ERROR_AUTH
        }
    } else {
        for _, err := range valid.Errors {
            logging.Info(err.Key, err.Message)
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}

```



#### 4.验证功能

修改文件后，重启服务，我们来试试吧！

获取到 API 的 Token 后，我们故意传错误 URL 参数给接口，如：`http://127.0.0.1:8000/api/v1/articles?tag_id=0&state=9999999&token=eyJhbG..`

然后我们到`$GOPATH/gin-blog/runtime/logs`查看日志：

```
2022/05/16 18:28:06 [err.key:state, err.Msg:state在0到1之间]
2022/05/16 18:28:06 [err.key:tagId, err.Msg:tagId不能小于1]
```

日志结构一切正常，我们的记录模式都为`Info`，因此前缀是对的，并且我们是入参有问题，也把错误记录下来了，这样排错就很方便了！

至此，本节就完成了，这只是一个简单的扩展，实际上我们线上项目要使用的文件日志，是更复杂一些，开动你的大脑 举一反三吧！



### 七、优雅的重启服务

 在前面编写案例代码时，我相信你会想到，每次更新完代码，更新完配置文件后，就直接这么 `ctrl+c` 真的没问题吗，`ctrl+c`到底做了些什么事情呢？ 

 在这一节中我们简单讲述 `ctrl+c` 背后的**信号**以及如何在`Gin`中**优雅的重启服务**，也就是对 `HTTP` 服务进行热更新。 

#### 1.ctrl + c

> 内核在某些情况下发送信号，比如在进程往一个已经关闭的管道写数据时会产生`SIGPIPE`信号 

 在终端执行特定的组合键可以使系统发送特定的信号给此进程，完成一系列的动作 

| **命令** | **信号** | **含义**                                                     |
| -------- | -------- | ------------------------------------------------------------ |
| ctrl + c | SIGINT   | 强制进程结束                                                 |
| ctrl + z | SIGTSTP  | 任务中断，进程挂起                                           |
| ctrl + \ | SIGQUIT  | 进程结束 和 `dump core`                                      |
| ctrl + d |          | EOF                                                          |
|          | SIGHUP   | 终止收到该信号的进程。若程序中没有捕捉该信号，当收到该信号时，进程就会退出（常用于 重启、重新加载进程） |

 因此在我们执行`ctrl + c`关闭`gin`服务端时，**会强制进程结束，导致正在访问的用户等出现问题** 

 常见的 `kill -9 pid` 会发送 `SIGKILL` 信号给进程，也是类似的结果 



#### 2.信号

本段中反复出现**信号**是什么呢？

信号是 `Unix` 、类 `Unix` 以及其他 `POSIX` 兼容的操作系统中进程间通讯的一种有限制的方式

它是一种异步的通知机制，用来提醒进程一个事件（硬件异常、程序执行异常、外部发出信号）已经发生。当一个信号发送给一个进程，操作系统中断了进程正常的控制流程。此时，任何非原子操作都将被中断。如果进程定义了信号的处理函数，那么它将被执行，否则就执行默认的处理函数



#### 3.所有信号

```

$ kill -l
 1) SIGHUP   2) SIGINT   3) SIGQUIT  4) SIGILL   5) SIGTRAP
 6) SIGABRT  7) SIGBUS   8) SIGFPE   9) SIGKILL 10) SIGUSR1
11) SIGSEGV 12) SIGUSR2 13) SIGPIPE 14) SIGALRM 15) SIGTERM
16) SIGSTKFLT   17) SIGCHLD 18) SIGCONT 19) SIGSTOP 20) SIGTSTP
21) SIGTTIN 22) SIGTTOU 23) SIGURG  24) SIGXCPU 25) SIGXFSZ
26) SIGVTALRM   27) SIGPROF 28) SIGWINCH    29) SIGIO   30) SIGPWR
31) SIGSYS  34) SIGRTMIN    35) SIGRTMIN+1  36) SIGRTMIN+2  37) SIGRTMIN+3
38) SIGRTMIN+4  39) SIGRTMIN+5  40) SIGRTMIN+6  41) SIGRTMIN+7  42) SIGRTMIN+8
43) SIGRTMIN+9  44) SIGRTMIN+10 45) SIGRTMIN+11 46) SIGRTMIN+12 47) SIGRTMIN+13
48) SIGRTMIN+14 49) SIGRTMIN+15 50) SIGRTMAX-14 51) SIGRTMAX-13 52) SIGRTMAX-12
53) SIGRTMAX-11 54) SIGRTMAX-10 55) SIGRTMAX-9  56) SIGRTMAX-8  57) SIGRTMAX-7
58) SIGRTMAX-6  59) SIGRTMAX-5  60) SIGRTMAX-4  61) SIGRTMAX-3  62) SIGRTMAX-2
63) SIGRTMAX-1  64) SIGRTMAX
```



#### 4.怎样算优雅

**目的**

- 不关闭现有连接（正在运行中的程序）
- 新的进程启动并替代旧进程
- 新的进程接管新的连接
- 连接要随时响应用户的请求，当用户仍在请求旧进程时要保持连接，新用户应请求新进程，不可以出现拒绝请求的情况

**流程**

1、替换可执行文件或修改配置文件

2、发送信号量 `SIGHUP`

3、拒绝新连接请求旧进程，但要保证已有连接正常

4、启动新的子进程

5、新的子进程开始 `Accet`

6、系统将新的请求转交新的子进程

7、旧进程处理完所有旧连接后正常结束



#### 5.实现优雅重启

> Zero downtime restarts for golang HTTP and HTTPS servers. (for golang 1.3+) 

我们借助 fvbock/endless 来实现 `Golang HTTP/HTTPS` 服务重新启动的零停机

`endless server` 监听以下几种信号量：

- syscall.SIGHUP：触发 `fork` 子进程和重新启动
- syscall.SIGUSR1/syscall.SIGTSTP：被监听，但不会触发任何动作
- syscall.SIGUSR2：触发 `hammerTime`
- syscall.SIGINT/syscall.SIGTERM：触发服务器关闭（会完成正在运行的请求）

`endless` 正正是依靠监听这些**信号量**，完成管控的一系列动作

**安装**

```
go get -u github.com/fvbock/endless
```

**编写**

 打开 gin-blog 的 `main.go`文件，修改文件： 

```
package main

import (
    "fmt"
    "github.com/fvbock/endless"
    "github.com/jamesluo111/gin-blog/pkg/setting"
    "github.com/jamesluo111/gin-blog/routers"
    "log"
    "syscall"
)

func main() {
    endless.DefaultReadTimeOut = setting.ReadTimeOut
    endless.DefaultWriteTimeOut = setting.WriteTimeOut
    endless.DefaultMaxHeaderBytes = 1 << 20
    endpoint := fmt.Sprintf(":%d", setting.HttpPort)
    server := endless.NewServer(endpoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getegid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err:%v", err)
    }
}

```

`endless.NewServer` 返回一个初始化的 `endlessServer` 对象，在 `BeforeBegin` 时输出当前进程的 `pid`，调用 `ListenAndServe` 将实际“启动”服务

**验证**

**编译**

```
$ go build main.go
```

**执行**

```
$ ./main
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
...
Actual pid is 48601
```

启动成功后，输出了`pid`为 48601；在另外一个终端执行 `kill -1 48601`，检验先前服务的终端效果

```
[root@localhost go-gin-example]# ./main
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> ...
[GIN-debug] GET    /api/v1/tags              --> ...
...

Actual pid is 48601

...

Actual pid is 48755
48601 Received SIGTERM.
48601 [::]:8000 Listener closed.
48601 Waiting for connections to finish...
48601 Serve() returning...
Server err: accept tcp [::]:8000: use of closed network connection
```

可以看到该命令已经挂起，并且 `fork` 了新的子进程 `pid` 为 `48755`

```
48601 Received SIGTERM.
48601 [::]:8000 Listener closed.
48601 Waiting for connections to finish...
48601 Serve() returning...
Server err: accept tcp [::]:8000: use of closed network connection
```

大致意思为主进程（`pid`为 48601）接受到 `SIGTERM` 信号量，关闭主进程的监听并且等待正在执行的请求完成；这与我们先前的描述一致

**唤醒**

这时候在 `postman` 上再次访问我们的接口，你可以惊喜的发现，他“复活”了！

```

Actual pid is 48755
48601 Received SIGTERM.
48601 [::]:8000 Listener closed.
48601 Waiting for connections to finish...
48601 Serve() returning...
Server err: accept tcp [::]:8000: use of closed network connection


$ [GIN] 2018/03/15 - 13:00:16 | 200 |     188.096µs |   192.168.111.1 | GET      /api/v1/tags...
```

这就完成了一次正向的流转了

你想想，每次更新发布、或者修改配置文件等，只需要给该进程发送**SIGTERM 信号**，而不需要强制结束应用，是多么便捷又安全的事！

**问题**

`endless` 热更新是采取创建子进程后，将原进程退出的方式，这点不符合守护进程的要求



#### 6.http.Server - Shutdown()

如果你的`Golang >= 1.8`，也可以考虑使用 `http.Server` 的 Shutdown 方法

```
    router := routers.InitRouter()

    s := &http.Server{
        Addr:           fmt.Sprintf(":%d", setting.HttpPort),
        Handler:        router,
        ReadTimeout:    setting.ReadTimeOut,
        WriteTimeout:   setting.WriteTimeOut,
        MaxHeaderBytes: 1 << 20,
    }

    go func() {
        if err := s.ListenAndServe(); err != nil {
            log.Printf("Listen:%v", err)
        }
    }()
    //创建暂停信号量chan，用来接收服务停止请求
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit

    log.Println("ShutDown Server...")
    //当服务停止后争取5秒钟来处理停止前的所有请求
    ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cannel()
    if err := s.Shutdown(ctx); err != nil {
        log.Fatal("Server Shutdown:", err)
    }

    log.Println("Server exiting")
```

**小结**

在日常的服务中，优雅的重启（热更新）是非常重要的一环。而 `Golang` 在 `HTTP` 服务方面的热更新也有不少方案了，我们应该根据实际应用场景挑选最合适的



### 八、Swagger

一个好的 `API's`，必然离不开一个好的`API`文档，如果要开发纯手写 `API` 文档，不存在的（很难持续维护），因此我们要自动生成接口文档。

#### 1.安装 swag

```
$ go get -u github.com/swaggo/swag/cmd/swag@v1.6.5
```

若 `$GOROOT/bin` 没有加入`$PATH`中，你需要执行将其可执行文件移动到`$GOBIN`下

```
mv $GOPATH/bin/swag /usr/local/go/bin
```



#### 2.验证是否安装成功

检查 $GOBIN 下是否有 swag 文件，如下：

```
$ swag -v
swag version v1.6.5
```



#### 3.安装 gin-swagger

```
$ go get -u github.com/swaggo/gin-swagger@v1.2.0 
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```

注：若无科学上网，请务必配置 Go modules proxy。



#### 4.初始化,编写 API 注释

`Swagger` 中需要将相应的注释或注解编写到方法上，再利用生成器自动生成说明文件

`gin-swagger` 给出的范例：

```
// @Summary Add a new pet to the store
// @Description get string by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    int     true        "Some ID"
// @Success 200 {string} string "ok"
// @Failure 400 {object} web.APIError "We need ID!!"
// @Failure 404 {object} web.APIError "Can not find ID"
// @Router /testapi/get-string-by-int/{some_id} [get]
```

我们可以参照 `Swagger` 的注解规范和范例去编写

```
// @Summary 获取文章标签
// @Produce json
// @Param name query string true "Name"
// @Param state query int false "State"
// @Success 200 {string} string "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
```

参考的注解请参见 go-gin-example。以确保获取最新的 swag 语法

**路由**

在完成了注解的编写后，我们需要针对 swagger 新增初始化动作和对应的路由规则，才可以使用。打开 routers/router.go 文件，新增内容如下：

```
package routers

import (
    "github.com/gin-gonic/gin"
    ...

    _ "github.com/jamesluo111/gin-blog/docs"
)

func InitRouter() *gin.Engine {
    ...

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    apivi := r.Group("/api/v1")
    apivi.Use(jwt.JWT())
    {
        ...
    }

    return r
}
```

**生成**

我们进入到`gin-blog`的项目根目录中，执行初始化命令

```
root@60aad428b48d:/wwwroot/gin-blog# swag init
2022/05/17 06:58:48 Generate swagger docs....
2022/05/17 06:58:48 Generate general API Info, search dir:./
2022/05/17 06:58:48 create docs.go at  docs/docs.go
2022/05/17 06:58:48 create swagger.json at  docs/swagger.json
2022/05/17 06:58:48 create swagger.yaml at  docs/swagger.yaml
```

完毕后会在项目根目录下生成`docs`

```
docs/
├── docs.go
└── swagger
    ├── swagger.json
    └── swagger.yaml
```

我们可以检查 `docs.go` 文件中的 `doc` 变量，详细记载中我们文件中所编写的注解和说明

![image-20220517163222194](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220517163222194.png?lastModify=1653785198)

**验证**

大功告成，访问一下 `http://127.0.0.1:8000/swagger/index.html`， 查看 `API` 文档生成是否正确

![image-20220517163303438](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220517163303438.png?lastModify=1653785198)



### 九、将Golang应用部署到Docker

将我们的 `gin-blog` 应用部署到一个 Docker 里，你需要先准备好如下东西：

- 你需要安装好 `docker`。
- 如果上外网比较吃力，需要配好镜像源。



#### 1.Docker

在这里简单介绍下 Docker，建议深入学习。Docker 是一个开源的轻量级容器技术，让开发者可以打包他们的应用以及应用运行的上下文环境到一个可移植的镜像中，然后发布到任何支持 Docker 的系统上运行。 通过容器技术，在几乎没有性能开销的情况下，Docker 为应用提供了一个隔离运行环境

- 简化配置
- 代码流水线管理
- 提高开发效率
- 隔离应用
- 快速、持续部署



#### 2.Golang

##### 一、编写 Dockerfile

在 `gin-blog` 项目根目录创建 Dockerfile 文件，写入内容

```
FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/jamesluo111/gin-blog
COPY . $GOPATH/src/github.com/jamesluo111/gin-blog
RUN go build .
EXPOSE 8000
ENTRYPOINT ["./gin-blog"]
```

**作用**

`golang:latest` 镜像为基础镜像，将工作目录设置为 `$GOPATH/src/go-gin-example`，并将当前上下文目录的内容复制到 `$GOPATH/src/go-gin-example` 中

在进行 `go build` 编译完毕后，将容器启动程序设置为 `./go-gin-example`，也就是我们所编译的可执行文件

注意 `go-gin-example` 在 `docker` 容器里编译，并没有在宿主机现场编译

**说明**

Dockerfile 文件是用于定义 Docker 镜像生成流程的配置文件，文件内容是一条条指令，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建；这些指令应用于基础镜像并最终创建一个新的镜像

你可以认为用于快速创建自定义的 Docker 镜像

**1、 FROM**

指定基础镜像（必须有的指令，并且必须是第一条指令）

**2、 WORKDIR**

格式为 `WORKDIR` <工作目录路径>

使用 `WORKDIR` 指令可以来**指定工作目录**（或者称为当前目录），以后各层的当前目录就被改为指定的目录，如果目录不存在，`WORKDIR` 会帮你建立目录

**3、COPY**

格式：

```
COPY <源路径>... <目标路径>
COPY ["<源路径1>",... "<目标路径>"]
```

`COPY` 指令将从构建上下文目录中 <源路径> 的文件/目录**复制**到新的一层的镜像内的 <目标路径> 位置

**4、RUN**

用于执行命令行命令

格式：`RUN` <命令>

**5、EXPOSE**

格式为 `EXPOSE` <端口 1> [<端口 2>...]

`EXPOSE` 指令是**声明运行时容器提供服务端口，这只是一个声明**，在运行时并不会因为这个声明应用就会开启这个端口的服务

在 Dockerfile 中写入这样的声明有两个好处

- 帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射
- 运行时使用随机端口映射时，也就是 `docker run -P` 时，会自动随机映射 `EXPOSE` 的端口

**6、ENTRYPOINT**

`ENTRYPOINT` 的格式和 `RUN` 指令格式一样，分为两种格式

- `exec` 格式：
- 

```
<ENTRYPOINT> "<CMD>"
```

- `shell` 格式：
- 

```
ENTRYPOINT [ "curl", "-s", "http://ip.cn" ]
```

`ENTRYPOINT` 指令是**指定容器启动程序及参数**



##### 二、构建镜像

```
gin-blog` 的项目根目录下**执行** `docker build -t gin-blog-docker
```

该命令作用是创建/构建镜像，`-t` 指定名称为 `gin-blog-docker`，`.` 构建内容为当前上下文目录

```
root@PC-20200715ZEOD:/mnt/d/dnmp/wwwroot/gin-blog# docker build -t gin-blog-docker .
Sending build context to Docker daemon  138.2kB
Step 1/7 : FROM golang:latest
latest: Pulling from library/golang
67e8aa6c8bbc: Pull complete
627e6c1e1055: Pull complete
0670968926f6: Pull complete
5a8b0e20be4b: Pull complete
10f766b17f53: Pull complete
21e7497335c1: Pull complete
1e452e64228d: Pull complete
Digest: sha256:02c05351ed076c581854c554fa65cb2eca47b4389fb79a1fc36f21b8df59c24f
Status: Downloaded newer image for golang:latest
 ---> 7d1902a99d63
Step 2/7 : ENV GOPROXY https://goproxy.cn,direct
 ---> Running in f68a1cf1d687
Removing intermediate container f68a1cf1d687
 ---> 589b80ac7a72
Step 3/7 : WORKDIR $GOPATH/src/github.com/jamesluo111/gin-blog
 ---> Running in 3b41f608e814
Removing intermediate container 3b41f608e814
 ---> 6f5198c2fa39
Step 4/7 : COPY . $GOPATH/src/github.com/jamesluo111/gin-blog
 ---> 40c674f5016f
Step 5/7 : RUN go build .
 ---> Running in 64fe5eb352a5
...
 ---> 0c30bcec8d66
Step 6/7 : EXPOSE 8000
 ---> Running in 9f68e2b7506f
Removing intermediate container 9f68e2b7506f
 ---> d4ebb3ea8be1
Step 7/7 : ENTRYPOINT ["./gin-blog"]
 ---> Running in 25b1fa0a6216
Removing intermediate container 25b1fa0a6216
 ---> b803fa2dcd3d
Successfully built b803fa2dcd3d
Successfully tagged gin-blog-docker:latest
```



##### 三、验证镜像

查看所有的镜像，确定刚刚构建的 `gin-blog-docker` 镜像是否存在

```
root@PC-20200715ZEOD:/mnt/d/dnmp/wwwroot/gin-blog# docker images
REPOSITORY            TAG       IMAGE ID       CREATED          SIZE
gin-blog-docker       latest    b803fa2dcd3d   18 seconds ago   1.39GB
golang                latest    7d1902a99d63   5 days ago       964MB
...
```



##### 四、创建并运行一个新容器

执行命令 `docker run -p 8000:8000 gin-blog-docker`

```
$ docker run -p 8000:8000 gin-blog-docker
dial tcp 127.0.0.1:3306: connect: connection refused
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

...
Actual pid is 1
```

运行成功，你以为大功告成了吗？

你想太多了，仔细看看控制台的输出了一条错误 `dial tcp 127.0.0.1:3306: connect: connection refused`

我们研判一下，发现是 `Mysql` 的问题，接下来第二项我们将解决这个问题



#### 3.Mysql

##### 一、拉取镜像

从 `Docker` 的公共仓库 `Dockerhub` 下载 `MySQL` 镜像（国内建议配个镜像）

```
$ docker pull mysql
```

##### 二、创建并运行一个新容器

运行 `Mysql` 容器，并设置执行成功后返回容器 ID

```
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -d mysql
5831467a8d86fb283ccea13a15e4bb75c0e664c0efb552f653bdb7148d579ddb
```

**连接 Mysql**



#### 4.Golang + Mysql

##### 一、删除镜像

由于原本的镜像存在问题，我们需要删除它，此处有几种做法

- 删除原本有问题的镜像，重新构建一个新镜像
- 重新构建一个不同 `name`、`tag` 的新镜像

删除原本的有问题的镜像，`-f` 是强制删除及其关联状态

若不执行 `-f`，你需要执行 `docker ps -a` 查到所关联的容器，将其 `rm` 解除两者依赖关系

```
root@PC-20200715ZEOD:/mnt/d/dnmp/wwwroot/gin-blog# docker rmi -f gin-blog-docker
Untagged: gin-blog-docker:latest
Deleted: sha256:b803fa2dcd3d2346aaa96bbed7f19f2b986e9f34a4786852389c840c7bba010f
Deleted: sha256:d4ebb3ea8be1c0ab54a4c7ba2286bc72de58453bb6eeca4d977cc60f2093f84c
Deleted: sha256:0c30bcec8d666067d20af9a0149d0bcc2bdd52c0f55bb358f48655e6f93ef222
Deleted: sha256:40c674f5016f2b8aa6635da9171afcdd16a2eba9fb78bc3b6b66ae3509fad28d
Deleted: sha256:6f5198c2fa39accd1167202787c9cf9768a6df47a51542828e438482593e37b4
Deleted: sha256:589b80ac7a72ff2c54440d21fff2febb10c519ce0df89c0d4ca69d3699153c4a
```



##### 二、修改配置文件

将项目的配置文件 `conf/app.ini`，内容修改为

```
#debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 233

[server]
HTTP_PORT = 8000
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60

[database]
TYPE = mysql
USER = root
PASSWORD = root
HOST = mysql:3306
NAME = blog
TABLE_PREFIX = blog_
```

##### 三、重新构建镜像

重复先前的步骤，回到 `gin-blog` 的项目根目录下**执行** `docker build -t gin-blog-docker .`



##### 四、创建并运行一个新容器

**关联**

Q：我们需要将 `Golang` 容器和 `Mysql` 容器关联起来，那么我们需要怎么做呢？

A：增加命令 `--link mysql:mysql` 让 `Golang` 容器与 `Mysql` 容器互联；通过 `--link`，**可以在容器内直接使用其关联的容器别名进行访问**，而不通过 IP，但是`--link`只能解决单机容器间的关联，在分布式多机的情况下，需要通过别的方式进行连接

**运行**

执行命令 `docker run --link mysql:mysql -p 8000:8000 gin-blog-docker`

```

$ docker run --link mysql:mysql -p 8000:8000 gin-blog-docker
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)
...
Actual pid is 1
```

**结果**

检查启动输出、接口测试、数据库内数据，均正常；我们的 `Golang` 容器和 `Mysql` 容器成功关联运行，大功告成 :)



#### 5.Review

**思考**

虽然应用已经能够跑起来了

但如果对 `Golang` 和 `Docker` 有一定的了解，我希望你能够想到至少 2 个问题

- 为什么 `gin-blog-docker` 占用空间这么大？（可用 `docker ps -as | grep gin-blog-docker` 查看）
- `Mysql` 容器直接这么使用，数据存储到哪里去了？

**创建超小的 Golang 镜像**

Q：第一个问题，为什么这么镜像体积这么大？

A：`FROM golang:latest` 拉取的是官方 `golang` 镜像，包含 Golang 的编译和运行环境，外加一堆 GCC、build 工具，相当齐全

这是有问题的，**我们可以不在 Golang 容器中现场编译的**，压根用不到那些东西，我们只需要一个能够运行可执行文件的环境即可

**构建 Scratch 镜像**

Scratch 镜像，简洁、小巧，基本是个空镜像

##### 一、修改 Dockerfile

```
FROM scratch

WORKDIR $GOPATH/src/github.com/jamesluo111/gin-blog
COPY . $GOPATH/src/github.com/jamesluo111/gin-blog

EXPOSE 8000
ENTRYPOINT ["./gin-blog"]
```

##### 二、编译可执行文件

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gin-blog .
```

编译所生成的可执行文件会依赖一些库，并且是动态链接。在这里因为使用的是 `scratch` 镜像，它是空镜像，因此我们需要将生成的可执行文件静态链接所依赖的库

##### 三、构建镜像

```
$ docker build -t gin-blog-docker-scratch .
Sending build context to Docker daemon 133.1 MB
Step 1/5 : FROM scratch
 --->
Step 2/5 : WORKDIR $GOPATH/src/github.com/EDDYCJY/go-gin-example
 ---> Using cache
 ---> ee07e166a638
Step 3/5 : COPY . $GOPATH/src/github.com/EDDYCJY/go-gin-example
 ---> 1489a0693d51
Removing intermediate container e3e9efc0fe4d
Step 4/5 : EXPOSE 8000
 ---> Running in b5630de5544a
 ---> 6993e9f8c944
Removing intermediate container b5630de5544a
Step 5/5 : CMD ./go-gin-example
 ---> Running in eebc0d8628ae
 ---> 5310bebeb86a
Removing intermediate container eebc0d8628ae
Successfully built 5310bebeb86a
```

注意，假设你的 Golang 应用没有依赖任何的配置等文件，是可以直接把可执行文件给拷贝进去即可，其他都不必关心

这里可以有好几种解决方案

- 依赖文件统一管理挂载
- go-bindata 一下

...

因此这里如果**解决了文件依赖的问题**后，就不需要把目录给 `COPY` 进去了

##### 四、运行

```
$ docker run --link mysql:mysql -p 8000:8000 gin-blog-docker-scratch
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/EDDYCJY/go-gin-example/routers/api.GetAuth (3 handlers)
...
```

成功运行，程序也正常接收请求

接下来我们再看看占用大小，执行 `docker ps -as` 命令

```
$ docker ps -as
CONTAINER ID        IMAGE                     COMMAND                  ...         SIZE
9ebdba5a8445        gin-blog-docker-scratch   "./go-gin-example"       ...     0 B (virtual 132 MB)
427ee79e6857        gin-blog-docker           "./go-gin-example"       ...     0 B (virtual 946 MB)
```

从结果而言，占用大小以`Scratch`镜像为基础的容器完胜，完成目标



#### 6.Mysql 挂载数据卷

倘若不做任何干涉，在每次启动一个 `Mysql` 容器时，数据库都是空的。另外容器删除之后，数据就丢失了（还有各类意外情况），非常糟糕！

**数据卷**

数据卷 是被设计用来持久化数据的，它的生命周期独立于容器，Docker 不会在容器被删除后自动删除 数据卷，并且也不存在垃圾回收这样的机制来处理没有任何容器引用的 数据卷。如果需要在删除容器的同时移除数据卷。可以在删除容器的时候使用 `docker rm -v` 这个命令

数据卷 是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性：

- 数据卷 可以在容器之间共享和重用
- 对 数据卷 的修改会立马生效
- 对 数据卷 的更新，不会影响镜像
- 数据卷 默认会一直存在，即使容器被删除

> 注意：数据卷 的使用，类似于 Linux 下对目录或文件进行 mount，镜像中的被指定为挂载点的目录中的文件会隐藏掉，能显示看的是挂载的 数据卷。

**如何挂载**

首先创建一个目录用于存放数据卷；示例目录 `/data/docker-mysql`，注意 `--name` 原本名称为 `mysql` 的容器，需要将其删除 `docker rm`

```
$ docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v /data/docker-mysql:/var/lib/mysql -d mysql
54611dbcd62eca33fb320f3f624c7941f15697d998f40b24ee535a1acf93ae72
```

创建成功，检查目录 `/data/docker-mysql`，下面多了不少数据库文件

**验证**

接下来交由你进行验证，目标是创建一些测试表和数据，然后删除当前容器，重新创建的容器，数据库数据也依然存在（当然了数据卷指向要一致）



### 十、定制 GORM Callbacks

GORM 本身是由回调驱动的，所以我们可以根据需要完全定制 GORM，以此达到我们的目的，如下：

- 注册一个新的回调
- 删除现有的回调
- 替换现有的回调
- 注册回调的顺序

在 GORM 中包含以上四类 Callbacks，我们结合项目选用 “替换现有的回调” 来解决一个小痛点。

#### 1.问题

在 models 目录下，我们包含 tag.go 和 article.go 两个文件，他们有一个问题，就是 BeforeCreate、BeforeUpdate 重复出现了，那难道 100 个文件，就要写一百次吗？

![image-20220524150552621](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220524150552621.png?lastModify=1653785198)

显然这是不可能的，如果先前你已经意识到这个问题，那挺 OK，但没有的话，现在开始就要改

#### 2.解决

在这里我们通过 Callbacks 来实现功能，不需要一个个文件去编写

#### 3.实现 Callbacks

打开 models 目录下的 models.go 文件，实现以下两个方法：

1、updateTimeStampForCreateCallback

```
// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        nowTime := time.Now().Unix()
        if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
            if createTimeField.IsBlank {
                createTimeField.Set(nowTime)
            }
        }

        if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
            if modifyTimeField.IsBlank {
                modifyTimeField.Set(nowTime)
            }
        }
    }
}
```

在这段方法中，会完成以下功能

- 检查是否有含有错误（db.Error）

- `scope.FieldByName` 通过 `scope.Fields()` 获取所有字段，判断当前是否包含所需字段

  ```
  for _, field := range scope.Fields() {
  ```

  ```
      if field.Name == name || field.DBName == name {
  ```

  ```
          return field, true
  ```

  ```
      }
  ```

  ```
      if field.DBName == dbName {
  ```

  ```
          mostMatchedField = field
  ```

  ```
      }
  ```

  ```
  }
  ```

- `field.IsBlank` 可判断该字段的值是否为空

  ```
  func isBlank(value reflect.Value) bool {
  ```

  ```
      switch value.Kind() {
  ```

  ```
      case reflect.String:
  ```

  ```
          return value.Len() == 0
  ```

  ```
      case reflect.Bool:
  ```

  ```
          return !value.Bool()
  ```

  ```
      case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
  ```

  ```
          return value.Int() == 0
  ```

  ```
      case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
  ```

  ```
          return value.Uint() == 0
  ```

  ```
      case reflect.Float32, reflect.Float64:
  ```

  ```
          return value.Float() == 0
  ```

  ```
      case reflect.Interface, reflect.Ptr:
  ```

  ```
          return value.IsNil()
  ```

  ```
      }
  ```

  ```
  
  ```

  ```
      return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
  ```

  ```
  }
  ```

- 若为空则 `field.Set` 用于给该字段设置值，参数为 `interface{}`

2、updateTimeStampForUpdateCallback

```
// updateTimeStampForUpdateCallback will set `ModifyTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
    if _, ok := scope.Get("gorm:update_column"); !ok {
        scope.SetColumn("ModifiedOn", time.Now().Unix())
    }
}
```

- `scope.Get(...)` 根据入参获取设置了字面值的参数，例如本文中是 `gorm:update_column` ，它会去查找含这个字面值的字段属性
- `scope.SetColumn(...)` 假设没有指定 `update_column` 的字段，我们默认在更新回调设置 `ModifiedOn` 的值



#### 4.注册 Callbacks

在上面小节我已经把回调方法编写好了，接下来需要将其注册进 GORM 的钩子里，但其本身自带 Create 和 Update 回调，因此调用替换即可

在 models.go 的 init 函数中，增加以下语句

```
db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
```



#### 5.验证

访问 AddTag 接口，成功后检查数据库，可发现 `created_on` 和 `modified_on` 字段都为当前执行时间

访问 EditTag 接口，可发现 `modified_on` 为最后一次执行更新的时间



#### 6.拓展

我们想到，在实际项目中硬删除是较少存在的，那么是否可以通过 Callbacks 来完成这个功能呢？

答案是可以的，我们在先前 `Model struct` 增加 `DeletedOn` 变量

```
type Model struct {
    ID int `gorm:"primary_key" json:"id"`
    CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
    DeletedOn int `json:"deleted_on"`
}
```

#### 7.实现 Callbacks

打开 models 目录下的 models.go 文件，实现以下方法：

```
func deleteCallback(scope *gorm.Scope) {
    if !scope.HasError() {
        var extraOption string
        if str, ok := scope.Get("gorm:delete_option"); ok {
            extraOption = fmt.Sprint(str)
        }

        deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")

        if !scope.Search.Unscoped && hasDeletedOnField {
            scope.Raw(fmt.Sprintf(
                "UPDATE %v SET %v=%v%v%v",
                scope.QuotedTableName(),
                scope.Quote(deletedOnField.DBName),
                scope.AddToVars(time.Now().Unix()),
                addExtraSpaceIfExist(scope.CombinedConditionSql()),
                addExtraSpaceIfExist(extraOption),
            )).Exec()
        } else {
            scope.Raw(fmt.Sprintf(
                "DELETE FROM %v%v%v",
                scope.QuotedTableName(),
                addExtraSpaceIfExist(scope.CombinedConditionSql()),
                addExtraSpaceIfExist(extraOption),
            )).Exec()
        }
    }
}

func addExtraSpaceIfExist(str string) string {
    if str != "" {
        return " " + str
    }
    return ""
}
```

- `scope.Get("gorm:delete_option")` 检查是否手动指定了 delete_option

- `scope.FieldByName("DeletedOn")` 获取我们约定的删除字段，若存在则 `UPDATE` 软删除，若不存在则 `DELETE` 硬删除

- `scope.QuotedTableName()` 返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理

- `scope.CombinedConditionSql()` 返回组合好的条件 SQL，看一下方法原型很明了

  ```
  
  ```

  ```
  func (scope *Scope) CombinedConditionSql() string {
  ```

  ```
      joinSQL := scope.joinsSQL()
  ```

  ```
      whereSQL := scope.whereSQL()
  ```

  ```
      if scope.Search.raw {
  ```

  ```
          whereSQL = strings.TrimSuffix(strings.TrimPrefix(whereSQL, "WHERE ("), ")")
  ```

  ```
      }
  ```

  ```
      return joinSQL + whereSQL + scope.groupSQL() +
  ```

  ```
          scope.havingSQL() + scope.orderSQL() + scope.limitAndOffsetSQL()
  ```

  ```
  }
  ```

- `scope.AddToVars` 该方法可以添加值作为 SQL 的参数，也可用于防范 SQL 注入

  ```
  func (scope *Scope) AddToVars(value interface{}) string {
  ```

  ```
      _, skipBindVar := scope.InstanceGet("skip_bindvar")
  ```

  ```
  
  ```

  ```
      if expr, ok := value.(*expr); ok {
  ```

  ```
          exp := expr.expr
  ```

  ```
          for _, arg := range expr.args {
  ```

  ```
              if skipBindVar {
  ```

  ```
                  scope.AddToVars(arg)
  ```

  ```
              } else {
  ```

  ```
                  exp = strings.Replace(exp, "?", scope.AddToVars(arg), 1)
  ```

  ```
              }
  ```

  ```
          }
  ```

  ```
          return exp
  ```

  ```
      }
  ```

  ```
  
  ```

  ```
      scope.SQLVars = append(scope.SQLVars, value)
  ```

  ```
  
  ```

  ```
      if skipBindVar {
  ```

  ```
          return "?"
  ```

  ```
      }
  ```

  ```
      return scope.Dialect().BindVar(len(scope.SQLVars))
  ```

  ```
  }
  ```

#### 8.注册 Callbacks

在 models.go 的 init 函数中，增加以下删除的回调

```
db.Callback().Delete().Replace("gorm:delete", deleteCallback)
```

#### 9.验证

重启服务，访问 DeleteTag 接口，成功后即可发现 deleted_on 字段有值

#### 10.小结

在这一章节中，我们结合 GORM 完成了新增、更新、查询的 Callbacks，在实际项目中常常也是这么使用

毕竟，一个钩子的事，就没有必要自己手写过多不必要的代码了

（注意，增加了软删除后，先前的代码需要增加 `deleted_on` 的判断）



### 十一、Cron定时任务

在实际的应用项目中，定时任务的使用是很常见的。你是否有过 Golang 如何做定时任务的疑问，莫非是轮询，在本文中我们将结合我们的项目讲述 Cron。

我们将使用 cron 这个包，它实现了 cron 规范解析器和任务运行器，简单来讲就是包含了定时任务所需的功能

#### 1.Cron 表达式格式

| 字段名                         | **是否必填** | **允许的值**    | **允许的特殊字符** |
| ------------------------------ | ------------ | --------------- | ------------------ |
| 秒（Seconds）                  | Yes          | 0-59            | * / , -            |
| 分（Minutes）                  | Yes          | 0-59            | * / , -            |
| 时（Hours）                    | Yes          | 0-23            | * / , -            |
| 一个月中的某天（Day of month） | Yes          | 1-31            | * / , - ?          |
| 月（Month）                    | Yes          | 1-12 or JAN-DEC | * / , -            |
| 星期几（Day of week）          | Yes          | 0-6 or SUN-SAT  | * / , - ?          |

Cron 表达式表示一组时间，使用 6 个空格分隔的字段

可以留意到 Golang 的 Cron 比 Crontab 多了一个秒级，以后遇到秒级要求的时候就省事了



#### 2.Cron 特殊字符

1、星号 ( * )

星号表示将匹配字段的所有值

2、斜线 ( / )

斜线用户 描述范围的增量，表现为 “N-MAX/x”，first-last/x 的形式，例如 3-59/15 表示此时的第三分钟和此后的每 15 分钟，到 59 分钟为止。即从 N 开始，使用增量直到该特定范围结束。它不会重复

3、逗号 ( , )

逗号用于分隔列表中的项目。例如，在 Day of week 使用“MON，WED，FRI”将意味着星期一，星期三和星期五

4、连字符 ( - )

连字符用于定义范围。例如，9 - 17 表示从上午 9 点到下午 5 点的每个小时

5、问号 ( ? )

不指定值，用于代替 “ * ”，类似 “ _ ” 的存在，不难理解



#### 3.预定义的 Cron 时间表

| 输入                   | 简述                                   | **相当于**  |
| ---------------------- | -------------------------------------- | ----------- |
| @yearly (or @annually) | 1 月 1 日午夜运行一次                  | 0 0 0 1 1 * |
| @monthly               | 每个月的午夜，每个月的第一个月运行一次 | 0 0 0 1 * * |
| @weekly                | 每周一次，周日午夜运行一次             | 0 0 0 * * 0 |
| @daily (or @midnight)  | 每天午夜运行一次                       | 0 0 0 * * * |
| @hourly                | 每小时运行一次                         | 0 0 * * * * |



#### 4.安装

```
$ go get -u github.com/robfig/cron
```



#### 5.实践

在上一章节 Gin 实践 连载十 定制 GORM Callbacks 中，我们使用了 GORM 的回调实现了软删除，同时也引入了另外一个问题

就是我怎么硬删除，我什么时候硬删除？这个往往与业务场景有关系，大致为

- 另外有一套硬删除接口
- 定时任务清理（或转移、backup）无效数据

在这里我们选用第二种解决方案来进行实践



#### 6.编写硬删除代码

打开 models 目录下的 tag.go、article.go 文件，分别添加以下代码

```
func CleanAllTag() bool {
    db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

    return true
}
```



```
func CleanAllArticle() bool {
    db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})
    return true
}
```

注意硬删除要使用 `Unscoped()`，这是 GORM 的约定



#### 7.编写 Cron

在 项目根目录下新建 cron.go 文件，用于编写定时任务的代码，写入文件内容

```
package main

import (
    "github.com/jamesluo111/gin-blog/models"
    "github.com/robfig/cron"
    "log"
    "time"
)

func main() {
    log.Println("Starting...")

    c := cron.New()
    c.AddFunc("* * * * * *", func() {
        log.Println("Run models.CleanAllTag...")
        models.CleanAllTag()
    })

    c.AddFunc("* * * * * *", func() {
        log.Println("Run models.CleanAllArticle")
        models.CleanAllArticle()
    })

    c.Start()

    t1 := time.NewTimer(time.Second * 10)

    for {
        select {
        case <-t1.C:
            t1.Reset(time.Second * 10)
        }
    }

}
```

在这段程序中，我们做了如下的事情

**cron.New()**

会根据本地时间创建一个新（空白）的 Cron job runner

```
func New() *Cron {
    return NewWithLocation(time.Now().Location())
}

// NewWithLocation returns a new Cron job runner.
func NewWithLocation(location *time.Location) *Cron {
    return &Cron{
        entries:  nil,
        add:      make(chan *Entry),
        stop:     make(chan struct{}),
        snapshot: make(chan []*Entry),
        running:  false,
        ErrorLog: nil,
        location: location,
    }
}
```

**c.AddFunc()**

AddFunc 会向 Cron job runner 添加一个 func ，以按给定的时间表运行

```
func (c *Cron) AddJob(spec string, cmd Job) error {
    schedule, err := Parse(spec)
    if err != nil {
        return err
    }
    c.Schedule(schedule, cmd)
    return nil
}
```

会首先解析时间表，如果填写有问题会直接 err，无误则将 func 添加到 Schedule 队列中等待执行

```
func (c *Cron) Schedule(schedule Schedule, cmd Job) {
    entry := &Entry{
        Schedule: schedule,
        Job:      cmd,
    }
    if !c.running {
        c.entries = append(c.entries, entry)
        return
    }

    c.add <- entry
}
```

**c.Start()**

在当前执行的程序中启动 Cron 调度程序。其实这里的主体是 goroutine + for + select + timer 的调度控制哦

```
func (c *Cron) Run() {
    if c.running {
        return
    }
    c.running = true
    c.run()
}
```



**time.NewTimer + for + select + t1.Reset**

如果你是初学者，大概会有疑问，这是干嘛用的？

**（1）time.NewTimer** 

会创建一个新的定时器，持续你设定的时间 d 后发送一个 channel 消息

**（2）for + select**

阻塞 select 等待 channel

**（3）t1.Reset**

会重置定时器，让它重新开始计时

注：本文适用于 “t.C 已经取走，可直接使用 Reset”。

总的来说，这段程序是为了阻塞主程序而编写的，希望你带着疑问来想，有没有别的办法呢？

有的，你直接 `select{}` 也可以完成这个需求 :)



#### 8.验证

```
root@3a3932991f13:/wwwroot/gin-blog# go run cron.go

[2022-05-18 07:15:33] [info] replacing callback `gorm:update_time_stamp` from /wwwroot/gin-blog/models/models.go:38

[2022-05-18 07:15:33] [info] replacing callback `gorm:update_time_stamp` from /wwwroot/gin-blog/models/models.go:39

[2022-05-18 07:15:33] [info] replacing callback `gorm:delete` from /wwwroot/gin-blog/models/models.go:40
2022/05/18 07:15:33 Starting...
2022/05/18 07:15:34 Run models.CleanAllArticle
2022/05/18 07:15:34 Run models.CleanAllTag...
2022/05/18 07:15:34

(/wwwroot/gin-blog/models/models.go:116)
[2022-05-18 07:15:34]  [0.52ms]  DELETE FROM `blog_article`  WHERE (deleted_on != 0 )
[0 rows affected or returned ]
2022/05/18 07:15:34

(/wwwroot/gin-blog/models/models.go:116)
[2022-05-18 07:15:34]  [0.64ms]  DELETE FROM `blog_tag`  WHERE (deleted_on != 0 )
[0 rows affected or returned ]
2022/05/18 07:15:35 Run models.CleanAllArticle
2022/05/18 07:15:35 Run models.CleanAllTag...
```

检查输出日志正常，模拟已软删除的数据，定时任务工作 OK



#### 9.小结

定时任务很常见，希望你通过本文能够熟知 Golang 怎么实现一个简单的定时任务调度管理

可以不依赖系统的 Crontab 设置，指不定哪一天就用上了呢



#### 10.问题

如果你手动修改计算机的系统时间，是会导致定时任务错乱的，所以一般不要乱来。



### 十二、化配置结构及实现图片上传

一天，产品经理突然跟你说文章列表，没有封面图，不够美观，！）&￥*！&）#&￥*！加一个吧，几分钟的事

你打开你的程序，分析了一波写了个清单：

- 优化配置结构（因为配置项越来越多）
- 抽离 原 logging 的 File 便于公用（logging、upload 各保有一份并不合适）
- 实现上传图片接口（需限制文件格式、大小）
- 修改文章接口（需支持封面地址参数）
- 增加 blog_article （文章）的数据库字段
- 实现 http.FileServer

嗯，你发现要较优的话，需要调整部分的应用程序结构，因为功能越来越多，原本的设计也要跟上节奏

也就是在适当的时候，及时优化



#### 1.优化配置结构

##### 一、讲解

在先前章节中，采用了直接读取 KEY 的方式去存储配置项，而本次需求中，需要增加图片的配置项，总体就有些冗余了

我们采用以下解决方法：

- 映射结构体：使用 MapTo 来设置配置参数
- 配置统管：所有的配置项统管到 setting 中

**映射结构体（示例）**

在 go-ini 中可以采用 MapTo 的方式来映射结构体，例如：

```
type Server struct {
    RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func main() {
    Cfg, err := ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    err = Cfg.Section("server").MapTo(ServerSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }
}
```

在这段代码中，可以注意 ServerSetting 取了地址，为什么 MapTo 必须地址入参呢？

```
// MapTo maps section to given struct.
func (s *Section) MapTo(v interface{}) error {
    typ := reflect.TypeOf(v)
    val := reflect.ValueOf(v)
    if typ.Kind() == reflect.Ptr {
        typ = typ.Elem()
        val = val.Elem()
    } else {
        return errors.New("cannot map to non-pointer struct")
    }

    return s.mapTo(val, false)
}
```

在 MapTo 中 `typ.Kind() == reflect.Ptr` 约束了必须使用指针，否则会返回 `cannot map to non-pointer struct` 的错误。这个是表面原因

更往内探究，可以认为是 `field.Set` 的原因，当执行 `val := reflect.ValueOf(v)` ，函数通过传递 `v` 拷贝创建了 `val`，但是 `val` 的改变并不能更改原始的 `v`，要想 `val` 的更改能作用到 `v`，则必须传递 `v` 的地址

显然 go-ini 里也是包含修改原始值这一项功能的，你觉得是什么原因呢？

**配置统管**

在先前的版本中，models 和 file 的配置是在自己的文件中解析的，而其他在 setting.go 中，因此我们需要将其在 setting 中统一接管

你可能会想，直接把两者的配置项复制粘贴到 setting.go 的 init 中，一下子就完事了，搞那么麻烦？

但你在想想，先前的代码中存在多个 init 函数，执行顺序存在问题，无法达到我们的要求，你可以试试

（此处是一个基础知识点）

在 Go 中，当存在多个 init 函数时，执行顺序为：

- 相同包下的 init 函数：按照源文件编译顺序决定执行顺序（默认按文件名排序）
- 不同包下的 init 函数：按照包导入的依赖关系决定先后顺序

所以要避免多 init 的情况，**尽量由程序把控初始化的先后顺序**



##### 二、落实

**修改配置文件**

打开 conf/app.ini 将配置文件修改为大驼峰命名，另外我们增加了 5 个配置项用于上传图片的功能，4 个文件日志方面的配置项

```
#debug or release
RUN_MODE = debug

[app]
PAGE_SIZE = 10
JWT_SECRET = 233

RuntimeRootPath = runtime/

ImagePrefixUrl = http://127.0.0.1:8000
ImageSavePath = upload/images/
# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
#debug or release
RunMode = debug
HttpPort = 8000
ReadTimeout = 60
WriteTimeout = 60

[database]
TYPE = mysql
USER = root
PASSWORD = root
HOST = mysql:3306
NAME = blog
TABLE_PREFIX = blog_
```

**优化配置读取及设置初始化顺序**

将散落在其他文件里的配置都删掉，**统一在 setting 中处理**以及**修改 init 函数为 Setup 方法**

打开 pkg/setting/setting.go 文件，修改如下：

```

package setting

import (
    "log"
    "time"

    "github.com/go-ini/ini"
)

type App struct {
    JwtSecret string
    PageSize int
    RuntimeRootPath string

    ImagePrefixUrl string
    ImageSavePath string
    ImageMaxSize int
    ImageAllowExts []string

    LogSavePath string
    LogSaveName string
    LogFileExt string
    TimeFormat string
}

var AppSetting = &App{}

type Server struct {
    RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
    Type string
    User string
    Password string
    Host string
    Name string
    TablePrefix string
}

var DatabaseSetting = &Database{}

func Setup() {
    Cfg, err := ini.Load("conf/app.ini")
    if err != nil {
        log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
    }

    err = Cfg.Section("app").MapTo(AppSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
    }

    AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024

    err = Cfg.Section("server").MapTo(ServerSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
    }

    ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
    ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

    err = Cfg.Section("database").MapTo(DatabaseSetting)
    if err != nil {
        log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
    }
}
```

在这里，我们做了如下几件事：

- 编写与配置项保持一致的结构体（App、Server、Database）
- 使用 MapTo 将配置项映射到结构体上
- 对一些需特殊设置的配置项进行再赋值

**需要你去做的事：**

- 将 models.go、setting.go、pkg/logging/log.go 的 init 函数修改为 Setup 方法
- 将 models/models.go 独立读取的 DB 配置项删除，改为统一读取 setting
- 将 pkg/logging/file 独立的 LOG 配置项删除，改为统一读取 setting

这几项比较基础，并没有贴出来，我希望你可以自己动手，有问题的话可右拐 项目地址



在这一步我们要设置初始化的流程，打开 main.go 文件，修改内容：

```
func main() {
    setting.Setup()
    models.Setup()
    logging.Setup()

    endless.DefaultReadTimeOut = setting.ServerSetting.ReadTimeout
    endless.DefaultWriteTimeOut = setting.ServerSetting.WriteTimeout
    endless.DefaultMaxHeaderBytes = 1 << 20
    endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

    server := endless.NewServer(endPoint, routers.InitRouter())
    server.BeforeBegin = func(add string) {
        log.Printf("Actual pid is %d", syscall.Getpid())
    }

    err := server.ListenAndServe()
    if err != nil {
        log.Printf("Server err: %v", err)
    }
}
```

修改完毕后，就成功将多模块的初始化函数放到启动流程中了（先后顺序也可以控制）

**验证**

在这里为止，针对本需求的配置优化就完毕了，你需要执行 `go run main.go` 验证一下你的功能是否正常哦

顺带留个基础问题，大家可以思考下

```
ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
```

若将 setting.go 文件中的这两行删除，会出现什么问题，为什么呢？



#### 2.抽离 File

在先前版本中，在 logging/file.go 中使用到了 os 的一些方法，我们通过前期规划发现，这部分在上传图片功能中可以复用

##### 第一步

```

package file

import (
    "os"
    "path"
    "mime/multipart"
    "io/ioutil"
)

func GetSize(f multipart.File) (int, error) {
    content, err := ioutil.ReadAll(f)

    return len(content), err
}

func GetExt(fileName string) string {
    return path.Ext(fileName)
}

func CheckNotExist(src string) bool {
    _, err := os.Stat(src)

    return os.IsNotExist(err)
}

func CheckPermission(src string) bool {
    _, err := os.Stat(src)

    return os.IsPermission(err)
}

func IsNotExistMkDir(src string) error {
    if notExist := CheckNotExist(src); notExist == true {
        if err := MkDir(src); err != nil {
            return err
        }
    }

    return nil
}

func MkDir(src string) error {
    err := os.MkdirAll(src, os.ModePerm)
    if err != nil {
        return err
    }

    return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
    f, err := os.OpenFile(name, flag, perm)
    if err != nil {
        return nil, err
    }

    return f, nil
}
```

在这里我们一共封装了 7 个 方法

- GetSize：获取文件大小
- GetExt：获取文件后缀
- CheckNotExist：检查文件是否存在
- CheckPermission：检查文件权限
- IsNotExistMkDir：如果不存在则新建文件夹
- MkDir：新建文件夹
- Open：打开文件

在这里我们用到了 `mime/multipart` 包，它主要实现了 MIME 的 multipart 解析，主要适用于 HTTP 和常见浏览器生成的 multipart 主体

multipart 又是什么，rfc2388 的 multipart/form-data 了解一下



##### 第二步

在 pkg 目录下新建 file/file.go ，写入文件内容如下：

我们在第一步已经将 file 重新封装了一层，在这一步我们将原先 logging 包的方法都修改掉

1、打开 pkg/logging/file.go 文件，修改文件内容：

```
package logging

import (
    "fmt"
    "os"
    "time"

    "github.com/EDDYCJY/go-gin-example/pkg/setting"
    "github.com/EDDYCJY/go-gin-example/pkg/file"
)

func getLogFilePath() string {
    return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.LogSavePath)
}

func getLogFileName() string {
    return fmt.Sprintf("%s%s.%s",
        setting.AppSetting.LogSaveName,
        time.Now().Format(setting.AppSetting.TimeFormat),
        setting.AppSetting.LogFileExt,
    )
}

func openLogFile(fileName, filePath string) (*os.File, error) {
    dir, err := os.Getwd()
    if err != nil {
        return nil, fmt.Errorf("os.Getwd err: %v", err)
    }

    src := dir + "/" + filePath
    perm := file.CheckPermission(src)
    if perm == true {
        return nil, fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
    }

    err = file.IsNotExistMkDir(src)
    if err != nil {
        return nil, fmt.Errorf("file.IsNotExistMkDir src: %s, err: %v", src, err)
    }

    f, err := file.Open(src + fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return nil, fmt.Errorf("Fail to OpenFile :%v", err)
    }

    return f, nil
}
```

我们将引用都改为了 file/file.go 包里的方法

2、打开 pkg/logging/log.go 文件，修改文件内容:

```
package logging

...

func Setup() {
    var err error
    filePath := getLogFilePath()
    fileName := getLogFileName()
    F, err = openLogFile(fileName, filePath)
    if err != nil {
        log.Fatalln(err)
    }

    logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

...
```

由于原方法形参改变了，因此 openLogFile 也需要调整



#### 3.实现图片上传功能

这一小节，我们开始实现上次图片相关的一些方法和功能

首先需要在 blog_article 中增加字段 `cover_image_url`，格式为 `varchar(255) DEFAULT '' COMMENT '封面图片地址'`

##### 第零步

一般不会直接将上传的图片名暴露出来，因此我们对图片名进行 MD5 来达到这个效果

在 util 目录下新建 md5.go，写入文件内容：

```
package util

import (
    "crypto/md5"
    "encoding/hex"
)

func EncodeMD5(value string) string {
    m := md5.New()
    m.Write([]byte(value))

    return hex.EncodeToString(m.Sum(nil))
}
```



##### 第一步

在先前我们已经把底层方法给封装好了，实质这一步为封装 image 的处理逻辑

在 pkg 目录下新建 upload/image.go 文件，写入文件内容：

```
package upload

import (
    "os"
    "path"
    "log"
    "fmt"
    "strings"
    "mime/multipart"

    "github.com/EDDYCJY/go-gin-example/pkg/file"
    "github.com/EDDYCJY/go-gin-example/pkg/setting"
    "github.com/EDDYCJY/go-gin-example/pkg/logging"
    "github.com/EDDYCJY/go-gin-example/pkg/util"
)

func GetImageFullUrl(name string) string {
    return setting.AppSetting.ImagePrefixUrl + "/" + GetImagePath() + name
}

func GetImageName(name string) string {
    ext := path.Ext(name)
    fileName := strings.TrimSuffix(name, ext)
    fileName = util.EncodeMD5(fileName)

    return fileName + ext
}

func GetImagePath() string {
    return setting.AppSetting.ImageSavePath
}

func GetImageFullPath() string {
    return setting.AppSetting.RuntimeRootPath + GetImagePath()
}

func CheckImageExt(fileName string) bool {
    ext := file.GetExt(fileName)
    for _, allowExt := range setting.AppSetting.ImageAllowExts {
        if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
            return true
        }
    }

    return false
}

func CheckImageSize(f multipart.File) bool {
    size, err := file.GetSize(f)
    if err != nil {
        log.Println(err)
        logging.Warn(err)
        return false
    }

    return size <= setting.AppSetting.ImageMaxSize
}

func CheckImage(src string) error {
    dir, err := os.Getwd()
    if err != nil {
        return fmt.Errorf("os.Getwd err: %v", err)
    }

    err = file.IsNotExistMkDir(dir + "/" + src)
    if err != nil {
        return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
    }

    perm := file.CheckPermission(src)
    if perm == true {
        return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
    }

    return nil
}
```

在这里我们实现了 7 个方法，如下：

- GetImageFullUrl：获取图片完整访问 URL
- GetImageName：获取图片名称
- GetImagePath：获取图片路径
- GetImageFullPath：获取图片完整路径
- CheckImageExt：检查图片后缀
- CheckImageSize：检查图片大小
- CheckImage：检查图片

这里基本是对底层代码的二次封装，为了更灵活的处理一些图片特有的逻辑，并且方便修改，不直接对外暴露下层



##### 第二步

这一步将编写上传图片的业务逻辑，在 routers/api 目录下 新建 upload.go 文件，写入文件内容:

```

package api

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/EDDYCJY/go-gin-example/pkg/e"
    "github.com/EDDYCJY/go-gin-example/pkg/logging"
    "github.com/EDDYCJY/go-gin-example/pkg/upload"
)

func UploadImage(c *gin.Context) {
    code := e.SUCCESS
    data := make(map[string]string)

    file, image, err := c.Request.FormFile("image")
    if err != nil {
        logging.Warn(err)
        code = e.ERROR
        c.JSON(http.StatusOK, gin.H{
            "code": code,
            "msg":  e.GetMsg(code),
            "data": data,
        })
    }

    if image == nil {
        code = e.INVALID_PARAMS
    } else {
        imageName := upload.GetImageName(image.Filename)
        fullPath := upload.GetImageFullPath()
        savePath := upload.GetImagePath()

        src := fullPath + imageName
        if ! upload.CheckImageExt(imageName) || ! upload.CheckImageSize(file) {
            code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
        } else {
            err := upload.CheckImage(fullPath)
            if err != nil {
                logging.Warn(err)
                code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
            } else if err := c.SaveUploadedFile(image, src); err != nil {
                logging.Warn(err)
                code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
            } else {
                data["image_url"] = upload.GetImageFullUrl(imageName)
                data["image_save_url"] = savePath + imageName
            }
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "code": code,
        "msg":  e.GetMsg(code),
        "data": data,
    })
}
```

所涉及的错误码（需在 pkg/e/code.go、msg.go 添加）：

```
// 保存图片失败
ERROR_UPLOAD_SAVE_IMAGE_FAIL = 30001
// 检查图片失败
ERROR_UPLOAD_CHECK_IMAGE_FAIL = 30002
// 校验图片错误，图片格式或大小有问题
ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
```

在这一大段的业务逻辑中，我们做了如下事情：

- c.Request.FormFile：获取上传的图片（返回提供的表单键的第一个文件）
- CheckImageExt、CheckImageSize 检查图片大小，检查图片后缀
- CheckImage：检查上传图片所需（权限、文件夹）
- SaveUploadedFile：保存图片

总的来说，就是 入参 -> 检查 -》 保存 的应用流程

##### 第三步

打开 routers/router.go 文件，增加路由 `r.POST("/upload", api.UploadImage)` ，如：

```
func InitRouter() *gin.Engine {
    r := gin.New()
    ...
    r.GET("/auth", api.GetAuth)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/upload", api.UploadImage)

    apiv1 := r.Group("/api/v1")
    apiv1.Use(jwt.JWT())
    {
        ...
    }

    return r
}
```

##### 验证

最后我们请求一下上传图片的接口，测试所编写的功能

![image-20220520115146609](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220520115146609.png?lastModify=1653785198)



#### 3.实现 http.FileServer

在完成了上一小节后，我们还需要让前端能够访问到图片，一般是如下：

- CDN
- http.FileSystem

在公司的话，CDN 或自建分布式文件系统居多，也不需要过多关注。而在实践里的话肯定是本地搭建了，Go 本身对此就有很好的支持，而 Gin 更是再封装了一层，只需要在路由增加一行代码即可

##### r.StaticFS

打开 routers/router.go 文件，增加路由 `r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))`，如：

```
func InitRouter() *gin.Engine {
    ...
    r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

    r.GET("/auth", api.GetAuth)
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.POST("/upload", api.UploadImage)
    ...
}
```

##### 它做了什么

当访问host/upload/images时，他会读取到 GOPATH/src/github.com/jamesluo111/gin-blog/runtime/upload/images 下的文件

而这行代码又做了什么事呢，我们来看看方法原型

```
// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
// Gin by default user: gin.Dir()
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem) IRoutes {
    if strings.Contains(relativePath, ":") || strings.Contains(relativePath, "*") {
        panic("URL parameters can not be used when serving a static folder")
    }
    handler := group.createStaticHandler(relativePath, fs)
    urlPattern := path.Join(relativePath, "/*filepath")

    // Register GET and HEAD handlers
    group.GET(urlPattern, handler)
    group.HEAD(urlPattern, handler)
    return group.returnObj()
}
```

首先在暴露的 URL 中禁止了 * 和 : 符号的使用，通过 `createStaticHandler` 创建了静态文件服务，实质最终调用的还是 `fileServer.ServeHTTP` 和一些处理逻辑了

```
func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
    absolutePath := group.calculateAbsolutePath(relativePath)
    fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
    _, nolisting := fs.(*onlyfilesFS)
    return func(c *Context) {
        if nolisting {
            c.Writer.WriteHeader(404)
        }
        fileServer.ServeHTTP(c.Writer, c.Request)
    }
}
```

**http.StripPrefix**

我们可以留意下 `fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))` 这段语句，在静态文件服务中很常见，它有什么作用呢？

```
http.StripPrefix` 主要作用是从请求 URL 的路径中删除给定的前缀，最终返回一个 `Handler
```

通常 http.FileServer 要与 http.StripPrefix 相结合使用，否则当你运行：

```
http.Handle("/upload/images", http.FileServer(http.Dir("upload/images")))
```

会无法正确的访问到文件目录，因为 `/upload/images` 也包含在了 URL 路径中，必须使用：

```
http.Handle("/upload/images", http.StripPrefix("upload/images", http.FileServer(http.Dir("upload/images"))))
```

**/\*filepath**

到下面可以看到 `urlPattern := path.Join(relativePath, "/*filepath")`，`/*filepath` 你是谁，你在这里有什么用，你是 Gin 的产物吗?

通过语义可得知是路由的处理逻辑，而 Gin 的路由是基于 httprouter 的，通过查阅文档可得到以下信息

```
Pattern: /src/*filepath

 /src/                     match
 /src/somefile.go          match
 /src/subdir/somefile.go   match
```

`*filepath` 将匹配所有文件路径，并且 `*filepath` 必须在 Pattern 的最后

##### 验证

重新执行 `go run main.go` ，去访问刚刚在 upload 接口得到的图片地址，检查 http.FileSystem 是否正常

![image-20220520141215250](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220520141215250.png?lastModify=1653785198)



#### 4.修改文章接口

接下来，需要你修改 routers/api/v1/article.go 的 AddArticle、EditArticle 两个接口

- 新增、更新文章接口：支持入参 cover_image_url
- 新增、更新文章接口：增加对 cover_image_url 的非空、最长长度校验

这块前面文章讲过，如果有问题可以参考项目的代码 👌

#### 5.总结

在这章节中，我们简单的分析了下需求，对应用做出了一个小规划并实施

完成了清单中的功能点和优化，在实际项目中也是常见的场景，希望你能够细细品尝并针对一些点进行深入学习



### 十三、优化应用结构和实现Redis缓存

在本章节，将介绍以下功能的整理：

- 抽离、分层业务逻辑：减轻 routers.go 内的 api 方法的逻辑（但本文暂不分层 repository，这块逻辑还不重）。
- 增加容错性：对 gorm 的错误进行判断。
- Redis 缓存：对获取数据类的接口增加缓存设置。
- 减少重复冗余代码



#### 1.问题在哪？

在规划阶段我们发现了一个问题，这是目前的伪代码：

```
if ! HasErrors() {
    if ExistArticleByID(id) {
        DeleteArticle(id)
        code = e.SUCCESS
    } else {
        code = e.ERROR_NOT_EXIST_ARTICLE
    }
} else {
    for _, err := range valid.Errors {
        logging.Info(err.Key, err.Message)
    }
}

c.JSON(http.StatusOK, gin.H{
    "code": code,
    "msg":  e.GetMsg(code),
    "data": make(map[string]string),
})
```

如果加上规划内的功能逻辑呢，伪代码会变成：

```
if ! HasErrors() {
    exists, err := ExistArticleByID(id)
    if err == nil {
        if exists {
            err = DeleteArticle(id)
            if err == nil {
                code = e.SUCCESS
            } else {
                code = e.ERROR_XXX
            }
        } else {
            code = e.ERROR_NOT_EXIST_ARTICLE
        }
    } else {
        code = e.ERROR_XXX
    }
} else {
    for _, err := range valid.Errors {
        logging.Info(err.Key, err.Message)
    }
}

c.JSON(http.StatusOK, gin.H{
    "code": code,
    "msg":  e.GetMsg(code),
    "data": make(map[string]string),
})
```

如果缓存的逻辑也加进来，后面慢慢不断的迭代，岂不是会变成如下图一样？

![image-20220520144302524](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220520144302524.png?lastModify=1653785198)

现在我们发现了问题，应及时解决这个代码结构问题，同时把代码写的清晰、漂亮、易读易改也是一个重要指标



#### 2.如何改？

这类代码被称为 “箭头型” 代码，有如下几个问题：

1、我的显示器不够宽，箭头型代码缩进太狠了，需要我来回拉水平滚动条，这让我在读代码的时候，相当的不舒服

2、除了宽度外还有长度，有的代码的 if-else 里的 if-else 里的 if-else 的代码太多，读到中间你都不知道中间的代码是经过了什么样的层层检查才来到这里的

总而言之，“箭头型代码”如果嵌套太多，代码太长的话，会相当容易让维护代码的人（包括自己）迷失在代码中，因为看到最内层的代码时，你已经不知道前面的那一层一层的条件判断是什么样的，代码是怎么运行到这里的，所以，箭头型代码是非常难以维护和 Debug 的。

简单的来说，就是**让出错的代码先返回，前面把所有的错误判断全判断掉，然后就剩下的就是正常的代码了**

##### 落实

本项目将对既有代码进行优化和实现缓存，希望你习得方法并对其他地方也进行优化

第一步：完成 Redis 的基础设施建设（需要你先装好 Redis）

第二步：对现有代码进行拆解、分层（不会贴上具体步骤的代码，希望你能够实操一波，加深理解 🤔）



#### 3.redis

##### 一、配置

打开 conf/app.ini 文件，新增配置：

```
...
[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
```

##### 二、缓存 Prefix

打开 pkg/e 目录，新建 cache.go，写入内容：

```
package e

const (
    CACHE_ARTICLE = "ARTICLE"
    CACHE_TAG     = "TAG"
)
```

##### 三、缓存 Key

（1）、打开 service 目录，新建 cache_service/article.go

```
package cache_service

import (
    "github.com/jamesluo111/gin-blog/pkg/e"
    "strconv"
    "strings"
)

type Article struct {
    ID    int
    TagId int
    State int

    PageNum  int
    PageSize int
}

func (a *Article) GetArticleKey() string {
    return e.CACHE_ARTICLE + "_" + strconv.Itoa(a.ID)
}

func (a *Article) GetArticlesKey() string {
    keys := []string{
        e.CACHE_ARTICLE,
        "LIST",
    }

    if a.ID > 0 {
        keys = append(keys, strconv.Itoa(a.ID))
    }

    if a.TagId > 0 {
        keys = append(keys, strconv.Itoa(a.TagId))
    }

    if a.State >= 0 {
        keys = append(keys, strconv.Itoa(a.State))
    }

    if a.PageNum > 0 {
        keys = append(keys, strconv.Itoa(a.PageNum))
    }

    if a.PageSize > 0 {
        keys = append(keys, strconv.Itoa(a.PageSize))
    }

    return strings.Join(keys, "_")
}

```



（2）、打开 service 目录，新建 cache_service/tag.go

```
package cache_service

import (
    "github.com/jamesluo111/gin-blog/pkg/e"
    "strconv"
    "strings"
)

type Tag struct {
    Id    int
    Name  string
    State int

    PageNum  int
    PageSize int
}

func (t *Tag) GetTagsKey() string {
    keys := []string{
        e.CACHE_TAG,
        "LIST",
    }

    if t.Name != "" {
        keys = append(keys, t.Name)
    }

    if t.State >= 0 {
        keys = append(keys, strconv.Itoa(t.State))
    }

    if t.PageNum > 0 {
        keys = append(keys, strconv.Itoa(t.PageNum))
    }

    if t.PageSize > 0 {
        keys = append(keys, strconv.Itoa(t.PageSize))
    }

    return strings.Join(keys, "_")
}
```

##### 四、Redis 工具包

打开 pkg 目录，新建 gredis/redis.go，写入内容：

```
package gredis

import (
    "encoding/json"
    "time"

    "github.com/gomodule/redigo/redis"

    "github.com/EDDYCJY/go-gin-example/pkg/setting"
)

var RedisConn *redis.Pool

func Setup() error {
    RedisConn = &redis.Pool{
        MaxIdle:     setting.RedisSetting.MaxIdle,
        MaxActive:   setting.RedisSetting.MaxActive,
        IdleTimeout: setting.RedisSetting.IdleTimeout,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", setting.RedisSetting.Host)
            if err != nil {
                return nil, err
            }
            if setting.RedisSetting.Password != "" {
                if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
                    c.Close()
                    return nil, err
                }
            }
            return c, err
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
    }

    return nil
}

func Set(key string, data interface{}, time int) error {
    conn := RedisConn.Get()
    defer conn.Close()

    value, err := json.Marshal(data)
    if err != nil {
        return err
    }

    _, err = conn.Do("SET", key, value)
    if err != nil {
        return err
    }

    _, err = conn.Do("EXPIRE", key, time)
    if err != nil {
        return err
    }

    return nil
}

func Exists(key string) bool {
    conn := RedisConn.Get()
    defer conn.Close()

    exists, err := redis.Bool(conn.Do("EXISTS", key))
    if err != nil {
        return false
    }

    return exists
}

func Get(key string) ([]byte, error) {
    conn := RedisConn.Get()
    defer conn.Close()

    reply, err := redis.Bytes(conn.Do("GET", key))
    if err != nil {
        return nil, err
    }

    return reply, nil
}

func Delete(key string) (bool, error) {
    conn := RedisConn.Get()
    defer conn.Close()

    return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
    conn := RedisConn.Get()
    defer conn.Close()

    keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
    if err != nil {
        return err
    }

    for _, key := range keys {
        _, err = Delete(key)
        if err != nil {
            return err
        }
    }

    return nil
}
```

在这里我们做了一些基础功能封装

1、设置 RedisConn 为 redis.Pool（连接池）并配置了它的一些参数：

- Dial：提供创建和配置应用程序连接的一个函数
- TestOnBorrow：可选的应用程序检查健康功能
- MaxIdle：最大空闲连接数
- MaxActive：在给定时间内，允许分配的最大连接数（当为零时，没有限制）
- IdleTimeout：在给定时间内将会保持空闲状态，若到达时间限制则关闭连接（当为零时，没有限制）

2、封装基础方法

文件内包含 Set、Exists、Get、Delete、LikeDeletes 用于支撑目前的业务逻辑，而在里面涉及到了如方法：

（1）`RedisConn.Get()`：在连接池中获取一个活跃连接

（2）`conn.Do(commandName string, args ...interface{})`：向 Redis 服务器发送命令并返回收到的答复

（3）`redis.Bool(reply interface{}, err error)`：将命令返回转为布尔值

（4）`redis.Bytes(reply interface{}, err error)`：将命令返回转为 Bytes

（5）`redis.Strings(reply interface{}, err error)`：将命令返回转为 []string

在 redigo 中包含大量类似的方法，万变不离其宗，建议熟悉其使用规则和 Redis 命令 即可

到这里为止，Redis 就可以愉快的调用啦。另外受篇幅限制，这块的深入讲解会另外开设！



#### 4.拆解、分层

在先前规划中，引出几个方法去优化我们的应用结构

- 错误提前返回
- 统一返回方法
- 抽离 Service，减轻 routers/api 的逻辑，进行分层
- 增加 gorm 错误判断，让错误提示更明确（增加内部错误码）



#### 5.编写返回方法

要让错误提前返回，c.JSON 的侵入是不可避免的，但是可以让其更具可变性，指不定哪天就变 XML 了呢？

1、打开 pkg 目录，新建 app/request.go，写入文件内容：

```
package app

import (
    "github.com/astaxie/beego/validation"

    "github.com/EDDYCJY/go-gin-example/pkg/logging"
)

func MarkErrors(errors []*validation.Error) {
    for _, err := range errors {
        logging.Info(err.Key, err.Message)
    }

    return
}
```

2、打开 pkg 目录，新建 app/response.go，写入文件内容：

```
package app

import (
    "github.com/gin-gonic/gin"

    "github.com/EDDYCJY/go-gin-example/pkg/e"
)

type Gin struct {
    C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}) {
    g.C.JSON(httpCode, gin.H{
        "code": errCode,
        "msg":  e.GetMsg(errCode),
        "data": data,
    })

    return
}
```

这样子以后如果要变动，直接改动 app 包内的方法即可



#### 6.修改既有逻辑

打开 routers/api/v1/article.go，查看修改 GetArticle 方法后的代码为：

```
func GetArticle(c *gin.Context) {
    appG := app.Gin{c}
    id := com.StrTo(c.Param("id")).MustInt()
    valid := validation.Validation{}
    valid.Min(id, 1, "id").Message("ID必须大于0")

    if valid.HasErrors() {
        app.MarkErrors(valid.Errors)
        appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
        return
    }

    articleService := article_service.Article{ID: id}
    exists, err := articleService.ExistByID()
    if err != nil {
        appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
        return
    }
    if !exists {
        appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
        return
    }

    article, err := articleService.Get()
    if err != nil {
        appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
        return
    }

    appG.Response(http.StatusOK, e.SUCCESS, article)
}
```

这里有几个值得变动点，主要是在内部增加了错误返回，如果存在错误则直接返回。另外进行了分层，业务逻辑内聚到了 service 层中去，而 routers/api（controller）显著减轻，代码会更加的直观

例如 service/article_service 下的 `articleService.Get()` 方法：

```
func (a *Article) Get() (*models.Article, error) {
    var cacheArticle *models.Article

    cache := cache_service.Article{ID: a.ID}
    key := cache.GetArticleKey()
    if gredis.Exists(key) {
        data, err := gredis.Get(key)
        if err != nil {
            logging.Info(err)
        } else {
            json.Unmarshal(data, &cacheArticle)
            return cacheArticle, nil
        }
    }

    article, err := models.GetArticle(a.ID)
    if err != nil {
        return nil, err
    }

    gredis.Set(key, article, 3600)
    return article, nil
}
```

而对于 gorm 的 错误返回设置，只需要修改 models/article.go 如下:

```
func GetArticle(id int) (*Article, error) {
	var article Article
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Related(&article.Tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &article, nil
}
```

习惯性增加 .Error，把控绝大部分的错误。另外需要注意一点，在 gorm 中，查找不到记录也算一种 “错误” 哦



#### 7.最后

重构路由，service，model方法。详情请查看https://github.com/jamesluo111/gin-blog



### 十四、实现导出、导入 Excel

在本节，我们将实现对标签信息的导出、导入功能，这是很标配功能了，希望你掌握基础的使用方式。

另外在本文我们使用了 2 个 Excel 的包，excelize 最初的 XML 格式文件的一些结构，是通过 tealeg/xlsx 格式文件结构演化而来的，因此特意在此都展示了，你可以根据自己的场景和喜爱去使用。

#### 1.配置

首先要指定导出的 Excel 文件的存储路径，在 app.ini 中增加配置：

```
[app]
...

ExportSavePath = export/
```

修改 setting.go 的 App struct：

```
type App struct {
	JwtSecret       string
	PageSize        int
	PrefixUrl       string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}
```

在这里需增加 ExportSavePath 配置项，另外将先前 ImagePrefixUrl 改为 PrefixUrl 用于支撑两者的 HOST 获取

（注意修改 image.go 的 GetImageFullUrl 方法）



#### 2.pkg

新建 pkg/export/excel.go 文件，如下：

```
package export

import "github.com/EDDYCJY/go-gin-example/pkg/setting"

func GetExcelFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetExcelPath() + name
}

func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
```

这里编写了一些常用的方法，以后取值方式如果有变动，直接改内部代码即可，对外不可见



#### 3.尝试一下标准库

使用单元测试：在当前文件下创建测试文件export_test.go并写入内容：

```
package export

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestExport(t *testing.T) {
	f, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(f)
	data := [][]string{
		{"1", "test1", "test1-1"},
		{"2", "test2", "test2-1"},
		{"3", "test3", "test3-1"},
	}

	w.WriteAll(data)
}
```

在 Go 提供的标准库 encoding/csv 中，天然的支持 csv 文件的读取和处理，在本段代码中，做了如下工作：

1、os.Create：

创建了一个 test.csv 文件

2、f.WriteString("\xEF\xBB\xBF")：

`\xEF\xBB\xBF` 是 UTF-8 BOM 的 16 进制格式，在这里的用处是标识文件的编码格式，通常会出现在文件的开头，因此第一步就要将其写入。如果不标识 UTF-8 的编码格式的话，写入的汉字会显示为乱码

3、csv.NewWriter：

```
func NewWriter(w io.Writer) *Writer {
	return &Writer{
		Comma: ',',
		w:     bufio.NewWriter(w),
	}
}
```

4、w.WriteAll：

```
func (w *Writer) WriteAll(records [][]string) error {
	for _, record := range records {
		err := w.Write(record)
		if err != nil {
			return err
		}
	}
	return w.w.Flush()
}
```

WriteAll 实际是对 Write 的封装，需要注意在最后调用了 `w.w.Flush()`，这充分了说明了 WriteAll 的使用场景，你可以想想作者的设计用意

#### 4.导出

##### 1.service方法

打开 service/tag.go，增加 Export 方法，如下：

```

func (t *Tag) Export() (string, error) {
    tags, err := t.GetAll()
    if err != nil {
        return "", err
    }

    file := xlsx.NewFile()
    sheet, err := file.AddSheet("标签信息")
    if err != nil {
        return "", err
    }

    titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
    row := sheet.AddRow()

    var cell *xlsx.Cell
    for _, title := range titles {
        cell = row.AddCell()
        cell.Value = title
    }

    for _, v := range tags {
        values := []string{
            strconv.Itoa(v.ID),
            v.Name,
            v.CreatedBy,
            strconv.Itoa(v.CreatedOn),
            v.ModifiedBy,
            strconv.Itoa(v.ModifiedOn),
        }

        row = sheet.AddRow()
        for _, value := range values {
            cell = row.AddCell()
            cell.Value = value
        }
    }

    time := strconv.Itoa(int(time.Now().Unix()))
    filename := "tags-" + time + ".xlsx"

    fullPath := export.GetExcelFullPath() + filename
    err = file.Save(fullPath)
    if err != nil {
        return "", err
    }

    return filename, nil
}
```

##### 2.routers 入口

打开 routers/api/v1/tag.go，增加如下方法：

```
func ExportTag(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.PostForm("name")
	state := -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
	}

	tagService := tag_service.Tag{
		Name:  name,
		State: state,
	}

	filename, err := tagService.Export()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_EXPORT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"export_url":      export.GetExcelFullUrl(filename),
		"export_save_url": export.GetExcelPath() + filename,
	})
}
```

##### 3.路由

在 routers/router.go 文件中增加路由方法，如下

```
apiv1 := r.Group("/api/v1")
apiv1.Use(jwt.JWT())
{
	...
	//导出标签
	r.POST("/tags/export", v1.ExportTag)
}
```

##### 4.验证接口

访问 `http://127.0.0.1:8000/tags/export`，结果如下：

```
{
    "code": 200,
    "data": {
        "export_save_url": "export/tags-1653442382.xlsx",
        "export_url": "http://127.0.0.1:8000/export/tags-1653442382.xlsx"
    },
    "msg": "ok"
}
```

最终通过接口返回了导出文件的地址和保存地址

##### 5.StaticFS

那你想想，现在直接访问地址肯定是无法下载文件的，那么该如何做呢？

打开 router.go 文件，增加代码如下：

```
r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
```

若你不理解，强烈建议温习下前面的章节，举一反三

##### 6.验证下载

再次访问上面的 export_url ，如：`http://127.0.0.1:8000/export/tags-1653442382.xlsx`，是不是成功了呢？



#### 5.导入

##### 1.Service 方法

需要提前引入excelize依赖：

```
go get -u github.com/360EntSecGroup-Skylar/excelize
```

此文档学习地址：[介绍 · Excelize 简体字文档 (xuri.me)](https://xuri.me/excelize/zh-hans/)

打开 service/tag.go，增加 Import 方法，如下：

```
func (t *Tag) Import(r io.Reader) error {
	//打开excel文件流
	xlsx1, err := excelize.OpenReader(r)
	if err != nil {
		return err
	}

	//打开标签信息工作表
	rows := xlsx1.GetRows("标签信息")
	for irow, row := range rows {
		//判断excel表头是否按照要求
		if irow == 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			if data[1] != "名称" && data[2] != "创建人" {
				return errors.New("格式错误!")
			}
		}
		//irow行数据,判断工作表内是否为空
		if irow > 0 {
			var data []string
			for _, cell := range row {
				data = append(data, cell)
			}
			var maps = map[string]interface{}{
				"name":       data[1],
				"state":      1,
				"created_by": data[2],
			}
			if err := models.AddTag(maps); err != nil {
				return err
			}
		}

	}
	return nil
}
```

##### 2.路由

在 routers/router.go 文件中增加路由方法，如下：

```
apiv1 := r.Group("/api/v1")
apiv1.Use(jwt.JWT())
{
	...
	//导入标签
	r.POST("/tags/import", v1.ImportTag)
}
```

##### 3.验证

![image-20220525135436811](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220525135436811.png?lastModify=1653785198)

在这里我们将先前导出的 Excel 文件作为入参，访问 `http://127.0.0.01:8000/tags/import`，检查返回和数据是否正确入库



#### 6.总结

在本文中，简单介绍了 Excel 的导入、导出的使用方式，使用了以下 2 个包：

- tealeg/xlsx
- 360EntSecGroup-Skylar/excelize

你可以细细阅读一下它的实现和使用方式，对你的把控更有帮助

我做了导出功能的重写(使用excelize进行重写，可以用来参考)

```
func (t *Tag) Export() (string, error) {
	//先从数据库中获取数据
	tags, err := t.Get()
	if err != nil {
		return "", err
	}
	x1 := excelize.NewFile()
    //重置工作表名
	x1.SetSheetName("Sheet1", "标签信息")
	row := 1
	for index, v := range tags {
		if index == 0 {
			//设置表头信息
			title := []string{
				"ID",
				"标签名",
				"创建人",
				"创建时间",
				"修改人",
				"修改时间",
			}
            //根据工作表名及行第一个单元格写入切片内容
			x1.SetSheetRow("标签信息", "A1", &title)
		}
		values := []string{
			strconv.Itoa(v.Id),
			v.Name,
			v.CreatedBy,
			strconv.Itoa(v.CreatedOn),
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}
		row++
		x1.SetSheetRow("标签信息", "A"+strconv.Itoa(row), &values)
	}
	timeNow := strconv.Itoa(int(time.Now().Unix()))
	fileName := "tags-" + timeNow + ".xlsx"
	fullPath := export.GetExcelFullPath() + fileName
	err = x1.SaveAs(fullPath)
	if err != nil {
		return "", err
	}
	return fileName, nil

}
```



### 十五、生成二维码、合并海报

#### 1.实现

首先，你需要在 App 配置项中增加二维码及其海报的存储路径，我们约定配置项名称为 `QrCodeSavePath`，值为 `qrcode/`，经过多节连载的你应该能够完成，若有不懂可参照 go-gin-example。



#### 2.生成二维码

##### 1.安装

```
$ go get -u github.com/boombuler/barcode
```

##### 2.工具包

考虑生成二维码这一动作贴合工具包的定义，且有公用的可能性，新建 pkg/qrcode/qrcode.go 文件，写入内容：

```
package qrcode

import (
	"image/jpeg"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"github.com/EDDYCJY/go-gin-example/pkg/file"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
)

type QrCode struct {
	URL    string
	Width  int
	Height int
	Ext    string
	Level  qr.ErrorCorrectionLevel
	Mode   qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width, height int, level qr.ErrorCorrectionLevel, mode qr.Encoding) *QrCode {
	return &QrCode{
		URL:    url,
		Width:  width,
		Height: height,
		Level:  level,
		Mode:   mode,
		Ext:    EXT_JPG,
	}
}

func GetQrCodePath() string {
	return setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.QrCodeSavePath
}

func GetQrCodeFullUrl(name string) string {
	return setting.AppSetting.PrefixUrl + "/" + GetQrCodePath() + name
}

func GetQrCodeFileName(value string) string {
	return util.EncodeMD5(value)
}

func (q *QrCode) GetQrCodeExt() string {
	return q.Ext
}

func (q *QrCode) CheckEncode(path string) bool {
	src := path + GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	if file.CheckNotExist(src) == true {
		return false
	}

	return true
}

func (q *QrCode) Encode(path string) (string, string, error) {
	name := GetQrCodeFileName(q.URL) + q.GetQrCodeExt()
	src := path + name
	if file.CheckNotExist(src) == true {
		code, err := qr.Encode(q.URL, q.Level, q.Mode)
		if err != nil {
			return "", "", err
		}

		code, err = barcode.Scale(code, q.Width, q.Height)
		if err != nil {
			return "", "", err
		}

		f, err := file.MustOpen(name, path)
		if err != nil {
			return "", "", err
		}
		defer f.Close()

		err = jpeg.Encode(f, code, nil)
		if err != nil {
			return "", "", err
		}
	}

	return name, path, nil
}
```

这里主要聚焦 `func (q *QrCode) Encode` 方法，做了如下事情：

- 获取二维码生成路径
- 创建二维码
- 缩放二维码到指定大小
- 新建存放二维码图片的文件
- 将图像（二维码）以 JPEG 4：2：0 基线格式写入文件

另外在 `jpeg.Encode(f, code, nil)` 中，第三个参数可设置其图像质量，默认值为 75

```
// DefaultQuality is the default quality encoding parameter.
const DefaultQuality = 75

// Options are the encoding parameters.
// Quality ranges from 1 to 100 inclusive, higher is better.
type Options struct {
	Quality int
}
```

##### 3.路由方法

1、第一步

在 routers/api/v1/article.go 新增 GenerateArticlePoster 方法用于接口开发

2、第二步

在 routers/router.go 的 apiv1 中新增 `apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)` 路由

3、第三步

修改 GenerateArticlePoster 方法，编写对应的生成逻辑，如下：

```
const (
	QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

func GenerateArticlePoster(c *gin.Context) {
	appG := app.Gin{c}
	qrc := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto)
	path := qrcode.GetQrCodeFullPath()
	_, _, err := qrc.Encode(path)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
```

##### 4.验证

通过 POST 方法访问 `http://127.0.0.1:8000/api/v1/articles/poster/generate?token=$token`（注意 $token）

![image-20220526110924332](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220526110924332.png?lastModify=1653785198)

通过检查两个点确定功能是否正常，如下：

1、访问结果是否 200

2、本地目录是否成功生成二维码图片

![image-20220526111022073](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220526111022073.png?lastModify=1653785198)

#### 3.合并海报

在这一节，将实现二维码图片与背景图合并成新的一张图，可用于常见的宣传海报等业务场景

##### 1.背景图

![image-20220526140528084](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220526140528084.png?lastModify=1653785198)

将背景图另存为 runtime/qrcode/bg.jpg（实际应用，可存在 OSS 或其他地方）

##### 2.service 方法

打开 service/article_service 目录，新建 article_poster.go 文件，写入内容：

```
func (a *ArticlePosterBg) Generate() (string, string, error) {
	//获取二维码储存路径
	fullPath := qrcode.GetQrCodeFullPath()
	//生成二维码图像
	fileName, path, err := a.Qr.EnCode(fullPath)
	if err != nil {
		return "", "", err
	}

	//检查合并后图像是否存在
	if !a.CheckMergedImage(path) {
		//生成合并后图像文件
		mergedF, err := a.OpenMergedImage(path)
		if err != nil {
			return "", "", err
		}
		defer mergedF.Close()
		//打开背景图片
		bgF, err := file.MustOpen(a.Name, path)
		if err != nil {
			return "", "", err
		}
		defer bgF.Close()
		//打开生成的二维码图片
		qrF, err := file.MustOpen(fileName, path)
		if err != nil {
			return "", "", err
		}
		defer qrF.Close()
		//解码背景图片,由于我的背景图片是png格式的，所以使用png解析
		bgImage, err := png.Decode(bgF)
		if err != nil {
			return "", "", err
		}
		//解码二维码图片
		qrImage, err := jpeg.Decode(qrF)

		if err != nil {
			return "", "", err
		}
		//创建一个新的rgba图像
		jpg := image.NewRGBA(image.Rect(a.Rect.X0, a.Rect.Y0, a.Rect.X1, a.Rect.Y1))
		//在rgba图像上绘制背景图
		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		//在已绘制背景图的rgba上指定Point上绘制二维码图像
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)
		//将绘制好的 RGBA 图像以 JPEG 4：2：0 基线格式写入合并后的图像文件
		jpeg.Encode(mergedF, jpg, nil)
	}
	return fileName, path, nil
}
```

这里重点留意 `func (a *ArticlePosterBg) Generate()` 方法，做了如下事情：

- 获取二维码存储路径
- 生成二维码图像
- 检查合并后图像（指的是存放合并后的海报）是否存在
- 若不存在，则生成待合并的图像 mergedF
- 打开事先存放的背景图 bgF
- 打开生成的二维码图像 qrF
- 解码 bgF 和 qrF 返回 image.Image
- 创建一个新的 RGBA 图像
- 在 RGBA 图像上绘制 背景图（bgF）
- 在已绘制背景图的 RGBA 图像上，在指定 Point 上绘制二维码图像（qrF）
- 将绘制好的 RGBA 图像以 JPEG 4：2：0 基线格式写入合并后的图像文件（mergedF）

##### 3.错误码

新增 错误码，错误提示

##### 4.路由方法

打开 routers/api/v1/article.go 文件，修改 GenerateArticlePoster 方法，编写最终的业务逻辑（含生成二维码及合并海报），如下：

```
const (
    QRCODE_URL = "https://github.com/EDDYCJY/blog#gin%E7%B3%BB%E5%88%97%E7%9B%AE%E5%BD%95"
)

func GenerateArticlePoster(c *gin.Context) {
    appG := app.Gin{c}
    article := &article_service.Article{}
    qr := qrcode.NewQrCode(QRCODE_URL, 300, 300, qr.M, qr.Auto) // 目前写死 gin 系列路径，可自行增加业务逻辑
    posterName := article_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
    articlePoster := article_service.NewArticlePoster(posterName, article, qr)
    articlePosterBgService := article_service.NewArticlePosterBg(
        "bg.jpg",
        articlePoster,
        &article_service.Rect{
            X0: 0,
            Y0: 0,
            X1: 550,
            Y1: 700,
        },
        &article_service.Pt{
            X: 125,
            Y: 298,
        },
    )

    _, filePath, err := articlePosterBgService.Generate()
    if err != nil {
        appG.Response(http.StatusOK, e.ERROR_GEN_ARTICLE_POSTER_FAIL, nil)
        return
    }

    appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
        "poster_url":      qrcode.GetQrCodeFullUrl(posterName),
        "poster_save_url": filePath + posterName,
    })
}
```

这块涉及到大量知识，强烈建议阅读下，如下：

- image.Rect
- image.Pt
- image.NewRGBA
- jpeg.Encode
- jpeg.Decode
- draw.Op
- draw.Draw
- go-imagedraw-package

其所涉及、关联的库都建议研究一下



##### 5.StaticFS

在 routers/router.go 文件，增加如下代码:

```
r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
```

##### 6.验证

![image-20220526163420474](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220526163420474.png?lastModify=1653785198)

访问完整的 URL 路径，返回合成后的海报并扫除二维码成功则正确

![image-20220526163452710](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220526163452710.png?lastModify=1653785198)



#### 4.总结

在本章节实现了两个很常见的业务功能，分别是生成二维码和合并海报。希望你能够仔细阅读我给出的链接，这块的知识量不少，想要用好图像处理的功能，必须理解对应的思路，举一反三



### 十六、在图片上绘制文字

#### 1.实现

这里使用的是 微软雅黑 的字体，请点击进行下载并**存放到 runtime/fonts 目录**下（字体文件占 16 MB 大小）

#### 2.安装

```
$ go get -u github.com/golang/freetype
```

#### 3.绘制文字

打开 service/article_service/article_poster.go 文件，增加绘制文字的业务逻辑，如下：

```
func (a *ArticlePosterBg) DrawPoster(d *DrawText, fontName string) error {
	//获取源字体路径
	fontSource := setting.AppSetting.RuntimeRootPath + setting.AppSetting.FontSavePath + fontName
	//读取源字体
	fontSourceBytes, err := ioutil.ReadFile(fontSource)
	if err != nil {
		return err
	}

	//解析字体库
	trueTypeFont, err := freetype.ParseFont(fontSourceBytes)
	if err != nil {
		return err
	}
	//创建新的context设置一些默认值
	fc := freetype.NewContext()
	//设置屏幕每英寸的分辨率
	fc.SetDPI(72)
	//设置用于绘制文本的字体
	fc.SetFont(trueTypeFont)
	//设置文本字体的大小,以磅为单位
	fc.SetFontSize(d.Size0)
	//设置裁剪矩形以进行绘制
	fc.SetClip(d.JPG.Bounds())
	//设置目标图像
	fc.SetDst(d.JPG)
	//置绘制操作的源图像，通常为 image.Uniform
	fc.SetSrc(image.Black)
	//设置绘制的起始坐标
	pt := freetype.Pt(d.X0, d.Y0)
	//在pt位置画title
	_, err = fc.DrawString(d.Title, pt)
	//重新设置文字大小用来画二级标题
	fc.SetFontSize(d.Size1)
	//画二级标题
	_, err = fc.DrawString(d.SubTitle, freetype.Pt(d.X1, d.Y1))
	if err != nil {
		return err
	}

	err = jpeg.Encode(d.Merged, d.JPG, nil)
	if err != nil {
		return err
	}

	return nil
}
```

这里主要使用了 freetype 包，分别涉及如下细项：

1、freetype.NewContext：创建一个新的 Context，会对其设置一些默认值

```
func NewContext() *Context {
    return &Context{
        r:        raster.NewRasterizer(0, 0),
        fontSize: 12,
        dpi:      72,
        scale:    12 << 6,
    }
}
```

2、fc.SetDPI：设置屏幕每英寸的分辨率

3、fc.SetFont：设置用于绘制文本的字体

4、fc.SetFontSize：以磅为单位设置字体大小

5、fc.SetClip：设置剪裁矩形以进行绘制

6、fc.SetDst：设置目标图像

7、fc.SetSrc：设置绘制操作的源图像，通常为 image.Uniform

```
var (
        // Black is an opaque black uniform image.
        Black = NewUniform(color.Black)
        // White is an opaque white uniform image.
        White = NewUniform(color.White)
        // Transparent is a fully transparent uniform image.
        Transparent = NewUniform(color.Transparent)
        // Opaque is a fully opaque uniform image.
        Opaque = NewUniform(color.Opaque)
)
```

8、fc.DrawString：根据 Pt 的坐标值绘制给定的文本内容



#### 4.业务逻辑

打开 service/article_service/article_poster.go 方法，在 Generate 方法增加绘制文字的代码逻辑，如下：

```
func (a *ArticlePosterBg) Generate() (string, string, error) {
	fullPath := qrcode.GetQrCodeFullPath()
	fileName, path, err := a.Qr.Encode(fullPath)
	if err != nil {
		return "", "", err
	}

	if !a.CheckMergedImage(path) {
		...

		draw.Draw(jpg, jpg.Bounds(), bgImage, bgImage.Bounds().Min, draw.Over)
		draw.Draw(jpg, jpg.Bounds(), qrImage, qrImage.Bounds().Min.Sub(image.Pt(a.Pt.X, a.Pt.Y)), draw.Over)

		err = a.DrawPoster(&DrawText{
			JPG:    jpg,
			Merged: mergedF,

			Title: "Golang Gin 系列文章",
			X0:    80,
			Y0:    160,
			Size0: 42,

			SubTitle: "---煎鱼",
			X1:       320,
			Y1:       220,
			Size1:    36,
		}, "msyhbd.ttc")

		if err != nil {
			return "", "", err
		}
	}

	return fileName, path, nil
}
```



#### 5.验证

访问生成文章海报的接口 `$HOST/api/v1/articles/poster/generate?token=$token`，检查其生成结果，如下图

![image-20220527154857278](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220527154857278.png?lastModify=1653785198)



#### 6.总结

在本章节在 连载十五 的基础上增加了绘制文字，在实现上并不困难，而这两块需求一般会同时出现，大家可以多加练习，了解里面的逻辑和其他 API



### 十七、用Nginx部署Go应用

在本章节，我们将简单介绍 Nginx 以及使用 Nginx 来完成对 go-gin-example 的部署，会实现反向代理和简单负载均衡的功能。

#### 1.Nginx是什么

Nginx 是一个 Web Server，可以用作反向代理、负载均衡、邮件代理、TCP / UDP、HTTP 服务器等等，它拥有很多吸引人的特性，例如：

- 以较低的内存占用率处理 10,000 多个并发连接（每 10k 非活动 HTTP 保持活动连接约 2.5 MB ）
- 静态服务器（处理静态文件）
- 正向、反向代理
- 负载均衡
- 通过 OpenSSL 对 TLS / SSL 与 SNI 和 OCSP 支持
- FastCGI、SCGI、uWSGI 的支持
- WebSockets、HTTP/1.1 的支持
- Nginx + Lua

#### 2.Nginx安装

请右拐谷歌或百度，安装好 Nginx 以备接下来的使用



#### 3.简单讲解

##### 1.常用命令

- nginx：启动 Nginx
- nginx -s stop：立刻停止 Nginx 服务
- nginx -s reload：重新加载配置文件
- nginx -s quit：平滑停止 Nginx 服务
- nginx -t：测试配置文件是否正确
- nginx -v：显示 Nginx 版本信息
- nginx -V：显示 Nginx 版本信息、编译器和配置参数的信息

##### 2.涉及配置

1、 proxy_pass：配置**反向代理的路径**。需要注意的是如果 proxy_pass 的 url 最后为 /，则表示绝对路径。否则（不含变量下）表示相对路径，所有的路径都会被代理过去

2、 upstream：配置**负载均衡**，upstream 默认是以轮询的方式进行负载，另外还支持**四种模式**，分别是：

（1）weight：权重，指定轮询的概率，weight 与访问概率成正比

（2）ip_hash：按照访问 IP 的 hash 结果值分配

（3）fair：按后端服务器响应时间进行分配，响应时间越短优先级别越高

（4）url_hash：按照访问 URL 的 hash 结果值分配



#### 4.部署

在这里需要对 nginx.conf 进行配置，如果你不知道对应的配置文件是哪个，可执行 `nginx -t` 看一下

```
$ nginx -t
nginx: the configuration file /usr/local/etc/nginx/nginx.conf syntax is ok
nginx: configuration file /usr/local/etc/nginx/nginx.conf test is successful
```

显然，我的配置文件在 `/usr/local/etc/nginx/` 目录下，并且测试通过



#### 5.反向代理

反向代理是指以代理服务器来接受网络上的连接请求，然后将请求转发给内部网络上的服务器，并将从服务器上得到的结果返回给请求连接的客户端，此时代理服务器对外就表现为一个反向代理服务器。（来自百科）

![image-20220527164527060](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220527164527060.png?lastModify=1653785198)

##### 1.配置 hosts

由于需要用本机作为演示，因此先把映射配上去，打开 `/etc/hosts`，增加内容：

```
127.0.0.1       api.blog.com
```

##### 2.配置 nginx.conf

打开 nginx 的配置文件 nginx.conf（我的是 /usr/local/etc/nginx/nginx.conf），我们做了如下事情：

增加 server 片段的内容，设置 server_name 为 api.blog.com 并且监听 8081 端口，将所有路径转发到 `http://127.0.0.1:8000/` 下

```
worker_processes  1;

events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       8081;
        server_name  api.blog.com;

        location / {
            proxy_pass http://127.0.0.1:8000/;
        }
    }
}
```

##### 3.验证

**启动 go-gin-example**

项目里创建Makefile文件

```
.PHONY: build clean tool lint help

all: build

build:
    @go build -v .

tool:
    go vet ./...; true
    gofmt -w .

lint:
    golint ./...

clean:
    rm -rf gin-blog
    go clean -i .

help:
    @echo "make: compile packages and dependencies"
    @echo "make tool: run specified go tool"
    @echo "make lint: golint ./..."
    @echo "make clean: remove object files and cached files"
```

项目下，执行 make，再运行 ./gin-blog

```
$ make
github.com/jamesluo111/gin-blog
$ ls
LICENSE        README.md      conf           go-gin-example middleware     pkg            runtime        vendor
Makefile       README_ZH.md   docs           main.go        models         routers        service
$ ./gin-blog
...
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/EDDYCJY/go-gin-example/routers/api/v1.DeleteArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles/poster/generate --> github.com/EDDYCJY/go-gin-example/routers/api/v1.GenerateArticlePoster (4 handlers)
Actual pid is 14672
```

##### 4.重启 nginx

```
$ nginx -t
nginx: the configuration file /usr/local/etc/nginx/nginx.conf syntax is ok
nginx: configuration file /usr/local/etc/nginx/nginx.conf test is successful
$ nginx -s reload
```

##### 5.访问接口

![image-20220527182254585](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/image-20220527182254585.png?lastModify=1653785198)

如此，就实现了一个简单的反向代理了，是不是很简单呢

#### 6.负载均衡

负载均衡，英文名称为 Load Balance（常称 LB），其意思就是分摊到多个操作单元上进行执行（来自百科）

你能从运维口中经常听见，XXX 负载怎么突然那么高。 那么它到底是什么呢？

其背后一般有多台 server，系统会根据配置的策略（例如 Nginx 有提供四种选择）来进行动态调整，尽可能的达到各节点均衡，从而提高系统整体的吞吐量和快速响应

##### 1.如何演示

前提条件为多个后端服务，那么势必需要多个 go-gin-example，为了演示我们可以启动多个端口，达到模拟的效果

为了便于演示，分别在启动前将 conf/app.ini 的应用端口修改为 8001 和 8002（也可以做成传入参数的模式），达到启动 2 个监听 8001 和 8002 的后端服务

##### 2.配置 nginx.conf

 回到 nginx.conf 的老地方，增加负载均衡所需的配置。新增 upstream 节点，设置其对应的 2 个后端服务，最后修改了 proxy_pass 指向（格式为 http:// + upstream 的节点名称） 

```
worker_processes  1;

events {
    worker_connections  1024;
}


http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    upstream api.blog.com {
        server 127.0.0.1:8000;
        server 127.0.0.1:8001;
    }

    server {
        listen       8081;
        server_name  api.blog.com;

        location / {
            proxy_pass http://api.blog.com/;
        }
    }
}
```

##### 3.重启 nginx

```
$ nginx -t
nginx: the configuration file /usr/local/etc/nginx/nginx.conf syntax is ok
nginx: configuration file /usr/local/etc/nginx/nginx.conf test is successful
$ nginx -s reload
```

##### 4.验证

再重复访问 `http://api.blog.com:8081/auth?username={USER_NAME}}&password={PASSWORD}`，多访问几次便于查看效果

目前 Nginx 没有进行特殊配置，那么它是轮询策略，而 go-gin-example 默认开着 debug 模式，看看请求 log 就明白了。

 服务A： 

![1653752865037](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/1653752865037.png?lastModify=1653785198)

 服务B:

![1653752888885](file://D:/%E6%A1%8C%E9%9D%A2/%E7%AC%94%E8%AE%B0/gin/gin%E5%AD%A6%E4%B9%A0%E7%AC%94%E8%AE%B0.assets/1653752888885.png?lastModify=1653785198)



#### 7.总结

 在本章节，希望您能够简单习得日常使用的 Web Server 背后都是一些什么逻辑，Nginx 是什么？反向代理？负载均衡？ 



### 十八、入门makefile

含一定复杂度的软件工程，基本上都是先编译 A，再依赖 B，再编译 C...，最后才执行构建。如果每次都人为编排，又或是每新来一个同事就问你项目 D 怎么构建、重新构建需要注意什么...等等情况，岂不是要崩溃？

我们常常会在开源项目中发现 Makefile，你是否有过疑问？

本章节会简单介绍 Makefile 的使用方式，最后建议深入学习。



#### 1.怎么解决

 对于构建编排，Docker 有 Dockerfile ，在 Unix 中有神器 Make .... 

#### 2.Make

Make 是一个构建自动化工具，会在当前目录下寻找 Makefile 或 makefile 文件。如果存在，会依据 Makefile 的**构建规则**去完成构建

当然了，实际上 Makefile 内都是你根据 make 语法规则，自己编写的特定 Shell 命令等

它是一个工具，规则也很简单。在支持的范围内，编译 A， 依赖 B，再编译 C，完全没问题

#### 3.规则

Makefile 由多条规则组成，每条规则都以一个 target（目标）开头，后跟一个 : 冒号，冒号后是这一个目标的 prerequisites（前置条件）

紧接着新的一行，必须以一个 tab 作为开头，后面跟随 command（命令），也就是你希望这一个 target 所执行的构建命令

```
[target] ... : [prerequisites] ...
<tab>[command]
    ...
    ...
```

- target：一个目标代表一条规则，可以是一个或多个文件名。也可以是某个操作的名字（标签），称为**伪目标（phony）**
- prerequisites：前置条件，这一项是**可选参数**。通常是多个文件名、伪目标。它的作用是 target 是否需要重新构建的标准，如果前置条件不存在或有过更新（文件的最后一次修改时间）则认为 target 需要重新构建
- command：构建这一个 target 的具体命令集

#### 4.简单的例子

 本文将以 go-gin-example 去编写 Makefile 文件，请跨入 make 的大门 

##### 1.分析

 在编写 Makefile 前，需要先分析构建先后顺序、依赖项，需要解决的问题等 

##### 2.编写

```
.PHONY: build clean tool lint help

all: build

build:
    go build -v .

tool:
    go tool vet . |& grep -v vendor; true
    gofmt -w .

lint:
    golint ./...

clean:
    rm -rf go-gin-example
    go clean -i .

help:
    @echo "make: compile packages and dependencies"
    @echo "make tool: run specified go tool"
    @echo "make lint: golint ./..."
    @echo "make clean: remove object files and cached files"
```

1、在上述文件中，使用了 `.PHONY`，其作用是声明 build / clean / tool / lint / help 为**伪目标**，声明为伪目标会怎么样呢？

- 声明为伪目标后：在执行对应的命令时，make 就不会去检查是否存在 build / clean / tool / lint / help 其对应的文件，而是每次都会运行标签对应的命令
- 若不声明：恰好存在对应的文件，则 make 将会认为 xx 文件已存在，没有重新构建的必要了

2、这块比较简单，在命令行执行即可看见效果，实现了以下功能：

1. make: make 就是 make all
2. make build: 编译当前项目的包和依赖项
3. make tool: 运行指定的 Go 工具集
4. make lint: golint 一下
5. make clean: 删除对象文件和缓存文件
6. make help: help

#### 5.为什么会打印执行的命令

 如果你实际操作过，可能会有疑问。明明只是执行命令，为什么会打印到标准输出上了？ 

##### 1.原因

 make 默认会打印每条命令，再执行。这个行为被定义为**回声** 

##### 2.解决

 可以在对应命令前加上 @，可指定该命令不被打印到标准输出上 

```
build:
    @go build -v .
```

 那么还有其他的特殊符号吗？有的，请课后去了解下 +、- 的用途。 