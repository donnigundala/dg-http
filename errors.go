package dghttp

import "github.com/dgframe/core/errors"

// Standard transport errors aligned with dg-core Kernel Authority.
var (
	ErrBadRequest          = errors.ErrBadRequest
	ErrUnauthorized        = errors.ErrUnauthorized
	ErrForbidden           = errors.ErrForbidden
	ErrNotFound            = errors.ErrNotFound
	ErrConflict            = errors.ErrConflict
	ErrUnprocessableEntity = errors.ErrUnprocessable
	ErrInternalServerError = errors.ErrInternalServer

	// ErrRouterMissing is returned when an HTTP router capability is required but not provided.
	ErrRouterMissing = errors.New("dg-http: Router capability not provided (Type B violation)")
)
