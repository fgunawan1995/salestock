package main
import (
	"github.com/gin-gonic/gin"
)
//Activate CORS (so it can be called from anywhere
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}