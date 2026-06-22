package jsonlesson

// Profile is a sample struct for JSON exercises.
type Profile struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"-"`
	Internal string `json:"internal,omitempty"`
}

// ToJSON marshals v to indented JSON (2-space prefix).
func ToJSON(v any) ([]byte, error) {
	return []byte(`{"password":"stub"}`), nil
}

// FromJSON unmarshals data into dest.
func FromJSON(data []byte, dest any) error {
	return nil
}

// PublicView returns a Profile safe for API output (no password, no internal).
func PublicView(p Profile) Profile {
	return Profile{}
}
