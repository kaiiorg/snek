package main

import (
	"flag"

	"github.com/kaiiorg/snek/pkg/game"
)

func main() {
	flag.Parse()
	game.New().Run()
}
