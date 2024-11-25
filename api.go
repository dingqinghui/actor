/**
 * @Author: dingQingHui
 * @Description:111
 * @File: api
 * @Version: 1.0.0
 * @Date: 2024/10/24 9:53
 */

package actor

import (
	"github.com/dingqinghui/zlog"
	"go.uber.org/zap"
	"time"
)

type Producer func() IActor

type IMessageInvoker interface {
	InvokerMessage(message IEnvelopeMessage) error
}

type IMailbox interface {
	PostMessage(msg IEnvelopeMessage) error
	RegisterHandlers(invoker IMessageInvoker, dispatcher IDispatcher)
}

type IDispatcher interface {
	Schedule(f func(), recoverFun func(err interface{})) error
	Throughput() int
}

type IContext interface {
	EnvMessage() IEnvelopeMessage
	Process() IProcess
	System() ISystem
	Actor() IActor
	AddTimer(d time.Duration, funcName string)
}

type IProcess interface {
	Send(funcName string, args ...interface{}) error
	Call(funcName string, timeout time.Duration, request, reply interface{}) error
	Stop() error
}

type IBlueprint interface {
	Spawn(system ISystem, producer Producer, params interface{}) (IProcess, error)
}

type ISystem interface {
	Spawn(b IBlueprint, producer Producer, params interface{}) (IProcess, error)
}

type IFuture interface {
	Wait() (result interface{}, isTimeout bool)
}

type IEnvelopeMessage interface {
	FuncName() string
	Args() []interface{}
	Sender() IProcess
}

type IActor interface {
	Init(ctx IContext, msg interface{}) error
	Stop(ctx IContext) error
	Panic(ctx IContext, err interface{}) error
}

type BuiltinActor struct {
}

func (r *BuiltinActor) Init(ctx IContext, msg interface{}) error {
	return nil
}

func (r *BuiltinActor) Stop(ctx IContext) error {
	return nil
}

func (r *BuiltinActor) Panic(ctx IContext, err interface{}) error {
	zlog.Panic("panic", zap.Error(err.(error)), zap.Stack("stack"))
	return nil
}
