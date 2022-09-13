package path

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/response"
	"github.com/manojnakp/splitup/pkg/schema"
)

func GetAllExpenses(c *gin.Context) {
	list, err := FetchAllExpenses()
	if err != nil {
		response.InternalServerError(c)
		return
	}
	c.JSON(http.StatusOK, list)
}

func GetExpense(c *gin.Context) {
	id := c.Param("id")
	eid := c.Param("eid")
	if id == "" || eid == "" {
		response.NotFound(c)
		return
	}
	xp, got, err := FetchExpense(id, eid)
	if err != nil {
		response.InternalServerError(c)
	}
	if !got {
		response.NotFound(c)
	}
	c.JSON(http.StatusOK, xp)
}

func UpdateExpense(c *gin.Context) {
	var xp schema.ExpenseUnit
	id := c.Param("id")
	eid := c.Param("eid")
	if id == "" || eid == "" {
		response.NotFound(c)
		return
	}
	if err := c.ShouldBindJSON(&xp); err != nil {
		response.BadReq(c)
		return
	}
	ok := CheckBind(xp)
	if !ok {
		response.BadReq(c)
		return
	}
	response.OK(c)
}

func DelExpense(c *gin.Context) {
	id := c.Param("id")
	eid := c.Param("eid")
	if id == "" || eid == "" {
		response.NotFound(c)
		return
	}
	response.OK(c)
}

func NewExpense(c *gin.Context) {
	var xp schema.ExpenseUnit
	if err := c.ShouldBindJSON(&xp); err != nil {
		response.BadReq(c)
		return
	}
	ok := CheckBind(xp)
	if !ok {
		response.BadReq(c)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
		"id": "31",
	})
}

func FetchAllExpenses() ([]schema.ExpenseUnit, error) {
	xp, err := schema.DefaultExpense()
	if err != nil {
		return nil, err
	}
	return []schema.ExpenseUnit{xp}, nil
}

func FetchExpense(id, eid string) (schema.ExpenseUnit, bool, error) {
	if id == "" || eid == "" {
		return schema.ExpenseUnit{}, false, nil
	}
	xp, err := schema.DefaultExpense()
	if err != nil {
		return schema.ExpenseUnit{}, false, err
	}
	return xp, true, nil
}

func routeExpenses(r *gin.Engine) {
	r.GET("/:id/expenses/all", GetAllExpenses)
	r.POST("/:id/expenses/new", NewExpense)
	r.GET("/:id/expenses/:eid", GetExpense)
	r.POST("/:id/expenses/:eid", UpdateExpense)
	r.DELETE("/:id/expenses/:eid", DelExpense)
}
