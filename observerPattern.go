package observerPattern

import (
	"sync"
)

type Observer interface {
	Notify(args ...interface{})
}

type ObservableSubject interface {
	NotifyObservers(args ...interface{})
	Register(listenerChannel chan []interface{})
	Unregister(listenerChannel chan []interface{})
}

type ConcreteObservable struct {
	observerMutex *sync.Mutex
	observers     [](chan []interface{})
}

func NewConcreteObservable() *ConcreteObservable {
	return &ConcreteObservable{
		observerMutex: new(sync.Mutex),
	}
}

func (c *ConcreteObservable) NotifyObservers(args ...interface{}) {
	for _, registeredListenerChannel := range c.observers {
		registeredListenerChannel <- args
	}
}

func (c *ConcreteObservable) Register(listenerChannel chan []interface{}) {
	c.observerMutex.Lock()
	defer c.observerMutex.Unlock()
	c.observers = append(c.observers, listenerChannel)
}

func (c *ConcreteObservable) Unregister(listenerChannel chan []interface{}) {
	c.observerMutex.Lock()
	defer c.observerMutex.Unlock()

	for i, registeredListenerChannel := range c.observers {
		if registeredListenerChannel == listenerChannel {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			return
		}
	}
}
