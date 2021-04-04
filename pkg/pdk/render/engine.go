package render

import (
	"context"
	"sync"
	"time"
)

// Engine is responsible for timing the rendering of renderable.
type Engine interface {
	// Start starts the engine rendering loop.
	Start(ctx context.Context)

	// Surface returns the underlying Surface used by this Engine.
	Surface() Surface

	// Enqueue enqueues the given Renderable for the next frame.
	Enqueue(renderable Renderable)
}

var _ Engine = &engine{}

type engine struct {
	sync.Mutex
	surface Surface
	queue   []Renderable
	clock   *time.Ticker
}

// Surface implements the Engine interface.
func (s *engine) Surface() Surface {
	return s.surface
}

// NewEngine returns a new Engine with the given Surface.
func NewEngine(surface Surface) Engine {
	return &engine{
		surface: surface,
		queue:   make([]Renderable, 0, 32),
		clock:   time.NewTicker(time.Millisecond * 16),
	}
}

func (s *engine) Enqueue(renderable Renderable) {
	s.Lock()
	defer s.Unlock()

	s.queue = append(s.queue, renderable)
}

func (s *engine) Start(ctx context.Context) {
	for {
		select {
		case <-s.clock.C:
			s.Lock()
			for _, renderable := range s.queue {
				if renderable.NeedRendering() {
					patch := renderable.Render()
					s.surface.Apply(patch)
				}
			}
			s.queue = make([]Renderable, 0, 32)
			s.Unlock()
			s.surface.Flush()

		case <-ctx.Done():
			return
		}
	}
}
