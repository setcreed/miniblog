// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/pkg/core"
	"github.com/setcreed/miniblog/internal/pkg/errno"
	"github.com/setcreed/miniblog/internal/pkg/known"
	"github.com/setcreed/miniblog/internal/pkg/log"
)

// Author 用来定义授权接口实现.
// sub: 操作主题，obj：操作对象, act：操作
type Author interface {
	Authorize(sub, obj, act string) (bool, error)
}

// Authz 是 Gin 中间件，用来进行请求授权.
func Authz(a Author) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetString(known.XUsernameKey)
		obj := c.Request.URL.Path
		act := c.Request.Method

		log.Debugw("Build authorize context", "sub", sub, "obj", obj, "act", act)
		if allowed, _ := a.Authorize(sub, obj, act); !allowed {
			core.WriteResponse(c, errno.ErrUnauthorized, nil)
			c.Abort()
			return
		}
	}
}
