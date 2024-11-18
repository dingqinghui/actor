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
	actor.BuiltinActor
}

func (t *testActor) Init(ctx actor.IContext, msg interface{}) {
	fmt.Printf("==================Init %v==================\n", msg)
	ctx.AddTimer(time.Second*1, "OnTimer")
}

func (t *testActor) TestHandler(ctx actor.IContext, msg interface{}) {
	fmt.Printf("==================TestHandler %v==================\n", msg)
}
func (t *testActor) OnTimer(ctx actor.IContext, msg interface{}) {
	fmt.Printf("==================OnTimer:%v==================\n", time.Now().Unix())
	ctx.AddTimer(time.Second*1, "OnTimer")
}

func (t *testActor) Stop(ctx actor.IContext, msg interface{}) {
	fmt.Printf("==================Stop:%v==================\n", time.Now().Unix())
	ctx.AddTimer(time.Second*1, "OnTimer")
}

func TestActor(t *testing.T) {
	system := actor.NewSystem()

	blueprint := actor.NewBlueprint()
	pid, _ := system.Spawn(blueprint, func() actor.IActor { return &testActor{} }, "init params")
	system.Named("1", pid)
	_pid, _ := system.GetProcessByName("1")
	_ = _pid
	system.DelName("1")

	pid.Send("TestHandler", "test msg")
	time.Sleep(time.Second * 2)
	pid.Stop()

	fmt.Printf("================================================================\n")
}
