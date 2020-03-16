package main

import (
	"log"
	"os"

	"github.com/rconway/goapp/cmdline"
)

func main() {
	cmdLine := cmdline.NewCmdLine(os.Args)

	log.Println("START goapp...")

	log.Printf("server.port = %v\n", cmdLine.Server.Port)

	log.Println("...COMPLETE goapp")
}
