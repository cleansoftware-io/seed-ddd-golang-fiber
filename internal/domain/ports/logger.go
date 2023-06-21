package ports

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Error(args ...interface{})
	Warning(args ...interface{})
	Fatal(args ...interface{})
}
