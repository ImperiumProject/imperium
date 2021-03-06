package dummy

import (
	"github.com/ImperiumProject/imperium/context"
	"github.com/ImperiumProject/imperium/log"
	"github.com/ImperiumProject/imperium/types"
)

type DummyStrategy struct {
	ctx *context.RootContext
	*types.BaseService
}

var _ types.Service = &DummyStrategy{}

func NewDummyStrategy(ctx *context.RootContext) *DummyStrategy {
	return &DummyStrategy{
		ctx:         ctx,
		BaseService: types.NewBaseService("dummyStrategy", ctx.Logger),
	}
}

func (d *DummyStrategy) Start() error {
	d.BaseService.StartRunning()
	return nil
}

func (d *DummyStrategy) Stop() error {
	d.BaseService.StopRunning()
	return nil
}

func (d *DummyStrategy) Step(event *types.Event) []*types.Message {
	if !event.IsMessageSend() {
		return []*types.Message{}
	}
	messageID, _ := event.MessageID()
	message, ok := d.ctx.MessageStore.Get(messageID)
	if ok {
		d.Logger.With(log.LogParams{
			"message_id": messageID,
			"from":       message.From,
			"to":         message.To,
		}).Debug("Delivering message")
		return []*types.Message{message}
	}
	return []*types.Message{}
}
