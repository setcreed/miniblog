// Copyright 2025 chengwz <cwzdzg@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/setcreed/miniblog.

package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
	"time"
)

// Logger 定义了 miniblog 项目的日志接口
// 该接口包含了项目中支持的日志记录方法，提供对不同日志级别的支持。
type Logger interface {
	// Debugw 用于记录调试级别的日志，通常用于开发阶段，包含详细的调试信息。
	Debugw(msg string, kvs ...any)

	// Infow 用于记录信息级别的日志，表示系统的正常运行状态。
	Infow(msg string, kvs ...any)

	// Warnw 用于记录警告级别的日志，表示可能存在问题但不影响系统正常运行。
	Warnw(msg string, kvs ...any)

	// Errorw 用于记录错误级别的日志，表示系统运行中出现的错误，需要开发人员介入处理。
	Errorw(msg string, kvs ...any)

	// Panicw 用于记录严重错误级别的日志，表示系统无法继续运行，记录日志后会触发 panic。
	Panicw(msg string, kvs ...any)

	// Fatalw 用于记录致命错误级别的日志，表示系统无法继续运行，记录日志后会直接退出程序。
	Fatalw(msg string, kvs ...any)

	// Sync 用于刷新日志缓冲区，确保日志被完整写入目标存储。
	Sync()
}

// zapLogger 是 Logger 接口的具体实现. 它底层封装了 zap.Logger.
type zapLogger struct {
	z *zap.Logger
}

// 确保 *zapLogger 实现了 Logger 接口. 以下变量赋值，可以使错误在编译期被发现.
var _ Logger = (*zapLogger)(nil)

var (
	mu sync.Mutex

	// std 定义了默认的全局 Logger.
	std = New(NewOptions())
)

// Init 初始化全局的日志对象.
func Init(opts *Options) {
	// 因为会给全局变量 std 赋值，所以这里对 std 变量加锁，防止出现并发问题.
	mu.Lock()
	defer mu.Unlock()

	std = New(opts)
}

// New 根据提供的 Options 参数创建一个自定义的 zapLogger 对象.
// 如果 Options 参数为空，则会使用默认的 Options 配置。
func New(opts *Options) *zapLogger {
	// 如果 opts 为空，则使用默认配置
	if opts == nil {
		opts = NewOptions()
	}

	// 将 Options 中的日志级别（字符串）转换为 zapcore.Level 类型
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		// 如果指定了非法的日志级别，则默认使用 info 级别
		zapLevel = zapcore.InfoLevel
	}

	// 创建 encoder 配置，用于控制日志的输出格式
	encoderConfig := zap.NewProductionEncoderConfig()
	// 自定义 MessageKey 为 message，message 语义更明确
	encoderConfig.MessageKey = "message"
	// 自定义 TimeKey 为 timestamp，timestamp 语义更明确
	encoderConfig.TimeKey = "timestamp"
	// 指定时间序列化函数，将时间序列化为 `2006-01-02 15:04:05.000` 格式，更易读
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	// 指定 time.Duration 序列化函数，将 time.Duration 序列化为经过的毫秒数的浮点数
	// 毫秒数比默认的秒数更精确
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	// 创建构建 zap.Logger 需要的配置
	cfg := &zap.Config{
		// 是否在日志中显示调用日志所在的文件和行号，例如：`"caller":"miniblog/miniblog.go:75"`
		DisableCaller: opts.DisableCaller,
		// 是否禁止在 panic 及以上级别打印堆栈信息
		DisableStacktrace: opts.DisableStacktrace,
		// 指定日志级别
		Level: zap.NewAtomicLevelAt(zapLevel),
		// 指定日志显示格式，可选值：console, json
		Encoding:      opts.Format,
		EncoderConfig: encoderConfig,
		// 指定日志输出位置
		OutputPaths: opts.OutputPaths,
		// 设置 zap 内部错误输出位置
		ErrorOutputPaths: []string{"stderr"},
	}

	// 使用 cfg 创建 *zap.Logger 对象
	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	// 将标准库的 log 输出重定向到 zap.Logger
	zap.RedirectStdLog(z)

	return &zapLogger{z: z}
}

// Sync 调用底层 zap.Logger 的 Sync 方法，将缓存中的日志刷新到磁盘文件中. 主程序需要在退出前调用 Sync.
func Sync() {
	std.Sync()
}

// Debugw 输出 debug 级别的日志.
func Debugw(msg string, kvs ...any) {
	std.Debugw(msg, kvs...)
}

func (l *zapLogger) Debugw(msg string, kvs ...any) {
	l.z.Sugar().Debugw(msg, kvs...)
}

// Infow 输出 info 级别的日志.
func Infow(msg string, kvs ...any) {
	std.Infow(msg, kvs...)
}

func (l *zapLogger) Infow(msg string, kvs ...any) {
	l.z.Sugar().Infow(msg, kvs...)
}

// Warnw 输出 warning 级别的日志.
func Warnw(msg string, kvs ...any) {
	std.Warnw(msg, kvs...)
}

func (l *zapLogger) Warnw(msg string, kvs ...any) {
	l.z.Sugar().Warnw(msg, kvs...)
}

// Errorw 输出 error 级别的日志.
func Errorw(msg string, kvs ...any) {
	std.Errorw(msg, kvs...)
}

func (l *zapLogger) Errorw(msg string, kvs ...any) {
	l.z.Sugar().Errorw(msg, kvs...)
}

// Panicw 输出 panic 级别的日志.
func Panicw(msg string, kvs ...any) {
	std.Panicw(msg, kvs...)
}

func (l *zapLogger) Panicw(msg string, kvs ...any) {
	l.z.Sugar().Panicw(msg, kvs...)
}

// Fatalw 输出 fatal 级别的日志.
func Fatalw(msg string, kvs ...any) {
	std.Fatalw(msg, kvs...)
}

func (l *zapLogger) Fatalw(msg string, kvs ...any) {
	l.z.Sugar().Fatalw(msg, kvs...)
}

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}
