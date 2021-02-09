package match

import (
	"github.com/gin-gonic/gin"
	"match-go/app/model/mysql"
	"net/http"
)

func Create(ctx *gin.Context) {
	res := mysql.MatchSQL.GetAll()
	ctx.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
