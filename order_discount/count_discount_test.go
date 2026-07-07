package main

import "testing"

func TestHitungDiskonBelanja(t *testing.T) {
	tests := []struct {
		name                  string
		belanja               int
		isMember              bool
		expectedDiskonBelanja int
		expectedDiskonMember  int
		expectedTotalDiskon   int
	}{
		{
			name:                  "belanja di atas 100rb, bukan member",
			belanja:               200000,
			isMember:              false,
			expectedDiskonBelanja: 20000,
			expectedDiskonMember:  0,
			expectedTotalDiskon:   20000,
		},
		{
			name:                  "belanja di bawah 100rb, member",
			belanja:               50000,
			isMember:              true,
			expectedDiskonBelanja: 0,
			expectedDiskonMember:  2500,
			expectedTotalDiskon:   2500,
		},
		{
			name:                  "belanja di atas 100rb, member",
			belanja:               200000,
			isMember:              true,
			expectedDiskonBelanja: 20000,
			expectedDiskonMember:  10000,
			expectedTotalDiskon:   30000,
		},
		{
			name:                  "belanja di bawah 100rb, bukan member",
			belanja:               50000,
			isMember:              false,
			expectedDiskonBelanja: 0,
			expectedDiskonMember:  0,
			expectedTotalDiskon:   0,
		},
		{
			name:                  "belanja di sama dengan 100rb, bukan member",
			belanja:               100000,
			isMember:              false,
			expectedDiskonBelanja: 0,
			expectedDiskonMember:  0,
			expectedTotalDiskon:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diskonBelanja, diskonMember, totalDiskon := HitungDiskonBelanja(tt.belanja, tt.isMember)
			if diskonBelanja != tt.expectedDiskonBelanja {
				t.Errorf("Expected diskon belanja %d, got %d", tt.expectedDiskonBelanja, diskonBelanja)
			}
			if diskonMember != tt.expectedDiskonMember {
				t.Errorf("Expected diskon member %d, got %d", tt.expectedDiskonMember, diskonMember)
			}
			if totalDiskon != tt.expectedTotalDiskon {
				t.Errorf("Expected total diskon %d, got %d", tt.expectedTotalDiskon, totalDiskon)
			}
		})
	}
}

func TestMaksDiskon(t *testing.T) {
	tests := []struct {
		name               string
		belanja            int
		expectedMaksDiskon int
	}{
		{
			name:               "belanja 100000",
			belanja:            100000,
			expectedMaksDiskon: 30000,
		},
		{
			name:               "belanja 200000",
			belanja:            200000,
			expectedMaksDiskon: 60000,
		},
		{
			name:               "belanja 50000",
			belanja:            50000,
			expectedMaksDiskon: 15000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maksDiskon := MaksDiskon(tt.belanja)
			if maksDiskon != tt.expectedMaksDiskon {
				t.Errorf("Expected maksimal diskon %d, got %d", tt.expectedMaksDiskon, maksDiskon)
			}
		})
	}
}
