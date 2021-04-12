package http

import (
	"HDTwG/internal/network"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutLocation(cmd network.PutCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := cmd(c.Request.Context())
		if err != nil {
			//TODO
			// switch err {
			// case model.ErrNotFound:
			// 	c.Status(http.StatusNotFound)
			// 	return
			// default:
			// 	c.Status(http.StatusInternalServerError)
			// 	return
			// }
		}
		c.JSON(http.StatusOK, err)
	}
}
