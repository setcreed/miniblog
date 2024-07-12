// # Copyright 2024 setcreed <cwzdzg@foxmail.com>. All rights reserved.
// # Use of this source code is governed by a MIT style
// # license that can be found in the LICENSE file. The original repo for
// # this file is https://github.com/setcreed/miniblog.

package post

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/setcreed/miniblog/internal/miniblog/store"
)

func TestNew(t *testing.T) {
	type args struct {
		ds store.IStore
	}
	tests := []struct {
		name string
		args args
		want *PostController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, New(tt.args.ds))
		})
	}
}
