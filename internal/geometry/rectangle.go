package geometry

var _ Sized = &Rectangle{}

// Rectangle define a rectangle with two points
// for the minimum (top-left corner) and the
// maximum (bottom-right corner).
type Rectangle struct {
	Min, Max Point
}

// Rect returns always a well-formed rectangle with the
// given dimension.
func Rect(x0, y0, x1, y1 int) Rectangle {
	if x0 > x1 {
		x0, x1 = x1, x0
	}
	if y0 > y1 {
		y0, y1 = y1, y0
	}
	return Rectangle{
		Min: Pt(x0, y0),
		Max: Pt(x1, y1),
	}
}

// RectFromCenter returns a Rectangle with the given
// Size and the given center Point.
func RectFromCenter(center Point, size Size) Rectangle {
	return Rectangle{
		Min: Pt(
			center.x-size.Width()/2,
			center.y-size.Height()/2,
		),
		Max: Pt(
			center.x+size.Width()/2,
			center.y+size.Height()/2,
		),
	}
}

// Bottom returns offset of the bottom edge from the
// y axis.
func (r Rectangle) Bottom() int {
	return r.Max.y
}

// BottomCenter returns the offset to the center of the
// bottom center of this rectangle.
func (r Rectangle) BottomCenter() Point {
	return Pt(r.Width()/2, r.Max.y)
}

// BottomLeft returns the offset to the bottom left
// corner of the bottom center of this rectangle.
func (r Rectangle) BottomLeft() Point {
	return Pt(r.Max.y, r.Max.x)
}

// BottomRight returns the offset to the bottom right
// corner of the bottom center of this rectangle.
func (r Rectangle) BottomRight() Point {
	return Pt(r.Right(), r.Bottom())
}

// Center returns the center of the rectangle.
func (r Rectangle) Center() Point {
	return Pt(r.Min.x+r.Width()/2, r.Min.y+r.Height()/2)
}

// CenterLeft returns the offset to the center of the left edge
// of this rectangle.
func (r Rectangle) CenterLeft() Point {
	return Pt(r.Min.x, r.Min.y+r.Height()/2)
}

// CenterRight returns the offset to the center of the right edge
// of this rectangle.
func (r Rectangle) CenterRight() Point {
	return Pt(r.Max.x, r.Min.y+r.Height()/2)
}

// Left returns the offset of the left edge of this
// rectangle from the x axis.
func (r Rectangle) Left() int {
	return r.Min.x
}

// Right returns the offset of the right edge of
// this rectangle from the x axis.
func (r Rectangle) Right() int {
	return r.Max.x
}

// Top return offset of the top edge from the
// y axis.
func (r Rectangle) Top() int {
	return r.Min.y
}

// TopCenter returns the offset to the center of the
// top center of this rectangle.
func (r Rectangle) TopCenter() Point {
	return Pt(r.Width()/2, r.Min.y)
}

// TopLeft returns the offset to the bottom left
// corner of the top center of this rectangle.
func (r Rectangle) TopLeft() Point {
	return Pt(r.Left(), r.Top())
}

// TopRight returns the offset to the bottom right
// corner of the bottom center of this rectangle.
func (r Rectangle) TopRight() Point {
	return r.Max
}

// Width returns the width of the rectangle
func (r Rectangle) Width() int {
	return r.Max.x - r.Min.x
}

// Height returns the height of the rectangle
func (r Rectangle) Height() int {
	return r.Max.y - r.Min.y
}

