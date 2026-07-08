package services

import "testing"

func TestCalculateDiscount(t *testing.T) {
	tests := []struct {
		name                     string
		purchaseAmount           int
		isMember                 bool
		expectedPurchaseDiscount int
		expectedMemberDiscount   int
		expectedTotalDiscount    int
	}{
		{
			name:                     "belanja di atas 100rb, bukan member",
			purchaseAmount:           200000,
			isMember:                 false,
			expectedPurchaseDiscount: 20000,
			expectedMemberDiscount:   0,
			expectedTotalDiscount:    20000,
		},
		{
			name:                     "belanja di bawah 100rb, member",
			purchaseAmount:           50000,
			isMember:                 true,
			expectedPurchaseDiscount: 0,
			expectedMemberDiscount:   2500,
			expectedTotalDiscount:    2500,
		},
		{
			name:                     "belanja di atas 100rb, member",
			purchaseAmount:           200000,
			isMember:                 true,
			expectedPurchaseDiscount: 20000,
			expectedMemberDiscount:   10000,
			expectedTotalDiscount:    30000,
		},
		{
			name:                     "belanja di bawah 100rb, bukan member",
			purchaseAmount:           50000,
			isMember:                 false,
			expectedPurchaseDiscount: 0,
			expectedMemberDiscount:   0,
			expectedTotalDiscount:    0,
		},
		{
			name:                     "belanja di sama dengan 100rb, bukan member",
			purchaseAmount:           100000,
			isMember:                 false,
			expectedPurchaseDiscount: 0,
			expectedMemberDiscount:   0,
			expectedTotalDiscount:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			purchaseDiscount, memberDiscount, totalDiscount := CalculateDiscount(tt.purchaseAmount, tt.isMember)
			if purchaseDiscount != tt.expectedPurchaseDiscount {
				t.Errorf("Expected purchase discount %d, got %d", tt.expectedPurchaseDiscount, purchaseDiscount)
			}
			if memberDiscount != tt.expectedMemberDiscount {
				t.Errorf("Expected member discount %d, got %d", tt.expectedMemberDiscount, memberDiscount)
			}
			if totalDiscount != tt.expectedTotalDiscount {
				t.Errorf("Expected total discount %d, got %d", tt.expectedTotalDiscount, totalDiscount)
			}
		})
	}
}