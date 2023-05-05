package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

type dailyFilterWriter struct {
	// 日志文件名称
	fileName string
	// 上次写入的日期
	lastVarDay int
	// 输出文件
	outputFile *os.File
	// 文件交换锁
	fileSwitchLock *sync.Mutex
}

func (w *dailyFilterWriter) Write(byteArray []byte) (n int, err error) {
	if nil == byteArray || len(byteArray) <= 0 {
		return 0, nil
	}
	outputFile, err := w.getOutputFile()
	if nil != err {
		return 0, err
	}
	_, _ = os.Stderr.Write(byteArray)
	_, _ = outputFile.Write(byteArray)
	return len(byteArray), nil
}

/**
获取输出问价你，每天创建一个新的日志文件
*/
func (w *dailyFilterWriter) getOutputFile() (io.Writer, error) {
	yearDay := time.Now().YearDay()
	if w.lastVarDay == yearDay && w.outputFile != nil {
		return w.outputFile, nil
	}

	w.fileSwitchLock.Lock()
	defer w.fileSwitchLock.Unlock()
	w.lastVarDay = yearDay
	if err := os.MkdirAll(path.Dir(w.fileName), os.ModePerm); err != nil {
		return nil, err
	}

	newDalyFile := w.fileName + "." + time.Now().Format("20060102")
	outfile, err := os.OpenFile(newDalyFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if nil != err || nil == outfile {
		return nil, fmt.Errorf("打开文件%s失败,error=%v", newDalyFile, err)
	}
	if nil != w.outputFile {
		// 关闭原来的文件
		_ = w.outputFile.Close()
	}
	w.outputFile = outfile
	return outfile, nil
}
