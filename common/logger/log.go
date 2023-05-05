package logger

import (
	"fmt"
	"log"
	"sync"
)

var writer *dailyFilterWriter

func Config(outputFile string) {
	writer = &dailyFilterWriter{
		fileName:       outputFile,
		lastVarDay:     -1,
		fileSwitchLock: &sync.Mutex{},
	}
	infoLogger = log.New(writer, "[ INFO ]", log.Ltime|log.Lmicroseconds|log.Lshortfile)
	errorLogger = log.New(writer, "[ ERROR ]", log.Ltime|log.Lmicroseconds|log.Lshortfile)

}

var infoLogger, errorLogger *log.Logger

func Info(format string, valArray ...interface{}) {
	_ = infoLogger.Output(2, fmt.Sprintf(format, valArray...))
}

func Error(format string, valArray ...interface{}) {
	_ = errorLogger.Output(2, fmt.Sprintf(format, valArray...))
}
