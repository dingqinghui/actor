/**
 * @Author: dingQingHui
 * @Description:
 * @File: process
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:43
 */

package actor

import (
	"github.com/duke-git/lancet/v2/convertor"
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
	cloneArgs := make([]interface{}, len(args), len(args))
	for i, arg := range args {
		cloneArgs[i] = convertor.DeepClone(arg)
	}
	env := WrapEnvMessage(funcName, nil, cloneArgs...)
	return p.mailBox.PostMessage(env)
}

func (p *ProcessActor) Call(funcName string, timeout time.Duration, request, reply interface{}) error {
	if p.mailBox == nil {
		return ErrMailBoxNil
	}
	if p.isStop.CompareAndSwap(true, true) {
		return ErrActorStopped
	}
	cloneRequest := convertor.DeepClone(request)
	cloneReply := convertor.DeepClone(reply)

	fut := newFuture(timeout)
	env := WrapEnvMessage(funcName, fut.Process(), cloneRequest, cloneReply)
	if err := p.mailBox.PostMessage(env); err != nil {
		return err
	}
	if err := fut.Wait(); err != nil {
		return err
	}
	if err := convertor.CopyProperties(reply, cloneReply); err != nil {
		return err
	}
	return nil
}

func (p *ProcessActor) Stop() error {
	if !p.isStop.CompareAndSwap(false, true) {
		return ErrActorStopped
	}
	fut := newFuture(time.Millisecond * 10)
	env := WrapEnvMessage(StopFuncName, fut.Process())
	if err := p.mailBox.PostMessage(env); err != nil {
		return err
	}
	fut.Wait()
	return nil
}
