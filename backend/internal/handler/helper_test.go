package handler_test

import (
	"abondoe/spond-assignment/internal/models/dto"
	"abondoe/spond-assignment/internal/types"
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func assertJSONContentType(t *testing.T, w *httptest.ResponseRecorder) {
	t.Helper() // Gjør feilsøking mye enklere!
	got := w.Header().Get("Content-Type")
	want := "application/json"
	if got != want {
		t.Errorf("Content-Type: forventet %q, fikk %q", want, got)
	}
}

func assertJSONBody(t *testing.T, body *bytes.Buffer, target any) {
	t.Helper()
	if err := json.NewDecoder(body).Decode(target); err != nil {
		t.Fatalf("kunne ikke dekode JSON body: %v", err)
	}
}

func CompactUUIDFromString(s string) types.CompactUUID {
	id, _ := uuid.Parse(s)
	return types.CompactUUID(id)
}

func validFormId() uuid.UUID {
	return uuid.MustParse("B171388180BC457D9887AD92B6CCFC86")
}

func validCompactFormId() types.CompactUUID {
	return types.CompactUUID(validFormId())
}

func validCreateRegistrationRequest() dto.CreateRegistrationRequest {
	return dto.CreateRegistrationRequest{
		FormId: validCompactFormId(),
		Name:   "John Doe",
	}
}
