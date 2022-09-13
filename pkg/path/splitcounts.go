package path

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/response"
	"github.com/manojnakp/splitup/pkg/schema"
)

func GetAllSplitcounts(c *gin.Context) {
	list, err := FetchAllSplitcounts()
	if err != nil {
		response.InternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetSplitcount(c *gin.Context) {
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
	if got == false {
		response.NotFound(c)
		return
	}
	c.JSON(http.StatusOK, spl)
}

func UpdateSplitcount(c *gin.Context) {
	var minispl schema.MiniSplitCount
	id := c.Param("id")
	if id == "" {
		response.NotFound(c)
		return
	}
	if err := c.ShouldBindJSON(&minispl); err != nil {
		response.BadReq(c)
		return
	}
	ok := CheckBind(minispl)
	if ok {
		response.OK(c)
	} else {
		response.BadReq(c)
	}
}

func DelSplitcount(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.NotFound(c)
		return
	}
	response.OK(c)
}

func NewSplitcount(c *gin.Context) {
	var spl schema.SplitCount
	if err := c.ShouldBindJSON(&spl); err != nil {
		response.BadReq(c)
		return
	}
	ok := CheckBind(spl)
	if !ok {
		response.BadReq(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"id": "kb9t18wbboczf8xecdiy",
	})
}

func FetchAllSplitcounts() ([]schema.SplitCount, error) {
	spl, err := schema.DefaultSplitCount()
	if err != nil {
		return nil, err
	}
	return []schema.SplitCount{spl}, nil
}

func FetchSplitcount(id string) (schema.SplitCount, bool, error) {
	if id == "" {
		return schema.SplitCount{}, false, nil
	}
	spl, err := schema.DefaultSplitCount()
	if err != nil {
		return schema.SplitCount{}, false, err
	}
	return spl, true, nil
}

func routeSplitCounts(r *gin.Engine) {
	r.GET("/all", GetAllSplitcounts)
	r.GET("/:id", GetSplitcount)
	r.POST("/:id", UpdateSplitcount)
	r.DELETE("/:id", DelSplitcount)
	r.POST("/new", NewSplitcount)
}
