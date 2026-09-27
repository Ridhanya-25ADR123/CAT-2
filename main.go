package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"learngo/utility"
)

func main() {

	args := os.Args[1:]

	if len(args) < 2 {
		panic("please provide application name and operation")
	}

	appName := strings.Join(args[:len(args)-1], " ")
	operation := args[len(args)-1]

	utility.LoadMeters()

	fmt.Println("====================================")
	fmt.Println("Welcome to", appName)
	fmt.Println("====================================")

	var wg sync.WaitGroup

	switch operation {

	case "add":

		wg.Add(1)
		go registerMeter(&wg)
		wg.Wait()

	case "find":

		getMeter()

	case "record":

		recordNewReading()

	case "discard":

		removeMeter()

	default:

		fmt.Println("Invalid operation!")
		fmt.Println("Use: add, find, record, discard")
	}
}

func registerMeter(wg *sync.WaitGroup) {

	defer wg.Done()

	var meter utility.Meter

	fmt.Print("Enter Meter Number: ")
	fmt.Scanln(&meter.MeterNumber)

	fmt.Print("Enter Consumer Name: ")
	fmt.Scanln(&meter.ConsumerName)

	fmt.Print("Enter Previous Reading: ")
	fmt.Scanln(&meter.PreviousReading)

	fmt.Print("Enter Current Reading: ")
	fmt.Scanln(&meter.CurrentReading)

	err := utility.RegisterMeter(meter)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Meter added successfully!")
}

func getMeter() {

	var meterNum string

	fmt.Print("Enter Meter Number: ")
	fmt.Scanln(&meterNum)

	meter, err := utility.GetMeter(meterNum)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nMeter Found!")
	fmt.Println("Meter Number:", meter.MeterNumber)
	fmt.Println("Consumer Name:", meter.ConsumerName)
	fmt.Println("Previous Reading:", meter.PreviousReading)
	fmt.Println("Current Reading:", meter.CurrentReading)
}

func recordNewReading() {

	var meterNum string
	var reading float64

	fmt.Print("Enter Meter Number: ")
	fmt.Scanln(&meterNum)

	fmt.Print("Enter New Reading: ")
	fmt.Scanln(&reading)

	err := utility.RecordNewReading(meterNum, reading)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("New reading recorded successfully!")
}

func removeMeter() {

	var meterNum string

	fmt.Print("Enter Meter Number: ")
	fmt.Scanln(&meterNum)

	err := utility.RemoveMeter(meterNum)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Meter discarded successfully!")
}
