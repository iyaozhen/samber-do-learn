package car

import (
	"github.com/iyaozhen/samber-do-learn/service/engine"
	"github.com/samber/do/v2"
)

// Provider
func NewCar(i do.Injector) (*Car, error) {
	return &Car{
		// import dependency
		Engine: do.MustInvoke[*engine.Engine](i),
	}, nil
}

type Car struct {
	Engine *engine.Engine
}

func (c *Car) Start() {
	c.Engine.Started = true
	println("vroooom")
}
