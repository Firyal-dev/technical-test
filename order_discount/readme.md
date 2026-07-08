# Order Discount

Sistem perhitungan diskon belanja yang menghitung diskon pembelian, diskon member, dan potongan voucher dengan validasi maksimal diskon.

## Deskripsi

Aplikasi ini menghitung total diskon dan pembayaran akhir berdasarkan:
- Diskon belanja 10% untuk pembelian di atas Rp. 100.000
- Diskon member 5% untuk pelanggan yang memiliki status member
- Potongan voucher Rp. 20.000 yang otomatis diterapkan
- Validasi maksimal diskon 30% (tidak termasuk voucher)

## Struktur Proyek

```
order_discount/
├── main.go                          # Entry point aplikasi
├── read_input.go                    # Fungsi untuk membaca input dari user
├── services/
│   ├── count_discount.go            # Logika perhitungan diskon belanja dan member
│   ├── count_discount_test.go       # Unit test untuk count_discount
│   ├── count_max_discount.go        # Logika validasi diskon maksimal 30%
│   └── count_max_discount_test.go   # Unit test untuk count_max_discount
└── readme.md                        # File ini
```

## Penjelasan File

### `main.go`
Entry point aplikasi yang mengorchestrasi keseluruhan flow program:
- Membaca input jumlah belanja dan status member
- Memanggil fungsi perhitungan diskon dari services
- Menampilkan detail diskon dan total pembayaran

### `read_input.go`
Berisi fungsi untuk membaca dan memvalidasi input dari user:
- Input jumlah belanja (integer)
- Input status member (ya/tidak)

### `services/count_discount.go`
Berisi fungsi `CalculateDiscount()` yang menghitung:
- Diskon belanja: 10% jika pembelian > Rp. 100.000
- Diskon member: 5% jika user adalah member
- Total diskon: gabungan diskon belanja dan diskon member
- Return: `(purchaseDiscount, memberDiscount, totalDiscount)`

### `services/count_max_discount.go`
Berisi fungsi validasi untuk memastikan total diskon tidak melebihi 30% dari total belanja (tidak termasuk voucher).

### `services/*_test.go`
Unit test untuk setiap fungsi di services menggunakan table-driven test pattern.

## Cara Menjalankan

### Menjalankan Aplikasi
```bash
go run ./
```

### Menjalankan Unit Test
```bash
go test -v ./services
```

### Menjalankan Test dengan Coverage
```bash
go test -v -cover ./services
```

## Asumsi

1. Input belanja berupa bilangan bulat (integer) tanpa pemisah ribuan atau desimal. contoh: 100000, sedangkan 100.000 atau 100.000.00 tidak valid. hal ini dilakukan karna dalam penggunaan sehari hari, Rupiah kebanyakan menggunakan bilangan bulat

2. Diskon belanja dan diskon member dihitung secara independent, bukan secara compounding. contoh: belanja Rp. 200.000 dengan status member. diskon belanja dihitung 10% dari total belanja, dan diskon member dihitung 5% dari total belanja juga, bukan 5% dari sisa belanja setelah diskon belanja diterapkan

3. Diskon maksimal 30% hanya berlaku untuk total diskon belanja + diskon member, tidak termasuk dengan potongan voucher 

4. Voucher Rp. 20.000 diasumsikan otomatis diterapkan disetiap transaksi