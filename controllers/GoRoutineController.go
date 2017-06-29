package controllers

import (
	"log"
	"fmt"
)

type GoRoutineController struct {
	BaseController
}


func say(s string) {
	for i := 0; i < 5; i++ {
		//time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func (c *GoRoutineController)AddGoRoutine() {
	//go sample1()
	log.Println("1")
	log.Println("2")
	go say("world")
	say("hello")
	log.Println("3")
	log.Println("4")

}



