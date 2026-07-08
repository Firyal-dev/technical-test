package utils

import (
	"fmt"
	"strings"
	"strconv"
)

// ParseTime parses time string in HH:MM format and validates the input
func ParseTime(input string) (int, int, error) {
	timeInput := strings.TrimSpace(input)
	parts := strings.Split(timeInput, ":")

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("format tidak valid. gunakan format HH:MM")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("input harus berupa angka")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("input harus berupa angka")
	}

	// Validate hour (0-23) and minute (0-59) ranges
	if hour < 0 || hour > 23 || minute < 0 || minute > 59 {
		return 0, 0, fmt.Errorf("jam harus antara 0-23 dan menit antara 0-59")
	}

	return hour, minute, nil
}