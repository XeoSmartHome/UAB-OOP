package main

import "fmt"

type Motor struct {
	subject Subject
	name    string
}

func NewMotor(name string) *Motor {
	return &Motor{
		name: name,
	}
}

func (m *Motor) getName() string {
	return m.name
}

func (m *Motor) update(value int) {
	fmt.Printf("Motor %s was updated with value %d\n", m.name, value)
}
