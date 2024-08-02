package bubble

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MsgBoxModel struct {
	Message string
	Yesno   bool
	Quit    bool
	width   int
}

func (m *MsgBoxModel) SetWidth(width int) {
	m.width = width
}

func NewMsgBoxModel(message string, defaultNo bool) MsgBoxModel {
	return MsgBoxModel{
		Message: message,
		Yesno:   !defaultNo,
	}
}

func (m MsgBoxModel) Init() tea.Cmd {
	return nil
}

func (m MsgBoxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case tea.KeyLeft.String(), tea.KeyRight.String(), tea.KeyTab.String():
			m.Yesno = !m.Yesno
		case tea.KeyEsc.String(), tea.KeyEscape.String():
			m.Quit = true
			return m, tea.Quit
		case tea.KeyEnter.String():
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MsgBoxModel) View() string {
	style := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), true, true, false, true).
		Padding(1).
		Width(m.width - 2).
		Align(lipgloss.Center)

	selectionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Border(lipgloss.NormalBorder(), false, true, true, true).
		Width(m.width - 2).
		Align(lipgloss.Center)

	normalStyle := lipgloss.NewStyle()
	selectedStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("14"))

	okText := "<  OK  >"
	exitText := "< Exit >"
	okStyle := normalStyle
	exitStyle := normalStyle
	if m.Yesno {
		okStyle = selectedStyle
	} else {
		exitStyle = selectedStyle
	}
	control := okStyle.Render(okText) + "        " + exitStyle.Render(exitText)

	msgBox := style.Render(m.Message)
	selection := selectionStyle.Render(control)
	return msgBox + "\n" + selection
}
