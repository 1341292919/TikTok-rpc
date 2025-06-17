## websocket传入json请求的格式
### 私聊
```json
{
  "type": "to_user_message",
  "data": {
    "content": "Hello, this is a test message",
    "target_id": 10003
  }
}
```
### 群聊
```json
{
  "type": "to_group_message",
  "data": {
    "content": "Hello, this is a test message",
    "target_id": 1
  }
}
```
### 获取私聊信息
```json
{
  "type": "get_private_history",
  "data": {
    "target_id": 10002
  },
  "param": {
    "page_size": 20,
    "page_num": 1
  }
}
```
### 获取群聊信息
```json
{
  "type": "get_group_history",
  "data": {
    "target_id": 1
  },
  "param": {
    "page_size": 20,
    "page_num": 1
  }
}
```