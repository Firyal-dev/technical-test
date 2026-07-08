package services

// CalculateDiscount calculates purchase and member discounts with a maximum cap
// Returns: purchaseDiscount, memberDiscount, totalDiscount (after applying cap)
func CalculateDiscount(purchaseAmount int, isMember bool) (int, int, int) {
	purchaseDiscount := 0
	memberDiscount := 0
	purchaseThreshold := 100000

	// 10% discount for purchases above 100k
	if purchaseAmount > purchaseThreshold {
		purchaseDiscount = purchaseAmount * 10 / 100
	}
	
	// 5% discount for members
	if isMember {
		memberDiscount = purchaseAmount * 5 / 100
	}

	totalDiscount := purchaseDiscount + memberDiscount

	// Apply 30% maximum discount cap
	maxDiscount := CalculateMaxDiscount(purchaseAmount)
	if totalDiscount > maxDiscount {
		totalDiscount = maxDiscount
	}

	return purchaseDiscount, memberDiscount, totalDiscount
}