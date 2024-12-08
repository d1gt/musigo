package tui

import (
	"context"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

type Tui struct {
	tea   *tea.Program
	model Model
}

func NewTui() *Tui {
	return &Tui{}
}

func (tui *Tui) Run(ctx context.Context) {
	tui.model = NewModel()

	tui.tea = tea.NewProgram(tui.model)

	_, err := tui.tea.Run()
	if err != nil {
		log.Fatal(err)
	}

}
