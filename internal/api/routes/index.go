package routes

import "go.uber.org/fx"

type Routes interface {
	Setup()
}

type Impl []Routes

func (rgw *Impl) Setup() {
	for _, r := range *rgw {
		r.Setup()
	}
}

func NewRoutes() *Impl {
	return &Impl{}
}

var GateWayModule = fx.Provide(NewRoutes)
