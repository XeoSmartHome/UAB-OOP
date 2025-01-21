package main

import "fmt"

type Robot struct {
	name   string
	width  int
	length int
	height int
	speed  float32
}

type RobotBuilder interface {
	setName(name string) RobotBuilder
	setWidth(width int) RobotBuilder
	setLength(length int) RobotBuilder
	setHeight(height int) RobotBuilder
	setSpeed(speed float32) RobotBuilder
	build() Robot
}

type ConcreteRobotBuilder struct {
	robot Robot
}

func NewConcreteRobotBuilder() *ConcreteRobotBuilder {
	return &ConcreteRobotBuilder{}
}

func (c *ConcreteRobotBuilder) setName(name string) RobotBuilder {
	c.robot.name = name
	return c
}

func (c *ConcreteRobotBuilder) setWidth(width int) RobotBuilder {
	c.robot.width = width
	return c
}

func (c *ConcreteRobotBuilder) setLength(length int) RobotBuilder {
	c.robot.length = length
	return c
}

func (c *ConcreteRobotBuilder) setHeight(height int) RobotBuilder {
	c.robot.height = height
	return c
}

func (c *ConcreteRobotBuilder) setSpeed(speed float32) RobotBuilder {
	c.robot.speed = speed
	return c
}

func (c *ConcreteRobotBuilder) build() Robot {
	return c.robot
}

func main() {
	builder := NewConcreteRobotBuilder()

	robot := builder.setName("Robot").setWidth(23).setLength(30).setHeight(50).setSpeed(5.0).build()

	fmt.Printf("Robot name: %s\n", robot.name)
	fmt.Printf("Robot width: %d\n", robot.width)
	fmt.Printf("Robot length: %d\n", robot.length)
	fmt.Printf("Robot height: %d\n", robot.height)
	fmt.Printf("Robot speed: %f\n", robot.speed)
}
