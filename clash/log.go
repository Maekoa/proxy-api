package clash

// log
type Event struct {
	LogLevel LogLevel
	Payload  string
}
type LogLevel uint8

const (
	DEBUG LogLevel = iota
	INFO
	WARNING
	ERROR
	SILENT
)

var LogLevelName = map[LogLevel]string{
	DEBUG:   "debug",
	INFO:    "info",
	WARNING: "warning",
	ERROR:   "error",
	SILENT:  "silent",
}

func (l LogLevel) String() string {
	return LogLevelName[l]
}
