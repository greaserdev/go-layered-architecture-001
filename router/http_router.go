package router

import (
	"be-test/wire"

	"github.com/gin-gonic/gin"
)

func HttpRouter() *gin.Engine {
	router := gin.Default()
	controllers := wire.AppController()

	v1 := router.Group("/api/v1")
	{

		redirect := v1.Group("/redirects")
		{
			redirect.GET(
				"/google-oauth",
				controllers.Redirect.HandleGoogleOauthRedirect,
			)
		}
	}

	return router
}
