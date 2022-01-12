package testlib

import (
	"fmt"

	"github.com/ImperiumProject/imperium/types"
)

// CountWrapper encapsulates the function to fetch counter (CounterFunc) from state dynamically.
// CountWrapper is used to define actions and condition based on the counter.
type CountWrapper struct {
	CounterFunc func(*types.Event, *Context) (*Counter, bool)
}

// Count returns a CountWrapper where the CounterFunc fetches the counter based on the label
func Count(label string) *CountWrapper {
	return &CountWrapper{
		CounterFunc: func(_ *types.Event, c *Context) (*Counter, bool) {
			return c.Vars.GetCounter(label)
		},
	}
}

// CountF returns a CountWrapper where the label is also fetched based on the event and context
func CountF(labelFunc func(*types.Event, *Context) (string, bool)) *CountWrapper {
	return &CountWrapper{
		CounterFunc: func(e *types.Event, c *Context) (*Counter, bool) {
			label, ok := labelFunc(e, c)
			if !ok {
				return nil, false
			}
			return c.Vars.GetCounter(label)
		},
	}
}

// CountTo returns a CountWrapper where the counter label is `label` appended with message.To, if the event is a message send or receive
func CountTo(label string) *CountWrapper {
	return &CountWrapper{
		CounterFunc: func(e *types.Event, c *Context) (*Counter, bool) {
			message, ok := c.GetMessage(e)
			if !ok {
				return nil, false
			}
			counter, ok := c.Vars.GetCounter(fmt.Sprintf("%s_%s", label, message.To))
			if !ok {
				return nil, false
			}
			return counter, true
		},
	}
}

// SetWrapper encapsulates the mechanism to fetch a message set from the state.
// SetFunc should return a message set given the current event and context.
type SetWrapper struct {
	SetFunc func(*types.Event, *Context) (*types.MessageStore, bool)
}

// Set returns a SetWrapper where the set is fetched based on the label
func Set(label string) *SetWrapper {
	return &SetWrapper{
		SetFunc: func(e *types.Event, c *Context) (*types.MessageStore, bool) {
			return c.Vars.GetMessageSet(label)
		},
	}
}

// SetF returns a SetWrapper where the label is determined dynamically by the event and context
func SetF(labelFunc func(*types.Event, *Context) (string, bool)) *SetWrapper {
	return &SetWrapper{
		SetFunc: func(e *types.Event, c *Context) (*types.MessageStore, bool) {
			label, ok := labelFunc(e, c)
			if !ok {
				return nil, false
			}
			return c.Vars.GetMessageSet(label)
		},
	}
}

// Count returns a counter where the value is size of the message set
func (s *SetWrapper) Count() *CountWrapper {
	return &CountWrapper{
		CounterFunc: func(e *types.Event, c *Context) (*Counter, bool) {
			set, ok := s.SetFunc(e, c)
			if !ok {
				return nil, false
			}
			counter := NewCounter()
			counter.SetValue(set.Size())
			return counter, true
		},
	}
}
