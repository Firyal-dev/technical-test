package main

import "fmt"

func main() {
	belanja, isMember := ReadInput()

	diskonBelanja, diskonMember, totalDiskon := HitungDiskonBelanja(belanja, isMember)
	voucher := 20000
	totalBelanja := belanja - totalDiskon - voucher

	fmt.Println("------------------------")
	fmt.Println("Diskon Belanja: Rp.", diskonBelanja)
	fmt.Println("Diskon Member : Rp.", diskonMember)
	fmt.Println("Total Diskon  : Rp.", totalDiskon)
	fmt.Println("Voucher       : Rp.", voucher)

	fmt.Println("------------------------")
	fmt.Println("Total Pembayaran : Rp.", totalBelanja)
}
