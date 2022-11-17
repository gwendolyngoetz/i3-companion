package main

import (
	"flag"
	"fmt"
	"gwendolyngoetz/i3-companion/pkg/workspaceloader"
	"gwendolyngoetz/i3-companion/pkg/workspaceswap"
	"log"
	"os"
)

type Config struct {
	ShowVersion bool
}

var (
	required string
	swapCmd  = flag.NewFlagSet("swap", flag.ExitOnError)
	loadCmd  = flag.NewFlagSet("load", flag.ExitOnError)
)

var Version = "development"

var subcommands = map[string]*flag.FlagSet{
	swapCmd.Name(): swapCmd,
	loadCmd.Name(): loadCmd,
}

func setupCommonFlags() {
	for _, fs := range subcommands {
		fs.StringVar(&required, "required", "", "required for all commands")
	}
}

func buildSwapConfig() *workspaceswap.SwapConfig {
	config := workspaceswap.SwapConfig{}

	swapCmd.StringVar(&config.Output1, "output1", "", "Output (DisplayPort-?, HDMI-?)")
	swapCmd.StringVar(&config.Output2, "output2", "", "Output (DisplayPort-?, HDMI-?)")
	swapCmd.BoolVar(&config.Debug, "debug", false, "debug")
	swapCmd.Parse(os.Args[2:])

	return &config
}

func buildLoadConfig() *workspaceloader.LoadConfig {
	config := workspaceloader.LoadConfig{}

	loadCmd.Int64Var(&config.Workspace, "workspace", -1, "Workspace to load")
	loadCmd.StringVar(&config.Output, "output-to", "", "Output to place loaded workspace")
	loadCmd.BoolVar(&config.Debug, "debug", false, "debug")
	loadCmd.Parse(os.Args[2:])

	return &config
}

func buildVersionConfig() *Config {
	config := Config{}

	flag.BoolVar(&config.ShowVersion, "version", false, "version")
	flag.Parse()

	return &config
}

func main() {
	setupCommonFlags()

	switch os.Args[1] {
	case "swap":
		workspaceswap.Swap(buildSwapConfig())
	case "load":
		workspaceloader.LoadWorkspace(buildLoadConfig())
	case "--version":
		buildVersionConfig()
		fmt.Println(Version)
	default:
		log.Fatalf("Unknown command: %s", os.Args[1])
	}
}
