package g

import (
	"log"
	"runtime"
)

// changelog:
// 0.0.1: init project
// 0.0.2: fmt import path, modify control and debug scripts
// 0.0.3: use pfc, rm vg

const (
	VERSION = "0.0.3"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
