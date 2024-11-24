package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"archetype/app/shared/configuration"
	"archetype/app/shared/infrastructure/labstackecho/httpserver"
)

func TestNewTemplatePatch(t *testing.T) {
	conf := configuration.Conf{}
	wrapper := httpserver.New(conf)

	newTemplatePatch(wrapper)

	req := httptest.NewRequest(http.MethodPatch, "/insert-your-custom-pattern-here", nil)
	rec := httptest.NewRecorder()

	wrapper.Manager.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	var resp struct {
		Message string `json:"message"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	expectedMessage := "Unimplemented"
	if resp.Message != expectedMessage {
		t.Errorf("expected message %q, got %q", expectedMessage, resp.Message)
	}
}
