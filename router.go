package dghttp

import "github.com/gin-gonic/gin"

// NewRouter creates a new Gin Engine instance without any default middleware.
// This gives applications full control over their middleware stack.
//
// Applications should add their own middleware, for example:
//
//	router := http.NewRouter()
//	router.Use(http.RequestIDWithDefault())
//	router.Use(http.LoggerWithDefault())
//	router.Use(http.RecoveryWithDefault())
func NewRouter() *gin.Engine {
	// Create Gin engine without any default middleware
	// Using gin.New() instead of gin.Default() for full control
	return gin.New()
}
