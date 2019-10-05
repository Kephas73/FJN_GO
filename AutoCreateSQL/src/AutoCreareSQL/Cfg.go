package main

import (
	"github.com/goframework/gf/cfg"
	"github.com/goframework/gf/ext"
	"github.com/goframework/gf/exterror"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var config *cfg.Cfg

type IFTable struct {
	Table string
	Path  string
	Sql   string
}

// Init information file config
func init() {
	config = &cfg.Cfg{}
	configFile := CONFIG_FILE
	if !ext.FileExists(configFile) {
		exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configFile = filepath.Join(exeDir, CONFIG_FILE)
	}
	config.Load(configFile)
}

func StrArray(this *cfg.Cfg, key string) []string {

	if v, ok := this.ArrData[key]; ok {
		return v
	}
	return nil
}

func CfgVarArray(this *cfg.Cfg, key string) ([]IFTable, error) {

	var tableIF []IFTable
	value := StrArray(this, key)
	for _, v := range value {
		v1 := strings.Split(v, SYMBOL_ARROW)
		exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		filePath := filepath.Join(exeDir, strings.Trim(v1[1], CHAR_SPACE))
		sql, err := ioutil.ReadFile(filePath)
		if err != nil {
			return tableIF, exterror.WrapExtError(err)
		}

		item := IFTable{strings.Trim(v1[0], CHAR_SPACE), strings.Trim(v1[1], CHAR_SPACE), string(sql)}
		tableIF = append(tableIF, item)
	}

	return tableIF, nil
}
