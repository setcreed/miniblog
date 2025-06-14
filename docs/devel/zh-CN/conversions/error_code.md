## 错误码规范

错误码是用户定位及解决问题的重要手段之一，当应用发生异常时，用户根据错误码及文档中的错误码描述和解决方案就可以快速定位并解决问题。更多错误信息返回规范请参考：[错误信息返回规范](https://konglingfei.com/onex/convention/error.html)

### 错误码命名规范

- 遵循驼峰命名法；
- 错误码分为两级。例如，`NotFound.UserNotFound`，以 `.` 号分隔。其中，第一级错误码为平台级，第二级错误码为资源级别，可根据场景自定义；
- 第二级错误码只能使用英文字母或数字（`[a-zA-Z0-9]`），应使用英文单词规范拼写、规范缩写、RFC 术语缩写等；
- 错误码应避免同一语义多种定义，例如：`InvalidParameter.ErrorBind`, `InvalidParameter.BindError`。

### 第一级公共错误码

| 错误码 | 错误描述 | 错误类型 |
| ---- | ---- | ---- |
| OK | 请求成功 | - |
| InternalError | 内部错误 | 1 |
| NotFound | 资源不存在 | 0 |
| BindError | 绑定失败，解析请求体失败 | 0 |
| InvalidArgument | 参数错误（包括参数类型、格式、值等错误） | 0 |
| Unauthenticated | 认证失败 | 0 |
| PermissionDenied | 授权失败 | 0 |
| OperationFailed | 操作失败 | 2 |

> 错误类型：0 代表客户端，1 代表服务端，2 代表客户端 / 服务端，- 代表请求成功。