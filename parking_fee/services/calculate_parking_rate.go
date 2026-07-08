package services

// CalculateParkingFee calculates parking fee based on duration with rounding
// Any partial hour is rounded up to full hour for billing
// Returns the total parking fee capped at maximum
func CalculateParkingFee(hours, minutes int) int {
	// Round up: any partial hour counts as full hour for billing
	if minutes > 0 {
		hours++
	}
	
	firstHourFee := 5000
	subsequentHourRate := 3000
	maxFee := 25000
	
	// Calculate fee: first hour + subsequent hours
	fee := firstHourFee + (hours-1)*subsequentHourRate
	
	// Apply maximum fee cap
	if fee > maxFee {
		fee = maxFee
	}
	
	return fee
}
