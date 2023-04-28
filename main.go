package main

import (
	"github.com/myuser/myrepo/getway"
	"github.com/myuser/myrepo/server"
)

func main() {
	server.Run()
	go func() {
		server.Run()
	}()
	getway.Run()
}
