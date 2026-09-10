package logger

type LoggerOpts struct{}
type Logger struct{}

func New() Logger {
	return Logger{}
}
