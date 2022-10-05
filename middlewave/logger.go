package middlewave

import (
	"github.com/BaiMeow/HduHelpLogin/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func Logger() func(ctx *gin.Context) {
	l := log.Logger.WithField("type", "gin")
	return func(r *gin.Context) {
		// Start timer
		start := time.Now()
		path := r.Request.URL.Path
		raw := r.Request.URL.RawQuery

		// Process request
		r.Next()
		// Stop timer
		n := time.Now()
		if raw != "" {
			path = path + "?" + raw
		}
		e := l.WithFields(logrus.Fields{
			"timeStamp":  n.Format("2006/01/02 - 15:04:05"),
			"Latency":    n.Sub(start),
			"statusCode": r.Writer.Status(),
			"method":     r.Request.Method,
			"clientIP":   r.ClientIP(),
			"bodySize":   r.Writer.Size(),
			"path":       path,
			"traceId":    r.Value("traceId"),
		})

		if r.Errors != nil {
			e.Error(r.Errors.String())
		} else {
			e.Info()
		}
	}
}
