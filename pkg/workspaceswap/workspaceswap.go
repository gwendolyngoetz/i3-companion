package workspaceswap

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

func Swap() {
	fmt.Println("swap-workspace")
	// workspaces, focusedWorkspace := getWorkspaces()

	//for i, s := range workspaces {
	//	fmt.Println(i, s.Num, s.Name, s.Output)
	//}

	// fmt.Println("F", focusedWorkspace.Num, focusedWorkspace.Name, focusedWorkspace.Output)
	// fmt.Println("")

	// swapWorkspace(workspaces[0], workspaces[1])
	// swapWorkspace(workspaces[1], workspaces[0])
	// focusWorkspace(focusedWorkspace)
}

func getWorkspaces() ([]i3.Workspace, i3.Workspace) {
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

func swapWorkspace(workspace1 i3.Workspace, workspace2 i3.Workspace) {
	command1 := fmt.Sprintf("[workspace=%d] move workspace to output %s", workspace1.Num, workspace2.Output)
	command2 := fmt.Sprintf("workspace %s", workspace1.Name)

	fmt.Println(command1)
	fmt.Println(command2)

	//_, err := i3.RunCommand(command1)
	//if err != nil {
	//	fmt.Println(err)
	//}

	//_, err = i3.RunCommand(command2)
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func focusWorkspace(workspace i3.Workspace) {
	command := fmt.Sprintf("workspace %s", workspace.Name)
	fmt.Println(command)

	//_, err := i3.RunCommand(command)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
