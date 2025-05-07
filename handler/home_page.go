package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html")
	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "authenticator",
	})

}
