package controllers

import (
	"be-test/services"
	"be-test/utils"
	"context"

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

	token, err := r.googleOauthService.ExchangeCode(
		context.Background(),
		code,
	)
	if err != nil {
		utils.HttpErrorBadRequest(
			c,
			err.Error(),
		)
	}

	user, err := r.googleOauthService.GetGoogleUserData(token.AccessToken)
	if err != nil {
		utils.HttpErrorBadRequest(
			c,
			err.Error(),
		)
	}

	utils.HttpSuccessOK(
		c,
		"OK",
		user,
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
