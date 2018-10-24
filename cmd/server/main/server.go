package main

import (
	"github.com/lastsys/bbs/internal/font"
	"github.com/lastsys/bbs/internal/server"
)

func init() {
	font.GenerateFont()
}

func main() {
	server.StartServer("0.0.0.0:9000")
}
