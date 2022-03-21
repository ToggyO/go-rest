package ioc

import (
	"github.com/google/wire"
	api "go-rest/internal/api/http"
)

// TODO: return type
type Test struct {
}

func BuildIoc() (*Test, error) {
	wire.Build(
		api.BindRouter,
		api.BindControllers,
		api.BindRouterGroups)

	return &Test{}, nil
}
