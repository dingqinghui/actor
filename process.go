/**
 * @Author: dingQingHui
 * @Description:
 * @File: process
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:43
 */

package actor

import (
	"sync/atomic"
	"time"
)

func NewBaseProcess(mailBox IMailbox) IProcess {
	process := &ProcessActor{
		mailBox: mailBox,
	}
	process.isStop.Store(false)
	return process
}

var _ IProcess = (*ProcessActor)(nil)

type ProcessActor struct {
	isStop  atomic.Bool
	mailBox IMailbox
}

func (p *ProcessActor) Send(message interface{}) error {
	if p.mailBox == nil {
		return ErrMailBoxNil
	}
	if p.isStop.CompareAndSwap(true, true) {
		return ErrActorStopped
	}

	return p.mailBox.PostMessage(message)
}

func (p *ProcessActor) Stop() error {
	if p.isStop.CompareAndSwap(false, true) {
		return p.mailBox.PostMessage(&StopMessage{})
	}
	return nil
}

func (p *ProcessActor) Call(message interface{}, timeout time.Duration) (IFuture, error) {
	if p.mailBox == nil {
		return nil, ErrMailBoxNil
	}
	if p.isStop.CompareAndSwap(true, true) {
		return nil, ErrActorStopped
	}

	fut := newFuture(timeout)
	env := WrapEnvelope(fut.Process(), message)
	return fut, p.mailBox.PostMessage(env)
}
