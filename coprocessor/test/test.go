package tes

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"data":{"hello":"world"}}`))
}

func TestApolloRouterEndpoint(t *testing.T) {
	body := []byte(`{"query":"{ hello }"}`)
	req := httptest.NewRequest(http.MethodPost, "/graphql", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	GraphQLHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	expected := `{"data":{"hello":"world"}}`
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != expected {
		t.Errorf("unexpected response body: got %s, want %s", buf.String(), expected)
	}
}
