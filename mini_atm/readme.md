# Mini ATM

Simulasi ATM sederhana dengan fitur withdraw, deposit, dan transfer dengan validasi saldo dan nominal transaksi.

## Deskripsi

Aplikasi ini menyediakan menu interaktif untuk melakukan transaksi perbankan dasar:
- Withdraw (tarik tunai) dengan validasi saldo mencukupi
- Deposit (setor tunai) untuk menambah saldo
- Transfer ke rekening lain dengan validasi saldo
- Validasi nominal transaksi harus lebih besar dari 0

## Struktur Proyek

```
mini_atm/
├── main.go                      # Entry point aplikasi dengan menu interaktif
├── read_input.go                # Fungsi untuk membaca input dari user
├── services/
│   ├── transaction.go           # Logika bisnis untuk Withdraw, Deposit, Transfer
│   └── transaction_test.go      # Unit test untuk semua transaksi
└── readme.md                    # File ini
```

## Penjelasan File

### `main.go`
Entry point aplikasi dengan menu interaktif yang menyediakan:
- Tampilan saldo saat ini
- Menu pilihan: Withdraw, Deposit, Transfer, Keluar
- Loop interaktif untuk melakukan transaksi berulang
- Handling input user dan menampilkan hasil transaksi

### `read_input.go`
Berisi fungsi untuk membaca dan memvalidasi input dari user:
- `ReadInput()`: Membaca pilihan menu (1-4)
- `ReadAmount()`: Membaca nominal transaksi
- `ReadAccountDestination()`: Membaca nomor rekening tujuan transfer

### `services/transaction.go`
Berisi struct `Account` dan method-method transaksi:

**Struct:**
- `Account`: Memiliki field `Saldo` untuk menyimpan saldo rekening

**Methods:**
- `Withdraw(amount int)`: Tarik tunai dengan validasi:
  - Nominal harus lebih besar dari 0
  - Saldo harus mencukupi
  - Return: `(saldoAkhir, error)`

- `Deposit(amount int)`: Setor tunai dengan validasi:
  - Nominal harus lebih besar dari 0
  - Return: `(saldoAkhir, error)`

- `Transfer(amount int, destinationAccount string)`: Transfer ke rekening lain dengan validasi:
  - Nominal harus lebih besar dari 0
  - Saldo harus mencukupi
  - Return: `(saldoAkhir, rekeningTujuan, error)`

### `services/transaction_test.go`
Unit test menggunakan table-driven test pattern untuk semua fungsi transaksi:
- `TestWithdraw`: 6 test cases (tarik tunai normal, seluruh saldo, melebihi saldo, nominal 0, negatif, dari saldo kosong)
- `TestDeposit`: 5 test cases (setor tunai normal, ke saldo kosong, nominal 0, negatif, nominal besar)
- `TestTransfer`: 6 test cases (transfer normal, seluruh saldo, melebihi saldo, nominal 0, negatif, dari saldo kosong)

## Cara Menjalankan

### Menjalankan Aplikasi
```bash
go run ./
```

Aplikasi akan menampilkan menu interaktif:
```
======= MINI ATM =======
Saldo saat ini : Rp. 1000000
Pilih transaksi:
1. Withdraw (tarik tunai)
2. Deposit (setor tunai)
3. Transfer ke rekening lain
4. Keluar
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

1. Rekening tujuan pada transfer diinput sebagai string bebas dan tidak divalidasi ada atau tidaknya. selama jumlah saldo cukup, maka transfer berhasil
2. 