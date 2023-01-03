package service

import (
	"github.com/gin-gonic/gin"
	"github.com/Alang0r/vypolnyator/pkg/error"
)

type Request interface {
	Execute(c *gin.Context) (Reply, error.Error)
}

type Reply interface {
}
