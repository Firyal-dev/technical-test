package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput reads and validates user's menu choice (1-4)
func ReadInput() int {
	reader := bufio.NewReader(os.Stdin)

	var (
		pilihan int
		err     error
	)
	for {
		fmt.Print("Pilihan: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		// Parse input to integer
		pilihan, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("input harus berupa angka")
			continue
		}
		
		// Validate menu choice is within valid range
		if pilihan < 1 || pilihan > 4 {
			fmt.Println("pilihan tidak valid. Silakan pilih antara 1 hingga 4")
			continue
		}

		return pilihan
	}
}

// ReadAmount reads and validates transaction amount from user input
func ReadAmount() int {
	reader := bufio.NewReader(os.Stdin)

	var (
		amount int
		err    error
	)
	for {
		fmt.Print("Jumlah: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		// Parse input to integer
		amount, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("input harus berupa angka")
			continue
		}
		
		// Validate amount is positive
		if amount <= 0 {
			fmt.Println("nominal harus diatas 0")
			continue
		}

		return amount
	}
}

// ReadAccountDestination reads and validates destination account number for transfer
func ReadAccountDestination() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Nomor rekening tujuan: ")

		input, _ := reader.ReadString('\n')
		account := strings.TrimSpace(input)
		
		// Validate account number is not empty
		if account == "" {
			fmt.Println("rekening tujuan tidak boleh kosong")
			continue
		}

		return account
	}
}
