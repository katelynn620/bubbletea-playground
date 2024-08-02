package main

import (
	"bubbletea-playground/pkg/tui"
	"log"
)

func main() {

	msgOpt := tui.NewMsgboxOpt("Are you sure you want to proceed?").WithDefualtNo()

	result, err := tui.NewMsgbox(msgOpt)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	log.Printf("User selected: %t\n", result)
}
