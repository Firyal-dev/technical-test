package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/Firyal-dev/parking-fee/utils"
)

func ReadInput() (int, int, int, int) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("======= Mall Parking =======")

		fmt.Print("Jam Masuk: ")
		entryTimeInput, _ := reader.ReadString('\n')
		entryHour, entryMinute, err := utils.ParseTime(entryTimeInput)
		if err != nil {
			fmt.Println(err)
			continue
		}

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