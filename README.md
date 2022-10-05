# HduHelpLogin
杭电助手后端二面小任务

## 设计

### token认证

考虑到本次小任务的体量与形式，通过token来认证应该是最为简洁好用的方式了

在service层简单地定义了一个map

```golang
var token = make(map[uuid.UUID]uint)
```

这样子，在内存中存储token

具体认证时，利用一个名为`authentic`的中间件，专门负责完成token认证，失败时直接abort，成功则保存认证后得到的UserId到context中以便后续使用

### traceId追踪

为了便于后续定位问题和调试（虽然就这个项目本身而言，这个需求非常有限），设计了traceId对请求的处理进行追踪。

每个请求都会分配到自己的traceId，在请求真正开始处理前，利用中间件`AddTraceId`生成一个traceId并保存在context中，此后每一步函数调用都会携带context作为其第一个参数。从而做到让所有日志信息都携带traceId

与此同时，traceId也作为请求的response的一部分返回提供给用户，便于用户提交traceId给开发者定位问题。

### 错误处理

本项目大部分的错误处理都来自数据库操作，数据库操作又是项目中安全风险最大的。

因此我选择把所有的数据库操作产生的错误压在数据库层，直接打印log，对外返回一个“database error”这样一个模糊的描述，避免用户接触这些信息。

当需要定位错误时，可以使用traceId来快速寻找log，不影响调试。

### 日志

因为要实现traceId，gin和gorm默认的日志处理都无法直接使用，需要重新写过。对这两边的日志，均采取了先抄默认再修改的做法，从而加入traceId。

日志库选用了logrus，比较经典。对于gin和gorm的日志的处理都是一致的，用`logrus.WithFields`替代了默认的`printf`的格式化输出，虽然好像改丑了但是格式上确实是更加统一了，也便于后续对日志的分析处理（如果有的话）

日志输出用了两种格式，在终端打印给人看的彩色`TextFormatter`，在文件中输出json格式便于后续处理。

关于日志的记录，借用了一些别人的库实现了日志文件随时间的切换

## 接口

### POST `/login`

#### 简介

登录账号

#### 请求

application/x-www-form-urlencoded

| key      | type   |
|----------|--------|
| username | string |
| password | string |

#### 响应

application/json

| key     | type   | description                                          |
|---------|--------|------------------------------------------------------|
| traceId | uuid   |                                                      |
| token   | uuid   | 登录成功时（200）返回，在后续请求中需要放置在Authorization里并加上"Bearer "前缀 |
| msg     | string | 登陆成功时（200）返回登录成功，失败时描述失败原因                           |


### POST `/register`

#### 简介

注册账号

#### 请求

application/x-www-form-urlencoded

| key      | type   |
|----------|--------|
| username | string |
| password | string |

#### 响应

application/json

| key     | type   | description                |
|---------|--------|----------------------------|
| traceId | uuid   |                            |
| msg     | string | 注册成功时（200）返回注册成功，失败时描述失败原因 |

### DELETE `/logout/:token`

登出账号

#### 请求

请求路径中传入token

#### 响应

application/json

| key     | type   | description                |
|---------|--------|----------------------------|
| traceId | uuid   |                            |
| msg     | string | 登出成功时（200）返回登出成功，失败时描述失败原因 |

注意：token符合uuid格式时，应该是总是成功

### GET `/api/user`

#### 简介

获取用户信息

#### 请求

需要Authorization

#### 响应

application/json

| key     | type   | description          |
|---------|--------|----------------------|
| traceId | uuid   |                      |
| msg     | string | 成功时（200）返回ok，失败时描述原因 |
| age     | uint   | age<200              |
| phone   | int64  | 11位，1开头的手机号          |
| email   | string | 正则校验格式               |

注意：用户信息在数据库中不存在时，这里返回零值

### PUT `/api/user`

#### 简介

更新用户信息

#### 请求

需要Authorization

application/x-www-form-urlencoded

| key     | type   | description          |
|---------|--------|----------------------|
| age     | uint   | age<200              |
| phone   | int64  | 11位，1开头的手机号          |
| email   | string | 正则校验格式               |

#### 响应

application/json

| key     | type   | description                   |
|---------|--------|-------------------------------|
| traceId | uuid   |                               |
| msg     | string | 成功时（200）返回put success，失败时描述原因 |

### PUT `/api/user/password`
todo
