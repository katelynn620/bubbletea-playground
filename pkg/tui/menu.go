package tui

import (
	"bubbletea-playground/pkg/tui/bubble"
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func NewMenu(title string, items []string) (string, error) {
	p := tea.NewProgram(bubble.NewMenuModel(title, items), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return "", fmt.Errorf("error running program: %s", err)
	}

	if m.(bubble.MenuModel).Quitting {
		return "", errors.New("user quit the program")
	}

	return m.(bubble.MenuModel).Choice, nil
}
