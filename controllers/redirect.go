package controllers

import (
	"be-test/services"

	"github.com/gin-gonic/gin"
)

type RedirectController interface {
	HandleGoogleOauthRedirect(c *gin.Context)
}

type RedirectControllerImpl struct {
	userService        services.UserService
	googleOauthService services.GoogleOauthService
}

func (r *RedirectControllerImpl) HandleGoogleOauthRedirect(c *gin.Context) {
	code := c.Query("code")

	c.JSON(
		200,
		gin.H{"code": code},
	)

}

func NewRedirectController(
	user services.UserService,
	googleOauth services.GoogleOauthService,
) *RedirectControllerImpl {
	return &RedirectControllerImpl{
		userService:        user,
		googleOauthService: googleOauth,
	}
}
