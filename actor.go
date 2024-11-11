/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor
 * @Version: 1.0.0
 * @Date: 2023/12/8 14:19
 */

package actor

import (
	"sync"
	"time"
)

type baseActorContext struct {
	actor      interface{}
	process    IProcess
	system     ISystem
	routers    IRoutes
	message    IEnvelope
	initParams []interface{}
	initOnce   sync.Once
}

var _ IContext = &baseActorContext{}
var _ IMessageInvoker = &baseActorContext{}

func NewBaseActorContext() *baseActorContext {
	return new(baseActorContext)
}
func (a *baseActorContext) initialize() {
	// 注册定时器回调
	a.registerTimerHandler()
}
func (a *baseActorContext) registerTimerHandler() {
	// 注册定时器处理函数
	_ = a.Routes().Add(timerMessageId, func(ctx IContext, msg IEnvelope) {
		handler := msg.Body().(MessageHandler)
		if handler == nil {
			return
		}
		handler(ctx, msg)
	})
}

func (a *baseActorContext) InvokerMessage(message IEnvelope) error {
	// 执行消息回调
	a.message = message
	handler := a.Routes().Get(a.message.ID())
	if handler == nil {
		return nil
	}
	handler(a, message)
	return nil
}
func (a *baseActorContext) AddTimer(d time.Duration, handler MessageHandler) {
	tw.AfterFunc(d, func() {
		msg := NewMessage(timerMessageId, handler)
		_ = a.Process().Send(msg)
	})
}

func (a *baseActorContext) Message() IEnvelope {
	return a.message
}

func (a *baseActorContext) Process() IProcess {
	return a.process
}

func (a *baseActorContext) System() ISystem {
	return a.system
}

func (a *baseActorContext) Routes() IRoutes {
	return a.routers
}

func (a *baseActorContext) Actor() IActor {
	return a.actor
}
