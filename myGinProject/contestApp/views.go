package contestApp

import "github.com/gin-gonic/gin"

func hello(c *gin.Context)  {
	c.String(200,"Hello ContestApp")
}