/**
 * @Author: dingQingHui
 * @Description:
 * @File: mailbox
 * @Version: 1.0.0
 * @Date: 2024/10/15 14:27
 */

package actor

import (
	"runtime"
	"sync/atomic"
)

const (
	idle int32 = iota
	running
)

var _ IMailbox = &mailbox{}

type mailbox struct {
	invoker      IMessageInvoker
	queue        *Queue
	dispatch     IDispatcher
	dispatchStat atomic.Int32
}

var _ IMailbox = &mailbox{}

func NewMailbox() *mailbox {
	mailbox := &mailbox{
		queue: NewQueue(),
	}
	return mailbox
}

func (m *mailbox) RegisterHandlers(invoker IMessageInvoker, dispatcher IDispatcher) {
	m.invoker = invoker
	m.dispatch = dispatcher
}

func (m *mailbox) PostMessage(msg IEnvelope) error {
	m.queue.Push(msg)
	return m.schedule()
}

func (m *mailbox) schedule() error {
	if !m.dispatchStat.CompareAndSwap(idle, running) {
		return nil
	}
	if err := m.dispatch.Schedule(m.process, func(err interface{}) {
		_ = m.invoker.InvokerMessage(panicMessage)
	}); err != nil {
		return err
	}
	return nil
}

func (m *mailbox) process() {
	m.run()
	m.dispatchStat.Store(idle)
}

func (m *mailbox) run() {
	throughput := m.dispatch.Throughput()
	var i int
	for true {
		if i > throughput {
			i = 0
			m.dispatchStat.Store(idle)
			runtime.Gosched()
			continue
		}
		i++
		msg := m.queue.Pop()
		if msg != nil {
			_ = m.invokerMessage(msg.(IEnvelope))
		} else {
			return
		}
	}
}

// invokerMessage
// @Description: 从队列中读取消息，并调用invoker处理
// @receiver m
// @return error
func (m *mailbox) invokerMessage(msg IEnvelope) error {
	if err := m.invoker.InvokerMessage(msg); err != nil {
		return err
	}
	return nil
}
