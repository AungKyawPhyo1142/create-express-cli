package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Title styling with gradient-like effect
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(1, 4).
			MarginBottom(2).
			Align(lipgloss.Center).
			Width(50)

	// Subtitle for instructions
	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A0A0A0")).
			MarginBottom(1).
			Italic(true)

	// Help text styling
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(2).
			Italic(true)

	// Label styling
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#D0D0D0")).
			MarginBottom(1)
)

// Init initializes the TUI model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			projectName := strings.TrimSpace(m.projectName.Value())
			if projectName != "" {
				return m, tea.Quit
			}
		}
		var cmd tea.Cmd
		m.projectName, cmd = m.projectName.Update(msg)
		return m, cmd
	}

	return m, nil
}

// View renders the UI
func (m Model) View() string {
	var b strings.Builder

	// Creative title with emoji
	title := "✨ Create Express Project CLI ✨"
	b.WriteString(titleStyle.Render(title))
	b.WriteString("\n\n")

	b.WriteString(subtitleStyle.Render("Let's set up your TypeScript Express.js application"))
	b.WriteString("\n\n")
	b.WriteString(labelStyle.Render("📁 Project name:"))
	b.WriteString("\n")
	// Style the input field
	inputView := m.projectName.View()
	b.WriteString(inputView)
	b.WriteString("\n")
	b.WriteString(helpStyle.Render("💡 Press Enter to continue, Esc to cancel"))

	return b.String()
}

// GetResult returns the collected form data
func (m Model) GetResult() Result {
	return Result{
		ProjectName: strings.TrimSpace(m.projectName.Value()),
		Template:    "express-ts", // Always TypeScript now
	}
}

// Run starts the TUI program and returns the result
func Run() (Result, error) {
	m := NewModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		return Result{}, fmt.Errorf("error running TUI: %w", err)
	}

	if model, ok := finalModel.(Model); ok {
		return model.GetResult(), nil
	}

	return Result{}, fmt.Errorf("unexpected model type")
}
