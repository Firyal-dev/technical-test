package services

import "testing"

func TestCalculateMaxDiscount(t *testing.T) {
	tests := []struct {
		name                  string
		purchaseAmount        int
		expectedMaxDiscount   int
	}{
		{
			name:                "belanja 100000",
			purchaseAmount:      100000,
			expectedMaxDiscount: 30000,
		},
		{
			name:                "belanja 200000",
			purchaseAmount:      200000,
			expectedMaxDiscount: 60000,
		},
		{
			name:                "belanja 50000",
			purchaseAmount:      50000,
			expectedMaxDiscount: 15000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			maxDiscount := CalculateMaxDiscount(tt.purchaseAmount)
			if maxDiscount != tt.expectedMaxDiscount {
				t.Errorf("Expected maximum discount %d, got %d", tt.expectedMaxDiscount, maxDiscount)
			}
		})
	}
}
