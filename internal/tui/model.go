package tui

import (
	"strings"
	"time"

	"github.com/d1gt/musigo/internal/youtube"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	prevTab int
	currTab int
	tabs    []string

	prevInputTs time.Time
	prevInput   string
	textInput   textinput.Model

	searchSuggestions []string
	currSuggestion    int

	searchResults    []youtube.Song
	currSearchResult int

	isPlaying    bool
	songDuration time.Duration

	youtubeClient youtube.Client
}

func NewModel() Model {
	m := Model{}

	m.tabs = []string{"search", "songs", "favorites", "offline", "options"}

	m.textInput = textinput.New()
	m.textInput.Prompt = ""
	m.textInput.Placeholder = ""
	m.textInput.Focus()

	m.currSearchResult = -1
	m.currSuggestion = -1

	m.youtubeClient = youtube.New(nil)

	return m
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.textInput, cmd = m.textInput.Update(msg)

	switch msg := msg.(type) {
	case suggestionsType:
		m.searchSuggestions = msg
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit

		case tea.KeyEnter:
			return m.handleEnter(cmd)

		case tea.KeyTab, tea.KeyRight:
			return m.handleRight(cmd)

		case tea.KeyShiftTab, tea.KeyLeft:
			return m.handleLeft(cmd)

		case tea.KeyDown:
			return m.handleDown(cmd)

		case tea.KeyUp:
			return m.handleUp(cmd)

		case tea.KeyEscape:
			return m.handleEsc(cmd)

		case tea.KeyBackspace, tea.KeyDelete:
			return m.handleDefault(cmd)

		default:
			return m.handleDefault(cmd)
		}
	}

	return m, cmd
}

func (m Model) View() string {
	var b strings.Builder

	b.WriteString(m.renderTabs())

	switch m.currTab {
	case 0:
		b.WriteString(m.renderSearch())

	case 1:
		b.WriteString(m.renderSearchResults())

	}

	b.WriteString(m.renderPlayer())

	return b.String()
}
