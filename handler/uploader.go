package handler

import (
	"net/http"
	"thunder_hoster/config"

	"github.com/gin-gonic/gin"
)

func UploadPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.tmpl", gin.H{
		"title": "File Uploader",
	})
}

func UploadHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "message.tmpl", gin.H{
			"title":       "Upload Failed",
			"message":     "Upload Failed",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	err = ctx.SaveUploadedFile(file, config.Cfg.FilePath)
	if err != nil {
		ctx.HTML(http.StatusInternalServerError, "message.tmpl", gin.H{
			"title":       "Upload Failed",
			"message":     "Failed when save file",
			"description": "Error: " + err.Error(),
			"color":       "red",
		})
		return
	}

	ctx.HTML(http.StatusOK, "message.tmpl", gin.H{
		"title":       "Upload Successed",
		"message":     "Upload Successed",
		"description": "",
		"color":       "green",
	})
}
