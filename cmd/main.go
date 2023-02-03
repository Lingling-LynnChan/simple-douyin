package main

import (
	"douyin-simple/internal"
	"douyin-simple/internal/conf"
	"flag"
	"fmt"
)

var addr *string

func init() {
	addr = flag.String("addr", "", "启动时手动配置监听端口")
	flag.Parse()
}

func main() {
	conf := conf.GetConf("configs/config.yaml")
	bootstrap := internal.WireApp(conf.Server, conf.Data)
	bootstrap.Bootstrap()
	fmt.Printf("addr: %v\n", addr)
	if err := bootstrap.Listen(*addr); err != nil {
		panic(err)
	}
}
