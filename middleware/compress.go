package middleware

import (
	"compress/gzip"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
)

// CompressConfig defines the configuration for compression middleware.
type CompressConfig struct {
	// Level is the compression level (0-9).
	// Default: gzip.DefaultCompression
	Level int

	// MinSize is the minimum response size to compress (in bytes).
	// Default: 1024 (1KB)
	MinSize int

	// ContentTypes is a list of content types to compress.
	// Default: ["text/", "application/json", "application/javascript", "application/xml"]
	ContentTypes []string
}

// DefaultCompressConfig returns the default compression configuration.
func DefaultCompressConfig() CompressConfig {
	return CompressConfig{
		Level:   gzip.DefaultCompression,
		MinSize: 1024,
		ContentTypes: []string{
			"text/",
			"application/json",
			"application/javascript",
			"application/xml",
		},
	}
}

// Compress returns a middleware that compresses HTTP responses.
func Compress(config CompressConfig) gin.HandlerFunc {
	// Apply defaults
	if config.Level == 0 {
		config.Level = gzip.DefaultCompression
	}
	if config.MinSize == 0 {
		config.MinSize = 1024
	}
	if len(config.ContentTypes) == 0 {
		config.ContentTypes = []string{
			"text/",
			"application/json",
			"application/javascript",
			"application/xml",
		}
	}

	return func(c *gin.Context) {
		// Check if client accepts gzip
		if !strings.Contains(c.Request.Header.Get("Accept-Encoding"), "gzip") {
			c.Next()
			return
		}

		// Create gzip writer
		gz, err := gzip.NewWriterLevel(c.Writer, config.Level)
		if err != nil {
			c.Next()
			return
		}
		defer gz.Close()

		// Wrap response writer
		grw := &gzipResponseWriter{
			ResponseWriter: c.Writer,
			Writer:         gz,
			config:         config,
		}

		// Set Content-Encoding header
		c.Header("Content-Encoding", "gzip")
		c.Header("Vary", "Accept-Encoding") // Using Vary is best practice
		c.Writer = grw

		c.Next()

		// Remove Content-Length as it's invalid for chunked/compressed responses usually
		c.Writer.Header().Del("Content-Length")
	}
}

// CompressWithDefault returns a Compress middleware with default configuration.
func CompressWithDefault() gin.HandlerFunc {
	return Compress(DefaultCompressConfig())
}

// gzipResponseWriter wraps gin.ResponseWriter to compress the response.
// Note: We embed gin.ResponseWriter to implement the interface.
type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
	config CompressConfig
}

func (w *gzipResponseWriter) Write(b []byte) (int, error) {
	// Check if we should compress based on content type
	contentType := w.Header().Get("Content-Type")
	shouldCompress := false

	if len(w.config.ContentTypes) == 0 {
		shouldCompress = true
	} else {
		for _, ct := range w.config.ContentTypes {
			if strings.HasPrefix(contentType, ct) {
				shouldCompress = true
				break
			}
		}
	}

	// If content type doesn't match or size is too small, write directly
	if !shouldCompress || len(b) < w.config.MinSize {
		return w.ResponseWriter.Write(b)
	}

	// Write compressed
	return w.Writer.Write(b)
}

func (w *gzipResponseWriter) WriteString(s string) (int, error) {
	return w.Write([]byte(s))
}
