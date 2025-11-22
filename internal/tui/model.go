package tui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

// Model represents the TUI state
type Model struct {
	projectName textinput.Model
	width       int
	height      int
}

// Result contains the collected form data
type Result struct {
	ProjectName string
	Template    string
}

// NewModel creates a new TUI model
func NewModel() Model {
	// Initialize text input for project name
	ti := textinput.New()
	ti.Placeholder = "my-express-app"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 50
	ti.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))
	ti.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))

	return Model{
		projectName: ti,
		width:       80,
		height:      24,
	}
}
