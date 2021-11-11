package paon

import (
	"context"
	"os"
	"runtime/debug"
	"time"

	"github.com/negrel/debuggo/pkg/log"
	"github.com/negrel/paon/internal/metrics"
	"github.com/negrel/paon/pdk/backend"
	"github.com/negrel/paon/pdk/backend/tcell"
	pdkevents "github.com/negrel/paon/pdk/events"
	pdkwidgets "github.com/negrel/paon/pdk/widgets"
)

// Application define a TUI application object.
type Application struct {
	terminal backend.Terminal
	clock    *time.Ticker
	do       chan func()
	root     *pdkwidgets.Root
	target   pdkevents.Target
	evch     chan pdkevents.Event
}

// NewApp returns a new Application object.
func NewApp() (*Application, error) {
	terminal, err := tcell.NewTerminal()
	if err != nil {
		return nil, err
	}

	app := &Application{
		terminal: terminal,
		clock:    time.NewTicker(time.Millisecond * 16), // 60 fps
		do:       make(chan func()),
		target:   pdkevents.NewTarget(),
		evch:     make(chan pdkevents.Event),
	}

	return app, nil
}

func (app *Application) recover() {
	if r := recover(); r != nil {
		log.Error(r)

		stack := debug.Stack()
		t := time.AfterFunc(time.Second, func() {
			log.Error(string(stack))
			os.Exit(1)
		})

		app.stop()
		t.Stop()

		panic(r)
	}
}

// DoChannel returns a write-only channel that can be used to execute
// function on the main thread.
func (app *Application) DoChannel() chan<- func() {
	return app.do
}

// Start starts the application console, event loop and render loop.
func (app *Application) Start(ctx context.Context, widget pdkwidgets.Widget) error {
	defer app.recover()

	err := app.terminal.Start(app.evch)
	if err != nil {
		return err
	}

	if app.root == nil {
		app.root = pdkwidgets.NewRoot(app.target, app.terminal)
	}
	app.root.SetChild(widget)

	app.eventLoop(ctx)

	return nil
}

func (app *Application) eventLoop(ctx context.Context) {
	for {
		select {
		case <-app.clock.C:
			app.root.PerformRender()
			go app.terminal.Flush()

		case ev := <-app.evch:
			app.target.DispatchEvent(ev)

		case fn := <-app.do:
			fn()

		case <-ctx.Done():
			metrics.Report(os.Stderr)
			app.stop()
			return
		}
	}
}

func (app *Application) stop() {
	app.terminal.Stop()

	if app.root != nil {
		app.root.SetChild(nil)
	}
}
