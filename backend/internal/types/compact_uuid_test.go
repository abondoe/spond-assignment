package types_test

import (
	"abondoe/spond-assignment/internal/types"
	"encoding/json"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestCompactUUID(t *testing.T) {
	// We use a fixed UUID to know exactly what we expect
	rawID := "550e8400-e29b-41d4-a716-446655440000"
	u, _ := uuid.Parse(rawID)
	compact := types.CompactUUID(u)

	// Expected format: No hyphens, only uppercase letters
	expectedCompact := "550E8400E29B41D4A716446655440000"

	t.Run("String() method", func(t *testing.T) {
		// Check if lowercase
		if compact.String() != strings.ToUpper(compact.String()) {
			t.Errorf("expected uppercase, got %s", compact.String())
		}
		// Check if contains -
		if strings.Contains(compact.String(), "-") {
			t.Errorf("expected no hyphens, got %s", compact.String())
		}
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		data, err := json.Marshal(compact)
		if err != nil {
			t.Fatalf("MarshalJSON feilet: %v", err)
		}

		expectedJSON := `"` + expectedCompact + `"`
		if string(data) != expectedJSON {
			t.Errorf("forventet %s, fikk %s", expectedJSON, string(data))
		}
	})

	t.Run("UnmarshalJSON Table Test", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr bool
		}{
			{
				name:    "Vellykket fra kompakt format",
				input:   `"550E8400E29B41D4A716446655440000"`,
				wantErr: false,
			},
			{
				name:    "Vellykket fra standard format",
				input:   `"550e8400-e29b-41d4-a716-446655440000"`,
				wantErr: false,
			},
			{
				name:    "Vellykket fra lowercase kompakt",
				input:   `"550e8400e29b41d4a716446655440000"`,
				wantErr: false,
			},
			{
				name:    "Feil ved ugyldig streng",
				input:   `"dette-er-ikke-en-uuid"`,
				wantErr: true,
			},
			{
				name:    "Feil ved ugyldig JSON-type (tall)",
				input:   `12345`,
				wantErr: true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				var c types.CompactUUID
				err := json.Unmarshal([]byte(tt.input), &c)

				if (err != nil) != tt.wantErr {
					t.Errorf("UnmarshalJSON() feil = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if !tt.wantErr {
					// Check that the value got correctly parsed back to the UUID
					if uuid.UUID(c) != u {
						t.Errorf("forventet %v, fikk %v", u, uuid.UUID(c))
					}
				}
			})
		}
	})
}
