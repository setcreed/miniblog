// Copyright 2025 chengwz <cwzdzg@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/setcreed/miniblog.

package apiserver

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/viper"
)

// Config 配置结构体，用于存储应用相关的配置.
// 不用 viper.Get，是因为这种方式能更加清晰的知道应用提供了哪些配置项.
type Config struct {
	ServerMode string
	JWTKey     string
	Expiration time.Duration
}

// UnionServer 定义一个联合服务器. 根据 ServerMode 决定要启动的服务器类型.
type UnionServer struct {
	cfg *Config
}

// NewUnionServer 根据配置创建联合服务器.
func (cfg *Config) NewUnionServer() (*UnionServer, error) {
	return &UnionServer{cfg: cfg}, nil
}

// Run 运行应用.
func (s *UnionServer) Run() error {
	fmt.Printf("JWTKey from ServerOptions: %s\n", s.cfg.JWTKey)
	fmt.Printf("JWTKey from Viper: %s\n\n", viper.GetString("jwt-key"))

	jsonData, _ := json.MarshalIndent(s.cfg, "", "  ")
	fmt.Println(string(jsonData))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	return nil
}
