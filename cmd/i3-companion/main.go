package main

import (
	"flag"
	"fmt"
	"gwendolyngoetz/i3-companion/pkg/workspaceloader"
	"gwendolyngoetz/i3-companion/pkg/workspaceswap"
	"log"
	"os"
)

var (
	required string
	swapCmd  = flag.NewFlagSet("swap", flag.ExitOnError)
	loadCmd  = flag.NewFlagSet("load", flag.ExitOnError)
)

var subcommands = map[string]*flag.FlagSet{
	swapCmd.Name(): swapCmd,
	loadCmd.Name(): loadCmd,
}

func setupCommonFlags() {
	for _, fs := range subcommands {
		fs.StringVar(
			&required,
			"required",
			"",
			"required for all commands",
		)
	}
}

func main() {
	setupCommonFlags()
	a := swapCmd.String("a", "default", "a")
	b := swapCmd.String("b", "default", "b")
	debug := swapCmd.Bool("debug", false, "debug")
	flag.Parse()
	// swapCmd.Parse()

	switch os.Args[1] {
	case "swap":
		// swapCmd.Parse(os.Args[2:])
		workspaceswap.Swap()
	case "load":
		loadCmd.Parse(os.Args[2:])
		workspaceloader.LoadWorkspace()
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}

	fmt.Println("")
	fmt.Println(flag.Args())
	fmt.Println(*a, *b, *debug)
	// fmt.Println(os.Args[2:])
	fmt.Println("done")
}
