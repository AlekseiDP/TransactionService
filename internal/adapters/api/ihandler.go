package api

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(engine *gin.Engine)
}
