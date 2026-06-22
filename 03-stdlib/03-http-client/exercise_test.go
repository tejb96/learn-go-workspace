// DO NOT EDIT — implement the solution in solution.go

package httpclient_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yourname/go-course/stdlib/httpclient"
)

type userResponse struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

func TestFetchJSON_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(userResponse{Name: "Ada", ID: 1})
	}))
	defer srv.Close()

	var got userResponse
	if err := httpclient.FetchJSON(context.Background(), srv.URL, &got); err != nil {
		t.Fatal(err)
	}
	if got.Name != "Ada" || got.ID != 1 {
		t.Fatalf("got %+v, want Ada/1", got)
	}
}

func TestFetchJSON_Non2xx(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nope", http.StatusTeapot)
	}))
	defer srv.Close()

	var dest userResponse
	err := httpclient.FetchJSON(context.Background(), srv.URL, &dest)
	if err == nil {
		t.Fatal("expected error for 418")
	}
}

func TestFetchJSON_InvalidJSON(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("not json"))
	}))
	defer srv.Close()

	var dest userResponse
	if err := httpclient.FetchJSON(context.Background(), srv.URL, &dest); err == nil {
		t.Fatal("expected JSON decode error")
	}
}

func TestGetStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	code, err := httpclient.GetStatus(context.Background(), srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if code != http.StatusNoContent {
		t.Fatalf("status = %d, want 204", code)
	}
}
