/**
 * @Author: dingQingHui
 * @Description:
 * @File: system
 * @Version: 1.0.0
 * @Date: 2023/12/7 14:54
 */

package actor

type System struct {
}

func NewSystem() ISystem {
	s := &System{}
	return s
}

func (s *System) Spawn(b IBlueprint, params ...interface{}) (IProcess, error) {
	return b.Spawn(s, params...)
}
