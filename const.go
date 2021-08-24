package eyes

const (
	DEBUG int = iota
	INFO
	WARN
	ERROR
	FATAL
)

const (
	DateFormat        = "2006-01-02 15:04:05"
	DefaultFilePrefix = "kite" // "kite" is just an example, please modify it as your project name.
	DefaultFileSuffix = "log"
	DefaultPath       = "logs"
	IsFile            = true
	TimeZone          = "Asia/Shanghai"
)
