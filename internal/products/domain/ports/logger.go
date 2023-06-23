package ports

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	InfoP(format string, args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
}
