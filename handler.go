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

type method struct {
	fun     reflect.Value
	typ     reflect.Type
	inTypes []reflect.Type
	inNum   int
}

func (m *method) call(ctx IContext, msg interface{}) {
	args := make([]reflect.Value, 2, 2)
	args[0] = reflect.ValueOf(ctx.Actor())
	args[1] = reflect.ValueOf(ctx)
	if m.inNum == 2 {
		m.fun.Call(args)
		return
	}
	args = append(args, valueOf(m.inTypes[2], msg))
	m.fun.Call(args)
}

func newHandlers(actor IActor) *handlers {
	h := &handlers{
		dict: make(map[string]*method),
	}
	h.init(actor)
	return h
}

type handlers struct {
	dict  map[string]*method
	actor reflect.Value
}

func (h *handlers) init(actor IActor) {
	typ := reflect.TypeOf(actor)
	for i := 0; i < typ.NumMethod(); i++ {
		m := typ.Method(i)
		mt := m.Type

		if !h.isHandlerMethod(m) {
			continue
		}
		_m := &method{
			fun:   m.Func,
			typ:   mt,
			inNum: m.Type.NumIn(),
		}
		for j := 0; j < m.Type.NumIn(); j++ {
			t := m.Type.In(j)
			_m.inTypes = append(_m.inTypes, t)
		}
		h.set(m.Name, _m)
	}
}
func (h *handlers) call(ctx IContext, env IEnvelopeMessage) {
	funcName, _, msg := UnwrapEnvMessage(env)
	m, ok := h.dict[funcName]
	if !ok {
		return
	}
	m.call(ctx, msg)
}
func (h *handlers) set(methodName string, m *method) {
	h.dict[methodName] = m
}

func (h *handlers) isHandlerMethod(method reflect.Method) bool {
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

func newType(t reflect.Type) interface{} {
	var argv reflect.Value

	if t.Kind() == reflect.Ptr { // reply must be ptr
		argv = reflect.New(t.Elem())
	} else {
		argv = reflect.New(t)
	}
	return argv.Interface()
}

func valueOf(t reflect.Type, msg interface{}) reflect.Value {
	if msg != nil {
		return reflect.ValueOf(msg)
	}
	v := newType(t)
	if t.Kind() != reflect.Ptr {
		return reflect.ValueOf(v).Elem()
	} else {
		return reflect.ValueOf(v)
	}
}
