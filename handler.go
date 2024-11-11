/**
 * @Author: dingQingHui
 * @Description:
 * @File: route
 * @Version: 1.0.0
 * @Date: 2024/11/5 15:24
 */

package actor

import (
	"reflect"
)

func defaultHandlerFunc(actor IActor, ctx IContext, msg interface{}) {

}

type handler struct {
	method reflect.Method
	typ    reflect.Type
}

func newHandlerContainer(actor IActor, dict map[string]*handler) *handlerContainer {
	h := &handlerContainer{
		dict:  dict,
		actor: reflect.ValueOf(actor),
	}
	return h
}

type handlerContainer struct {
	dict  map[string]*handler
	actor reflect.Value
}

func (h *handlerContainer) Call(ctx IContext, env IEnvelopeMessage) {
	funcName, _, msg := UnwrapEnvMessage(env)
	handle, ok := h.dict[funcName]
	if !ok {
		return
	}
	v := reflect.ValueOf(msg)
	t := v.Kind()
	_ = t
	args := make([]reflect.Value, 3, 3)
	args[0] = h.actor
	args[1] = reflect.ValueOf(ctx)
	if msg == nil {
		args[2] = reflect.ValueOf(NilMsg)
	} else {
		args[2] = reflect.ValueOf(msg)
	}
	handle.method.Func.Call(args)
}

func getActorHandler(actor IActor) map[string]*handler {
	dict := make(map[string]*handler)
	typ := reflect.TypeOf(actor)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		mt := method.Type
		mn := method.Name
		if !isHandlerMethod(method) {
			continue
		}
		dict[mn] = &handler{
			method: method,
			typ:    mt,
		}
	}
	return dict
}

func isHandlerMethod(method reflect.Method) bool {
	if !method.IsExported() {
		return false
	}
	mt := method.Type
	if mt.NumIn() != 3 {
		return false
	}

	v := reflect.TypeOf(defaultHandlerFunc)

	for i := 0; i < v.NumMethod(); i++ {
		if mt.In(1).Name() != v.In(1).Name() {
			return false
		}
	}
	return true
}
