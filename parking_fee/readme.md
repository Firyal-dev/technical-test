# Parking Fee

Sistem perhitungan biaya parkir berdasarkan durasi parkir dengan tarif berbeda untuk jam pertama dan jam selanjutnya.

## Deskripsi

Aplikasi ini menghitung biaya parkir berdasarkan:
- Durasi parkir dihitung dari jam masuk hingga jam keluar
- Tarif jam pertama: Rp. 5.000
- Tarif jam selanjutnya: Rp. 3.000 per jam
- Pembulatan ke atas: setiap menit lebih dari 0 dihitung sebagai 1 jam penuh

## Struktur Proyek

```
parking_fee/
├── main.go                                # Entry point aplikasi
├── read_input.go                          # Fungsi untuk membaca input dari user
├── services/
│   ├── count_parking_duration.go          # Logika perhitungan durasi parkir
│   ├── count_parking_duration_test.go     # Unit test untuk durasi parkir
│   ├── count_parking_rate.go              # Logika perhitungan tarif parkir
│   └── count_parking_rate_test.go         # Unit test untuk tarif parkir
├── utils/
│   └── parse_time.go                      # Helper untuk parsing waktu
└── readme.md                              # File ini
```

## Penjelasan File

### `main.go`
Entry point aplikasi yang mengorchestrasi flow perhitungan parkir:
- Membaca input jam masuk dan jam keluar
- Menghitung durasi parkir
- Menghitung tarif parkir dengan pembulatan
- Menampilkan durasi dan total bayar

### `read_input.go`
Berisi fungsi untuk membaca dan memvalidasi input dari user:
- Input jam masuk (format: HH:MM)
- Input jam keluar (format: HH:MM)

### `services/count_parking_duration.go`
Berisi fungsi `CalculateParkingDuration()` yang menghitung:
- Durasi parkir dalam jam dan menit
- Validasi jam keluar harus lebih besar dari jam masuk
- Return: `(hours, minutes, error)`

### `services/count_parking_rate.go`
Berisi fungsi `CalculateParkingFee()` yang menghitung:
- Tarif parkir berdasarkan durasi
- Jam pertama: Rp. 5.000
- Jam kedua dan seterusnya: Rp. 3.000/jam
- Return: `totalFee`

### `utils/parse_time.go`
Utility helper untuk:
- Parsing format waktu HH:MM
- Validasi input waktu
- Konversi waktu ke format yang dapat dihitung

### `services/*_test.go`
Unit test untuk setiap fungsi di services menggunakan table-driven test pattern dengan berbagai test case.

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

1. Jam masuk dan jam keluar diasumsikan berada pada hari yang sama. contoh: jam masuk 08.30, jam keluar 13.00. jika jam masuk dan jam keluar beda hari, contoh: jam masuk 17.00, jam keluar 06.00, maka dianggap tidak valid

2. Tarif parkir satu jam pertama diasumsikan Rp. 5.000, lalu jam setelahnya Rp. 3.000