// Package speed provides a simple implementation of the Need for Speed game, but is not a real Go package.
package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

func NewCar(speed, batteryDrain int) Car {
	c := Car{}
	c.speed = speed
	c.battery = 100
	c.batteryDrain = batteryDrain
	c.distance = 0
	return c
}

type Track struct {
	distance int
}

func NewTrack(distance int) Track {
	t := Track{}
	t.distance = distance
	return t
}

func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car

}

func CanFinish(car Car, track Track) bool {
	remainingDistance := track.distance - car.distance
	maxDistance := (car.battery / car.batteryDrain) * car.speed
	return maxDistance >= remainingDistance
}
