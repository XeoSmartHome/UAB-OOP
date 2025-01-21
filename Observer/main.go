package main

func main() {
	// Create a sensor (Subject)
	sensor := NewSensor("Sensor", 100)

	// Create 2 motors (Observers)
	leftMotor := NewMotor("Left Motor")

	rightMotor := NewMotor("Right Motor")

	// Attach the motors to the sensor
	sensor.attach(leftMotor)
	sensor.attach(rightMotor)

	// Set the new value for the sensor
	sensor.setValue(10)

	// Detach the motors from the sensor
	sensor.detach(leftMotor)

	// Set the new value for the sensor
	sensor.setValue(20)

	// Detach the last motor from the sensor
	sensor.detach(rightMotor)

	// Automatic detachment using defer
	func() {
		motor := NewMotor("Motor")

		sensor.attach(motor)

		defer sensor.detach(motor)

		sensor.setValue(30)
	}()

	// Set the new value for the sensor, no motors attached
	sensor.setValue(40)
}
