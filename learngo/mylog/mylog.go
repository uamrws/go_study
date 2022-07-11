package mylog

// 自定义一个日志库
// 日志级别

const (
	DEBUG = iota
	TRACE
	INFO
	WARN
	ERROR
	FATAL
)

type Logger interface {
	Log(level string, message string)
	Debug(message string)
	Trace(message string)
	Info(message string)
	Warn(message string)
	Error(message string)
	Fatal(message string)
}
