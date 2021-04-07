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

		ipAddress := c.Request.URL.Query().Get("network")
		if ipAddress == "" {
			c.Status(http.StatusBadRequest)
			return
		}

		location, _ := cmd(c.Request.Context(), store.Options{IP: ipAddress, Lang: ""})
		//TODO
		/*if err != nil {
			switch err {
			case model.ErrNotFound:
				c.Status(http.StatusNotFound)
				return
			default:
				c.Status(http.StatusInternalServerError)
				return
			}
		}*/
		c.JSON(http.StatusOK, location)
	}
}
