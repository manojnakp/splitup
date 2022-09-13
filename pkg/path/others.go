package path

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/response"
	"github.com/manojnakp/splitup/pkg/schema"
)

func GetBalances(c *gin.Context) {
	id := c.Param("id")
	balances, got, err := FetchBalances(id)
	if err != nil {
		response.InternalServerError(c)
		return
	}
	if got == false {
		response.NotFound(c)
		return
	}
	c.JSON(http.StatusOK, balances)
}

func FetchBalances(id string) ([]schema.BalanceUnit, bool, error) {
	if id == "" {
		return nil, false, nil
	}
	balance, err := schema.DefaultBalance()
	if err != nil {
		return nil, false, err
	}
	return []schema.BalanceUnit{balance}, true, nil
}

func CheckBind(jsonobj any) bool {
	v := reflect.ValueOf(jsonobj)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.IsZero() {
			return false
		}
	}
	return true
}

func routeOthers(r *gin.Engine) {
	r.GET("/:id/balance", GetBalances)
}
