package controllers

import (
	"be-test/domain"
	"be-test/services"

	"github.com/gin-gonic/gin"
)

type RedirectController interface {
	HandleGoogleOauthRedirect(c *gin.Context)
}

type RedirectControllerImpl struct {
	userService services.UserService
}

func (r *RedirectControllerImpl) HandleGoogleOauthRedirect(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"ping": r.userService.FindUser(
				c.Request.Context(),
				domain.UserServiceFindUserArg{
					Email:     "ihsan@gmail.com",
					FirstName: "",
					LastName:  "",
				},
			),
		},
	)
}

func NewRedirectController(user services.UserService) *RedirectControllerImpl {
	return &RedirectControllerImpl{userService: user}
}
