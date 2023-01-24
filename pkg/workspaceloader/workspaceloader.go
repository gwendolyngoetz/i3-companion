package workspaceloader

import (
	"errors"
	"fmt"

	"go.i3wm.org/i3/v4"
)

type LoadConfig struct {
	Output    string
	Workspace int64
	Debug     bool
}

var config *LoadConfig

var workspace_nodes = map[int64][]string{
	1: {"alacritty", "firefox"},
	2: {"firefox", "firefox", "alacritty"},
	3: {"virtualbox"},
	4: {"slack", "teams", "firefox"},
	8: {"alacritty"},
	9: {"alacritty", "alacritty", "alacritty"},
}

var workspace_applications = map[string]string{
	"alacritty":  "alacritty --working-directory ~",
	"firefox":    "firefox",
	"virtualbox": "vboxmanage startvm \"Windows10\"",
	"slack":      "slack",
	"teams":      "/usr/lib/chromium-browser/chromium-browser --disable-features=TFLiteLanguageDetectionEnabled --enable-pinch --profile-directory=Default --app-id=cifhbcnohmdccbgoicgdjpfamggdegmo",
}

func LoadWorkspace(loadConfig *LoadConfig) {
	config = loadConfig

	workspace_node, _ := getWorkspaceNode(config.Workspace)
	cmds := buildCommand(config.Output, config.Workspace, workspace_node)

	for _, cmd := range cmds {
		if config.Debug {
			fmt.Println(cmd)
		} else {
			i3.RunCommand(cmd)
		}
	}
}

func getWorkspaceNode(workspace int64) ([]string, error) {
	for key, workspace_node := range workspace_nodes {
		if key == workspace {
			return workspace_node, nil
		}
	}

	return nil, errors.New("Could not find workspace node")
}

func buildCommand(output string, workspace int64, workspace_node []string) []string {
	return []string{
		buildApplicationCmd(workspace, workspace_node),
		fmt.Sprintf("[workspace=%d] move workspace to output %s;", config.Workspace, config.Output),
	}
}

func buildApplicationCmd(workspace int64, workspace_node []string) string {
	appendCmd := fmt.Sprintf("workspace \"%d:%d\"; append_layout ~/.config/i3/workspaces/workspace-%d.json; ", workspace, workspace, workspace)

	for _, application := range workspace_node {
		appendCmd = appendCmd + fmt.Sprintf("exec --no-startup-id %s; ", workspace_applications[application])
	}

	return appendCmd
}
