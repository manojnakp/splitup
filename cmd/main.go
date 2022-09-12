package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manojnakp/splitup/pkg/path"
)

func main() {
	r := gin.Default()
	path.Route(r)
	r.Run()
}
