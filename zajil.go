package main

import (
	"github.com/ammardev/zajil/components"
	tea "github.com/charmbracelet/bubbletea"
)

type Zajil struct {
	mode       string
	urlInput   components.Input
	windowSize tea.WindowSizeMsg
}

func NewApplicationModel() Zajil {
	return Zajil{
		mode:     "normal",
		urlInput: components.NewInput(10),
		windowSize: tea.WindowSizeMsg{
			Width:  4,
			Height: 4,
		},
	}

}

func (zajil Zajil) Init() tea.Cmd {
	return tea.ClearScreen
}

func (zajil Zajil) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return zajil, zajil.processKeyboardInput(msg.(tea.KeyMsg))
	case tea.WindowSizeMsg:
		zajil.windowSize = msg.(tea.WindowSizeMsg)
        zajil.urlInput.Resize(zajil.windowSize.Width)
		return zajil, nil
	}

	return zajil, nil
}

func (zajil Zajil) View() string {
	view := ""
    view = zajil.urlInput.Render()

	return view
}

func (zajil *Zajil) processKeyboardInput(key tea.KeyMsg) tea.Cmd {
	if zajil.mode == "normal" {
		switch key.String() {
		case "q", "esc":
			return tea.Quit
		case "i", "I":
			zajil.mode = "url"
			zajil.urlInput.Focus()
			return nil
		}
	} else if zajil.mode == "url" {
		switch key.Type {
		case tea.KeyEsc, tea.KeyCtrlC, tea.KeyEnter:
			zajil.mode = "normal"
			zajil.urlInput.Blur()
			return nil
		default:
            return zajil.urlInput.Insert(key)
		}
	}

	return nil
}