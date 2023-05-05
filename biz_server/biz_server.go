package main

import (
	"hero_story_server/common/logger"
	"os"
	"path"
)

func main() {

	ex, err := os.Executable()
	if nil != err {
		panic(err)
	}
	logger.Config(path.Dir(ex) + "/log/biz_server.log")
	logger.Info("helldf")
}
