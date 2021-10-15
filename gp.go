package main

import (
	_ "github.com/xiaozhang3/gp/boot"
	_ "github.com/xiaozhang3/gp/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
