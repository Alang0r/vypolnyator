package service

import (
	"github.com/Alang0r/vypolnyator/pkg/error"
	"github.com/gin-gonic/gin"
)

type Request interface {
	Execute(c *gin.Context) (Reply, error.Error)
	GetService()
}

type Reply interface {
}
