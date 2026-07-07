package main

// menghitung diskon belanja dan diskon member
func HitungDiskonBelanja(belanja int, isMember bool) (int, int, int) {
	diskonBelanja := 0
	diskonMember := 0
	minBelanja := 100000

	if belanja > minBelanja {
		diskonBelanja = belanja * 10 / 100 // diskon 10%
	}
	if isMember {
		diskonMember = belanja * 5 / 100 // diskon 5% untuk member
	}

	totalDiskon := diskonBelanja + diskonMember

	maxDiskon := MaksDiskon(belanja)
	if totalDiskon > maxDiskon {
		totalDiskon = maxDiskon
	}

	return diskonBelanja, diskonMember, totalDiskon
}

// menghitung maks diskon (30%)
func MaksDiskon(belanja int) int {
	maxDiskon := belanja * 30 / 100
	return maxDiskon
}
