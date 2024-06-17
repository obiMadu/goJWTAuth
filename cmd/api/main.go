package main

const webPort string = ":8080"

func main() {

	// start http server
	routes().Run(webPort)

}
