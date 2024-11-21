/**
 * @Author: dingQingHui
 * @Description:
 * @File: receiver
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:22
 */

package actor

import (
	"github.com/dingqinghui/zlog"
	"go.uber.org/zap"
)

type BuiltinActor struct {
	Ctx IContext
}

func (r *BuiltinActor) Init(ctx IContext, msg interface{}) error {
	r.Ctx = ctx
	return nil
}

func (r *BuiltinActor) Stop() error {
	return nil
}

func (r *BuiltinActor) Panic(errMsg string) error {
	zlog.Panic("panic", zap.String("errMsg", errMsg))
	return nil
}
