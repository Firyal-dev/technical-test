package main

import (
	"fmt"
	"github.com/Firyal-dev/mini-atm/services"
)

func main() {
	// Initialize account with starting balance of Rp. 1,000,000
	acc := services.Account{Saldo: 1000000}

	for {
		fmt.Println("======= MINI ATM =======")
		fmt.Println("Saldo saat ini : Rp.", acc.Saldo)
		fmt.Println("Pilih transaksi:")
		fmt.Println("1. Withdraw (tarik tunai)")
		fmt.Println("2. Deposit (setor tunai)")
		fmt.Println("3. Transfer ke rekening lain")
		fmt.Println("4. Keluar")

		choice := ReadInput()
		switch choice {
		case 1:
			// Handle withdraw transaction
			amount := ReadAmount()
			saldo, err := acc.Withdraw(amount)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("berhasil mengambil saldo. sisa saldo saat ini: Rp. %d\n", saldo)
			}
		case 2:
			// Handle deposit transaction
			amount := ReadAmount()
			saldo, err := acc.Deposit(amount)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("berhasil deposit saldo. jumlah saldo saat ini: Rp. %d\n", saldo)
			}
		case 3:
			// Handle transfer transaction
			amount := ReadAmount()
			destinationAccount := ReadAccountDestination()
			saldo, account, err := acc.Transfer(amount, destinationAccount)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("berhasil transfer saldo ke rekening %s. sisa saldo saat ini: Rp. %d\n", account, saldo)
			}
		case 4:
			// Exit application
			fmt.Println("terimakasih telah menggunakan Mini ATM. Sampai jumpa!")
			return
		default:
			fmt.Println("pilihan tidak valid")
		}
	}
}
