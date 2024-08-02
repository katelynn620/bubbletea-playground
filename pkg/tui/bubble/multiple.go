package bubble

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MultipleItem struct {
	Title    string
	Selected bool
}

func (i MultipleItem) FilterValue() string { return "" }

type MultipleItemDelegate struct{}

func (d MultipleItemDelegate) Height() int                             { return 1 }
func (d MultipleItemDelegate) Spacing() int                            { return 0 }
func (d MultipleItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d MultipleItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(MultipleItem)
	if !ok {
		return
	}

	status := "[ ]"
	if i.Selected {
		status = "[*]"
	}

	str := fmt.Sprintf("%s %s", status, i.Title)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = itemStyle.Foreground(lipgloss.Color("26")).Render
	}

	fmt.Fprint(w, fn(str))
}

type MutipleMenuModel struct {
	List     list.Model
	Choice   string
	Quitting bool
	title    titleModel
}

func (m MutipleMenuModel) Init() tea.Cmd {
	return nil
}

func (m MutipleMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetWidth(msg.Width)
		m.title.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case " ":
			if i, ok := m.List.SelectedItem().(MultipleItem); ok {
				i.Selected = !i.Selected
				m.List.SetItem(m.List.Index(), i)
			}
		case "q", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.List.SelectedItem().(item)
			if ok {
				m.Choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m MutipleMenuModel) View() string {
	return m.title.View() + "\n\n" + m.List.View()
}

func NewMutipleMenuModel(title string, menuItems []string) MutipleMenuModel {
	var items []list.Item
	for _, i := range menuItems {
		items = append(items, MultipleItem{Title: i})
	}
	const defaultWidth = 20

	l := list.New(items, MultipleItemDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	t := NewTitleBox("Hello Kitty")

	return MutipleMenuModel{List: l, title: t}
}
