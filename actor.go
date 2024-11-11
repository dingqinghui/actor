/**
 * @Author: dingQingHui
 * @Description:
 * @File: receiver
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:22
 */

package actor

type IActor interface {
	Init(ctx IContext, msg interface{})
	Stop(ctx IContext, msg interface{})
	Panic(ctx IContext, msg interface{})
}

type BuiltinActor struct {
}

func (r *BuiltinActor) Init(ctx IContext, msg interface{}) {

}

func (r *BuiltinActor) Stop(ctx IContext, msg interface{}) {

}

func (r *BuiltinActor) Panic(ctx IContext, msg interface{}) {
	println(msg)
}
