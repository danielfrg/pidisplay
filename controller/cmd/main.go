package main

import (
	"pidisplay/server"
)

func main() {
	server := server.New()
	server.Listen("localhost:3000")
}
