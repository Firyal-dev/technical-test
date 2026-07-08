package services

import "fmt"

// CalculateParkingDuration calculates parking duration from entry to exit time
func CalculateParkingDuration(entryHour, entryMinute, exitHour, exitMinute int) (int, int, error) {
	// Convert time to total minutes for easier calculation
	totalEntryMinutes := entryHour*60 + entryMinute
	totalExitMinutes := exitHour*60 + exitMinute

	if totalEntryMinutes >= totalExitMinutes {
		return 0, 0, fmt.Errorf("waktu masuk tidak boleh lebih besar dari waktu keluar")
	}

	// Calculate duration and split into hours and minutes
	totalDurationMinutes := totalExitMinutes - totalEntryMinutes
	durationHours := totalDurationMinutes / 60
	durationMinutes := totalDurationMinutes % 60

	return durationHours, durationMinutes, nil
}
