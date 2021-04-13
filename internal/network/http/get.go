package http

import (
	"HDTwG/internal/network"
	"HDTwG/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSale ...
func GetLocation(cmd network.GetCmd) gin.HandlerFunc {
	return func(c *gin.Context) {
		ipAddress := c.Request.URL.Query().Get("ip")
		lang := c.Request.URL.Query().Get("lang")

		if ipAddress == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		location, err := cmd(c.Request.Context(), store.Options{IP: ipAddress, Lang: lang})
		//TODO
		if err != nil {
			switch err {

			default:
				c.Status(http.StatusInternalServerError)
				return
			}
		}
		c.JSON(http.StatusOK, location)
	}
}
