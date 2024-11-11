/**
 * @Author: dingQingHui
 * @Description:111
 * @File: api
 * @Version: 1.0.0
 * @Date: 2024/10/24 9:53
 */

package actor

import (
	"time"
)

type Producer func() IActor

type IMessageInvoker interface {
	InvokerMessage(message IEnvelopeMessage) error
}

type IMailbox interface {
	//
	// PostMessage
	// @Description: 向邮箱投递消息(写入)
	// @param msg
	//
	PostMessage(msg IEnvelopeMessage) error
	//
	// RegisterHandlers
	// @Description: 注册消息处理函数(取出并处理)
	// @param invoker
	//
	RegisterHandlers(invoker IMessageInvoker, dispatcher IDispatcher)
}

type IDispatcher interface {
	//
	// Schedule
	// @Description: 调度
	// @param f
	// @param recoverFun
	//
	Schedule(f func(), recoverFun func(err interface{})) error
	//
	// Throughput
	// @Description: 单次调度最大吞吐量
	// @return int
	//
	Throughput() int
}

type IContext interface {
	EnvMessage() IEnvelopeMessage
	Process() IProcess
	System() ISystem
	Actor() IActor
	AddTimer(d time.Duration, funcName string)
}

type IProcess interface {
	//
	// Send
	// @Description: 发送异步消息
	// @param message
	// @return error
	//
	Send(funcName string, message interface{}) error
	//
	// Call
	// @Description:发送同步消息
	// @param message
	// @param timeout
	// @return IFuture
	// @return error
	//
	Call(funcName string, message interface{}, timeout time.Duration) (IFuture, error)
	//
	// Stop
	// @Description: 停止Actor
	// @param isGrace 处理完接收到消息后关闭
	//
	Stop() error
}

type IBlueprint interface {
	Spawn(system ISystem, producer Producer, params ...interface{}) (IProcess, error)
}

type INamedHub interface {
	Named(name string, p IProcess) error
	GetProcessByName(name string) (IProcess, error)
	DelName(name string) error
}

type ISystem interface {
	INamedHub
	Spawn(b IBlueprint, producer Producer, params ...interface{}) (IProcess, error)
}

type IFuture interface {
	Wait() (result interface{}, isTimeout bool)
}

type IEnvelopeMessage interface {
	FuncName() string
	Msg() interface{}
	Sender() IProcess
}
