package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/southwind/ainews/lib"
	"github.com/unknwon/com"
)

func GetValueFromMapWithDefaultValue(maps map[string]string, key string, defaultValue interface{}) (result interface{}) {

	value, ok := maps[key]

	if ok {
		result = value
		return
	}
	result = defaultValue
	return
}

func Getpage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * lib.LoadServerConfig().PageSize
	}
	return result
}
