// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package id

import (
	shortid "github.com/jasonsoft/go-short-id"
)

// GenShortID 生成 6 位字符长度的唯一 ID.
func GenShortID() string {
	opt := shortid.Options{
		Number:        4,
		StartWithYear: true,
		EndWithHost:   false,
	}

	return toLower(shortid.Generate(opt))
}

func toLower(ss string) string {
	var lower string
	for _, s := range ss {
		lower += string(s | ' ')
	}

	return lower
}
