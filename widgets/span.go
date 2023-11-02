package widgets

import (
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/pdk/draw"
	"github.com/negrel/paon/pdk/layout"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

type Span struct {
	*pdkwidgets.BaseWidget

	text string
}

func NewSpan(text string) *Span {
	w := &Span{
		text: text,
	}

	w.BaseWidget = pdkwidgets.NewBaseWidget(
		pdkwidgets.Wrap(w),
		pdkwidgets.LayoutFunc(func(co layout.Constraint) geometry.Size {
			return co.ApplyOnSize(geometry.NewSize(len(w.text), 1))
		}),
		pdkwidgets.DrawerFunc(func(surface draw.Surface) {
			// TODO: iterate over grapheme instead of runes.
			for i, c := range w.text {
				surface.Set(geometry.NewVec2D(i, 0), draw.Cell{
					Style:   draw.CellStyle{},
					Content: c,
				})
			}
		}),
	)

	return w
}
