// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package miniblog

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/miniblog/controller/v1/user"
	"github.com/setcreed/miniblog/internal/pkg/core"
	"github.com/setcreed/miniblog/internal/pkg/errno"
	"github.com/setcreed/miniblog/internal/pkg/log"
	mw "github.com/setcreed/miniblog/internal/pkg/middleware"
	"github.com/setcreed/miniblog/internal/pkg/store"
	"github.com/setcreed/miniblog/pkg/auth"
)

// installRouters 安装 miniblog 接口路由.
func installRouters(g *gin.Engine) error {
	// 注册 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		core.WriteResponse(c, errno.ErrPageNotFound, nil)
	})

	// 注册 /healthz handler.
	g.GET("/healthz", func(c *gin.Context) {
		log.C(c).Infow("Healthz function called")

		core.WriteResponse(c, nil, map[string]string{"status": "ok"})
	})

	authz, err := auth.NewAuthz(store.S.DB())
	if err != nil {
		return err
	}

	uc := user.New(store.S, authz)

	g.POST("/login", uc.Login)

	// 创建 v1 路由分组
	v1 := g.Group("/v1")
	{
		// 创建 users 路由分组
		userv1 := v1.Group("/users")
		{
			userv1.POST("", uc.Create)                             // 创建用户
			userv1.PUT(":name/change-password", uc.ChangePassword) // 修改用户密码
			userv1.Use(mw.Authn(), mw.Authz(authz))
			userv1.GET(":name", uc.Get) // 获取用户详情
		}
	}

	return nil
}
