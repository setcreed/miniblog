// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package user

import (
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/pkg/core"
	"github.com/setcreed/miniblog/internal/pkg/errno"
	"github.com/setcreed/miniblog/internal/pkg/log"
	v1 "github.com/setcreed/miniblog/pkg/api/miniblog/v1"
)

// Update 更新用户信息.
func (ctrl *UserController) Update(c *gin.Context) {
	log.C(c).Infow("Update user function called")

	var r v1.UpdateUserRequest
	processRequest(c, &r, func() error {
		return ctrl.b.Users().Update(c, c.Param("name"), &r)
	})
}

func processRequest(c *gin.Context, r interface{}, handler func() error) {
	if err := c.ShouldBindJSON(r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)
		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)
		return
	}

	if err := handler(); err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
}
