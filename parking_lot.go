package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Slot represents a parking slot
type Slot struct {
	occupied           bool
	registrationNumber string
	startTime          time.Time
}

// ParkingLot represents the parking lot
type ParkingLot struct {
	slots []Slot
}

// CreateParkingLot initializes the parking lot with n slots
func (pl *ParkingLot) CreateParkingLot(n int) {
	pl.slots = make([]Slot, n)
	fmt.Printf("Created a parking lot with %d slots\n", n)
}

// Park assigns the nearest available slot to a car
func (pl *ParkingLot) Park(registrationNumber string) {
	for i := range pl.slots {
		if !pl.slots[i].occupied {
			pl.slots[i] = Slot{
				occupied:           true,
				registrationNumber: registrationNumber,
				startTime:          time.Now(),
			}
			fmt.Printf("Allocated slot number: %d\n", i+1)
			return
		}
	}
	fmt.Println("Sorry, parking lot is full")
}

// Leave releases a slot and calculates the charge
func (pl *ParkingLot) Leave(slotNumber int) {
	if slotNumber <= 0 || slotNumber > len(pl.slots) || !pl.slots[slotNumber-1].occupied {
		fmt.Println("Invalid slot number")
		return
	}

	slot := &pl.slots[slotNumber-1]
	duration := time.Since(slot.startTime).Hours()
	charge := 10
	if duration > 2 {
		charge += int(duration-2) * 10
	}

	fmt.Printf("Slot number %d is free. Charge: $%d\n", slotNumber, charge)
	*slot = Slot{}
}

// LeaveByRegistrationNumber releases a slot by car's registration number and calculates the charge
func (pl *ParkingLot) LeaveByRegistrationNumber(registrationNumber string, hoursParked int) {
	if hoursParked <= 0 {
		fmt.Println("Invalid number of hours. Must be greater than 0.")
		return
	}

	for i := range pl.slots {
		slot := &pl.slots[i]
		if slot.occupied && slot.registrationNumber == registrationNumber {
			// Calculate charges
			charge := 10 // Base charge for the first 2 hours
			if hoursParked > 2 {
				charge += (hoursParked - 2) * 10 // $10 for each additional hour
			}

			// Free the slot
			fmt.Printf("Registration number %s with slot Number %d is free with Charge $%d\n", registrationNumber, i+1, charge)
			*slot = Slot{} // Reset the slot
			return
		}
	}
	fmt.Printf("Registration number %s not found\n", registrationNumber)
}

// Status displays the current parking lot status
func (pl *ParkingLot) Status() {
	fmt.Println("Slot No.    Registration No")
	for i, slot := range pl.slots {
		if slot.occupied {
			fmt.Printf("%-10d %-17s\n", i+1, slot.registrationNumber)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file with commands")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	pl := ParkingLot{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		command := strings.ToLower(parts[0])
		switch command {
		case "create_parking_lot":
			if len(parts) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			n, _ := strconv.Atoi(parts[1])
			pl.CreateParkingLot(n)
		case "park":
			if len(parts) != 2 {
				fmt.Println("Invalid command")
				continue
			}
			pl.Park(parts[1])
		case "leave":
			if len(parts) != 3 {
				fmt.Println("Invalid command")
				continue
			}
			hoursParked, err := strconv.Atoi(parts[2])
			if err != nil || hoursParked <= 0 {
				fmt.Println("Invalid number of hours")
				continue
			}
			pl.LeaveByRegistrationNumber(parts[1], hoursParked)
		case "status":
			pl.Status()
		default:
			fmt.Println("Unknown command:", command)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
