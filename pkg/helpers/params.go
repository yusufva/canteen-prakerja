package helpers

import (
	"canteen-prakerja/pkg/custerrs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamsId(c *gin.Context, key string) (int, custerrs.MessageErr) {
	value := c.Param(key)

	id, err := strconv.Atoi(value)

	if err != nil {
		return 0, custerrs.NewBadRequest("invalid parameter id")
	}

	return id, nil
}
