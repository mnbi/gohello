package main

import (
	"flag"
	"os"

	"github.com/mnbi/gohello"
	"github.com/mnbi/gohello/internal/greeting"
	"github.com/mnbi/gohello/internal/utils"
)

var versionFlag = flag.Bool("v", false, "show version")
var usageFlag = flag.Bool("h", false, "show usage")

func main() {
	flag.Usage = gohello.Usage
	flag.Parse()

	if *versionFlag {
		gohello.ShowVersion()
		os.Exit(1)
	}

	if *usageFlag {
		flag.Usage()
	}

	msg := greeting.Greeting()
	utils.Print(string(msg))
}
