package main

import (
	"fmt"
	"log"

	"go.i3wm.org/i3/v4"
)

func main() {
	workspaces, err := i3.GetWorkspaces()
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range workspaces {
		fmt.Println(i, s.Name)
	}

	fmt.Println("")
	fmt.Println("done")
}
