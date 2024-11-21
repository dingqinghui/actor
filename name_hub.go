/**
 * @Author: dingQingHui
 * @Description:
 * @File: name_hub
 * @Version: 1.0.0
 * @Date: 2024/11/11 10:01
 */

package actor

type namedMsg struct {
	name    string
	process IProcess
}

func newNameHubActor(system ISystem) IProcess {
	blueprint := NewBlueprint()
	pid, _ := system.Spawn(blueprint, func() IActor { return &NameHub{} }, nil)
	return pid
}

type NameHub struct {
	BuiltinActor
	dict map[string]IProcess
}

func (n *NameHub) Init(ctx IContext, msg interface{}) {
	n.dict = make(map[string]IProcess)
}

func (n *NameHub) Named(ctx IContext, msg interface{}) {
	nm := msg.(*namedMsg)
	n.dict[nm.name] = nm.process
}

func (n *NameHub) Get(ctx IContext, msg interface{}) {
	name := msg.(string)
	v, _ := n.dict[name]
	ctx.EnvMessage().Sender().Send("", v)
}

func (n *NameHub) Del(ctx IContext, msg interface{}) {
	name := msg.(string)
	delete(n.dict, name)
}
