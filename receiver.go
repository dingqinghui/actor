/**
 * @Author: dingQingHui
 * @Description:
 * @File: receiver
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:22
 */

package actor

func NewDefaultReceiver() IReceiver {
	return &BuiltinReceiver{}
}

type BuiltinReceiver struct {
}

var _ IReceiver = &BuiltinReceiver{}

func (r *BuiltinReceiver) Initialize(ctx IContext, params ...interface{}) {
}
