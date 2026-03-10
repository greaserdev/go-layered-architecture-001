//go:build wireinject
// +build wireinject

package wire

import (
	"be-test/controllers"
	"be-test/repositories"
	"be-test/services"

	googleWire "github.com/google/wire"
)

var serviceSet = googleWire.NewSet(
	services.NewUserServiceImpl,
	googleWire.Bind(
		new(services.UserService),
		new(*services.UserServiceImpl),
	),
)

var controllerSet = googleWire.NewSet(
	controllers.NewRedirectController,
	googleWire.Bind(
		new(controllers.RedirectController),
		new(*controllers.RedirectControllerImpl),
	),
)

var repositorySet = googleWire.NewSet(
	repositories.NewUserRepository,
	googleWire.Bind(
		new(repositories.UserRepository),
		new(*repositories.UserRepositoryImpl),
	),
)

func AppController() *controllers.ControllerRegistry {
	googleWire.Build(
		serviceSet,
		controllerSet,
		repositorySet,
		controllers.NewControllerRegistry,
	)
	return nil
}
