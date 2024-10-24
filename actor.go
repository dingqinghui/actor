/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor
 * @Version: 1.0.0
 * @Date: 2023/12/8 14:19
 */

package actor

type baseActorContext struct {
	receiver IReceiver
	process  IProcess
	system   ISystem
	message  interface{}
}

var _ IContext = &baseActorContext{}
var _ IMessageInvoker = &baseActorContext{}

func NewBaseActorContext(receiver IReceiver, system ISystem, process IProcess) *baseActorContext {
	a := &baseActorContext{
		receiver: receiver,
		system:   system,
		process:  process,
	}
	return a
}

func (a *baseActorContext) InvokerMessage(message interface{}) error {
	a.message = message
	switch message.(type) {
	case *StartedMessage:
		msg := message.(*StartedMessage)
		a.receiver.Init(a, msg.Params...)
		return nil
	case *StopMessage:
		a.receiver.Stop(a)
		return nil
	case *PanicMessage:
		a.receiver.Panic(a)
		return nil
	}
	return a.Receive()
}

func (a *baseActorContext) Receive() error {
	if a.receiver == nil {
		return ErrActorReceiveIsNil
	}
	return a.receiver.Receive(a)
}

func (a *baseActorContext) Message() interface{} {
	return a.message
}

func (a *baseActorContext) Process() IProcess {
	return a.process
}

func (a *baseActorContext) System() ISystem {
	return a.system
}
