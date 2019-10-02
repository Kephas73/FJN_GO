package main

import (
	"Common"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
fmt.Println(filepath.Dir(os.Args[0]))
pgName := config.Str(PG_NAME,"")
OKLog := OKLogger{pgName, Common.GetDateTimeCurrent()}
OKLog.Remove()
}
