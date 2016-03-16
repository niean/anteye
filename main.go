package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/niean/anteye/g"
	"github.com/niean/anteye/http"
	"github.com/niean/anteye/monitor"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)

	// monitor
	monitor.Start()

	// http
	http.Start()

	select {}
}
