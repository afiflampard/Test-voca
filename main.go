package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Car struct {
	Number string
	Status string
	Hour   int
	Charge int
}
type Parking struct {
	Capacity int
	Slots    []*Car
}

func NewParking(capacity int) *Parking {
	return &Parking{
		Capacity: capacity,
		Slots:    make([]*Car, capacity),
	}
}

func (p *Parking) IsFull() bool {
	for _, slots := range p.Slots {
		if slots == nil || slots.Status == "leave" {
			return false
		}
	}
	return true
}

func (p *Parking) ParkCar(carNumber string) {
	for i, slot := range p.Slots {
		if slot != nil {
			if slot.Number == carNumber && slot.Status != "leave" {
				fmt.Printf("Number already registered \n")
				return
			}

		}
		if slot == nil || slot.Status == "leave" {
			p.Slots[i] = &Car{Number: carNumber, Status: "park", Charge: 10}
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}

}

func (p *Parking) Leave(carNumber string, hours int) {
	if p != nil {
		for i, car := range p.Slots {
			if car != nil && (strings.ToLower(car.Number) == strings.ToLower(carNumber)) {
				p.Slots[i] = &Car{Number: carNumber, Status: "leave", Hour: hours, Charge: car.Charge}
				if hours > 2 {
					p.Slots[i].Charge += (hours - 2) * 10
				}
				fmt.Printf("Registration number %s with Slot Number %d is free with Charge $%d\n",
					carNumber, i+1, p.Slots[i].Charge)
				return
			}
		}
		fmt.Printf("Car With Number %s Not Found \n", carNumber)
	}
}

func (p *Parking) Status() {
	for i, car := range p.Slots {
		if car != nil {
			fmt.Printf("%d %s %s %d hours \n", i+1, car.Status, car.Number, car.Hour)
		}
	}
}

var menus = []string{"create_parking_lot", "park", "leave", "status", "exit"}

func main() {
	pilih := true
	var parking *Parking
	for pilih {
		var pilihan string
		fmt.Println("Menu : ")
		for _, menu := range menus {
			fmt.Println(menu)
		}
		fmt.Println("Pilih Menu : ")
		fmt.Scan(&pilihan)
		switch strings.ToLower(pilihan) {
		case strings.ToLower("create_parking_lot"):
			var t string
			fmt.Print("Input capacity : ")
			fmt.Scan(&t)
			capacity, _ := strconv.Atoi(t)
			parking = NewParking(capacity)
			fmt.Printf("parking lot created with capacity %d \n", parking.Capacity)
		case strings.ToLower("park"):
			if parking != nil {
				if !parking.IsFull() {
					var number string
					fmt.Print("Input Number of Car : ")
					fmt.Scan(&number)
					parking.ParkCar(number)
				} else {
					fmt.Println("Sorry, parking lot is full")
				}

			} else {
				fmt.Println("Please create parking lot first")
			}
		case strings.ToLower("leave"):
			var number, hourTemp string
			fmt.Printf("Input car Number : \n")
			fmt.Scan(&number)
			fmt.Print("Input hour : \n")
			fmt.Scan(&hourTemp)
			hour, err := strconv.Atoi(hourTemp)
			if err != nil {
				fmt.Println("Please insert with number type")
			}
			parking.Leave(number, hour)
		case strings.ToLower("status"):
			if parking != nil {
				parking.Status()
			}
		case strings.ToLower("exit"):
			pilih = false
		}

	}
}
