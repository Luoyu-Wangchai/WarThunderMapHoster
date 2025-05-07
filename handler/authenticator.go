package handler

import (
	"fmt"
	"net/http"
	"thunder_hoster/config"
	"time"

	"github.com/gin-gonic/gin"
)

var ValidTime time.Time

func PasswordAuthenticator(ctx *gin.Context) {
	passwd := ctx.PostForm("password")

	if passwd == config.Cfg.Password {
		ValidTime = time.Now().Add(time.Duration(config.Cfg.ValidMin) * time.Minute)
		ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
			"title":       "Verification Successed",
			"message":     "Verification Successed",
			"description": fmt.Sprintf("Now you can visit %s/map , before %s", ctx.Request.Host, ValidTime.Format("2006-01-02 15:04:05")),
			"color":       "green",
		})
	} else {
		ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
			"title":       "Verification Failed",
			"message":     "Verification Failed",
			"description": "Please try again",
			"color":       "red",
		})
	}

}
