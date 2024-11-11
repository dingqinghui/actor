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

func (t testActor) initialize() {
	println("================testActor Initialize=================")
}

func initRoute() actor.IRoutes {
	routers := actor.NewBuiltinRoutes()
	routers.Add(actor.StartMessageId, func(ctx actor.IContext, env actor.IEnvelope) {
		println("================actor start=================")
		t := ctx.Actor().(*testActor)
		t.initialize()
	})
	routers.Add(actor.StopMessageId, func(ctx actor.IContext, env actor.IEnvelope) {
		println("================actor stop=================")
	})

	routers.Add(1, func(ctx actor.IContext, env actor.IEnvelope) {
		fmt.Printf("================actor msg:%v=================\n", env.Body())
	})

	routers.Add(2, func(ctx actor.IContext, env actor.IEnvelope) {
		fmt.Printf("================actor msg:%v=================\n", env.Body())
		env.Sender().Send(actor.NewMessage(1, "1111"))
	})
	return routers
}

func TestActor(t *testing.T) {
	system := actor.NewSystem()

	blueprint := actor.NewBlueprint(actor.WithReceiver(&testActor{}), actor.WithRouter(initRoute()))
	pid, _ := system.Spawn(blueprint, func() actor.IActor { return &testActor{} })
	pid.Send(actor.NewMessage(1, "1111"))

	fut, _ := pid.Call(actor.NewMessage(2, "1111"), time.Second)
	result, isTimeout := fut.Wait()
	fmt.Printf("fut:%v %v\n", result, isTimeout)

	system.Named("testName", pid)
	p, _ := system.GetProcessByName("testName")

	system.DelName("testName")
	p, _ = system.GetProcessByName("testName")
	_ = p
	pid.Stop()
	time.Sleep(time.Second * 10)
}
