package tui

import (
	"bubbletea-playground/pkg/tui/bubble"
	"errors"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MsgboxOpt struct {
	Message   string
	DefualtNo bool
}

func NewMsgboxOpt(message string) *MsgboxOpt {
	return &MsgboxOpt{
		Message: message,
	}
}

func (o *MsgboxOpt) WithDefualtNo() *MsgboxOpt {
	o.DefualtNo = true
	return o
}

func NewMsgbox(opt *MsgboxOpt) (bool, error) {
	p := tea.NewProgram(bubble.NewMsgBoxModel(opt.Message, opt.DefualtNo), tea.WithAltScreen())
	m, err := p.Run()
	if err != nil {
		return false, fmt.Errorf("error running program: %s", err)
	}

	finalModel, ok := m.(bubble.MsgBoxModel)
	if !ok {
		return false, errors.New("error: could not assert model")
	}

	if finalModel.Quit {
		return false, errors.New("user quit the program")
	}

	return finalModel.Yesno, nil
}
