package painting

import (
	"errors"
	"image"
)

// Position is X, Y coordinate pair. The axes increase
// right and down.
type Position = image.Point

// Frame are rectangular update screen patch
type Frame struct {
	// Position is the relative position where the
	// frame painting must start (from top left corner).
	Position Position

	// Patch is screen patch relative to the frame
	// position on the screen.
	Patch *Matrix
}

// NewFrame return a new Frame object
func NewFrame(p Position, width, height int) *Frame {
	return &Frame{
		Position: p,
		Patch:    NewMatrix(width, height),
	}
}

/*****************************************************
 ***************** GETTERS & SETTERS *****************
 *****************************************************/
// ANCHOR Getters & setter

// Bounds getter return the bounds of the frame.
func (f *Frame) Bounds() *image.Rectangle {
	return &image.Rectangle{
		Min: image.Point{
			X: f.Position.X,
			Y: f.Position.Y,
		},
		Max: image.Point{
			X: (f.Position.X + f.Patch.Width()),
			Y: (f.Position.Y + f.Patch.Height()),
		},
	}
}

// MaxHeightIndex return the max height index used by frame.
// (height + y position)
func (f *Frame) maxHeightIndex() uint {
	// Avoid -1 on with position and height = 0
	return uint(f.Position.Y + f.Patch.Height() - 1)
}

// MaxWidthIndex return the max width index used by frame.
// (Width + x position)
func (f *Frame) maxWidthIndex() uint {
	// Avoid -1 on with position and width = 0
	return uint(f.Position.X + f.Patch.Width() - 1)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// Add method add the given frame patch into the frame itself
// using the patch relative position.
func (f *Frame) Add(o *Frame) error {
	if can := f.CanContain(o); !can {
		return errors.New("the given frame can't be added or contained")
	}

	for i := 0; i < o.Patch.Height(); i++ {
		yOffset := o.Position.Y + i

		for j := 0; j < o.Patch.Width(); j++ {
			xOffset := o.Position.X + j

			f.Patch.M[yOffset][xOffset] = o.Patch.M[i][j]
		}
	}

	return nil
}

// CanContain return wether or not it can contain
// the given frame.
func (f *Frame) CanContain(o *Frame) bool {

	if !o.Patch.isValid() {
		return false
	}

	if o.Position.Y < 0 ||
		o.maxHeightIndex() >= uint(f.Patch.Height()) {
		return false
	}

	if o.Position.X < 0 ||
		o.maxWidthIndex() >= uint(f.Patch.Width()) {
		return false
	}

	return true
}

func (f *Frame) isEqual(other *Frame) bool {

	if f.maxWidthIndex() != other.maxWidthIndex() ||
		f.maxHeightIndex() != other.maxHeightIndex() {
		return false
	}

	if !f.Patch.isEqual(other.Patch) {
		return false
	}

	return true
}
