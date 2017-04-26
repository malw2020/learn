package main

import (
	seelog "github.com/cihub/seelog"
)

func main() {
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		seelog.Critical("err parsing config log file, ", err)
		return
	}

	seelog.ReplaceLogger(logger)
	defer seelog.Flush()

	seelog.Info("tester start...")
}
