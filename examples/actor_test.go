/**
 * @Author: dingQingHui
 * @Description:
 * @File: actor_test
 * @Version: 1.0.0
 * @Date: 2024/10/24 11:25
 */

package examples

import (
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

func (t *testActor) TestHandler(ctx actor.IContext, msg *Message) {
	fmt.Printf("==================TestHandler %v==================\n", msg)
}

func TestActor(t *testing.T) {
	var msg *Message
	_t := reflect.TypeOf(msg).Elem()
	name := _t.Name()
	_ = name
	system := actor.NewSystem()

	blueprint := actor.NewBlueprint()
	pid, _ := system.Spawn(blueprint, func() actor.IActor { return &testActor{} }, "init params")

	pid.Send("TestHandler", &Message{A: 1})
	time.Sleep(time.Second * 2)
	pid.Stop()

	fmt.Printf("================================================================\n")
}
