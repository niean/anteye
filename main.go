package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/niean/anteye/g"
	"github.com/niean/anteye/http"
	"github.com/niean/anteye/monitor"
	"github.com/niean/anteye/proc"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)
	// proc
	proc.Start()

	// monitor
	monitor.Start()

	// http
	http.Start()

	select {}
}
