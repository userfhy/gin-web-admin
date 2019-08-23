## Gin Web Test

### Run

```bash
mv conf/app.ini.example conf/app.ini
go run main.go
```

### Swagger Docs

Access ```BASE_URL/swagger/index.html``` view docs.

Please check the instructions for use.
[gin-swagger](https://github.com/swaggo/gin-swagger)

#### Generate
```bash
$ swag init
2019/08/22 16:17:11 Generate swagger docs....
2019/08/22 16:17:11 Generate general API Info, search dir:./
2019/08/22 16:17:11 create docs.go at  docs/docs.go
2019/08/22 16:17:11 create swagger.json at  docs/swagger.json
2019/08/22 16:17:11 create swagger.yaml at  docs/swagger.yaml
```

## Test Route
```
[GET]
http://localhost:8080//v1/api/test/font?base64=W3sidGV4dCI6InN0cmluZzExMSIsIngiOjQxMCwieSI6OTAsImZvbnRTaXplIjoyMCwiY29sb3IiOjJ9LHsidGV4dCI6IuaWh+WtlzIiLCJ4Ijo0MTAsInkiOjE5MCwiZm9udFNpemUiOjIwLCJjb2xvciI6Mn0seyJ0ZXh0Ijoi5paH5a2X5LiJIiwieCI6NDEwLCJ5IjoyOTAsImZvbnRTaXplIjoyMCwiY29sb3IiOjF9LHsidGV4dCI6IuaWh+Wtl+WbmyIsIngiOjQxMCwieSI6MzkwLCJmb250U2l6ZSI6MjAsImNvbG9yIjoxfV0=
```

| *Parameters* | *Required* | *Description*               |
| ------------ | ---------- | --------------------------- |
| base64       | True       | Json array to base64 string |


Json array:

| ***Parameters*** | ***Required*** | ***Description*** |
| ---------------- | -------------- | ----------------- |
| text             | True           |                   |
| x                | True           |                   |
| y                | True           |                   |
| fontSize         | True           |                   |
| color            | True           | 1-White 2-Black   |

```json
{
    "code": 200,
    "msg": "文字解析成功",
    "data": [{
        "text": "string111",
        "x": 410,
        "y": 90,
        "fontSize": 20,
        "color": 2
    }, {
        "text": "文字2",
        "x": 410,
        "y": 190,
        "fontSize": 20,
        "color": 2
    }, {
        "text": "文字三",
        "x": 410,
        "y": 290,
        "fontSize": 20,
        "color": 1
    }, {
        "text": "文字四",
        "x": 410,
        "y": 390,
        "fontSize": 20,
        "color": 1
    }]
}
```

