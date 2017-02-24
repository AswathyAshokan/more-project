package main

import (
	_ "app/passporte/routers"
	"github.com/astaxie/beegae"
)

func init() {
	beegae.Run()
}
