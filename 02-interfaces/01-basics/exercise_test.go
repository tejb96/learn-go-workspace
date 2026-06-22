// DO NOT EDIT — implement the solution in solution.go

package basics_test

import (
	"reflect"
	"testing"

	"github.com/yourname/go-course/interfaces/basics"
)

func TestPerson_Speak(t *testing.T) {
	p := basics.Person{Name: "Ada"}
	if got := p.Speak(); got != "Hi, I'm Ada" {
		t.Fatalf("Speak() = %q, want %q", got, "Hi, I'm Ada")
	}
}

func TestRobot_Speak(t *testing.T) {
	r := basics.Robot{Model: "X1"}
	if got := r.Speak(); got != "Beep boop, model X1" {
		t.Fatalf("Speak() = %q, want %q", got, "Beep boop, model X1")
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		name    string
		speaker basics.Speaker
		want    string
	}{
		{
			name:    "person",
			speaker: basics.Person{Name: "Lin"},
			want:    "Greeting: Hi, I'm Lin",
		},
		{
			name:    "robot",
			speaker: basics.Robot{Model: "7"},
			want:    "Greeting: Beep boop, model 7",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := basics.Greet(tt.speaker)
			if got != tt.want {
				t.Fatalf("Greet() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSpeakAll(t *testing.T) {
	speakers := []basics.Speaker{
		basics.Person{Name: "A"},
		basics.Robot{Model: "B"},
	}
	got := basics.SpeakAll(speakers)
	want := []string{"Hi, I'm A", "Beep boop, model B"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("SpeakAll() = %v, want %v", got, want)
	}
}

func TestSpeakAll_Empty(t *testing.T) {
	got := basics.SpeakAll(nil)
	if got == nil || len(got) != 0 {
		t.Fatalf("SpeakAll(nil) = %v, want empty non-nil slice", got)
	}
}
