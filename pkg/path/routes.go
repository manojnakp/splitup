package path

import (
	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	routeOthers(r)
	routeSplitCounts(r)
	routeParticipants(r)
	routeExpenses(r)
}
