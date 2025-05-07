package handler

import (
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
		ctx.String(http.StatusOK, "Verification has been passed.")
	} else {
		ctx.String(http.StatusOK, "Verification failed.")
	}

}
