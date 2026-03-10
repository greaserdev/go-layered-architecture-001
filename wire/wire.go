//go:build wireinject
// +build wireinject

package wire

import (
	"be-test/config"
	"be-test/controllers"
	"be-test/repositories"
	"be-test/services"

	googleWire "github.com/google/wire"
)

var configSet = googleWire.NewSet(config.GoogleOauthConfigInit)

var serviceSet = googleWire.NewSet(
	services.NewUserServiceImpl,
	googleWire.Bind(
		new(services.UserService),
		new(*services.UserServiceImpl),
	),
	services.NewGoogleOauthService,
	googleWire.Bind(
		new(services.GoogleOauthService),
		new(*services.GoogleOauthServiceImpl),
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
