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
	"testing"
	"time"
)

type testActor struct {
}

func (t testActor) Init(ctx actor.IContext, params ...interface{}) {
	fmt.Printf("Init\n")
	ctx.TimerHub().AddTimer(time.Second*2, func() {
		t.OnTimer(ctx)
	})
}

func (t testActor) OnTimer(ctx actor.IContext) {
	fmt.Printf("OnTimer %d\n", time.Now().Unix())
	ctx.TimerHub().AddTimer(time.Second*2, func() {
		t.OnTimer(ctx)
	})
}

func (t testActor) Receive(context actor.IContext) error {
	fmt.Printf("Receive:%v\n", context.Message())

	switch context.Message().(type) {
	case actor.IEnvelope:
		env := context.Message().(actor.IEnvelope)
		env.Sender().Send("11111111111111111")
	}

	return nil
}

func (t testActor) Panic(context actor.IContext) {
	fmt.Printf("Stop\n")
}

func (t testActor) Stop(context actor.IContext) {
	fmt.Printf("Stop\n")
}

var _ actor.IReceiver = &testActor{}

func TestActor(t *testing.T) {
	system := actor.NewSystem()
	blueprint := actor.NewBlueprint(actor.WithReceiver(&testActor{}))
	pid, _ := system.Spawn(blueprint)
	pid.Send("1")
	pid.Send("2")
	pid.Send("3")
	pid.Stop()
	time.Sleep(time.Second * 10)
}

func TestActorFuture(t *testing.T) {
	system := actor.NewSystem()
	blueprint := actor.NewBlueprint(actor.WithReceiver(&testActor{}))
	pid, _ := system.Spawn(blueprint)
	fut, _ := pid.Call("1", time.Second*10)
	result, isTimeout := fut.Wait()
	fmt.Printf("fut:%v %v\n", result, isTimeout)
	pid.Send("2")
	pid.Send("3")
	time.Sleep(time.Second * 10)
}
