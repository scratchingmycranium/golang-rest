package utils

import "github.com/gin-gonic/gin"

type RouterSettings struct {
	// ...
}

func (u *utils) InitRouter() *gin.Engine {
	r := gin.Default()
	return r
}
