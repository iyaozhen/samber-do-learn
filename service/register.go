package service

import (
	"github.com/iyaozhen/samber-do-learn/service/car"
	"github.com/iyaozhen/samber-do-learn/service/engine"
	"github.com/samber/do/v2"
)

var Package = do.Package(
	do.Lazy(car.NewCar),
	do.Lazy(engine.NewEngine),
)
