/*
Package http provides the web server implementation for the framework.
It includes the Router, Kernel, Server, and Middleware chain.

# Key Features

  - Gin-based high-performance router
  - Middleware support (Global, Group, Route-specific)
  - Graceful shutdown
  - Request/Response helpers

# Basic Usage

	Create a new router (returns *gin.Engine):

		router := http.NewRouter()

	Register routes:

		router.GET("/ping", func(c *gin.Context) {
		    c.String(200, "pong")
		})

	Start the server:

		router.Run(":8080")

	# Middleware

	Middleware are standard `gin.HandlerFunc` functions:

		router.Use(middleware.LoggerWithDefault())
		router.Use(middleware.RecoveryWithDefault())

	# Groups

	Group routes using Gin's standard grouping:

		api := router.Group("/api/v1")
		api.Use(AuthMiddleware)
		{
		    api.GET("/users", GetUsers)
		}
*/
package dghttp
