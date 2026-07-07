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
			belanja  int
			isMember bool
		)

		fmt.Println("======= My Store =======")

		fmt.Print("Belanja: Rp. ")
		inputBelanja, _ := reader.ReadString('\n') // baca input dari user
		inputBelanja = strings.TrimSpace(inputBelanja) // hapus spasi pada input

		belanja, err := strconv.Atoi(inputBelanja) // parse input menjadi integer
		if err != nil {
			fmt.Println("input harus berupa angka.")
			continue
		}

		if belanja <= 0 {
			fmt.Println("input harus lebih besar dari 0.")
			continue
		}

		fmt.Print("Member (Ya/Tidak): ")
		inputMember, _ := reader.ReadString('\n') // baca input dari user
		inputMember = strings.TrimSpace(strings.ToLower(inputMember)) // hapus spasi dan ubah menjadi lowercase

		switch inputMember {
		case "ya", "y":
			isMember = true
		case "tidak", "t":
			isMember = false
		default:
			fmt.Println("input tidak valid, gunakan 'Ya' atau 'Tidak'")
			continue
		}

		return belanja, isMember
	}
}
