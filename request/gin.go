package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Param returns a Gin path parameter value.
//
// Example:
//
//	router.GET("/users/:id", func(c *gin.Context) {
//	    id := request.Param(c, "id")
//	})
func Param(c *gin.Context, key string) string {
	return c.Param(key)
}

// ParamInt returns a Gin path parameter as an integer.
func ParamInt(c *gin.Context, key string, defaultValue int) int {
	value := c.Param(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// ParamInt64 returns a Gin path parameter as an int64.
func ParamInt64(c *gin.Context, key string, defaultValue int64) int64 {
	value := c.Param(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// Query returns a Gin query parameter with a default value.
//
// Example:
//
//	page := request.Query(c, "page", "1")
func Query(c *gin.Context, key, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// QueryInt returns a Gin query parameter as an integer.
func QueryInt(c *gin.Context, key string, defaultValue int) int {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// QueryInt64 returns a Gin query parameter as an int64.
func QueryInt64(c *gin.Context, key string, defaultValue int64) int64 {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// QueryBool returns a Gin query parameter as a boolean.
func QueryBool(c *gin.Context, key string, defaultValue bool) bool {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}

// QueryFloat returns a Gin query parameter as a float64.
func QueryFloat(c *gin.Context, key string, defaultValue float64) float64 {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}

// QueryArray returns all values for a Gin query parameter.
func QueryArray(c *gin.Context, key string) []string {
	return c.QueryArray(key)
}

// PostForm returns a Gin form value with a default.
func PostForm(c *gin.Context, key, defaultValue string) string {
	value := c.PostForm(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// PostFormInt returns a Gin form value as an integer.
func PostFormInt(c *gin.Context, key string, defaultValue int) int {
	value := c.PostForm(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

// PostFormBool returns a Gin form value as a boolean.
func PostFormBool(c *gin.Context, key string, defaultValue bool) bool {
	value := c.PostForm(key)
	if value == "" {
		return defaultValue
	}
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return boolValue
}
