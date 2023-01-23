[![Build Only](https://github.com/gwendolyngoetz/i3-companion/actions/workflows/build.yml/badge.svg?branch=master)](https://github.com/gwendolyngoetz/i3-companion/actions/workflows/build.yml)

# README

## Version
### Example
`i3-companion --version`

## Swap Workspaces

### Example
`i3-companion swap --output1=DisplayPort-0 --output2=DisplayPort-1`

### Usage
```
Usage of swap:
  --debug
      Prints extra debug info
  --output1 string
      Output (DisplayPort-?, HDMI-?)
  --output2 string
      Output (DisplayPort-?, HDMI-?)
```

## Load Workspace

### Example

`i3-companion load --workspace=1 --output-to=DisplayPort-0`

### Usage
```
Usage of load:
  --debug
      Prints extra debug info
  --output-to string
      Output to place loaded workspace
  --workspace int
      Workspace to load (default -1)
```
