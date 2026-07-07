package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var (
		belanja  int
		isMember bool
		err      error
	)
	for {
		fmt.Println("======= My Store =======")
		
		// input nominal belanja
		fmt.Print("Belanja: ")
		inputBelanja, _ := reader.ReadString('\n')               // baca input user
		inputBelanja = strings.TrimSpace(inputBelanja)           // hapus spasi di awal dan akhir
		belanja, err = strconv.Atoi(inputBelanja)
		// validasi input belanja
		if err != nil {
			fmt.Println("Input harus berupa angka. Silakan coba lagi")
			continue
		}
		if belanja <= 0 {
			fmt.Println("Input harus lebih besar dari 0. Silakan coba lagi")
			continue
		}
		
		// input status member
		fmt.Print("Member: (Ya/Tidak): ")
		inputMember, _ := reader.ReadString('\n')    // baca input user
		inputMember = strings.TrimSpace(inputMember) // hapus spasi di awal dan akhir
		inputMember = strings.ToLower(inputMember)   // ubah input menjadi huruf kecil
		
		// validasi input member
		switch inputMember {
		case "ya", "y":
			isMember = true
		case "tidak", "t":
			isMember = false
		default:
			fmt.Println("Input tidak valid. Silakan coba lagi")
			continue
		}
		
		fmt.Println("------------------------")
		fmt.Println("Diskon Belanja:")
		fmt.Println("Diskon Member:")
		fmt.Println("Totak diskon:")
		
		fmt.Println("------------------------")
		fmt.Println("Total Belanja:")

		break
	}

	println("Belanja:", belanja)
	println("Member:", isMember)
}
