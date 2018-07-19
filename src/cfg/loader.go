package cfg

import (
	"bufio"
	"os"
	"util/log"

	"github.com/BurntSushi/toml"
)

func newConfig() *config {
	c := new(config)
	c.DBUrl = "name:passwd@tcp(db_ip:db_port)/sqlx_db?charset=utf8mb4"

	return c
}

func createCfgFile(path string) (*config, error) {
	log.I("create:", path)

	f, e := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0660)
	if e != nil {
		return nil, e
	}
	defer f.Close()

	fw := bufio.NewWriter(f)
	enc := toml.NewEncoder(fw)

	c := newConfig()
	return c, enc.Encode(c)
}

func loadCfgFile(path string) (*config, error) {
	log.I("read:", path)

	c := new(config)
	_, e := toml.DecodeFile(path, c)
	if e != nil {
		return nil, e
	}

	return c, nil
}
