package main

import (
	"fmt"
	"github.com/Firyal-dev/parking-fee/services"
)

func main() {
	entryHour, entryMinute, exitHour, exitMinute := ReadInput()

	durationHours, durationMinutes, err := services.CalculateParkingDuration(entryHour, entryMinute, exitHour, exitMinute)
	if err != nil {
		fmt.Println(err)
		return
	}

	parkingFee := services.CalculateParkingFee(durationHours, durationMinutes)

	fmt.Println("------------------------")
	fmt.Printf("Durasi Parkir: %d jam %d menit\n", durationHours, durationMinutes)
	fmt.Println("------------------------")
	fmt.Printf("Total Bayar: Rp. %d\n", parkingFee)
}