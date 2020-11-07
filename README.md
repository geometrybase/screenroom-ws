# 使用
```
screenroom-ws -mode release -port 8000
```

```
ws://localhost:8000/s/:topicID
ws://localhost:8000/c/:topicID
```

同一个topic ID的消息会相互广播


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
        }
    }
}
```

```javascript 
{
    action: "favorite",
    data: {
        user: {
        },
        topic: {
            id: "string"
        }
    }
}
```