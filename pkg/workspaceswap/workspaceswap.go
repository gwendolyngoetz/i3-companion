package workspaceswap

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

type SwapConfig struct {
	Output1 string
	Output2 string
	Debug   bool
}

var (
	config     *SwapConfig
	focusedKey = "Focused"
)

func Swap(swapConfig *SwapConfig) {
	config = swapConfig
	workspaces := getWorkspaces()

	workspace1 := workspaces[config.Output1]
	workspace2 := workspaces[config.Output2]
	workspaceFocused := workspaces[focusedKey]

	// TODO: Handle output range workspaces

	if config.Debug {
		log.Println("WS1", workspace1.Num, workspace1.Name, workspace1.Output)
		log.Println("WS2", workspace2.Num, workspace2.Name, workspace2.Output)
		log.Println("WSF", workspaceFocused.Num, workspaceFocused.Name, workspaceFocused.Output)
	}

	swapWorkspace(workspace1, workspace2)
	swapWorkspace(workspace2, workspace1)
	setFocusedWorkspace(workspaceFocused)
}

func getWorkspaces() map[string]i3.Workspace {
	allWorkspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	filteredWorkspaces := map[string]i3.Workspace{}

	for _, s := range allWorkspaces {
		if s.Focused {
			filteredWorkspaces[focusedKey] = s
		}

		if s.Visible {
			filteredWorkspaces[s.Output] = s
		}
	}

	return filteredWorkspaces
}

func swapWorkspace(workspace1 i3.Workspace, workspace2 i3.Workspace) {
	command1 := fmt.Sprintf("[workspace=%d] move workspace to output %s", workspace1.Num, workspace2.Output)
	command2 := fmt.Sprintf("workspace %s", workspace1.Name)

	if config.Debug {
		log.Println(command1)
		log.Println(command2)
	}

	_, err := i3.RunCommand(command1)
	if err != nil {
		log.Println(err)
	}

	_, err = i3.RunCommand(command2)
	if err != nil {
		log.Println(err)
	}
}

func setFocusedWorkspace(workspace i3.Workspace) {
	command := fmt.Sprintf("workspace %s", workspace.Name)

	if config.Debug {
		log.Println(command)
	}

	_, err := i3.RunCommand(command)
	if err != nil {
		log.Println(err)
	}
}
