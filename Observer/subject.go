package main

type Subject interface {
	attach(observer Observer)
	detach(observer Observer)
	notifyAll()
}
