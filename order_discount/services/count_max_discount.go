package services

// CalculateMaxDiscount returns the maximum allowed discount (30% of purchase amount)
func CalculateMaxDiscount(purchaseAmount int) int {
	maxDiscount := purchaseAmount * 30 / 100
	return maxDiscount
}
