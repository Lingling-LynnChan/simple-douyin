package main

import (
	"douyin-simple/internal"
	"douyin-simple/internal/conf"
	"flag"
)

var addr *string

func init() {
	addr = flag.String("addr", "", "启动时手动配置监听端口")
	flag.Parse()
}

func main() {
	conf := conf.GetConf("configs/config.yaml")
	bootstrap := internal.WireApp(
		conf.Bootstrap,
		conf.Server,
		conf.Data,
	)

	if err := bootstrap.Listen(*addr); err != nil {
		panic(err)
	}
}
