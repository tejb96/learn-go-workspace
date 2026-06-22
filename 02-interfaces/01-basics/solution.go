package basics

// Speaker can produce spoken output. Any type with a Speak() string method satisfies this
// interface implicitly — no "implements" keyword required.
type Speaker interface {
	Speak() string
}

// Person is a human Speaker.
type Person struct {
	Name string
}

func (p Person) Speak() string {
	return ""
}

// Robot is a non-human Speaker.
type Robot struct {
	Model string
}

func (r Robot) Speak() string {
	return ""
}

// Greet returns a greeting using the Speaker's output.
func Greet(s Speaker) string {
	return ""
}

// SpeakAll calls Speak on each Speaker and returns the joined lines in order.
func SpeakAll(speakers []Speaker) []string {
	return nil
}
