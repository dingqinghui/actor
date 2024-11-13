/**
 * @Author: dingQingHui
 * @Description:
 * @File: system
 * @Version: 1.0.0
 * @Date: 2023/12/7 14:54
 */

package actor

import "time"

type System struct {
	nameHub IProcess
}

func NewSystem() ISystem {
	s := new(System)
	s.nameHub = newNameHubActor(s)
	return s
}

func (s *System) Spawn(b IBlueprint, producer Producer, params interface{}) (IProcess, error) {
	return b.Spawn(s, producer, params)
}

func (s *System) Named(name string, p IProcess) error {
	return s.nameHub.Send("Named", &namedMsg{
		name:    name,
		process: p,
	})
}

func (s *System) GetProcessByName(name string) (IProcess, error) {
	res, _, err := s.nameHub.Call("Get", name, time.Second*3)
	if err != nil {
		return nil, err
	}
	return res.(IProcess), nil

}

func (s *System) DelName(name string) error {
	return s.nameHub.Send("Del", name)
}
