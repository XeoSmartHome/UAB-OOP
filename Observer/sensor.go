package main

type Sensor struct {
	name      string
	value     int
	observers []Observer
}

func NewSensor(name string, value int) *Sensor {
	return &Sensor{
		name:      name,
		value:     value,
		observers: make([]Observer, 0),
	}
}

func (s *Sensor) setValue(value int) {
	s.value = value
	s.notifyAll()
}

func (s *Sensor) attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Sensor) detach(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Sensor) notifyAll() {
	for _, o := range s.observers {
		o.update(s.value)
	}
}
