package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/donnigundala/dg-core/contracts/resilience"
	"github.com/gin-gonic/gin"
)

// ResilienceMiddleware provides a way to wrap routes with a resilience policy
type ResilienceMiddleware struct {
	manager resilience.Manager
}

func NewResilienceMiddleware(manager resilience.Manager) *ResilienceMiddleware {
	return &ResilienceMiddleware{manager: manager}
}

// Wrap returns a Gin middleware that executes the next handler inside the given policy
// This is useful for Circuit Breakers and Bulkheads on specific routes
func (rm *ResilienceMiddleware) Wrap(policyName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// We need to resolve the policy dynamically because it might be a Retry, CB, or Bulkhead
		// However, the common interface Runner allows us to execute it uniformally.

		// For simplicity in this middleware, we assume the user configured a CircuitBreaker
		// or Bulkhead with this name. A more advanced version might take the policy type.
		// Here we try to find a Runner by checking the types.

		var runner resilience.Runner

		// Try Circuit Breaker
		cb := rm.manager.CircuitBreaker(policyName)
		if cb != nil {
			runner = cb
		} else {
			// Try Bulkhead
			bh := rm.manager.Bulkhead(policyName)
			if bh != nil {
				runner = bh
			}
		}

		if runner == nil {
			// No policy found, proceed without resilience
			c.Next()
			return
		}

		_, err := runner.Execute(c.Request.Context(), func(ctx context.Context) (interface{}, error) {
			c.Next()

			// Check if the handler set an error in Gin context
			if len(c.Errors) > 0 {
				return nil, c.Errors.Last()
			}

			// Check status code for failure detection
			if c.Writer.Status() >= 500 {
				return nil, errors.New("internal server error")
			}

			return nil, nil
		})

		if err != nil {
			// If resilience failed (e.g. CB open, Bulkhead full), abort with appropriate error
			// Note: If the error came from the handler (c.Next), it's already written if Execute propagated it.
			// We mainly care about resilience errors preventing execution.

			// If the header hasn't been written yet:
			if !c.Writer.Written() {
				c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable", "detail": err.Error()})
			}
		}
	}
}
