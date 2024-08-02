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

func NewMultipleMenu(title string, items []string) ([]string, error) {
	p := tea.NewProgram(bubble.NewMutipleMenuModel(title, items), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return nil, fmt.Errorf("error running program: %s", err)
	}

	finalModel, ok := m.(bubble.MutipleMenuModel)
	if !ok {
		return nil, errors.New("Error: Could not assert model.")
	}

	if finalModel.Quitting {
		return nil, errors.New("user quit the program")
	}

	selectedItems := []string{}
	for _, i := range finalModel.List.Items() {
		if itm, ok := i.(bubble.MultipleItem); ok && itm.Selected {
			selectedItems = append(selectedItems, itm.Title)
		}
	}

	return selectedItems, nil
}
