package middleware

import (
	"github.com/Alang0r/vypolnyator/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewMiddleware(l *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// generate uuid
		uuid := uuid.New()

		c.Set("uuid", uuid)
		l.Infof("Incoming Request <%s> : %s", uuid.String(), c.Query("request"))
		c.Next()
		l.Infof("Done Request %s", uuid.String())
	}
}
