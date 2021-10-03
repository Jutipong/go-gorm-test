package main

import (
	"gorm-test/config"
	"gorm-test/prejoin"
)

//GIN_MODE=debug
//GIN_MODE=release
func init() {
	config.InitialConfig()
	config.InitialDB()
}

func main() {
	prejoin.Query()
}
