package main

import (
	"fmt"
	"github.com/Firyal-dev/order-discount/services"
)

func main() {
	purchaseAmount, isMember := ReadInput()

	// Calculate all applicable discounts
	purchaseDiscount, memberDiscount, totalDiscount := services.CalculateDiscount(purchaseAmount, isMember)
	voucherAmount := 20000

	// Apply discounts and voucher to get final payment amount
	finalAmount := purchaseAmount - totalDiscount - voucherAmount

	fmt.Println("------------------------")
	fmt.Println("Diskon Belanja: Rp.", purchaseDiscount)
	fmt.Println("Diskon Member : Rp.", memberDiscount)
	fmt.Println("Total Diskon  : Rp.", totalDiscount)
	fmt.Println("Voucher       : Rp.", voucherAmount)

	fmt.Println("------------------------")
	fmt.Println("Total Pembayaran : Rp.", finalAmount)
}