// Size returns the rectangle width and height in a
// Size object.
func (r Rectangle) Size() Size {
	return Size{
		width:  r.Width(),
		height: r.Height(),
	}
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() int {
	return r.Width() * r.Height()
}

// MoveBy returns a new Rectangle moved by the given offset.
func (r Rectangle) MoveBy(n Point) Rectangle {
	return Rectangle{
		Min: r.Min.Add(n),
		Max: r.Max.Add(n),
	}
}

// MoveTo returns a new Rectangle with the same dimensions with
// the given Point as origin.
func (r Rectangle) MoveTo(n Point) Rectangle {
	return Rect(n.X(), n.Y(), n.X()+r.Width(), n.Y()+r.Height())
}

// GrowLeft returns a new rectangle growing by n to the left.
func (r Rectangle) GrowLeft(n int) Rectangle {
	return Rectangle{
		Min: r.Min.Add(Pt(-n, 0)),
		Max: r.Max,
	}
}

// GrowTop returns a new rectangle growing by n to the top.
func (r Rectangle) GrowTop(n int) Rectangle {
	return Rectangle{
		Min: r.Min.Add(Pt(0, -n)),
		Max: r.Max,
	}
}

// GrowRight returns a new rectangle growing by n to the right.
func (r Rectangle) GrowRight(n int) Rectangle {
	return Rectangle{
		Min: r.Min,
		Max: r.Max.Add(Pt(n, 0)),
	}
}

// GrowBottom returns a new rectangle growing by n to the bottom.
func (r Rectangle) GrowBottom(n int) Rectangle {
	return Rectangle{
		Min: r.Min,
		Max: r.Max.Add(Pt(0, n)),
	}
}

// Empty returns true if the width or the height of the rectangle is 0.
func (r Rectangle) Empty() bool {
	return r.Min.X() == r.Max.X() || r.Min.Y() == r.Max.Y()
}

// Equals returns true if the given Rectangle is equal to this Rectangle.
func (r Rectangle) Equals(other Rectangle) bool {
	return r.Min.Equals(other.Min) && r.Max.Equals(other.Max)
}

// Contains returns whether or not the given Point is contained in the Rectangle.
func (r Rectangle) Contains(point Point) bool {
	return point.x >= r.Min.x && point.x < r.Max.x &&
		point.y >= r.Min.y && point.y < r.Max.y
}

// Intersect returns the largest rectangle contained by both this rectangle and the other. If the
// two rectangles do not overlap then the zero rectangle will be returned.
func (r Rectangle) Intersect(other Rectangle) Rectangle {
	if r.Empty() || !r.Overlaps(other) {
		return Rectangle{}
	}

	if r.Min.x < other.Min.x {
		r.Min.x = other.Min.x
	}
	if r.Min.y < other.Min.y {
		r.Min.y = other.Min.y
	}

	if r.Max.x < other.Max.x {
		r.Max.x = other.Max.x
	}
	if r.Max.y < other.Max.y {
		r.Max.y = other.Max.y
	}

	return r
}

// Mask returns a new rectangle of the overlapped part between this rectangle and
// the given one.
func (r Rectangle) Mask(other Rectangle) Rectangle {
	if r.Min.x < other.Min.x {
		r.Min.x = other.Min.x
	}
	if r.Min.y < other.Min.y {
		r.Min.y = other.Min.y
	}

	if r.Max.x > other.Max.x {
		r.Max.x = other.Max.x
	}
	if r.Max.y > other.Max.y {
		r.Max.y = other.Max.y
	}

	return r
}

// Union returns the smallest rectangle that contains both r and s.
func (r Rectangle) Union(other Rectangle) Rectangle {
	if r.Empty() {
		return other
	}
	if other.Empty() {
		return r
	}
	if r.Min.x > other.Min.x {
		r.Min.x = other.Min.x
	}
	if r.Min.y > other.Min.y {
		r.Min.y = other.Min.y
	}
	if r.Max.x < other.Max.x {
		r.Max.x = other.Max.x
	}
	if r.Max.y < other.Max.y {
		r.Max.y = other.Max.y
	}
	return r
}

// Overlaps returns true if the other Rectangle overlap this one.
func (r Rectangle) Overlaps(other Rectangle) bool {
	return !r.Empty() && !other.Empty() &&
		r.Min.X() < other.Max.X() && other.Min.X() < r.Max.X() &&
		r.Min.Y() < other.Max.Y() && other.Min.Y() < r.Max.Y()
}
