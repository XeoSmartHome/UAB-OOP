package main

import "fmt"

// IRobotFactory este o interfață pentru fabricile de roboți
type IRobotFactory interface {
	// makeRobot este metoda pentru crearea unui robot
	makeRobot() IRobot
}

// VacuumCleanerRobotFactory este o structură pentru fabrica de roboți aspirator
type VacuumCleanerRobotFactory struct{}

// makeRobot creează un robot aspirator
func (v VacuumCleanerRobotFactory) makeRobot() IRobot {
	return VacuumCleanerRobot{
		Robot: Robot{
			name:   "Vacuum Cleaner Robot",
			width:  50,
			length: 50,
			height: 20,
		},
	}
}

// FlameThrowerRobotFactory este o structură pentru fabrica de roboți cu aruncător de flăcări
type FlameThrowerRobotFactory struct{}

// makeRobot creează un robot cu aruncător de flăcări
func (f FlameThrowerRobotFactory) makeRobot() IRobot {
	return FlameThrowerRobot{
		Robot: Robot{
			name:   "Flame Thrower Robot",
			width:  100,
			length: 100,
			height: 100,
		},
	}
}

// getRobotFactory returnează o fabrică de roboți în funcție de tipul specificat
func getRobotFactory(robotType string) (IRobotFactory, error) {
	if robotType == "vacuum" {
		return VacuumCleanerRobotFactory{}, nil
	}

	if robotType == "flame" {
		return FlameThrowerRobotFactory{}, nil
	}

	// Returnează o eroare dacă tipul de fabrică nu este găsit
	return nil, fmt.Errorf("robot factory for %s not found", robotType)
}
