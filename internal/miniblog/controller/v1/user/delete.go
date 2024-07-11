// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/pkg/core"
	"github.com/setcreed/miniblog/internal/pkg/log"
)

// Delete 删除一个用户.
func (ctrl *UserController) Delete(c *gin.Context) {
	log.C(c).Infow("Delete user function called")

	username := c.Param("name")

	if err := ctrl.b.Users().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if _, err := ctrl.a.RemoveNamedPolicy("p", username, "", ""); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
