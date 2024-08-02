// copy from example
package main

import (
	"bubbletea-playground/pkg/tui"
	"log"
)

func main() {

	items := []string{"Ramen", "Tomato Soup", "Hamburgers", "Cheeseburgers", "Currywurst", "Okonomiyaki", "Pasta", "Fillet Mignon", "Caviar", "Just Wine"}
	title := "What do you want for dinner?"

	result, err := tui.NewMenu(title, items)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	log.Printf("You chose: %s", result)
}
