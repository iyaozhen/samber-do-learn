package engine

import "github.com/samber/do/v2"

// Provider
func NewEngine(i do.Injector) (*Engine, error) {
	return &Engine{
		Started: false,
	}, nil
}

type Engine struct {
	Started bool
}

func (e *Engine) Shutdown() error {
	// called on injector shutdown
	e.Started = false
	return nil
}
