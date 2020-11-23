package healthy

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if rec.Body.String() != OK {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rec.Body.String(), OK)
	}

}
