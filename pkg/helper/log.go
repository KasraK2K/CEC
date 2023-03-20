package helper

import (
	"fmt"
	"log"
	"os"
)

type logger struct{}

var Logger logger

/* -------------------------------------------------------------------------- */
/*                               Stdout Loggers                               */
/* -------------------------------------------------------------------------- */
func (l *logger) Info(v ...interface{}) {
	l.initLogger("\u001b[32;1mINFO\x1b[0m:    ", v...)
}

func (l *logger) Verbose(v ...interface{}) {
	l.initLogger("\u001b[34;1mVERBOSE\x1b[0m: ", v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.initLogger("\u001b[38;5;226mWARNING\x1b[0m: ", v...)
}

func (l *logger) Error(v ...interface{}) {
	l.initLogger("\x1b[31mERROR\x1b[0m:   ", v...)
}

/* -------------------------------------------------------------------------- */
/*                                File Loggers                                */
/* -------------------------------------------------------------------------- */
func (l *logger) InfoFile(v ...interface{}) {
	l.initFileLogger("info", v...)
}

func (l *logger) VerboseFile(v ...interface{}) {
	l.initFileLogger("verbose", v...)
}

func (l *logger) WarningFile(v ...interface{}) {
	l.initFileLogger("warning", v...)
}

func (l *logger) ErrorFile(v ...interface{}) {
	l.initFileLogger("error", v...)
}

/* -------------------------------------------------------------------------- */
/*                              Private Functions                             */
/* -------------------------------------------------------------------------- */
func (l *logger) folderExistence() {
	if _, err := os.Stat("log"); os.IsNotExist(err) {
		os.Mkdir("log", os.ModePerm)
	}
}

func (l *logger) initLogger(prefix string, v ...interface{}) {
	flags := log.Ldate | log.Ltime
	errorLogger := log.New(os.Stdout, prefix, flags)
	errorLogger.Println(v...)
}

func (l *logger) initFileLogger(fileName string, v ...interface{}) {
	l.folderExistence()
	flags := log.Ldate | log.Ltime
	file, _ := os.OpenFile(fmt.Sprintf("log/%s.log", fileName), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	infoFileLogger := log.New(file, "", flags)
	infoFileLogger.SetOutput(file)
	infoFileLogger.Println(v...)
}
