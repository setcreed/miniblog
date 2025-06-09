## 日志规范

更多关于日志的规范请参考：[日志规范](https://konglingfei.com/onex/convention/log.html)。以下是 miniblog 的差异化日志规范。

### 日志规范

- 日志包统一使用 `github.com/setcreed/miniblog/internal/pkg/log`;
- 使用结构化的日志打印格式：`log.Infow`, `log.Warnw`, `log.Errorw` 等; 例如：`log.Infow("Update post function called")`;
- 日志均以大写开头，结尾不跟 `.`，例如：`log.Infow("Update post function called")`;
- 使用过去时，例如：`Could not delete B` 而不是 `Cannot delete B`;
- 遵循日志级别规范：
    - Debug 级别的日志使用 `log.Debugw`;
    - Info 级别的日志使用 `log.Infow`;
    - Warning 级别的日志使用 `log.Warnw`;
    - Error 级别的日志使用 `log.Errorw`;
    - Panic 级别的日志使用 `log.Panicw`;
    - Fatal 级别的日志使用 `log.Fatalw`.
- 日志设置：
    - 开发测试环境：日志级别设置为 `debug`、日志格式可根据需要设置为 `console` / `json`、开启 caller；
    - 生产环境：日志级别设置为 `info`、日志格式设置为 `json`、开启 caller。（注意：上线初期，为了方便现网排障，日志级别可以设置为 `debug`）
- 在记录日志时，不要输出一些敏感信息，例如密码、密钥等。
- 如果在具有 `context.Context` 参数的函数/方法中，调用日志函数，建议使用 `log.W(ctx).Infow()` 进行日志记录。
- 【建议】Error 日志应该在最原始的报错位置打印，一是避免上层代码缺失部分入参，二是避免漏打。
- 【建议】在可供外部项目使用的第三包中，原则上不打印日志。如果需要尽量使用 Logger 接口；
- 【建议】 当调用第三方包的函数，且第三方包函数出错时，在错误处打印日志，例如：
```go
if err := os.Chdir("/root"); err != nil {
    log.Errorf("change dir failed: %v", err)
}
```

**提示**
在日志消息中，关于开头字母大小写的惯例因开发团队而异。一些团队更喜欢使用大写字母开头，这有助于强调重要性，以及使日志更易读，尤其是在较长的日志行中。而其他团队更倾向于使用小写字母，因为这样的日志看起来更加紧凑和一致；

** 在错误产生的最原始位置打印日？**

对于嵌套的 Error，可在 Error 产生的最初位置打印 Error 日志，上层如果不需要添加必要的信息，可以直接返回下层的 Error。例如：

```go
package main

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	if err := loadConfig(); err != nil {
		glog.Error(err)
	}
}

// 正例：直接返回错误
func loadConfig() error {
	return decodeConfig() // 直接返回
}

// 正例：如果需要基于函数返回的错误，封装更多的信息，可以封装返回的 err。否则，建议直接返回 err
func decodeConfig() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("could not decode configuration data for user %s: %v", "colin", err) // 添加必要的信息，用户名称
	}
	return nil
}

func readConfig() error {
	glog.Errorf("read: end of input.")
	return fmt.Errorf("read: end of input")
}
```

通过在最初产生错误的位置打印日志，可以很方便地追踪到错误产生的根源，并且错误日志只打印一次，可以减少重复的日志打印，减少排障时重复日志干扰，也可以提高代码的简洁度。

当然，在开发中也可以对错误补充一些有用的信息，以记录错误产生的其他影响。