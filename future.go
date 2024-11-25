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
	f.ch = make(chan IEnvelopeMessage, 1)
	f.after = time.After(f.timeout)
	return f
}

type future struct {
	ch      chan IEnvelopeMessage
	timeout time.Duration
	after   <-chan time.Time
}

func (f *future) Wait() (err error) {
	select {
	case env := <-f.ch:
		if len(env.Args()) > 0 && env.Args()[0] != nil {
			err = env.Args()[0].(error)
		}
	case <-f.after:
		err = ErrActorCallTimeout
	}
	return
}

func (f *future) Process() IProcess {
	return f
}

func (f *future) Call(funcName string, timeout time.Duration, reply, request interface{}) error {
	panic("future call not imp")
}

func (f *future) Send(funcName string, args ...interface{}) error {
	f.ch <- WrapEnvMessage(funcName, nil, args...)
	return nil
}

func (f *future) Stop() error {
	return nil
}
