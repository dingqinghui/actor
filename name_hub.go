/**
 * @Author: dingQingHui
 * @Description:
 * @File: name_hub
 * @Version: 1.0.0
 * @Date: 2024/11/11 10:01
 */

package actor

const (
	namedMsgId = iota
	getNameMsgId
	delNameMsgId
)

type namedBody struct {
	name    string
	process IProcess
}

func initNamedHubRoute() IRoutes {
	routers := NewBuiltinRoutes()
	routers.Add(StartMessageId, func(ctx IContext, env IEnvelope) {
		ctx.Actor().(*nameHub).init()
	})

	routers.Add(namedMsgId, func(ctx IContext, env IEnvelope) {
		body := env.Body().(*namedBody)
		ctx.Actor().(*nameHub).named(body.name, body.process)
	})

	routers.Add(getNameMsgId, func(ctx IContext, env IEnvelope) {
		body := env.Body().(string)
		v := ctx.Actor().(*nameHub).get(body)
		env.Sender().Send(NewMessage(getNameMsgId, v))
	})

	routers.Add(delNameMsgId, func(ctx IContext, env IEnvelope) {
		body := env.Body().(string)
		ctx.Actor().(*nameHub).del(body)
	})
	return routers
}

func newNameHubActor(system ISystem) IProcess {
	routes := initNamedHubRoute()
	blueprint := NewBlueprint(WithReceiver(&nameHub{}), WithRouter(routes))
	pid, _ := system.Spawn(blueprint, func() IActor { return &nameHub{} })
	return pid
}

type nameHub struct {
	dict map[string]IProcess
}

func (n *nameHub) init() {
	n.dict = make(map[string]IProcess)
}

func (n *nameHub) named(name string, process IProcess) {
	n.dict[name] = process
}

func (n *nameHub) get(name string) IProcess {
	v, _ := n.dict[name]
	return v
}

func (n *nameHub) del(name string) {
	delete(n.dict, name)
}
