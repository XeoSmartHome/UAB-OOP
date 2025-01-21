package main

import (
	"fmt"
	"log"
)

func main() {
	// Obține o fabrică de roboți pentru aspiratoare
	vacuumCleanerRobotFactory, err := getRobotFactory("vacuum")
	if err != nil {
		log.Fatal(err)
	}

	// Creează un robot aspirator folosind fabrica
	vacuumCleanerRobot := vacuumCleanerRobotFactory.makeRobot()
	fmt.Println(vacuumCleanerRobot.getName())
	fmt.Println(vacuumCleanerRobot.getDimensions())

	// Obține o fabrică de roboți cu aruncător de flăcări
	flameThrowerRobotFactory, err := getRobotFactory("flame")
	if err != nil {
		log.Fatal(err)
	}

	// Creează un robot cu aruncător de flăcări folosind fabrica
	flameThrowerRobot := flameThrowerRobotFactory.makeRobot()
	fmt.Println(flameThrowerRobot.getName())
	fmt.Println(flameThrowerRobot.getDimensions())

	// Încearcă să obțină o fabrică de roboți pentru gătit (nu există)
	_, err = getRobotFactory("cooking")
	if err != nil {
		// Afișează eroarea dacă fabrica nu este găsită
		fmt.Println(err)
	}
}
