package dghttp

// Logger is the interface for structured logging.
// It is designed to be compatible with slog.Logger and dg-core/logging.Logger.
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	With(args ...any) Logger
}
