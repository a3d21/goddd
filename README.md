# GoDDD

Go DDD实践Demo项目。

假设场景：一个账户服务，可以开/关户、存取钱。

## 设计思想
将测试、领域层、设施层解耦，使测试、领域服务、设施实现可组合，使测试可复用。

## 代码结构
```
├── application   app层。聚合领域服务，并对外提供接口
│   └── app.go
├── base    基础包。
│   └── baseerr  通用Err
├── di  DI工具
│   └── di.go
├── domain  领域层
│   └── account  账户服务
│       ├── entity.go
│       ├── repo.go
│       └── service.go
├── infra  设施层
│   ├── meminfra  mem设施层。使用内存实现，用于快速Domain建模
│   │   └── repo.go
│   └── prodinfra 生产设施层。使用DB实现
│       ├── db.go
│       ├── po.go
│       ├── repo.go
│       └── util.go
├── main.go  程序入口
└── tests  测试
    ├── domain_test.go  领域测试。一个领域测试对应一个业务用例，与infra实现无法，可以在任意infra运行
    ├── infra_test.go  设施层测试。保证infra一致性
    └── sql  prodinfra初始化sql
        └── 2021-12-15-create-account.sql
```

## 说明
使用环境变更`APP_ENV`区分环境（infra），可选值

|APP_ENV|说明|
|-|-|
|mem|内存环境（默认）。使用map实现infra，用于快速测试domain层|
|prod|生产环境。连接真实db|
|test|测试环境。使用sqlite fake db，用于开发过程中快速调试prodinfra|

## How To Run

```
### 运行程序
# 以mem infra运行
$ APP_ENV=mem go run main.go
# 以test infra运行
$ APP_ENV=test go run main.go
# 以prod infra运行
$ APP_ENV=prod APP_MYSQL_DSN="root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local" go run main.go

### 测试接口
$ curl http://127.0.0.1:8080/health
$ curl http://127.0.0.1:8080/v1/open?name=xxx
$ curl http://127.0.0.1:8080/v1/balance?name=xxx
```

## How To Test

测试与infra解耦，可复用。可以任何环境运行测试。
```
$ APP_ENV=mem go test -count=1 ./...
$ APP_ENV=test go test -count=1 ./...
$ APP_ENV=prod APP_MYSQL_DSN='root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local' go test -count=1 ./...
```