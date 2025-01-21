package main

import "fmt"

type Robot interface {
	getSpeed() float32
	setSpeed(speed float32)
	getAcceleration() float32
}

type VacuumCleanerRobot struct {
	speed float32
}

func NewVacuumCleanerRobot() *VacuumCleanerRobot {
	return &VacuumCleanerRobot{
		speed: 0,
	}
}

func (v *VacuumCleanerRobot) getSpeed() float32 {
	return v.speed
}

func (v *VacuumCleanerRobot) setSpeed(speed float32) {
	v.speed = speed
}

func (v *VacuumCleanerRobot) getAcceleration() float32 {
	return 1
}

type FlameThrowerRobot struct {
	speed float32
}

func NewFlameThrowerRobot() *FlameThrowerRobot {
	return &FlameThrowerRobot{
		speed: 0,
	}
}

func (f *FlameThrowerRobot) getSpeed() float32 {
	return f.speed
}

func (f *FlameThrowerRobot) setSpeed(speed float32) {
	f.speed = speed
}

func (f *FlameThrowerRobot) getAcceleration() float32 {
	return 5
}

type RemoteControl struct {
	robot Robot
}

func NewRemoteControl(robot Robot) *RemoteControl {
	return &RemoteControl{robot}
}

func (r RemoteControl) Accelerate() {
	r.robot.setSpeed(r.robot.getSpeed() + r.robot.getAcceleration())
}

func (r RemoteControl) Decelerate() {
	r.robot.setSpeed(r.robot.getSpeed() - r.robot.getAcceleration())
}

func main() {
	vacuumCleanerRobot := NewVacuumCleanerRobot()
	flameThrowerRobot := NewFlameThrowerRobot()

	vacuumCleanerRemoteControl := NewRemoteControl(vacuumCleanerRobot)

	vacuumCleanerRemoteControl.Accelerate()
	fmt.Printf("Vacuum Cleaner Robot speed: %f\n", vacuumCleanerRobot.getSpeed())

	vacuumCleanerRemoteControl.Accelerate()
	fmt.Printf("Vacuum Cleaner Robot speed: %f\n", vacuumCleanerRobot.getSpeed())

	vacuumCleanerRemoteControl.Decelerate()
	fmt.Printf("Vacuum Cleaner Robot speed: %f\n", vacuumCleanerRobot.getSpeed())

	flameThrowerRemoteControl := NewRemoteControl(flameThrowerRobot)

	flameThrowerRemoteControl.Accelerate()
	fmt.Printf("Flame Thrower Robot speed: %f\n", flameThrowerRobot.getSpeed())

	flameThrowerRemoteControl.Accelerate()
	fmt.Printf("Flame Thrower Robot speed: %f\n", flameThrowerRobot.getSpeed())
}
