package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	// Define HTTP requests for testing
	requests := []struct {
		name         string
		queryParam   string
		codeExpected int
	}{
		{
			name:         "Empty input form",
			queryParam:   "",
			codeExpected: http.StatusOK,
		},
		{
			name:         "A valid number in input form",
			queryParam:   "6",
			codeExpected: http.StatusOK,
		},
		{
			name:         "A valid number in input form",
			queryParam:   "A4_",
			codeExpected: http.StatusBadRequest,
		},
	}

	for _, rq := range requests {
		req, err := http.NewRequest("GET", "/?n="+rq.queryParam, nil)
		if err != nil {
			t.Fatalf("Error making http request : %v", err)
		}

		w := httptest.NewRecorder()
		r := http.HandlerFunc(GetNumber)
		r.ServeHTTP(w, req)
		if w.Code != rq.codeExpected {
			t.Log(rq.name)
			t.Fatalf("Code %d for response not expected (%d)", w.Code, rq.codeExpected)
		}
	}
}

func TestFibo(t *testing.T) {
	// First place
	wantPos := 1
	expected := 0

	if p := fibo(wantPos - 1); p != expected {
		t.Fatalf("This position <%v> number should be : %v and results : %v", wantPos, expected, p)
	}

	// Other
	wantPos = 12
	expected = 89

	if p := fibo(wantPos - 1); p != expected {
		t.Fatalf("This position <%v> number should be : %v and results : %v", wantPos, expected, p)
	}
}
