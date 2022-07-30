# Gin Starter
基于`gin`和`gorm`构建的一个RESTful实例。实现JWT认证，并提供生成CURD代码的`moc`命令。

## 如何运行
创建一个`starter`的数据库，编辑 config/config.yaml 配置数据库:
```
addr: :9000
gin_mode: "debug"
#gin_mode: "release"

dsn: "root:@tcp(127.0.0.1:3306)/starter?charset=utf8mb4&parseTime=True&loc=Local"
max_idle_conn: 100
sql_debug: true

jwt_secret: "your secret"
jwt_expire_day: 100

redis_host: "127.0.0.1:6379"
redis_password:
redis_max_idle: 30
redis_active: 30
redis_idle_timeout: 200
```

运行：
```
go run ./cmd/app/main.go

[GIN-debug] POST   /api/v1/login             --> moell/internal/app/api.(*Auth).Login-fm (3 handlers)
[GIN-debug] POST   /api/v1/register          --> moell/internal/app/api.(*Auth).Register-fm (3 handlers)
[GIN-debug] GET    /api/v1/check-app-upgrade/:platform --> moell/internal/app/api.(*AppUpgrade).Check-fm (3 handlers)
[GIN-debug] PATCH  /api/v1/change-password   --> moell/internal/app/api.(*Auth).ChangePassword-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/frozen-account    --> moell/internal/app/api.(*Auth).FrozenAccount-fm (5 handlers)
[GIN-debug] GET    /api/v1/user              --> moell/internal/app/api.(*User).Index-fm (5 handlers)
[GIN-debug] POST   /api/v1/user              --> moell/internal/app/api.(*User).Store-fm (5 handlers)
[GIN-debug] PATCH  /api/v1/user/:id          --> moell/internal/app/api.(*User).Update-fm (5 handlers)
[GIN-debug] DELETE /api/v1/user/:id          --> moell/internal/app/api.(*User).Destroy-fm (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :9000

```
## moc 命令行的使用
```shell
go run ./cmd/moc/main.go -h
NAME:
   moc - moell cli application

USAGE:
   moc [global options] command [command options] [arguments...]

COMMANDS:
   curd     生成CURD代码
   api      生成API代码
   model    生成Model代码
   repo     生成Repository代码
   service  生成Service代码
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```
实例：
```shell
go run ./cmd/moc/main.go curd --name Article
"./internal/app/requests/admin_user.go" Generated successfully
"./internal/app/api/admin_user.go" Generated successfully
"./internal/app/repositories/admin_user.go" Generated successfully
"./internal/app/models/admin_user.go" Generated successfully
"./internal/app/services/admin_user.go" Generated successfully

go run ./cmd/moc/main.go curd --name AdminUser --dir=admin
"./internal/admin/requests/admin_user.go" Generated successfully
"./internal/admin/api/admin_user.go" Generated successfully
"./internal/admin/repositories/admin_user.go" Generated successfully
"./internal/admin/models/admin_user.go" Generated successfully
"./internal/admin/services/admin_user.go" Generated successfully
```
更多使用方法，请参照帮助命令

## 项目结构
```
├── cmd 项目主干
│   ├── app 默认项目目录
│   └── moc 代码生成器入口
├── config 配置文件
├── internal 项目代码目录
│   ├── app 项目
│   │   ├── api 路由入口
│   │   ├── middleware 中间件
│   │   ├── models 模型
│   │   ├── repositories Repository
│   │   ├── requests 请求验证
│   │   ├── routes 路由
│   │   └── services Service
│   └── moc 
├── logs 日志存放位置
└── pkg 公共代码库
    ├── auth
    ├── database
    ├── logger
    ├── model
    ├── repository
    ├── response
    └── utils
```
## 特性
* gin
* gorm
* sirupsen/logrus
* thedevsaddam/govalidator
* urfave/cli/
* sirupsen/logrus
* gomodule/redigo
* golang-jwt/jwt

## 打赏

<p>
  <img src="http://ww1.sinaimg.cn/mw690/7a679ca1ly1fvxrfnvxa4j20dw0dwdic.jpg" width="250" />
  <img src="http://ww1.sinaimg.cn/mw690/7a679ca1ly1fvxrfnr0dhj20dw0dwgp0.jpg" width="250" />
</p>

## License

Apache License Version 2.0 see http://www.apache.org/licenses/LICENSE-2.0.html