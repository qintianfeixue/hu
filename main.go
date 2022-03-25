package main

import (
	_ "ceshi1/boot"
	_ "ceshi1/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
