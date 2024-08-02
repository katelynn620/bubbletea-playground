// copy from example
package main

import (
	"bubbletea-playground/pkg/tui"
	"fmt"
	"log"
	"strings"
)

func main() {

	items := []string{"Ramen", "Tomato Soup", "Hamburgers", "Cheeseburgers", "Currywurst", "Okonomiyaki", "Pasta", "Fillet Mignon", "Caviar", "Just Wine"}
	title := "What do you want for dinner?"

	result, err := tui.NewMultipleMenu(title, items)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("Selected items: %s\n", strings.Join(result, ", "))
}
