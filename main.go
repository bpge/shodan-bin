package main

import (
	"github.com/innovandalism/shodan"

	// module imports
	_ "github.com/innovandalism/shodan/modules/log"
	_ "github.com/innovandalism/shodan/modules/oauth"
	_ "github.com/bpge/mod_slutbot"
	_ "github.com/innovandalism/shodan/modules/version"
)

func main() {
	// this invokes the main function
	shodan.Run()
}
