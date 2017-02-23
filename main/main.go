package main

import (
	_ "app/passporte/routers"
	"github.com/astaxie/beegae"
	"log"
)

func init() {
	beegae.Run()
	log.Println(" ")
}
