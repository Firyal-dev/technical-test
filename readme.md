# Technical Test Repository

Repository ini berisi 3 proyek technical test yang berbeda, masing-masing mendemonstrasikan implementasi logika bisnis sederhana menggunakan Go.

## Daftar Proyek

### 1. Order Discount
Sistem perhitungan diskon belanja yang menghitung:
- Diskon belanja 10% (untuk pembelian di atas Rp. 100.000)
- Diskon member 5% (untuk pelanggan member)
- Potongan voucher Rp. 20.000 (otomatis diterapkan)
- Validasi maksimal diskon 30% (tidak termasuk voucher)

### 2. Parking Fee
Sistem perhitungan biaya parkir berdasarkan durasi parkir:
- Jam pertama: Rp. 5.000
- Jam selanjutnya: Rp. 3.000/jam
- Durasi dihitung dengan pembulatan ke atas (setiap menit lebih dari 0 dihitung 1 jam penuh)

### 3. Mini ATM
Simulasi ATM sederhana dengan fitur:
- Withdraw (tarik tunai)
- Deposit (setor tunai)
- Transfer ke rekening lain
- Validasi saldo dan nominal transaksi

## Teknologi

- **Bahasa:** Go
- **Testing Framework:** Go testing package (built-in)
- **Go Version:** 1.21+

## Cara Setup Projek

Clone repository ini:
```bash
git clone https://github.com/Firyal-dev/technical-test.git
cd technical-test
```

## Struktur Repository

```
technical-test/
├── order_discount/     # Proyek perhitungan diskon belanja
├── parking_fee/        # Proyek perhitungan biaya parkir
├── mini_atm/           # Proyek simulasi Mini ATM
└── readme.md           # File ini
```

## Cara Menjalankan Projek

### 1. Masuk ke direktori salah satu projek
```bash
# Contoh: masuk ke projek order_discount
cd order_discount
```

### 2. Jalankan projek
```bash
go run ./
```

### 3. Menjalankan Unit Test
Unit test berada di folder `./services` di masing-masing projek. Cara menjalankannya:
```bash
go test -v ./services
```

## Catatan

Setiap proyek memiliki README.md tersendiri yang menjelaskan struktur folder, file, dan fungsi secara detail. Silakan baca README di masing-masing direktori proyek untuk informasi lebih lanjut.