package global

import "github.com/topdeoo/codeprint/back/config"

var MyConfig *config.Config

func Init() {
	MyConfig = config.ConfigInit()
}
