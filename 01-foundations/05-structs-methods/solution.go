package structs

// Rectangle represents a 2D axis-aligned rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns Width * Height. Use a value receiver — reading fields does not need a pointer.
func (r Rectangle) Area() float64 {
	return 0
}

// Scale multiplies Width and Height by factor. Must use a pointer receiver so the caller's struct is modified.
func (r *Rectangle) Scale(factor float64) {
}

// NewRectangle returns a pointer to a Rectangle with validated dimensions.
// Returns error if width or height is negative.
func NewRectangle(width, height float64) (*Rectangle, error) {
	return &Rectangle{}, nil
}
