package main

import "fmt"

// IRobot este o interfață pentru roboți
type IRobot interface {
	// getName returnează numele robotului
	getName() string
	// getDimensions returnează dimensiunile robotului
	getDimensions() string
}

// Robot este o structură de bază pentru un robot
type Robot struct {
	name   string
	width  int
	length int
	height int
}

// getName returnează numele robotului
func (r Robot) getName() string {
	return r.name
}

// getDimensions returnează dimensiunile robotului
func (r Robot) getDimensions() string {
	return fmt.Sprintf("Width: %dcm, Length: %dcm, Height: %dcm", r.width, r.length, r.height)
}

// FlameThrowerRobot este o structură pentru un robot cu aruncător de flăcări
type FlameThrowerRobot struct {
	Robot
}

// getName returnează un mesaj specific pentru robotul cu aruncător de flăcări
func (f FlameThrowerRobot) getName() string {
	return "Flame Thrower Robot does not have a name, he throws flames"
}

// VacuumCleanerRobot este o structură pentru un robot aspirator
type VacuumCleanerRobot struct {
	Robot
}
