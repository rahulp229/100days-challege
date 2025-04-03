package main

import (
	"flag"
	"log"
	"runtime"	

	"github.com/veandco/go-sdl2/sdl"

)

func main() {
	runtime.LockOSThread()
	flag.Parse()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Fatal(err)
	}
	//day15
	//usecase :
	gin.Init
	
}