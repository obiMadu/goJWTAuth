package main

import (
	"github.com/obiMadu/goJWTAuth/internals/config"
)

const webPort string = ":8080"

func main() {
	// config
	config.Config()

	// start http server
	routes().Run(webPort)
}
