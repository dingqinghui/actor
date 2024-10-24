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

func (b *blueprint) Spawn(system ISystem, params ...interface{}) (IProcess, error) {
	mb := b.getMailBox()
	process := NewBaseProcess(mb)
	context := NewBaseActorContext(b.getReceiver(), system, process)
	mb.RegisterHandlers(context, b.getDispatcher())
	// notify actor start
	if err := process.Send(&StartedMessage{Params: params}); err != nil {
		return nil, err
	}
	return process, nil
}
