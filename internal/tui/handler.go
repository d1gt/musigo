package tui

import (
	"context"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) handleEnter(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	switch m.currTab {
	case searchTab:
		query := m.textInput.Value()
		if m.currSearchResult != -1 {
			query = m.searchResults[m.currSearchResult].Name
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		results, _ := m.Search(ctx, query)

		m.searchResults = results

		m.currSearchResult = -1
		m.currSuggestion = -1

		m.currTab = 1

	case songsTab:

	}

	return m, cmd
}

func (m Model) handleRight(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	tabLen := len(m.tabs) - 1

	tabIdx := m.currTab
	tabIdx++

	if tabIdx > tabLen {
		tabIdx = 0
	}

	m.currTab = tabIdx

	return m, cmd
}

func (m Model) handleLeft(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	tabLen := len(m.tabs) - 1

	tabIdx := m.currTab
	tabIdx--

	if tabIdx < 0 {
		tabIdx = tabLen
	}

	m.currTab = tabIdx

	return m, cmd
}

func (m Model) handleDown(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	switch m.currTab {
	case searchTab:
		suggestionsLen := len(m.searchSuggestions) - 1

		suggestionIdx := m.currSuggestion
		suggestionIdx++

		if suggestionIdx > suggestionsLen {
			suggestionIdx = 0
		}

		m.currSuggestion = suggestionIdx

	case songsTab:
		resultsLen := len(m.searchResults) - 1

		resultsIdx := m.currSearchResult
		resultsIdx++

		if resultsIdx > resultsLen {
			resultsIdx = 0
		}

		m.currSearchResult = resultsIdx
	}

	return m, cmd
}

func (m Model) handleUp(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	switch m.currTab {
	case searchTab:
		suggestionsLen := len(m.searchSuggestions) - 1

		suggestionIdx := m.currSuggestion
		suggestionIdx--

		if suggestionIdx < 0 {
			suggestionIdx = suggestionsLen
		}

		m.currSuggestion = suggestionIdx

	case songsTab:
		resultsLen := len(m.searchResults) - 1

		resultsIdx := m.currSearchResult
		resultsIdx--

		if resultsIdx < 0 {
			resultsIdx = resultsLen
		}

		m.currSearchResult = resultsIdx
	}

	return m, cmd
}

func (m Model) handleEsc(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	prevTab := m.prevTab
	currTab := m.currTab

	m.currTab = prevTab
	m.prevTab = currTab

	return m, cmd
}

func (m Model) handleDefault(cmd tea.Cmd) (tea.Model, tea.Cmd) {
	switch m.currTab {
	case searchTab:
		prevInputTs := m.prevInputTs
		m.prevInputTs = time.Now()

		if prevInputTs.IsZero() {
			return m, cmd
		}

		searchQuery := m.textInput.Value()
		searchQueryLen := len(searchQuery)
		prevQuery := m.prevInput

		if searchQueryLen < 3 {
			m.searchSuggestions = []string{}
			return m, cmd
		}

		if prevQuery == searchQuery {
			return m, cmd
		}

		return m, m.fetchSuggestionsCmd(searchQuery)
	}

	return m, cmd
}
