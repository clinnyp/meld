package main

import (
	"os"

	"github.com/charmbracelet/huh"
)

type MenuSelection struct {
	Selection string
}

func main() {

	form := huh.NewForm(huh.NewGroup(
		huh.NewNote().Title("Meld").Description("Lets organise everything about you"),
	))

	err := form.Run()
	if err != nil {
		os.Exit(1)
	}

}
