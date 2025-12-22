package dghttp

// Logger is the interface for structured logging.
// It is designed to be compatible with dg-core/logging.Logger.
type Logger interface {
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	With(args ...interface{}) Logger
}
