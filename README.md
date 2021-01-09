# 使用
```
screenroom-ws -mode release -port 8000
```

```
ws://localhost:8000/s/:deviceID
ws://localhost:8000/c/:deviceID
```

```bash
wss://play.screenroom.cn/s/:deviceID  #屏幕URL
wss://play.screenroom.cn/c/:deviceID/:shopNo  #手机URL
```

同一个device ID的消息会相互广播


# 数据格式

```javascript 
{
    action: "select",
    data: {
        user: {
            shopno: "string"
        },
        topic: {
            id: "string"
        }
    }
}
```

```javascript 
{
    action: "like",
    data: {
        user: {
          shopno: "string"
        },
        topic: {
            id: "string"
        }
    }
}
```

```javascript 
{
    action: "unlike",
    data: {
        user: {
            shopno: "string"
        },
        topic: {
            id: "string"
        }
    }
}
```

```javascript 
{
    action: "comment",
    data: {
        user: {
            shopno: "string"
        },
        topic: {
            id: "string"
               },
        comment: {
            text: "string",
            pos_x: "int",
            pos_y: "int"
        }
    }
}
```

```javascript 
{
    action: "collect",
    data: {
        user: {
            shopno: "string"
        },
        topic: {
            id: "string"
        }
    }
}
```
