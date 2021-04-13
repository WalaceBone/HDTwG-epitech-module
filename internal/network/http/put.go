package http

import (
	"HDTwG/internal/network"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutLocation(cmd network.PutCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		cmd(c.Request.Context())
		c.JSON(http.StatusOK, nil)
	}
}
