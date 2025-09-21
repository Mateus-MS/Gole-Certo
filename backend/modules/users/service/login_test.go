package user_service

import "testing"

func TestPasswordMatch(t *testing.T) {
	tests := []struct {
		entityPass  string
		requestPass string
		expected    bool
	}{
		{"secret123", "secret123", true},
		{"secret123", "wrongpass", false},
		{"", "", true},          // both empty
		{"notempty", "", false}, // mismatch
		{"", "notempty", false}, // mismatch
	}

	for _, tt := range tests {
		result := passwordMatch(tt.entityPass, tt.requestPass)
		if result != tt.expected {
			t.Errorf("passwordMatch(%q, %q) = %v; want %v",
				tt.entityPass, tt.requestPass, result, tt.expected)
		}
	}
}
