// DO NOT EDIT — implement the solution in solution.go

package restapi_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yourname/go-course/projects/restapi"
)

const apiKey = "dev-key"

func authReq(method, url string, body []byte) *http.Request {
	var req *http.Request
	if body == nil {
		req = httptest.NewRequest(method, url, nil)
	} else {
		req = httptest.NewRequest(method, url, bytes.NewReader(body))
	}
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestHealth_NoAuthRequired(t *testing.T) {
	srv := httptest.NewServer(restapi.NewServer())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status = %d", resp.StatusCode)
	}
}

func TestCRUD_RequiresAuth(t *testing.T) {
	h := restapi.NewServer()
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", rec.Code)
	}
}

func TestCRUD_Lifecycle(t *testing.T) {
	h := restapi.NewServer()

	// Create
	body, _ := json.Marshal(map[string]string{"name": "widget"})
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodPost, "/items", body))
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status = %d body=%s", rec.Code, rec.Body.String())
	}
	var created restapi.Item
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	if created.ID == "" || created.Name != "widget" {
		t.Fatalf("created = %+v", created)
	}

	// Read list
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodGet, "/items", nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("list status = %d", rec.Code)
	}

	// Read one
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodGet, "/items/"+created.ID, nil))
	if rec.Code != http.StatusOK {
		t.Fatalf("get status = %d", rec.Code)
	}

	// Update
	upd, _ := json.Marshal(map[string]string{"name": "gadget"})
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodPut, "/items/"+created.ID, upd))
	if rec.Code != http.StatusOK {
		t.Fatalf("update status = %d", rec.Code)
	}

	// Delete
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodDelete, "/items/"+created.ID, nil))
	if rec.Code != http.StatusNoContent {
		t.Fatalf("delete status = %d", rec.Code)
	}

	// Gone
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodGet, "/items/"+created.ID, nil))
	if rec.Code != http.StatusNotFound {
		t.Fatalf("after delete status = %d", rec.Code)
	}
}

func TestLoggingMiddleware(t *testing.T) {
	h := restapi.NewServer()
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, authReq(http.MethodGet, "/items", nil))
	if got := rec.Header().Get("X-Request-Logged"); got != "true" {
		t.Fatalf("X-Request-Logged = %q, want true", got)
	}
}
