package gohello

import (
	"flag"
	"fmt"
	"os"
)

const name = "gohello"
const description = "Gohello - print a greeting message"

var version = "LATEST_TAG"
var revision = "HEAD"

func Usage() {
	fmt.Fprintf(os.Stderr, "%s\n", description)
	fmt.Fprintf(os.Stderr, "usage: %s [options]\n", name)
	flag.PrintDefaults()
	os.Exit(2)
}

func ShowVersion() {
	fmt.Fprintf(os.Stderr, "%s version %s (rev:%s)\n", name, version, revision)
	os.Exit(2)
}
