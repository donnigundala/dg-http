package contracts

import (
	"github.com/donnigundala/dg-core/errors"
)

// Standard transport errors aligned with dg-core Kernel Authority.
var (
	ErrBadRequest          = errors.ErrBadRequest
	ErrUnauthorized        = errors.ErrUnauthorized
	ErrForbidden           = errors.ErrForbidden
	ErrNotFound            = errors.ErrNotFound
	ErrConflict            = errors.ErrConflict
	ErrUnprocessableEntity = errors.ErrUnprocessable
	ErrInternalServerError = errors.ErrInternalServer
)
