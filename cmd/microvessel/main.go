package main

import (
	"fmt"
	"os"

	"github.com/genc-murat/microvessel/internal/container"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: microvessel run <command> <args>")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "run":
		container.Run(os.Args[2:])
	case "child":
		container.RunContainer(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}
