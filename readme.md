## Gin Web Test

## Run

```bash
mv conf/app.ini.example conf/app.ini
go run main.go
```

### Hot Reload(Optional)
Use Fresh docs: [Fresh](https://github.com/gravityblast/fresh)

Recommended only for development environments.

#### Installation

```go get github.com/pilu/fresh```

#### Usage

Start fresh:

```fresh -c fresh.conf```

#### Logs
```bash
$ fresh -c fresh.conf
Loading settings from fresh.conf
16:11:57 runner      | InitFolders
16:11:57 runner      | mkdir ./tmp
16:11:57 runner      | mkdir ./tmp: file exists
16:11:57 watcher     | Watching .
16:11:57 watcher     | Watching app
......
16:11:57 watcher     | Ignoring vendor
16:11:57 main        | Waiting (loop 1)...
16:11:57 main        | receiving first event /
16:11:57 main        | sleeping for 600 milliseconds
16:11:58 main        | flushing events
16:11:58 main        | Started! (53 Goroutines)
16:11:58 main        | remove tmp/runner-build-errors.log: no such file or directory
16:11:58 build       | Building...
16:11:58 runner      | Running...
16:11:58 main        | --------------------
16:11:58 main        | Waiting (loop 2)...
16:11:58 app         | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

16:11:58 app         | [GIN-debug] GET    /v1/api/test/ping         --> gin-test/app/controllers/index.Ping (4 handlers)
[GIN-debug] GET    /v1/api/test/font         --> gin-test/app/controllers/index.Test (4 handlers)
[GIN-debug] GET    /v1/api/test/test_users   --> gin-test/app/controllers/index.GetTestUsers (4 handlers)
[GIN-debug] GET    /swagger                  --> gin-test/routers.InitRouter.func1 (3 handlers)
16:11:58 app         | [GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
16:11:58 app         | 2019/08/30 16:11:58 [info] start http server listening :8080
16:11:58 app         | 2019/08/30 16:11:58 [info] Actual pid is 2451

```

## Swagger Docs

### Preview

![swagger_preview](./img/swagger_preview.png)

Access ```BASE_URL/swagger/index.html``` view docs.

Please check the instructions for use.
[gin-swagger](https://github.com/swaggo/gin-swagger)

### Generate
```bash
$ swag init
2019/08/22 16:17:11 Generate swagger docs....
2019/08/22 16:17:11 Generate general API Info, search dir:./
2019/08/22 16:17:11 create docs.go at  docs/docs.go
2019/08/22 16:17:11 create swagger.json at  docs/swagger.json
2019/08/22 16:17:11 create swagger.yaml at  docs/swagger.yaml
```

## Parameter Verification

### 1.Defining structure

use `validator.v9` Docs: [validator.v9](https://godoc.org/gopkg.in/go-playground/validator.v9)

```golang
type Page struct {
    P uint `json:"p" form:"p" validate:"required,numeric,min=1"`
    N uint `json:"n" form:"n" validate:"required,numeric,min=1"`
}
```

### 2.Binding Request Parameters

```golang
    var p Page
    if err := c.ShouldBindQuery(&p); err != nil {
        return err, "参数绑定失败,请检查传递参数类型！", 0, 0
    }
```

### 3.Verify Binding Parameters

```golang
    err, parameterErrorStr := common.CheckBindStructParameter(p, c)
```

### Complete example

```golang

type Page struct {
    P uint `json:"p" form:"p" validate:"required,numeric,min=1"`
    N uint `json:"n" form:"n" validate:"required,numeric,min=1"`
}

// GetPage get page parameters
func GetPage(c *gin.Context) (error, string, int, int) {
    currentPage := 0

    // 绑定 query 参数到结构体
    var p Page
    if err := c.ShouldBindQuery(&p); err != nil {
        return err, "参数绑定失败,请检查传递参数类型！", 0, 0
    }

    // 验证绑定结构体参数
    err, parameterErrorStr := common.CheckBindStructParameter(p, c)
    if err != nil {
        return err, parameterErrorStr, 0, 0
    }

    page := com.StrTo(c.DefaultQuery("p", "0")).MustInt()
    limit := com.StrTo(c.DefaultQuery("n", "15")).MustInt()
    
    if page > 0 {
       currentPage = (page - 1) * limit
    }

    return nil, "", currentPage, limit
}
```

## Features

- [Gorm](https://github.com/jinzhu/gorm)
- [Swagger(swag)](https://github.com/swaggo/swag)
- [Gin-gonic](https://github.com/gin-gonic/gin)
- [Go-ini](https://github.com/go-ini/ini)
- [Redis](https://github.com/gomodule/redigo)
- [Fresh](https://github.com/gravityblast/fresh)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [Casbin](https://github.com/casbin/casbin)
- [Gorm-adapter](https://github.com/casbin/gorm-adapter)

## License

MIT