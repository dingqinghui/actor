/**
 * @Author: dingQingHui
 * @Description:
 * @File: route
 * @Version: 1.0.0
 * @Date: 2024/11/5 15:24
 */

package actor

import (
	"errors"
	"github.com/dingqinghui/zlog"
	"go.uber.org/zap"
	"reflect"
	"unicode"
	"unicode/utf8"
)

func DefaultMethod(actor IActor, ctx IContext) error { return nil }

type method struct {
	name     string
	fun      reflect.Value
	typ      reflect.Type
	argTypes []reflect.Type
	argNum   int
}

func (m *method) call(ctx IContext, args []interface{}) error {
	if len(args) != m.argNum {
		zlog.Error("actor method call args num wrong",
			zap.String("typeName", reflect.TypeOf(ctx.Actor()).String()),
			zap.String("methodName", m.name),
			zap.Int("argNum", m.argNum),
			zap.Int("inNum", len(args)))
		return errors.New("args count err")
	}
	argValues := make([]reflect.Value, 2+m.argNum, 2+m.argNum)
	argValues[0] = reflect.ValueOf(ctx.Actor())
	argValues[1] = reflect.ValueOf(ctx)
	for i, arg := range args {
		if !checkArgsType(reflect.TypeOf(arg), m.argTypes[i]) {
			zlog.Error("actor method call args type err",
				zap.String("typeName", reflect.TypeOf(ctx.Actor()).String()),
				zap.String("methodName", m.name),
				zap.String("parameter", m.argTypes[i].String()),
				zap.String("argument", reflect.TypeOf(arg).String()))
			return errors.New("args type err")
		}
		argValues[i+2] = valueOf(m.argTypes[i], arg)
	}
	returnValues := m.fun.Call(argValues)

	errInter := returnValues[0].Interface()
	if errInter != nil {
		return errInter.(error)
	}
	return nil
}

func newHandlers(actor IActor) *handlers {
	h := new(handlers)
	h.register(actor)
	return h
}

type handlers struct {
	dict  map[string]*method
	actor reflect.Value
}

func (h *handlers) register(actor IActor) {
	h.dict = suitableMethods(reflect.TypeOf(actor))
}
func (h *handlers) call(ctx IContext, env IEnvelopeMessage) error {
	funcName, _, args := UnwrapEnvMessage(env)
	m, ok := h.dict[funcName]
	if !ok {
		return nil
	}
	return m.call(ctx, args)
}

func newType(t reflect.Type) interface{} {
	var argv reflect.Value
	if t.Kind() == reflect.Ptr {
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

func suitableMethods(typ reflect.Type) map[string]*method {
	var defaultParamType []reflect.Type
	dt := reflect.TypeOf(DefaultMethod)
	for i := 0; i < dt.NumIn(); i++ {
		defaultParamType = append(defaultParamType, dt.In(i))
	}

	methods := make(map[string]*method)
	for index := 0; index < typ.NumMethod(); index++ {
		fun := typ.Method(index)
		funType := fun.Type
		funName := fun.Name
		if fun.PkgPath != "" {
			continue
		}
		if funType.NumIn() < 2 {
			continue
		}
		if funType.NumOut() != 1 {
			continue
		}
		if funType.Out(0) != dt.Out(0) {
			continue
		}
		// IActor,request,reply....
		if !funType.In(0).Implements(defaultParamType[0]) {
			continue
		}
		n1 := funType.In(1).Name()
		n2 := defaultParamType[1].Name()
		_, _ = n1, n2
		if funType.In(1) != defaultParamType[1] {
			continue
		}

		argNum := funType.NumIn() - 2
		// 检测是否所有参数都是导出类型
		isExported := true
		argTypes := make([]reflect.Type, argNum, argNum)
		for i := 2; i < funType.NumIn(); i++ {
			argType := funType.In(i)
			if !isExportedType(argType) {
				isExported = false
				break
			}
			argTypes[i-2] = argType
		}
		if !isExported {
			continue
		}
		methods[funName] = &method{
			fun:      fun.Func,
			typ:      funType,
			name:     funName,
			argNum:   argNum,
			argTypes: argTypes,
		}
	}
	return methods
}

func isExportedType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name := t.Name()
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune) || t.PkgPath() == ""
}

// checkArgsType
// @Description:
// @param argumentType 实参类型
// @param parameterType 形参类型
// @return bool
func checkArgsType(argumentType, parameterType reflect.Type) bool {
	if parameterType == argumentType {
		return true
	}
	if parameterType.Kind() == reflect.Interface {
		if argumentType.Implements(parameterType) {
			return true
		}
	}

	return false
}
