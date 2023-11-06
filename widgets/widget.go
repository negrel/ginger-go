package widgets

import (
	"github.com/negrel/paon/draw"
	"github.com/negrel/paon/events"
	"github.com/negrel/paon/geometry"
	"github.com/negrel/paon/layout"
	"github.com/negrel/paon/styles"
	"github.com/negrel/paon/tree"
)

// Widget is a generic interface that define any component part of the widget/element tree.
// Any types that implement the Widget interface can be added to the widget tree. However, it is strongly
// recommended to create custom widgets using the BaseWidget implementation.
type Widget interface {
	events.Target
	layout.Layout
	draw.Drawer
	styles.Styled

	Node() *tree.Node[Widget]
}

type eventTarget = events.Target

var _ Widget = &BaseWidget{}

// BaseWidget define a basic Widget object that implements the Widget interface.
// BaseWidget can either be used alone (see NewBaseWidget for the required options)
// or in composite struct.
type BaseWidget struct {
	eventTarget

	node *tree.Node[Widget]

	layout layout.Layout
	drawer draw.Drawer

	style styles.Style
}

// NewBaseWidget returns a new BaseWidget object configured with
// the given options.
// The LayoutAlgo and Drawer widget options are required.
// To embed this layout in composite struct, use the Wrap widget options.
func NewBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := newBaseWidget(options...)

	return widget
}

func newBaseWidget(options ...WidgetOption) *BaseWidget {
	widget := &BaseWidget{
		eventTarget: events.NewTarget(),
	}
	widgetConf := &baseWidgetOption{
		BaseWidget: widget,
		data:       widget,
	}

	for _, option := range options {
		option(widgetConf)
	}

	widget.node = tree.NewNode(widgetConf.data)

	for _, listener := range widgetConf.listeners {
		widget.AddEventListener(listener.EventType, listener.Handler)
	}

	return widget
}

// Layout implements Layout.
func (bw *BaseWidget) Layout(co layout.Constraint) geometry.Size {
	return bw.layout.Layout(co)
}

// Draw implements Drawer.
func (bw *BaseWidget) Draw(surface draw.Surface) {
	bw.drawer.Draw(surface)
}

// Node implements the Widget interface.
func (bw *BaseWidget) Node() *tree.Node[Widget] {
	return bw.node
}

// Style implements the styles.Styled interface.
func (bw *BaseWidget) Style() *styles.Style {
	return &bw.style
}
