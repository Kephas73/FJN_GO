package main

import (
	"github.com/goframework/gf/cfg"
	"github.com/goframework/gf/ext"
	"os"
	"path/filepath"
)

var config *cfg.Cfg

// Init information file config
func init() {
	config = &cfg.Cfg{}
	configFile := CONFIG_FILE
	if  !ext.FileExists(configFile) {
		exeDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
		configFile = filepath.Join(exeDir, CONFIG_FILE)
	}
	config.Load(configFile)
}




