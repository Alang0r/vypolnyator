package middleware

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/gin-gonic/gin"
)

func OldMiddleware(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		l.Infof("Incoming Request <%s> : %s", c.Query("reqID"), c.Query("request"))
		defer l.Infof("Done Request %s", c.Query("request"))
		c.Next()
	}
}

func NewMiddleware(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		rName := c.Request.Header.Get("Request")
		rID := c.Request.Header.Get("ReqID")
		l.Infof("new request id:%s, reqName:%s", rID, rName)
		c.Next()
	}
}
