/**
 * @Author: dingQingHui
 * @Description:
 * @File: process
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:43
 */

package actor

import (
	"errors"
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

func (p *ProcessActor) Send(funcName string, args ...interface{}) error {
	if p.mailBox == nil {
		return ErrMailBoxNil
	}
	if p.isStop.CompareAndSwap(true, true) {
		return ErrActorStopped
	}
	env := WrapEnvMessage(funcName, nil, args...)
	return p.mailBox.PostMessage(env)
}

func (p *ProcessActor) Call(funcName string, timeout time.Duration, reply, request interface{}) error {
	if p.mailBox == nil {
		return ErrMailBoxNil
	}
	if p.isStop.CompareAndSwap(true, true) {
		return ErrActorStopped
	}
	fut := newFuture(timeout)
	env := WrapEnvMessage(funcName, fut.Process(), reply, request)
	if err := p.mailBox.PostMessage(env); err != nil {
		return err
	}
	res, isTimeout := fut.Wait()
	if isTimeout {
		return errors.New("time out")
	}
	if res[0] == nil {
		return nil
	}
	return res[0].(error)
}

func (p *ProcessActor) Stop() error {
	if !p.isStop.CompareAndSwap(false, true) {
		return errors.New("actor stopped")
	}
	fut := newFuture(time.Millisecond * 10)
	env := WrapEnvMessage(StopFuncName, fut.Process())
	if err := p.mailBox.PostMessage(env); err != nil {
		return err
	}
	fut.Wait()
	return nil
}
