// DO NOT EDIT — implement the solution in solution.go

package jsonlesson_test

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/yourname/go-course/stdlib/jsonlesson"
)

func TestToJSON_OmitEmpty(t *testing.T) {
	data, err := jsonlesson.ToJSON(jsonlesson.Profile{Name: "Ada"})
	if err != nil {
		t.Fatal(err)
	}
	var m map[string]any
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatal(err)
	}
	if _, ok := m["age"]; ok {
		t.Fatal("age should be omitted when zero")
	}
	if m["name"] != "Ada" {
		t.Fatalf("name = %v", m["name"])
	}
}

func TestToJSON_Indented(t *testing.T) {
	data, err := jsonlesson.ToJSON(jsonlesson.Profile{Name: "x", Age: 1})
	if err != nil {
		t.Fatal(err)
	}
	if data[0] != '{' || data[1] != '\n' {
		t.Fatalf("expected indented JSON, got %q", string(data[:min(20, len(data))]))
	}
}

func TestFromJSON_UnknownFieldsIgnored(t *testing.T) {
	raw := []byte(`{"name":"Lin","age":30,"extra":"ignored"}`)
	var p jsonlesson.Profile
	if err := jsonlesson.FromJSON(raw, &p); err != nil {
		t.Fatal(err)
	}
	if p.Name != "Lin" || p.Age != 30 {
		t.Fatalf("got %+v", p)
	}
}

func TestFromJSON_Invalid(t *testing.T) {
	var p jsonlesson.Profile
	if err := jsonlesson.FromJSON([]byte("{"), &p); err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestProfile_PasswordNeverInJSON(t *testing.T) {
	data, err := jsonlesson.ToJSON(jsonlesson.Profile{Name: "a", Password: "secret"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(data), "secret") {
		t.Fatal("password leaked into JSON")
	}
}

func TestPublicView(t *testing.T) {
	in := jsonlesson.Profile{Name: "Ada", Age: 36, Password: "x", Internal: "hidden"}
	got := jsonlesson.PublicView(in)
	want := jsonlesson.Profile{Name: "Ada", Age: 36}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PublicView() = %+v, want %+v", got, want)
	}
}
