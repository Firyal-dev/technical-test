package services

// CalculateParkingFee calculates parking fee based on hours
// Rate: First hour Rp5,000, then Rp3,000 per subsequent hour, capped at Rp25,000
func CalculateParkingFee(hours int) int {
	firstHourFee := 5000
	maxFee := 25000
	fee := firstHourFee + (hours-1)*3000

	if fee > maxFee {
		fee = maxFee
	}

	return fee
}
