package main

import . "peterchu999/url-shorterner/servers"

func main() {
	server := SetupServer()
	server.Run()
}
