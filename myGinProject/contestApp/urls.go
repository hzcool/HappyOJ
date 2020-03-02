package contestApp

import "github.com/gin-gonic/gin"

func urls_init(r *gin.Engine) {
	r.GET("/contest",hello)
}