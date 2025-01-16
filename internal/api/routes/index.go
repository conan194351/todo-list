package routes

import "go.uber.org/fx"

type Routes interface {
	Setup()
}

type RoutesImpl []Routes

func (rgw *RoutesImpl) Setup() {
	for _, r := range *rgw {
		r.Setup()
	}
}

func NewRoutes() *RoutesImpl {
	return &RoutesImpl{}
}

var RoutesGateWayModule = fx.Option(fx.Provide(NewRoutes))
