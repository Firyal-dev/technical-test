package services

import "testing"

func TestCalculateParkingDuration(t *testing.T) {
	tests := []struct {
		name              string
		entryHour         int
		entryMinute       int
		exitHour          int
		exitMinute        int
		expectedHours     int
		expectedMinutes   int
		expectError       bool
	}{
		{
			name:            "parkir 2 jam tepat",
			entryHour:       8,
			entryMinute:     0,
			exitHour:        10,
			exitMinute:      0,
			expectedHours:   2,
			expectedMinutes: 0,
			expectError:     false,
		},
		{
			name:            "parkir 1 jam 30 menit",
			entryHour:       9,
			entryMinute:     15,
			exitHour:        10,
			exitMinute:      45,
			expectedHours:   1,
			expectedMinutes: 30,
			expectError:     false,
		},
		{
			name:            "parkir 3 jam 45 menit",
			entryHour:       7,
			entryMinute:     20,
			exitHour:        11,
			exitMinute:      5,
			expectedHours:   3,
			expectedMinutes: 45,
			expectError:     false,
		},
		{
			name:            "parkir 30 menit",
			entryHour:       14,
			entryMinute:     0,
			exitHour:        14,
			exitMinute:      30,
			expectedHours:   0,
			expectedMinutes: 30,
			expectError:     false,
		},
		{
			name:            "waktu masuk sama dengan waktu keluar",
			entryHour:       10,
			entryMinute:     0,
			exitHour:        10,
			exitMinute:      0,
			expectedHours:   0,
			expectedMinutes: 0,
			expectError:     true,
		},
		{
			name:            "waktu masuk lebih besar dari waktu keluar",
			entryHour:       15,
			entryMinute:     30,
			exitHour:        14,
			exitMinute:      20,
			expectedHours:   0,
			expectedMinutes: 0,
			expectError:     true,
		},
		{
			name:            "parkir dari pagi sampai sore",
			entryHour:       8,
			entryMinute:     15,
			exitHour:        17,
			exitMinute:      30,
			expectedHours:   9,
			expectedMinutes: 15,
			expectError:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hours, minutes, err := CalculateParkingDuration(tt.entryHour, tt.entryMinute, tt.exitHour, tt.exitMinute)
			
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if hours != tt.expectedHours {
					t.Errorf("Expected hours %d, got %d", tt.expectedHours, hours)
				}
				if minutes != tt.expectedMinutes {
					t.Errorf("Expected minutes %d, got %d", tt.expectedMinutes, minutes)
				}
			}
		})
	}
}
