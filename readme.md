## Gin Web Test

```bash
go run main.go
```

```
[GET]
http://localhost:8080/font?base64=W3sidGV4dCI6InN0cmluZzExMSIsIngiOjQxMCwieSI6OTAsImZvbnRTaXplIjoyMCwiY29sb3IiOjJ9LHsidGV4dCI6IuaWh+WtlzIiLCJ4Ijo0MTAsInkiOjE5MCwiZm9udFNpemUiOjIwLCJjb2xvciI6Mn0seyJ0ZXh0Ijoi5paH5a2X5LiJIiwieCI6NDEwLCJ5IjoyOTAsImZvbnRTaXplIjoyMCwiY29sb3IiOjF9LHsidGV4dCI6IuaWh+Wtl+WbmyIsIngiOjQxMCwieSI6MzkwLCJmb250U2l6ZSI6MjAsImNvbG9yIjoxfV0=
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
[
    {
        "text":"string111",
        "x":410,
        "y":90,
        "fontSize":20,
        "color":2
    },
    {
        "text":"文字2",
        "x":410,
        "y":190,
        "fontSize":20,
        "color":2
    },
    {
        "text":"文字三",
        "x":410,
        "y":290,
        "fontSize":20,
        "color":1
    },
    {
        "text":"文字四",
        "x":410,
        "y":390,
        "fontSize":20,
        "color":1
    }
]
```

