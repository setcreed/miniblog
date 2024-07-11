// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package post

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/miniblog/internal/pkg/core"
	"github.com/setcreed/miniblog/internal/pkg/known"
	"github.com/setcreed/miniblog/internal/pkg/log"
)

// Get 获取指定的博客.
func (ctrl *PostController) Get(c *gin.Context) {
	log.C(c).Infow("Get post function called")

	post, err := ctrl.b.Posts().Get(c, c.GetString(known.XUsernameKey), c.Param("postID"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, post)
}
