package airplane

import (
	"github.com/iyaozhen/samber-do-learn/service/engine"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(engine.NewEngine),
	do.Lazy(NewAirplane),
)

// Provider
func NewAirplane(i do.Injector) (*Airplane, error) {
	return &Airplane{
		// import dependency
		Engine: do.MustInvoke[*engine.Engine](i),
	}, nil
}

type Airplane struct {
	Engine *engine.Engine
}

func (c *Airplane) Start() {
	c.Engine.Started = true
	println("vroooom")
}
