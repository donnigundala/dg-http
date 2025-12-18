package response

import (
	"github.com/donnigundala/dg-core/errors"
	"github.com/gin-gonic/gin"
)

// JSON writes a JSON response using Gin.
//
// Example:
//
//	response.JSON(c, 200, gin.H{"message": "success"})
func JSON(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

// Success writes a successful JSON response using Gin.
//
// Example:
//
//	response.Success(c, user, "User created successfully")
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(200, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created writes a 201 Created response using Gin.
//
// Example:
//
//	response.Created(c, user, "/api/users/123")
func Created(c *gin.Context, data interface{}, location string) {
	if location != "" {
		c.Header("Location", location)
	}
	c.JSON(201, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

// NoContent writes a 204 No Content response using Gin.
func NoContent(c *gin.Context) {
	c.Status(204)
}

// Accepted writes a 202 Accepted response using Gin.
func Accepted(c *gin.Context, data interface{}, message string) {
	c.JSON(202, SuccessResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error writes an error response using Gin.
//
// Example:
//
//	response.Error(c, err, 500)
func Error(c *gin.Context, err error, status int) {
	// Check if it's our custom Error type
	if e, ok := err.(*errors.Error); ok {
		c.JSON(e.StatusCode(), ErrorResponse{
			Success: false,
			Error:   e.Message(),
			Code:    e.Code(),
			Fields:  e.Fields(),
		})
		return
	}

	// Default error response
	c.JSON(status, ErrorResponse{
		Success: false,
		Error:   err.Error(),
	})
}

// ValidationError writes a validation error response using Gin.
func ValidationError(c *gin.Context, validationErrors interface{}) {
	c.JSON(422, ErrorResponse{
		Success: false,
		Error:   "Validation failed",
		Code:    "VALIDATION_ERROR",
		Fields:  validationErrors,
	})
}

// NotFound writes a 404 Not Found response using Gin.
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Resource not found"
	}
	c.JSON(404, ErrorResponse{
		Success: false,
		Error:   message,
		Code:    "NOT_FOUND",
	})
}

// Unauthorized writes a 401 Unauthorized response using Gin.
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Unauthorized"
	}
	c.JSON(401, ErrorResponse{
		Success: false,
		Error:   message,
		Code:    "UNAUTHORIZED",
	})
}

// Forbidden writes a 403 Forbidden response using Gin.
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Forbidden"
	}
	c.JSON(403, ErrorResponse{
		Success: false,
		Error:   message,
		Code:    "FORBIDDEN",
	})
}

// BadRequest writes a 400 Bad Request response using Gin.
func BadRequest(c *gin.Context, message string) {
	if message == "" {
		message = "Bad request"
	}
	c.JSON(400, ErrorResponse{
		Success: false,
		Error:   message,
		Code:    "BAD_REQUEST",
	})
}

// InternalServerError writes a 500 Internal Server Error response using Gin.
func InternalServerError(c *gin.Context, message string) {
	if message == "" {
		message = "Internal server error"
	}
	c.JSON(500, ErrorResponse{
		Success: false,
		Error:   message,
		Code:    "INTERNAL_ERROR",
	})
}

// Paginated writes a paginated response using Gin.
//
// Example:
//
//	response.Paginated(c, users, 1, 10, 100)
func Paginated(c *gin.Context, data interface{}, page, perPage, total int) {
	c.JSON(200, PaginatedResponse{
		Success: true,
		Data:    data,
		Meta: PaginationMeta{
			Page:       page,
			PerPage:    perPage,
			Total:      total,
			TotalPages: (total + perPage - 1) / perPage,
		},
	})
}

// AbortWithError aborts the request chain and writes an error response.
//
// Example:
//
//	response.AbortWithError(c, errors.New("unauthorized"), 401)
func AbortWithError(c *gin.Context, err error, status int) {
	Error(c, err, status)
	c.Abort()
}

// AbortWithJSON aborts the request chain and writes a JSON response.
func AbortWithJSON(c *gin.Context, status int, data interface{}) {
	c.AbortWithStatusJSON(status, data)
}
