# 使用
```
screenroom-ws -mode release -port 8000
```

```
ws://localhost:8000/s/:deviceID
ws://localhost:8000/c/:deviceID
```

```bash
wss://a.screenroom.cn/s/:deviceID  #屏幕URL
wss://a.screenroom.cn/c/:deviceID  #手机URL
```

同一个device ID的消息会相互广播


# 数据格式

```javascript 
{
    action: "select",
    data: {
        user: {
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
        },
        topic: {
            id: "string"
               },
        content: {
            text: "string",
            loc_x: "int",
            loc_y: "int"
        }
    }
}
```

```javascript 
{
    action: "collect",
    data: {
        user: {
        },
        topic: {
            id: "string"
        }
    }
}
```