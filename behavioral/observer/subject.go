package observer

// an instance which publishes an event when anything happens

type Subject interface {
	Register(observer Observer)
	Deregister(bserver Observer)
	notifyAll()
}
