/**
 * @Author: dingQingHui
 * @Description:
 * @File: BlueprintOptionsFunc
 * @Version: 1.0.0
 * @Date: 2024/10/24 10:16
 */

package actor

import "sync"

type BlueprintOptionsFunc func(b *blueprint)

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

func NewBlueprint(opts ...BlueprintOptionsFunc) IBlueprint {
	b := new(blueprint)
	for _, opt := range opts {
		opt(b)
	}
	return b
}

type blueprint struct {
	dispatcher  IDispatcher
	mailbox     IMailbox
	onceHandler sync.Once
	h           *handlers
}

func (b *blueprint) getDispatcher() IDispatcher {
	if b.dispatcher == nil {
		b.dispatcher = NewDefaultDispatcher(50)
	}
	return b.dispatcher
}

func (b *blueprint) getMailBox() IMailbox {
	if b.mailbox == nil {
		b.mailbox = NewMailbox()
	}
	return b.mailbox
}

func (b *blueprint) getHandlers(actor IActor) *handlers {
	b.onceHandler.Do(func() {
		b.h = newHandlers(actor)
	})
	return b.h
}

func (b *blueprint) Spawn(system ISystem, producer Producer, params interface{}) (IProcess, error) {
	mb := b.getMailBox()
	process := NewBaseProcess(mb)
	actor := producer()
	h := b.getHandlers(actor)

	context := NewBaseActorContext()
	context.actor = actor
	context.system = system
	context.process = process
	context.handlers = h
	mb.RegisterHandlers(context, b.getDispatcher())
	// notify actor start
	if err := process.Send(InitFuncName, params); err != nil {
		return nil, err
	}
	return process, nil
}
