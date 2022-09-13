package path

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/response"
)

func GetParticipants(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NotFound(c)
		return
	}
	spl, got, err := FetchSplitcount(id)
	if err != nil {
		response.InternalServerError(c)
		return
	}
	if !got {
		response.NotFound(c)
		return
	}
	pslice := spl.Participants
	c.JSON(http.StatusOK, pslice)
}

func UpdateParticipants(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NotFound(c)
		return
	}
	var pslice []string
	err := c.ShouldBindJSON(&pslice)
	if err != nil {
		response.BadReq(c)
		return
	}
	response.OK(c)
}

func AddParticipant(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NotFound(c)
		return
	}
	var pstruct struct {
		Participant string `json:"participant"`
	}
	err := c.ShouldBindJSON(&pstruct)
	if err != nil {
		response.BadReq(c)
		return
	}
	ok := CheckBind(pstruct)
	if !ok {
		response.BadReq(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"id": "kb9t18wbboczf8xecdiy",
	})
}

func routeParticipants(r *gin.Engine) {
	r.GET("/:id/participants", GetParticipants)
	r.POST("/:id/participants", UpdateParticipants)
	r.POST("/:id/participants/new", AddParticipant)
}
