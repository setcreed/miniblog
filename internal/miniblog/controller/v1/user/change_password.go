// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/pkg/log"
	v1 "github.com/setcreed/miniblog/pkg/api/miniblog/v1"
)

// ChangePassword 用来修改指定用户的密码.
func (ctrl *UserController) ChangePassword(c *gin.Context) {
	log.C(c).Infow("Change password function called")

	var r v1.ChangePasswordRequest
	processRequest(c, &r, func() error {
		return ctrl.b.Users().ChangePassword(c, c.Param("name"), &r)
	})
}
