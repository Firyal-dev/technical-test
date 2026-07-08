package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput reads and validates purchase amount and member status from user
// Returns purchase amount as integer and member status as boolean
// Loops until valid input is provided
func ReadInput() (int, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var (
			purchaseAmount int
			isMember       bool
		)

		fmt.Println("======= My Store =======")

		// Prompt for purchase amount
		fmt.Print("Belanja: Rp. ")
		purchaseInput, _ := reader.ReadString('\n')
		purchaseInput = strings.TrimSpace(purchaseInput)

		// Parse purchase amount to integer
		purchaseAmount, err := strconv.Atoi(purchaseInput)
		if err != nil {
			fmt.Println("input harus berupa angka")
			continue
		}

		// Validate purchase amount is positive
		if purchaseAmount <= 0 {
			fmt.Println("input harus lebih besar dari 0")
			continue
		}

		// Prompt for member status
		fmt.Print("Member (Ya/Tidak): ")
		memberInput, _ := reader.ReadString('\n')
		memberInput = strings.TrimSpace(strings.ToLower(memberInput))

		// Parse and validate member status input
		switch memberInput {
		case "ya", "y":
			isMember = true
		case "tidak", "t":
			isMember = false
		default:
			fmt.Println("input tidak valid, gunakan 'Ya' atau 'Tidak'")
			continue
		}

		return purchaseAmount, isMember
	}
}
