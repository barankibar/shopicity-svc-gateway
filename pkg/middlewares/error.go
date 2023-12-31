package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		log.Fatal(err)
	}
	c.JSON(http.StatusInternalServerError, "")
}
