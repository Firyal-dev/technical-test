package services

import "testing"

func TestCalculateParkingFee(t *testing.T) {
	tests := []struct {
		name        string
		hours       int
		expectedFee int
	}{
		{
			name:        "parkir 1 jam",
			hours:       1,
			expectedFee: 5000,
		},
		{
			name:        "parkir 2 jam",
			hours:       2,
			expectedFee: 8000,
		},
		{
			name:        "parkir 3 jam",
			hours:       3,
			expectedFee: 11000,
		},
		{
			name:        "parkir 4 jam",
			hours:       4,
			expectedFee: 14000,
		},
		{
			name:        "parkir 5 jam",
			hours:       5,
			expectedFee: 17000,
		},
		{
			name:        "parkir 6 jam",
			hours:       6,
			expectedFee: 20000,
		},
		{
			name:        "parkir 7 jam",
			hours:       7,
			expectedFee: 23000,
		},
		{
			name:        "parkir 8 jam - mencapai cap",
			hours:       8,
			expectedFee: 25000,
		},
		{
			name:        "parkir 10 jam - tetap cap",
			hours:       10,
			expectedFee: 25000,
		},
		{
			name:        "parkir 24 jam - tetap cap",
			hours:       24,
			expectedFee: 25000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fee := CalculateParkingFee(tt.hours)
			if fee != tt.expectedFee {
				t.Errorf("Expected fee %d, got %d", tt.expectedFee, fee)
			}
		})
	}
}
