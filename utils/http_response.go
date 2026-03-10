package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpErrorNotFound(
	c *gin.Context,
	message string,
) {
	if message == "" {
		message = "Not Found"
	}
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"code":    http.StatusNotFound,
			"message": message,
		},
	)
}

func HttpErrorBadRequest(
	c *gin.Context,
	message string,
) {
	if message == "" {
		message = "Bad Request"
	}
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"code":    http.StatusBadRequest,
			"message": message,
		},
	)
}

func HttpSuccessOK(
	c *gin.Context,
	message string,
	data interface{},
) {
	if message == "" {
		message = "OK"
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": message,
			"data":    data,
		},
	)
}

func HttpSuccessCreated(
	c *gin.Context,
	message string,
	data interface{},
) {
	if message == "" {
		message = "Created"
	}
	c.JSON(
		http.StatusCreated,
		gin.H{
			"code":    http.StatusCreated,
			"message": message,
			"data":    data,
		},
	)
}
