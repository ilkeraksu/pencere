package pencere

import (
	"context"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type EventData struct {
	Subject string
	Event   interface{}
}

type eventHandler struct {
	Subject  string
	CallBack EventCallback
}

type EventCallback func(ctx context.Context, event interface{}) error

func NewEventBus() *EventBus {
	bus := &EventBus{
		subjectHandlers: make(map[string][]*eventHandler),
		ch:              make(chan EventData, 100),
	}

	return bus
}

type EventBus struct {
	lock            sync.Mutex
	handlers        []*eventHandler
	subjectHandlers map[string][]*eventHandler
	ch              chan EventData
}

func (this *EventBus) On(ctx context.Context, subject string, callback EventCallback) {
	this.lock.Lock()
	defer this.lock.Unlock()
	h := &eventHandler{
		Subject:  subject,
		CallBack: callback,
	}
	this.handlers = append(this.handlers, h)
}

func (this *EventBus) getSubjectHandlers(subject string) []*eventHandler {
	if handlers, ok := this.subjectHandlers[subject]; ok {
		return handlers
	}

	this.lock.Lock()
	handlers := this.calculateSubjectHandlers(subject)
	this.subjectHandlers[subject] = handlers
	this.lock.Unlock()
	return handlers
}

func (this *EventBus) calculateSubjectHandlers(subject string) []*eventHandler {
	handlers := make([]*eventHandler, 0)
	for _, handler := range this.handlers {
		if this.isSubjectMatch(subject, handler.Subject) {
			handlers = append(handlers, handler)
		}
	}

	return handlers
}

func (this *EventBus) PushEvent(ctx context.Context, subject string, event interface{}) error {
	handlers := this.getSubjectHandlers(subject)
	for _, handler := range handlers {
		err := handler.CallBack(ctx, event)
		if err != nil {
			return errors.Wrapf(err, "could not call callback for sucject:%v", subject)
		}
	}
	return nil
}

func (this *EventBus) isSubjectMatch(subject string, filter string) bool {
	subjects := strings.Split(subject, ".")
	filters := strings.Split(filter, ".")

	// don't support empty subject
	for _, s := range subjects {
		if s == "" {
			return false
		}
	}

	for _, f := range filters {
		// don't support empty filter
		if f == "" {
			return false
		}
	}

	// we expect filters less or equal length than subject parts
	if len(filters) > len(subjects) {
		return false
	}

	for i, s := range subjects {
		if i >= len(filters) {
			//filter is short and does not matched
			return false
		}

		f := filters[i]

		// if f is '>'  then match, we can stop checking
		if f == ">" {
			return true
		}

		// if f is '*' match this level, continue checking
		if f == "*" {
			continue
		}

		//is f!=s then it is not matched
		if f != s {
			return false
		}
	}

	// we checked all levels so it is match
	return true
}
