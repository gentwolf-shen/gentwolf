package gentwolf

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type logger struct {
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
}

var Logger *logger

func init() {
	Logger = &logger{}

	flag := log.Ldate | log.Ltime | log.Llongfile

	Logger.Trace = log.New(os.Stdout, "  TRACE: ", flag)
	Logger.Info = log.New(os.Stdout, "   INFO: ", flag)
	Logger.Warning = log.New(os.Stdout, "WARNING: ", flag)

	dir := filepath.Dir(os.Args[0]) + "/log/"
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0666)
	}

	filename := dir + "error-" + time.Now().Format("2006-01-02") + ".txt"
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	Logger.Error = log.New(io.MultiWriter(file, os.Stderr), "  ERROR: ", flag)
}
