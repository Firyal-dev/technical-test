package services

import "testing"

func TestCalculateParkingFee(t *testing.T) {
	tests := []struct {
		name        string
		hours       int
		minutes     int
		expectedFee int
	}{
		{
			name:        "parkir 1 jam",
			hours:       1,
			minutes:     0,
			expectedFee: 5000,
		},
		{
			name:        "parkir 2 jam",
			hours:       2,
			minutes:     0,
			expectedFee: 8000,
		},
		{
			name:        "parkir 3 jam",
			hours:       3,
			minutes:     0,
			expectedFee: 11000,
		},
		{
			name:        "parkir 4 jam",
			hours:       4,
			minutes:     0,
			expectedFee: 14000,
		},
		{
			name:        "parkir 5 jam",
			hours:       5,
			minutes:     0,
			expectedFee: 17000,
		},
		{
			name:        "parkir 6 jam",
			hours:       6,
			minutes:     0,
			expectedFee: 20000,
		},
		{
			name:        "parkir 7 jam",
			hours:       7,
			minutes:     0,
			expectedFee: 23000,
		},
		{
			name:        "parkir 8 jam - mencapai cap",
			hours:       8,
			minutes:     0,
			expectedFee: 25000,
		},
		{
			name:        "parkir 10 jam - tetap cap",
			hours:       10,
			minutes:     0,
			expectedFee: 25000,
		},
		{
			name:        "parkir 24 jam - tetap cap",
			hours:       24,
			minutes:     0,
			expectedFee: 25000,
		},
		{
			name:        "parkir 1 jam 30 menit - rounded up to 2 jam",
			hours:       1,
			minutes:     30,
			expectedFee: 8000,
		},
		{
			name:        "parkir 2 jam 1 menit - rounded up to 3 jam",
			hours:       2,
			minutes:     1,
			expectedFee: 11000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fee := CalculateParkingFee(tt.hours, tt.minutes)
			if fee != tt.expectedFee {
				t.Errorf("Expected fee %d, got %d", tt.expectedFee, fee)
			}
		})
	}
}
