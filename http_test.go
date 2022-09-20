package httperr

import (
	"net/http"
	"testing"
)

func TestNewHTTPErr(t *testing.T) {
	err := New(http.StatusOK, "resource available")

	if e, ok := err.(*HTTPErr); ok {
		if e.Code != http.StatusOK {
			t.Errorf("want %+v, got %+v", http.StatusOK, e.Code)
		}
	} else {
		t.Error("type check failed")
	}

	if err.Error() != "HTTP 200: resource available" {
		t.Errorf("want %+v, got %+v", err.Error(), "HTTP 200: resource available")
	}
}
