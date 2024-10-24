/**
 * @Author: dingQingHui
 * @Description:
 * @File: pool
 * @Version: 1.0.0
 * @Date: 2024/10/15 15:04
 */

package actor

import (
	"github.com/panjf2000/ants/v2"
)

var pool *ants.Pool

func init() {
	_pool, err := ants.NewPool(1000)
	if err != nil {
		panic("ants.NewPool, error")
	}
	pool = _pool
}

func Submit(fn func(), recoverFun func(err interface{})) error {
	return pool.Submit(func() {
		Try(fn, recoverFun)
	})
}

func Try(fn func(), recoverFun func(err interface{})) {
	defer func() {
		if err := recover(); err != nil {
			if recoverFun != nil {
				recoverFun(err)
			}
		}
	}()
	fn()
}
