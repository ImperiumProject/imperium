package testlib

import "github.com/ImperiumProject/imperium/types"

// HandlerFunc type to define a conditional handler
// returns false in the second return value if the handler is not concerned about the event
type HandlerFunc func(*types.Event, *Context) ([]*types.Message, bool)

// HandlerCascade implements Handler
// Executes handlers in the specified order until the event is handled
// If no handler handles the event then the default handler is called
type HandlerCascade struct {
	Handlers       []HandlerFunc
	DefaultHandler HandlerFunc
}

// HandlerCascadeOption changes the parameters of the HandlerCascade
type HandlerCascadeOption func(*HandlerCascade)

// WithDefault changes the HandlerCascade default handler
func WithDefault(d HandlerFunc) HandlerCascadeOption {
	return func(hc *HandlerCascade) {
		hc.DefaultHandler = d
	}
}

// NewHandlerCascade creates a new cascade handler with the specified state machine and options
func NewHandlerCascade(opts ...HandlerCascadeOption) *HandlerCascade {
	h := &HandlerCascade{
		Handlers:       make([]HandlerFunc, 0),
		DefaultHandler: If(IsMessageSend()).Then(DeliverMessage()),
	}
	for _, o := range opts {
		o(h)
	}
	return h
}

func (c *HandlerCascade) addStateMachine(sm *StateMachine) {
	c.Handlers = append(
		[]HandlerFunc{NewStateMachineHandler(sm)},
		c.Handlers...,
	)
}

// AddHandler adds a handler to the cascade
func (c *HandlerCascade) AddHandler(h HandlerFunc) {
	c.Handlers = append(c.Handlers, h)
}

// HandleEvent implements Handler
func (c *HandlerCascade) HandleEvent(e *types.Event, ctx *Context) []*types.Message {
	for _, h := range c.Handlers {
		ret, ok := h(e, ctx)
		if ok {
			return ret
		}
	}
	ret, _ := c.DefaultHandler(e, ctx)
	return ret
}
