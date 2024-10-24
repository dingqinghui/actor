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
	f.ch = make(chan interface{}, 1)
	f.after = time.After(f.timeout)
	return f
}

type future struct {
	ch      chan interface{}
	timeout time.Duration
	after   <-chan time.Time
}

func (f *future) Wait() (result interface{}, isTimeout bool) {
	select {
	case result = <-f.ch:
	case <-f.after:
		isTimeout = true
	}
	return
}

func (f *future) Process() IProcess {
	return f
}

func (f *future) Call(message interface{}, timeout time.Duration) (IFuture, error) {
	panic("future call not imp")
}

func (f *future) Send(message interface{}) error {
	f.ch <- message
	return nil
}

func (f *future) Stop() error {
	return nil
}
