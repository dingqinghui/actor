/**
 * @Author: dingQingHui
 * @Description:
 * @File: system_message
 * @Version: 1.0.0
 * @Date: 2023/12/8 16:13
 */

package actor

type StartedMessage struct {
	Params []interface{}
}
type StopMessage struct {
}
type PanicMessage struct {
	err error
}

var _ IEnvelope = (*Envelope)(nil)

type Envelope struct {
	message interface{}
	sender  IProcess
}

func (e Envelope) Message() interface{} {
	return e.message
}

func (e Envelope) Sender() IProcess {
	return e.sender
}

func WrapEnvelope(Sender IProcess, Message interface{}) IEnvelope {
	return &Envelope{
		message: Message,
		sender:  Sender,
	}
}
