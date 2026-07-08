package services

import "testing"

func TestWithdraw(t *testing.T) {
	tests := []struct {
		name          string
		initialSaldo  int
		withdrawAmount int
		expectedSaldo int
		expectError   bool
	}{
		{
			name:          "tarik tunai normal",
			initialSaldo:  100000,
			withdrawAmount: 50000,
			expectedSaldo: 50000,
			expectError:   false,
		},
		{
			name:          "tarik tunai seluruh saldo",
			initialSaldo:  100000,
			withdrawAmount: 100000,
			expectedSaldo: 0,
			expectError:   false,
		},
		{
			name:          "tarik tunai melebihi saldo",
			initialSaldo:  50000,
			withdrawAmount: 100000,
			expectedSaldo: 50000,
			expectError:   true,
		},
		{
			name:          "tarik tunai dengan nominal 0",
			initialSaldo:  100000,
			withdrawAmount: 0,
			expectedSaldo: 100000,
			expectError:   true,
		},
		{
			name:          "tarik tunai dengan nominal negatif",
			initialSaldo:  100000,
			withdrawAmount: -50000,
			expectedSaldo: 100000,
			expectError:   true,
		},
		{
			name:          "tarik tunai dari saldo kosong",
			initialSaldo:  0,
			withdrawAmount: 10000,
			expectedSaldo: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &Account{Saldo: tt.initialSaldo}
			saldo, err := account.Withdraw(tt.withdrawAmount)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if saldo != tt.expectedSaldo {
					t.Errorf("Expected saldo %d, got %d", tt.expectedSaldo, saldo)
				}
				if account.Saldo != tt.expectedSaldo {
					t.Errorf("Expected account saldo %d, got %d", tt.expectedSaldo, account.Saldo)
				}
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	tests := []struct {
		name          string
		initialSaldo  int
		depositAmount int
		expectedSaldo int
		expectError   bool
	}{
		{
			name:          "setor tunai normal",
			initialSaldo:  50000,
			depositAmount: 50000,
			expectedSaldo: 100000,
			expectError:   false,
		},
		{
			name:          "setor tunai ke saldo kosong",
			initialSaldo:  0,
			depositAmount: 100000,
			expectedSaldo: 100000,
			expectError:   false,
		},
		{
			name:          "setor tunai dengan nominal 0",
			initialSaldo:  100000,
			depositAmount: 0,
			expectedSaldo: 100000,
			expectError:   true,
		},
		{
			name:          "setor tunai dengan nominal negatif",
			initialSaldo:  100000,
			depositAmount: -50000,
			expectedSaldo: 100000,
			expectError:   true,
		},
		{
			name:          "setor tunai nominal besar",
			initialSaldo:  1000000,
			depositAmount: 5000000,
			expectedSaldo: 6000000,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &Account{Saldo: tt.initialSaldo}
			saldo, err := account.Deposit(tt.depositAmount)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if saldo != tt.expectedSaldo {
					t.Errorf("Expected saldo %d, got %d", tt.expectedSaldo, saldo)
				}
				if account.Saldo != tt.expectedSaldo {
					t.Errorf("Expected account saldo %d, got %d", tt.expectedSaldo, account.Saldo)
				}
			}
		})
	}
}

func TestTransfer(t *testing.T) {
	tests := []struct {
		name                  string
		initialSaldo          int
		transferAmount        int
		destinationAccount    string
		expectedSaldo         int
		expectedDestination   string
		expectError           bool
	}{
		{
			name:                "transfer normal",
			initialSaldo:        200000,
			transferAmount:      100000,
			destinationAccount:  "123456789",
			expectedSaldo:       100000,
			expectedDestination: "123456789",
			expectError:         false,
		},
		{
			name:                "transfer seluruh saldo",
			initialSaldo:        100000,
			transferAmount:      100000,
			destinationAccount:  "987654321",
			expectedSaldo:       0,
			expectedDestination: "987654321",
			expectError:         false,
		},
		{
			name:                "transfer melebihi saldo",
			initialSaldo:        50000,
			transferAmount:      100000,
			destinationAccount:  "123456789",
			expectedSaldo:       50000,
			expectedDestination: "",
			expectError:         true,
		},
		{
			name:                "transfer dengan nominal 0",
			initialSaldo:        100000,
			transferAmount:      0,
			destinationAccount:  "123456789",
			expectedSaldo:       100000,
			expectedDestination: "",
			expectError:         true,
		},
		{
			name:                "transfer dengan nominal negatif",
			initialSaldo:        100000,
			transferAmount:      -50000,
			destinationAccount:  "123456789",
			expectedSaldo:       100000,
			expectedDestination: "",
			expectError:         true,
		},
		{
			name:                "transfer dari saldo kosong",
			initialSaldo:        0,
			transferAmount:      10000,
			destinationAccount:  "123456789",
			expectedSaldo:       0,
			expectedDestination: "",
			expectError:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			account := &Account{Saldo: tt.initialSaldo}
			saldo, destination, err := account.Transfer(tt.transferAmount, tt.destinationAccount)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if saldo != tt.expectedSaldo {
					t.Errorf("Expected saldo %d, got %d", tt.expectedSaldo, saldo)
				}
				if account.Saldo != tt.expectedSaldo {
					t.Errorf("Expected account saldo %d, got %d", tt.expectedSaldo, account.Saldo)
				}
				if destination != tt.expectedDestination {
					t.Errorf("Expected destination %s, got %s", tt.expectedDestination, destination)
				}
			}
		})
	}
}
