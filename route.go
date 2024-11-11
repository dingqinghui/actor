/**
 * @Author: dingQingHui
 * @Description:
 * @File: route
 * @Version: 1.0.0
 * @Date: 2024/11/5 15:24
 */

package actor

func NewBuiltinRoutes() IRoutes {
	return &BuiltinRoutes{
		routers: make(map[int32]MessageHandler),
	}
}

type BuiltinRoutes struct {
	routers map[int32]MessageHandler
}

func (r *BuiltinRoutes) Add(msgId int32, fn MessageHandler) error {
	_, ok := r.routers[msgId]
	if ok {
		return ErrRouteAddRepeat
	}
	r.routers[msgId] = fn
	return nil
}

func (r *BuiltinRoutes) Get(msgId int32) MessageHandler {
	fnInfo, _ := r.routers[msgId]
	return fnInfo
}
