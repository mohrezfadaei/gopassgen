package generator

import (
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name             string
		length           int
		excludeUppercase bool
		excludeLowercase bool
		excludeDigits    bool
		excludeSymbols   bool
		includeChars     string
		expectedError    bool
	}{
		{"Default", 16, false, false, false, false, "", false},
		{"No Uppercase", 16, true, false, false, false, "", false},
		{"No Lowercase", 16, false, true, false, false, "", false},
		{"No Digits", 16, false, false, true, false, "", false},
		{"No Special Characters", 16, false, false, false, true, "", false},
		{"Include Specific Symbols", 16, false, false, false, false, "A#<d5", false},
		{"Only Include Characters", 16, true, true, true, true, "A#<d5", false},
		{"Invalid Length", -1, false, false, false, false, "", false},
		{"No Characters Available", 16, true, true, true, true, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := GeneratePassword(tt.length, tt.excludeUppercase, tt.excludeLowercase, tt.excludeDigits, tt.excludeSymbols, tt.includeChars)
			if (err != nil) != tt.expectedError {
				t.Errorf("GeneratePassword() error = %v, expectedError %v", err, tt.expectedError)
				return
			}
			if !tt.expectedError && len(password) != tt.length && tt.length > 0 {
				t.Errorf("GeneratePassword() length = %v, expected %v", len(password), tt.length)
			}
		})
	}
}

func TestRandomChar(t *testing.T) {
	charset := "abc"
	char, err := randomChar(charset)
	if err != nil {
		t.Fatalf("randomChar() error = %v", err)
	}
	if !strings.Contains(charset, string(char)) {
		t.Errorf("randomChar() char = %v, not in charset %v", char, charset)
	}
}
