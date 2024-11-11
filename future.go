/**
 * @Author: dingQingHui
 * @Description:
 * @File: future
 * @Version: 1.0.0
 * @Date: 2024/10/24 14:23
 */

package actor

import "time"

func newFuture(timeout time.Duration) *future {
	f := new(future)
	f.timeout = timeout
	f.ch = make(chan IEnvelope, 1)
	f.after = time.After(f.timeout)
	return f
}

type future struct {
	ch      chan IEnvelope
	timeout time.Duration
	after   <-chan time.Time
}

func (f *future) Wait() (result interface{}, isTimeout bool) {
	select {
	case env := <-f.ch:
		result = env.Body()
	case <-f.after:
		isTimeout = true
	}
	return
}

func (f *future) Process() IProcess {
	return f
}

func (f *future) Call(message IMessage, timeout time.Duration) (IFuture, error) {
	panic("future call not imp")
}

func (f *future) Send(message IMessage) error {
	f.ch <- WrapMessage(nil, message)
	return nil
}

func (f *future) Stop() error {
	return nil
}
