package tui

import (
	"context"
	"fmt"
	"strings"

	"github.com/d1gt/musigo/internal/youtube"
)

func (m Model) renderSearch() string {
	var b strings.Builder

	b.WriteString(m.textInput.View() + "\n")

	for i, suggestion := range m.searchSuggestions {
		if i == m.currSuggestion {
			fmt.Fprintf(&b, " > %v\n", suggestion)
		} else {
			fmt.Fprintf(&b, "   %v\n", suggestion)
		}
	}

	return b.String()
}

func (m Model) renderSearchResults() string {
	var b strings.Builder

	b.WriteString("\n")

	for i, res := range m.searchResults {
		var artists string

		for i, artist := range res.Artists {
			if i == 0 {
				artists = artist.Name
			} else {
				artists = " & " + artist.Name
			}
		}

		if m.currSearchResult == i {
			fmt.Fprintf(&b, " > %v ● %v ● %v ● %v\n", res.Name,
				artists, res.Plays, res.Duration)
		} else {
			fmt.Fprintf(&b, "   %v ● %v ● %v ● %v\n", res.Name,
				artists, res.Plays, res.Duration)
		}
	}

	return b.String()
}

func (m Model) Search(ctx context.Context, input string) ([]youtube.Song, error) {
	songs, err := m.youtubeClient.Search(ctx, input)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (m Model) GetSearchSuggestions(ctx context.Context, input string) ([]string, error) {
	suggestions, err := m.youtubeClient.GetSearchSuggestions(ctx, input)
	if err != nil {
		return nil, err
	}

	return suggestions, nil
}
