package draw

type Drawer interface {
	Draw(Context)
}

// DrawerFn define a function that implements the Drawable interface.
type DrawerFn func(Context)

// Draw implements the Drawable interface.
func (s DrawerFn) Draw(ctx Context) {
	s(ctx)
}
