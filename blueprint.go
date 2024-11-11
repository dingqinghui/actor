/**
 * @Author: dingQingHui
 * @Description:
 * @File: BlueprintOptionsFunc
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:16
 */

package actor

type BlueprintOptionsFunc func(b *blueprint)

func WithReceiver(receiver IReceiver) BlueprintOptionsFunc {
	return func(b *blueprint) {
		b.receiver = receiver
	}
}

func WithDispatcher(dispatcher IDispatcher) BlueprintOptionsFunc {
	return func(b *blueprint) {
		b.dispatcher = dispatcher
	}
}

func WithMailBox(mailbox IMailbox) BlueprintOptionsFunc {
	return func(b *blueprint) {
		b.mailbox = mailbox
	}
}

func WithRouter(routers IRoutes) BlueprintOptionsFunc {
	return func(b *blueprint) {
		b.routers = routers
	}
}

func NewBlueprint(opts ...BlueprintOptionsFunc) IBlueprint {
	b := new(blueprint)
	for _, opt := range opts {
		opt(b)
	}
	return b
}

type blueprint struct {
	receiver   IReceiver
	dispatcher IDispatcher
	mailbox    IMailbox
	routers    IRoutes
}

func (b *blueprint) getDispatcher() IDispatcher {
	if b.dispatcher == nil {
		b.dispatcher = NewDefaultDispatcher(50)
	}
	return b.dispatcher
}

func (b *blueprint) getReceiver() IReceiver {
	if b.receiver == nil {
		b.receiver = NewDefaultReceiver()
	}
	return b.receiver
}

func (b *blueprint) getMailBox() IMailbox {
	if b.mailbox == nil {
		b.mailbox = NewMailbox()
	}
	return b.mailbox
}

func (b *blueprint) getRouter() IRoutes {
	if b.routers == nil {
		b.routers = NewBuiltinRoutes()
	}
	return b.routers
}

func (b *blueprint) Spawn(system ISystem, producer Producer, initParams ...interface{}) (IProcess, error) {
	mb := b.getMailBox()
	process := NewBaseProcess(mb)

	context := NewBaseActorContext()
	context.routers = b.getRouter()
	context.actor = producer()
	context.system = system
	context.process = process
	context.initParams = initParams
	context.initialize()
	mb.RegisterHandlers(context, b.getDispatcher())
	// notify actor start
	if err := process.Send(startMessage); err != nil {
		return nil, err
	}
	return process, nil
}
