package main

import (
	"github.com/goframework/gf/ext"
	"github.com/goframework/gf/exterror"
	"os"
	"path/filepath"
)

type OKLogger struct {
	JobName string
	Date    string
}

// Check file exits
func (this *OKLogger) CheckExits() bool {
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	folderLog := config.Str(CFG_OK_FOLDER_LOG,"OK_LOG")
	filePath  := filepath.Join(exeDir, folderLog, this.JobName,this.Date + FILE_EXTENSION_OKLOG)

	if ext.FileExists(filePath) {
		return true
	}
	return  false
}

func (this *OKLogger) Set() error {
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	folderLog := config.Str(CFG_OK_FOLDER_LOG,"OK_LOG")
	filePath  := filepath.Join(exeDir, folderLog, this.JobName,this.Date + FILE_EXTENSION_OKLOG)
	fileDir   := filepath.Dir(filePath)
	os.MkdirAll(fileDir, os.ModePerm)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
	  return exterror.WrapExtError(err)
	}
	err = file.Close()
	if err != nil {
		return exterror.WrapExtError(err)
	}
	return nil
}

func (this *OKLogger) Remove() error {
	exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	folderLog := config.Str(CFG_OK_FOLDER_LOG,"OK_Log")
	filePath  := filepath.Join(exeDir, folderLog, this.JobName,this.Date + FILE_EXTENSION_OKLOG)

	err := os.Remove(filePath)
	if err != nil {
		return exterror.WrapExtError(err)
	}
	return nil
}