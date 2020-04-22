package render

import (
	"image"
)

// Func define the type of Rendable render function.
type Func = func(image.Rectangle) *Frame

// Rendable is implementeb by elements that have a Render
// method which return the rendered Frame.
type Rendable interface {
	Render(image.Rectangle) *Frame
}
