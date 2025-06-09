// Copyright 2025 chengwz <cwzdzg@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/setcreed/miniblog.

package log_test

import (
	"testing"
	"time"

	"github.com/setcreed/miniblog/internal/pkg/log"
)

func TestLogger(t *testing.T) {
	// 自定义日志配置
	opts := &log.Options{
		Level:             "debug",            // 设置日志级别为 debug
		Format:            "json",             // 设置日志格式为 JSON
		DisableCaller:     false,              // 显示调用日志的文件和行号
		DisableStacktrace: false,              // 允许打印堆栈信息
		OutputPaths:       []string{"stdout"}, // 将日志输出到标准输出
	}

	// 初始化全局日志对象
	log.Init(opts)

	// 测试不同级别的日志输出
	log.Debugw("This is a debug message", "key1", "value1", "key2", 123)
	log.Infow("This is an info message", "key", "value")
	log.Warnw("This is a warning message", "timestamp", time.Now())
	log.Errorw("This is an error message", "error", "something went wrong")

	// 注意：Panicw 和 Fatalw 会中断程序运行，因此在测试中应小心使用。
	// 可以注释掉以下两行进行测试，或者在单独的环境中运行。
	// log.Panicw("This is a panic message", "reason", "unexpected situation")
	// log.Fatalw("This is a fatal message", "reason", "critical failure")

	// 确保日志缓冲区被刷新
	log.Sync()
}
