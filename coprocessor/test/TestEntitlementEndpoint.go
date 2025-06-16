package tes

import (
	coprocessor "coprocessor/internal"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEntitlementsEndpoint(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/entitlements", nil)
	w := httptest.NewRecorder()

	// Call the handler directly
	coprocessor.RequestHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %d", resp.StatusCode)
	}
}
