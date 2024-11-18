package main

import (
	"fmt"
)

type Vehicle interface {
	CalculateTravelTime(distance float64) float64
	GetType() string
}

type Car struct {
	Speed    float64
	Type     string
	FuelType string
}

type Motorcycle struct {
	Speed float64
	Type  string
}

func (c Car) CalculateTravelTime(distance float64) float64 {
	return distance / c.Speed
}

func (m Motorcycle) CalculateTravelTime(distance float64) float64 {
	return distance / m.Speed
}

func (c Car) GetType() string {
	return c.Type
}

func (m Motorcycle) GetType() string {
	return m.Type
}

func EstimateTravelTime(vehicles []Vehicle, distance float64) map[string]float64 {
	transports := make(map[string]float64)
	for _, el := range vehicles {
		transports[fmt.Sprintf("%T", el)] = el.CalculateTravelTime(distance)
	}
	return transports
}

func main() {
	car := Car{Type: "Седан", Speed: 60.0, FuelType: "Бензин"}
	motorcycle := Motorcycle{Type: "Мотоцикл", Speed: 80.0}

	vehicles := []Vehicle{car, motorcycle}

	distance := 200.0

	travelTimes := EstimateTravelTime(vehicles, distance)

	expectedCarTime := distance / car.Speed

	if travelTimes["main.Car"] != expectedCarTime {
		fmt.Printf("Ожидается время для автомобиля %.2f часа, получено %.2f", expectedCarTime, travelTimes["main.Car"])
	}

	expectedMotorcycleTime := distance / motorcycle.Speed
	if travelTimes["main.Motorcycle"] != expectedMotorcycleTime {
		fmt.Printf("Ожидается время для мотоцикла %.2f часа, получено %.2f", expectedMotorcycleTime, travelTimes["main.Motorcycle"])
	}
}
