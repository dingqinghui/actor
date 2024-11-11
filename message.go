/**
 * @Author: dingQingHui
 * @Description:
 * @File: system_message
 * @Version: 1.0.0
 * @Date: 2023/12/8 16:13
 */

package actor

const (
	InitFuncName  = "Init"
	StopFuncName  = "Stop"
	PanicFuncName = "Panic"
)

var (
	NilMsg = struct{}{}
)

type EnvelopeMessage struct {
	msg      interface{}
	funcName string
	sender   IProcess
}

func (e *EnvelopeMessage) FuncName() string {
	return e.funcName
}

func (e *EnvelopeMessage) Msg() interface{} {
	return e.msg
}
func (e *EnvelopeMessage) Sender() IProcess {
	return e.sender
}

func WrapEnvMessage(funcName string, sender IProcess, msg interface{}) *EnvelopeMessage {
	return &EnvelopeMessage{
		msg:      msg,
		sender:   sender,
		funcName: funcName,
	}
}

func UnwrapEnvMessage(env IEnvelopeMessage) (funcName string, sender IProcess, msg interface{}) {
	if env == nil {
		return
	}
	return env.FuncName(), env.Sender(), env.Msg()
}
