package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Firyal-dev/parking-fee/utils"
)

// ReadInput reads and validates entry and exit times from user
// Returns entry hour, entry minute, exit hour, exit minute as integers
// Loops until valid time inputs are provided
func ReadInput() (int, int, int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display parking system banner
		fmt.Println("======= Mall Parking =======")

		// Prompt and parse entry time
		fmt.Print("Jam Masuk: ")
		entryTimeInput, _ := reader.ReadString('\n')
		entryHour, entryMinute, err := utils.ParseTime(entryTimeInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Prompt and parse exit time
		fmt.Print("Jam Keluar: ")
		exitTimeInput, _ := reader.ReadString('\n')
		exitHour, exitMinute, err := utils.ParseTime(exitTimeInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

		return entryHour, entryMinute, exitHour, exitMinute
	}
}