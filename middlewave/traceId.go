package middlewave

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTraceId(r *gin.Context) {
	traceId := uuid.New()
	r.Set("traceId", traceId)
}
