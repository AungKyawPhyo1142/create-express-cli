package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProgressStep int

const (
	StepGenerating ProgressStep = iota
	StepInstalling
	StepInitializingGit
	StepComplete
)

type ProgressModel struct {
	progress    progress.Model
	spinner     spinner.Model
	currentStep ProgressStep
	message     string
	percent     float64
	width       int
	height      int
	done        bool
}

type tickMsg time.Time
type stepMsg struct {
	step    ProgressStep
	message string
}
type progressMsg struct {
	percent float64
}

var (
	progressTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Background(lipgloss.Color("#7D56F4")).
				Padding(1, 4).
				MarginBottom(2).
				Align(lipgloss.Center).
				Width(50)

	progressMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#A0A0A0")).
				MarginTop(1).
				MarginBottom(1)
)

func NewProgressModel() ProgressModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4"))

	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(50),
		progress.WithoutPercentage(),
	)

	return ProgressModel{
		progress:    p,
		spinner:     s,
		currentStep: StepGenerating,
		message:     "Generating project files...",
		percent:     0.0,
		width:       80,
		height:      24,
		done:        false,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		progressTickCmd(),
	)
}

func (m ProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.progress.Width = msg.Width - 20
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tickMsg:
		if m.done {
			return m, nil
		}
		// Increment progress
		if m.percent < 1.0 {
			var increment float64
			switch m.currentStep {
			case StepGenerating:
				if m.percent < 0.4 {
					increment = 0.02
				} else {
					m.percent = 0.4
				}
			case StepInstalling:
				if m.percent < 0.4 {
					m.percent = 0.4
				}
				if m.percent < 0.9 {
					increment = 0.01
				} else {
					m.percent = 0.9
				}
			case StepInitializingGit:
				if m.percent < 0.9 {
					m.percent = 0.9
				}
				if m.percent < 1.0 {
					increment = 0.02
				} else {
					m.percent = 1.0
					m.done = true
				}
			case StepComplete:
				m.percent = 1.0
				m.done = true
			}
			m.percent += increment
			if m.percent > 1.0 {
				m.percent = 1.0
			}
		}
		cmd := m.progress.SetPercent(m.percent)
		return m, tea.Batch(cmd, progressTickCmd())

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case stepMsg:
		m.currentStep = msg.step
		m.message = msg.message
		if msg.step == StepComplete {
			m.percent = 1.0
			m.done = true
		}
		return m, nil

	case progressMsg:
		if msg.percent > m.percent {
			m.percent = msg.percent
		}
		cmd := m.progress.SetPercent(m.percent)
		return m, cmd
	}

	return m, nil
}

func (m ProgressModel) View() string {
	var b strings.Builder

	title := "🚀 Creating Your Express Project 🚀"
	b.WriteString(progressTitleStyle.Render(title))
	b.WriteString("\n\n")

	// Spinner and message
	b.WriteString(fmt.Sprintf(" %s %s\n\n", m.spinner.View(), progressMessageStyle.Render(m.message)))

	// Progress bar
	b.WriteString(m.progress.View())
	b.WriteString("\n\n")

	if m.done {
		b.WriteString(lipgloss.NewStyle().
			Foreground(lipgloss.Color("#4CAF50")).
			Bold(true).
			Render("✅ Project created successfully!"))
		b.WriteString("\n")
	}

	return b.String()
}

func progressTickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*50, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

// ProgressCallback is a function type for progress updates
type ProgressCallback func(step ProgressStep, message string, percent float64)

// RunProgressWithCallback starts the progress TUI and calls the provided function
func RunProgressWithCallback(fn func(ProgressCallback) error) error {
	m := NewProgressModel()
	p := tea.NewProgram(m, tea.WithAltScreen())

	done := make(chan error, 1)

	// Run the actual work in a goroutine
	go func() {
		err := fn(func(step ProgressStep, message string, percent float64) {
			p.Send(stepMsg{step: step, message: message})
			p.Send(progressMsg{percent: percent})
		})
		done <- err
	}()

	// Wait for completion or error
	go func() {
		err := <-done
		if err != nil {
			// Send error message
			p.Send(stepMsg{step: StepComplete, message: fmt.Sprintf("Error: %v", err)})
			time.Sleep(1 * time.Second)
		} else {
			// Mark as complete
			p.Send(stepMsg{step: StepComplete, message: "Complete!"})
			time.Sleep(500 * time.Millisecond)
		}
		p.Quit()
	}()

	_, err := p.Run()
	if err != nil {
		return err
	}
	
	// Check if there was an error in the work
	select {
	case err := <-done:
		return err
	default:
		return nil
	}
}

