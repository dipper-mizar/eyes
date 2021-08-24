package eyes

import (
	"io"
	"os"
	"time"
)

type Logger struct {
	Level    int
	IsFile   bool
	TimeZone *time.Location

	OutConsole io.Writer

	FilePrefix       string
	Path             string
	AbsoluteFileName string
	LogFile          *os.File
	FileWriter       io.Writer
}
