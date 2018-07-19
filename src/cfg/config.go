package cfg

import (
	"os"
	"util/log"
)

type config struct {
	DBUrl string
}

var configPath string = "config.ini"
var Val *config

func init() {
	// if config file not exist, create it.
	_, e := os.Stat(configPath)
	if e != nil {
		log.E("file not found:", configPath)
		Val, e = createCfgFile(configPath)
		if e != nil {
			panic(e.Error())
		}
		return
	}

	// if config file exist, read it
	Val, e = loadCfgFile(configPath)
}
