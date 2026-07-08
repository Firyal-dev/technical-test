package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput() (int, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var (
			purchaseAmount int
			isMember       bool
		)

		fmt.Println("======= My Store =======")

		fmt.Print("Belanja: Rp. ")
		purchaseInput, _ := reader.ReadString('\n')
		purchaseInput = strings.TrimSpace(purchaseInput)

		purchaseAmount, err := strconv.Atoi(purchaseInput)
		if err != nil {
			fmt.Println("input harus berupa angka.")
			continue
		}

		if purchaseAmount <= 0 {
			fmt.Println("input harus lebih besar dari 0.")
			continue
		}

		fmt.Print("Member (Ya/Tidak): ")
		memberInput, _ := reader.ReadString('\n')
		memberInput = strings.TrimSpace(strings.ToLower(memberInput))

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
