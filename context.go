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
	actor      IActor
	process    IProcess
	system     ISystem
	env        IEnvelopeMessage
	initParams []interface{}
	initOnce   sync.Once
	handlers   *handlers
}

var _ IContext = &baseActorContext{}
var _ IMessageInvoker = &baseActorContext{}

func NewBaseActorContext() *baseActorContext {
	return new(baseActorContext)
}

func (a *baseActorContext) InvokerMessage(env IEnvelopeMessage) error {
	// 执行消息回调
	a.env = env
	err := a.handlers.call(a, env)
	if IsSyncMessage(env) {
		return a.respond(err)
	}
	return nil
}

func (a *baseActorContext) AddTimer(d time.Duration, funcName string) {
	tw.AfterFunc(d, func() {
		_ = a.Process().Send(funcName)
	})
}

func (a *baseActorContext) EnvMessage() IEnvelopeMessage {
	return a.env
}

func (a *baseActorContext) Process() IProcess {
	return a.process
}

func (a *baseActorContext) System() ISystem {
	return a.system
}

func (a *baseActorContext) Actor() IActor {
	return a.actor
}

func (a *baseActorContext) respond(err error) error {
	if a.EnvMessage() == nil {
		return ErrActorRespondEnvIsNil
	}
	sender := a.EnvMessage().Sender()
	if sender == nil {
		return ErrActorRespondSenderIsNil
	}
	return a.EnvMessage().Sender().Send("", err)
}
