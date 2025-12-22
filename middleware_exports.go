package dghttp

import (
	"time"

	"github.com/donnigundala/dg-core/logging"
	"github.com/donnigundala/dg-http/middleware"
	"github.com/gin-gonic/gin"
)

// Middleware exports - make middleware accessible from http package

// Recovery middleware
func Recovery(logger *logging.Logger) gin.HandlerFunc {
	return middleware.Recovery(logger)
}

func RecoveryWithDefault() gin.HandlerFunc {
	return middleware.RecoveryWithDefault()
}

// CORS middleware
func CORS(config middleware.CORSConfig) gin.HandlerFunc {
	return middleware.CORS(config)
}

func CORSWithDefault() gin.HandlerFunc {
	return middleware.CORSWithDefault()
}

func DefaultCORSConfig() middleware.CORSConfig {
	return middleware.DefaultCORSConfig()
}

// Security Headers middleware
func SecurityHeaders(config middleware.SecurityConfig) gin.HandlerFunc {
	return middleware.SecurityHeaders(config)
}

func SecurityHeadersWithDefault() gin.HandlerFunc {
	return middleware.SecurityHeadersWithDefault()
}

func DefaultSecurityConfig() middleware.SecurityConfig {
	return middleware.DefaultSecurityConfig()
}

// Rate Limit middleware
func RateLimit(config middleware.RateLimitConfig) gin.HandlerFunc {
	return middleware.RateLimit(config)
}

func RateLimitWithDefault() gin.HandlerFunc {
	return middleware.RateLimitWithDefault()
}

func DefaultRateLimitConfig() middleware.RateLimitConfig {
	return middleware.DefaultRateLimitConfig()
}

// Timeout middleware
func Timeout(timeout time.Duration) gin.HandlerFunc {
	return middleware.Timeout(timeout)
}

// Body Size Limit middleware
func BodySizeLimit(maxBytes int64) gin.HandlerFunc {
	return middleware.BodySizeLimit(maxBytes)
}

func BodySizeLimitWithError(maxBytes int64) gin.HandlerFunc {
	return middleware.BodySizeLimitWithError(maxBytes)
}

// Compress middleware
func Compress(config middleware.CompressConfig) gin.HandlerFunc {
	return middleware.Compress(config)
}

func CompressWithDefault() gin.HandlerFunc {
	return middleware.CompressWithDefault()
}

func DefaultCompressConfig() middleware.CompressConfig {
	return middleware.DefaultCompressConfig()
}

// AccessLogger middleware
func AccessLogger(config middleware.LoggerConfig) gin.HandlerFunc {
	return middleware.Logger(config)
}

func LoggerWithDefault() gin.HandlerFunc {
	return middleware.LoggerWithDefault()
}

func DefaultLoggerConfig() middleware.LoggerConfig {
	return middleware.DefaultLoggerConfig()
}

// RequestID middleware
func RequestID() gin.HandlerFunc {
	return middleware.RequestID()
}

func RequestIDWithDefault() gin.HandlerFunc {
	return middleware.RequestIDWithDefault()
}

// Observability middleware
func Observability(config middleware.ObservabilityConfig) gin.HandlerFunc {
	return middleware.Observability(config)
}

func ObservabilityWithDefault() gin.HandlerFunc {
	return middleware.Observability(middleware.ObservabilityConfig{})
}
