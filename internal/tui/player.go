package tui

import (
	"fmt"
	"strings"
	"time"
)

func (m Model) renderPlayer() string {
	var b strings.Builder

	b.WriteString("\n")

	playStatus := "▶"
	if m.isPlaying {
		playStatus = "⏸ "
	}

	fmt.Fprintf(&b, "%s | [%s / %s]", playStatus, formatDuration(time.Duration(0)), formatDuration(time.Duration(0)))

	return b.String()
}

func formatDuration(d time.Duration) string {
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
