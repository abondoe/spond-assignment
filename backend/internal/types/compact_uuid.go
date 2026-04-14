package types

import (
	"encoding/json"
	"strings"

	"github.com/google/uuid"
)

type CompactUUID uuid.UUID

func compactUUIDToString(c CompactUUID) string {
	// Konverter til streng, fjern bindestreker, sett til UpperCase
	res := strings.ReplaceAll(uuid.UUID(c).String(), "-", "")
	res = strings.ToUpper(res)
	return res
}

// MarshalJSON sørger for at UUID-en blir "8FE411..." i stedet for "8fe411-..."
func (c CompactUUID) MarshalJSON() ([]byte, error) {
	res := compactUUIDToString(c)
	return json.Marshal(res)
}

// UnmarshalJSON gjør at du kan ta imot det kompakte formatet i en request
func (c *CompactUUID) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	id, err := uuid.Parse(s)
	if err != nil {
		return err
	}
	*c = CompactUUID(id)
	return nil
}

func (c CompactUUID) String() string {
	res := compactUUIDToString(c)
	return res
}
