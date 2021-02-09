package http

import (
	"github.com/gin-gonic/gin"
	"match-go/app/http/match"
	"net/http"
)

func SetRouters(r *gin.Engine) {
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello")
	})
	// match
	matchGroup := r.Group("/api/match")
	{
		matchGroup.POST("/create", match.Create) // 创建比赛
	}
}
