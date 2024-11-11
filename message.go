/**
 * @Author: dingQingHui
 * @Description:
 * @File: system_message
 * @Version: 1.0.0
 * @Date: 2023/12/8 16:13
 */

package actor

const (
	SystemMessageMax = 0
	StartMessageId   = -1
	StopMessageId    = -2
	PanicMessageId   = -3
	timerMessageId   = -4
	SystemMessageMin = -5
)

var (
	NilMessage   = new(BuiltinMessage)
	startMessage = newSysMessage(StartMessageId)
	stopMessage  = newSysMessage(StopMessageId)
	panicMessage = newSysMessage(PanicMessageId)
)

func newSysMessage(msgId int32) IEnvelope {
	msg := &BuiltinMessage{
		msgId: msgId,
	}
	return WrapMessage(nil, msg)
}

func WrapMessage(sender IProcess, message IMessage) IEnvelope {
	return &EnvelopeMessage{
		IMessage: message,
		sender:   sender,
	}
}

func NewMessage(msgId int32, body interface{}) IMessage {
	return &BuiltinMessage{
		msgId: msgId,
		body:  body,
	}
}

type BuiltinMessage struct {
	msgId int32
	body  interface{}
}

var _ IMessage = (*BuiltinMessage)(nil)

func (b *BuiltinMessage) ID() int32 {
	return b.msgId
}

func (b *BuiltinMessage) Body() interface{} {
	return b.body
}

type EnvelopeMessage struct {
	IMessage
	sender IProcess
}

func (e *EnvelopeMessage) Sender() IProcess {
	return e.sender
}
