package tui

import (
	"fmt"
	"strings"
)

const (
	searchTab int = iota
	songsTab
	favoritesTab
	offlineTab
	settingsTab
)

func (m Model) renderTabs() string {
	var b strings.Builder

	for i, tab := range m.tabs {
		if i == m.currTab {
			fmt.Fprintf(&b, ">%s  ", tab)
		} else {
			fmt.Fprintf(&b, " %s  ", tab)
		}
	}

	b.WriteString("\n")

	return b.String()
}
