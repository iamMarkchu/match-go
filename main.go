package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"match-go/app/http"
	"match-go/app/model"
)

func main() {
	model.Init()
	r := gin.New()
	http.SetRouters(r)
	if err := r.Run(":8888"); err != nil {
		log.Fatal("gin RUN error:", err.Error())
	}
}
