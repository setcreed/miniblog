// Copyright 2025 chengwz <cwzdzg@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/setcreed/miniblog.

package app

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/setcreed/miniblog/cmd/mb-apiserver/app/options"
	"github.com/setcreed/miniblog/internal/pkg/log"
	"github.com/setcreed/miniblog/pkg/version"
)

var configFile string // 配置文件路径

// NewMiniBlogCommand 创建一个 *cobra.Command 对象，用于启动应用程序.
func NewMiniBlogCommand() *cobra.Command {
	// 创建默认的应用命令行选项
	opts := options.NewServerOptions()

	cmd := &cobra.Command{
		// 指定命令的名字，该名字会出现在帮助信息中
		Use: "mb-apiserver",
		// 命令的简短描述
		Short: "A mini blog show best practices for develop a full-featured Go project",
		// 命令的详细描述
		Long: `A mini blog show best practices for develop a full-featured Go project.

The project features include:
• Utilization of a clean architecture;
• Use of many commonly used Go packages: gorm, casbin, govalidator, jwt, gin, 
  cobra, viper, pflag, zap, pprof, grpc, protobuf, grpc-gateway, etc.;
• A standardized directory structure following the project-layout convention;
• Authentication (JWT) and authorization features (casbin);
• Independently designed log and error packages;
• Management of the project using a high-quality Makefile;
• Static code analysis;
• Includes unit tests, performance tests, fuzz tests, and mock tests;
• Rich web functionalities (tracing, graceful shutdown, middleware, CORS, 
  recovery from panics, etc.);
• Implementation of HTTP, HTTPS, and gRPC servers;
• Implementation of JSON and Protobuf data exchange formats;
• The project adheres to numerous development standards: 
  code standards, versioning standards, API standards, logging standards, 
  error handling standards, commit standards, etc.;
• Access to MySQL with programming implementation;
• Implemented business functionalities: user management and blog management;
• RESTful API design standards;
• OpenAPI 3.0/Swagger 2.0 API documentation;
• High-quality code.`,
		// 命令出错时，不打印帮助信息。设置为 true 可以确保命令出错时一眼就能看到错误信息
		SilenceUsage: true,
		// 指定调用 cmd.Execute() 时，执行的 Run 函数
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(opts)
		},
		// 设置命令运行时的参数检查，不需要指定命令行参数。例如：./miniblog param1 param2
		Args: cobra.NoArgs,
	}

	// 初始化配置函数，在每个命令运行时调用
	cobra.OnInitialize(onInitialize)

	// cobra 支持持久性标志(PersistentFlag)，该标志可用于它所分配的命令以及该命令下的每个子命令
	// 推荐使用配置文件来配置应用，便于管理配置项
	cmd.PersistentFlags().StringVarP(&configFile, "config", "c", filePath(), "Path to the miniblog configuration file.")

	// 将 ServerOptions 中的选项绑定到命令标志
	opts.AddFlags(cmd.PersistentFlags())

	// 添加 --version 标志
	version.AddFlags(cmd.PersistentFlags())

	return cmd
}

// run 是主运行逻辑，负责初始化日志、解析配置、校验选项并启动服务器。
func run(opts *options.ServerOptions) error {
	// 如果传入 --version，则打印版本信息并退出
	version.PrintAndExitIfRequested()

	// 初始化日志
	log.Init(logOptions())
	defer log.Sync() // 确保日志在退出时被刷新到磁盘

	// 将 viper 中的配置解析到 opts.
	if err := viper.Unmarshal(opts); err != nil {
		return err
	}

	// 校验命令行选项
	if err := opts.Validate(); err != nil {
		return err
	}

	// 获取应用配置.
	// 将命令行选项和应用配置分开，可以更加灵活的处理 2 种不同类型的配置.
	cfg, err := opts.Config()
	if err != nil {
		return err
	}

	// 创建服务器实例.
	// 注意这里是联合服务器，因为可能同时启动多个不同类型的服务器.
	server, err := cfg.NewUnionServer()
	if err != nil {
		return err
	}

	// 启动服务器
	return server.Run()
}

// logOptions 从 viper 中读取日志配置，构建 *log.Options 并返回.
// 注意：viper.Get<Type>() 中 key 的名字需要使用 . 分割，以跟 YAML 中保持相同的缩进.
func logOptions() *log.Options {
	opts := log.NewOptions()
	if viper.IsSet("log.disable-caller") {
		opts.DisableCaller = viper.GetBool("log.disable-caller")
	}
	if viper.IsSet("log.disable-stacktrace") {
		opts.DisableStacktrace = viper.GetBool("log.disable-stacktrace")
	}
	if viper.IsSet("log.level") {
		opts.Level = viper.GetString("log.level")
	}
	if viper.IsSet("log.format") {
		opts.Format = viper.GetString("log.format")
	}
	if viper.IsSet("log.output-paths") {
		opts.OutputPaths = viper.GetStringSlice("log.output-paths")
	}
	return opts
}
