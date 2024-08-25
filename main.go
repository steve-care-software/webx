package main

import (
	"log"

	"github.com/steve-care-software/webx/gui"
)

func main() {
	title := "Steve Care"
	width := float32(1024.0)
	height := float32(1024.0)
	guiIns, err := gui.NewBuilder().
		Create().
		WithTitle(title).
		WithWidth(width).
		WithHeight(height).
		Now()

	if err != nil {
		panic(err)
	}

	err = guiIns.Execute()
	if err != nil {
		log.Print(err.Error())
	}
}
