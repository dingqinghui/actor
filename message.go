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
	args     []interface{}
	funcName string
	sender   IProcess
	reply    interface{}
}

func (e *EnvelopeMessage) FuncName() string {
	return e.funcName
}

func (e *EnvelopeMessage) Args() []interface{} {
	return e.args
}
func (e *EnvelopeMessage) Sender() IProcess {
	return e.sender
}

func WrapEnvMessage(funcName string, sender IProcess, args ...interface{}) *EnvelopeMessage {
	return &EnvelopeMessage{
		args:     args,
		sender:   sender,
		funcName: funcName,
	}
}

func UnwrapEnvMessage(env IEnvelopeMessage) (funcName string, sender IProcess, args []interface{}) {
	if env == nil {
		return
	}
	return env.FuncName(), env.Sender(), env.Args()
}

func IsSyncMessage(env IEnvelopeMessage) bool {
	return env.Sender() != nil
}
