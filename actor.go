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
}

func (r *BuiltinActor) Init(ctx IContext, msg interface{}) {

}

func (r *BuiltinActor) Stop(ctx IContext) {
}

func (r *BuiltinActor) Panic(ctx IContext, err interface{}) {
	zlog.Panic("panic", zap.Error(err.(error)))
}
