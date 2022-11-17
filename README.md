[![Build Only](https://github.com/gwendolyngoetz/i3-companion/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/gwendolyngoetz/i3-companion/actions/workflows/build.yml)

# README

## General

i3-companion COMMAND [OPTIONS]

## Version

i3-companion --version

## Swap Workspaces

i3-companion swap --output1=OUTPUT --output2=OUTPUT

Usage of swap:
-debug
debug
-output1 string
Output (DisplayPort-?, HDMI-?)
-output2 string
Output (DisplayPort-?, HDMI-?)

## Load Workspace

i3-companion load --workspace=WORKSPACE --output-to=OUTPUT

Usage of load:
-debug
debug
-output-to string
Output to place loaded workspace
-required string
required for all commands
-workspace int
Workspace to load (default -1)
