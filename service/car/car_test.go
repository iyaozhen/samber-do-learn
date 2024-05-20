package car

import (
	"testing"

	"github.com/iyaozhen/samber-do-learn/service/engine"
	"github.com/samber/do/v2"
	"github.com/stretchr/testify/assert"
)

func TestCar_Start_V1(t *testing.T) {
	injector := do.New()
	Package(injector)

	car, err := do.Invoke[*Car](injector)
	assert.NoError(t, err)
	car.Start()
	assert.True(t, car.Engine.Started)
}

func TestCar_Start_V2(t *testing.T) {
	injector := do.New()
	do.Provide(injector, engine.NewEngine)
	do.Provide(injector, NewCar)

	car, err := do.Invoke[*Car](injector)
	assert.NoError(t, err)
	car.Start()
	assert.True(t, car.Engine.Started)
}
