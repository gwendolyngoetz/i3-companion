package main

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

type LocalWorkspace struct {
	Num     int64
	Name    string
	Output  string
	Focused bool
}

func GetWorkspaces() ([]i3.Workspace, i3.Workspace) {
	allWorkspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	filteredWorkspaces := []i3.Workspace{}
	var focusedWorkspace i3.Workspace

	for _, s := range allWorkspaces {
		if s.Focused {
			focusedWorkspace = s
		}

		if s.Visible {
			filteredWorkspaces = append(filteredWorkspaces, s)
		}
	}

	return filteredWorkspaces, focusedWorkspace
}

func SwapWorkspace(workspace1 i3.Workspace, workspace2 i3.Workspace) {
	num := workspace1.Num
	screen := workspace2.Name
	command1 := fmt.Sprintf("[workspace=%d] move workspace to output %s", num, screen)
	command2 := fmt.Sprintf("workspace %d:%d", num, num)
	fmt.Println(command1)
	fmt.Println(command2)
	// i3.RunCommand(command1)
}

func main() {
	workspaces, focusedWorkspace := GetWorkspaces()

	for i, s := range workspaces {
		fmt.Println(i, s.Num, s.Name, s.Output)
	}

	SwapWorkspace(workspaces[0], workspaces[1])

	fmt.Println("F", focusedWorkspace.Num, focusedWorkspace.Name, focusedWorkspace.Output)
	fmt.Println("")
	fmt.Println("done")
}
