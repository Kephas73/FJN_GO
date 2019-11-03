package main

import (
	"Common"
	"bytes"
	"github.com/goframework/gf/exterror"
	"io"
	"log"
	"os"
	"path/filepath"
)

const MAX_LOG_BUF = 1024 * 1024 // 1MB

var gLog log.Logger
var gLogFile string
var gLogBuf = bytes.NewBuffer(nil)

type BufWriter struct {
	destWriter io.Writer
	bufWriter  *bytes.Buffer
}

func (b *BufWriter) Write(p []byte) (n int, err error) {
	if b.bufWriter.Len() < MAX_LOG_BUF {
		b.bufWriter.Write(p)
	}
	return b.destWriter.Write(p)
}

func init() {
	gLog = log.Logger{}
	gLog.SetFlags(log.Ldate | log.Ltime)
	if Common.IsWindowsRuntime() {
		gLog.SetOutput(os.Stdout)
	} else {
		exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		folderLog := config.Str(CFG_FOLDER_LOG, "LogFile")
		logPath := filepath.Join(exeDir, folderLog)
		os.MkdirAll(logPath, os.ModePerm)
		gLogFile = filepath.Join(logPath, Common.GetDateTimeCurrent()+FILE_EXTENSION_LOG)
		fLog, err := os.OpenFile(gLogFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		if err != nil {
			log.Fatalln(exterror.WrapExtError(err))
		}

		gLog.SetOutput(&BufWriter{bufWriter: gLogBuf, destWriter: fLog})
	}
	exterror.SetFilePathSpliter("RELEASE")
}
