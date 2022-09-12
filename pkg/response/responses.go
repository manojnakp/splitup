package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/schema"
)

type StatusOK struct {
	ok bool
}

func newError(msg string) schema.ErrorResponse {
	e := schema.ErrorString{msg}
	return schema.ErrorResponse{e}
}

func BadReq(c *gin.Context) {
	res := newError("Invalid API request: malformed request syntax")
	c.JSON(http.StatusBadRequest, res)
}

func InternalServerError(c *gin.Context) {
	res := newError("Internal Server Error")
	c.JSON(http.StatusInternalServerError, res)
}

func NotFound(c *gin.Context) {
	res := newError("Invalid API request: record not found")
	c.JSON(http.StatusNotFound, res)
}

func OK(c *gin.Context) {
	res := StatusOK{true}
	c.JSON(http.StatusOK, res)
}
