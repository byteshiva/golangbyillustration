package main

import (
	"fmt"
	"os"
)

// CLI - processes command line arguments
type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage: ")

}
func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

// Run parses command line arguments
func (cli *CLI) Run() {
	cli.validateArgs()
}

func main() {
	cli := CLI{}
	cli.Run()
}
