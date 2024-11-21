/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor_test
 * @Version: 1.0.0
 * @Date: 2024/10/24 11:25
 */

package examples

import (
	"errors"
	"fmt"
	"github.com/dingqinghui/actor"
	"reflect"
	"testing"
	"time"
)

type Message struct {
	A int
}

type testActor struct {
	actor.BuiltinActor
}

func (t *testActor) TestHandler(req *Message, reply *Message) error {
	reply.A += req.A
	fmt.Printf("==================TestHandler %v %v==================\n", req, reply)
	return errors.New("test error")
}

func (t *testActor) TestHandler2(req *Message) error {
	fmt.Printf("==================TestHandler2 %v ==================\n", req)
	return errors.New("test error")
}

func TestActor(t *testing.T) {
	var msg *Message
	_t := reflect.TypeOf(msg).Elem()
	name := _t.Name()
	_ = name
	system := actor.NewSystem()
	blueprint := actor.NewBlueprint()
	pid, _ := system.Spawn(blueprint, func() actor.IActor { return &testActor{} }, "init params")
	reply := &Message{A: 100}
	err := pid.Call("TestHandler", time.Second*1, &Message{A: 2}, reply)
	fmt.Printf("=======TestHandler respond:%v\n", err)
	pid.Send("TestHandler2", &Message{A: 2})

	fmt.Printf("========reply:%v========\n", reply)
	pid.Send("TestHandler", &Message{A: 1}, 2)
	time.Sleep(time.Second * 2)
	pid.Stop()
}
