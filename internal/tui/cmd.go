package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type suggestionsType []string

func (m Model) fetchSuggestionsCmd(query string) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		suggestions, _ := m.GetSearchSuggestions(ctx, query)
		return suggestionsType(suggestions)
	}
}
