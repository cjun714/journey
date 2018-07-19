package dao

import (
	"cfg"
	"util/log"
)

func init() {
	log.I("DB url:", cfg.Val.DBUrl)
}
