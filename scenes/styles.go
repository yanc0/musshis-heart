package scenes

import "github.com/charmbracelet/lipgloss"

var buttonStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#888B7E")).
	Padding(0, 3).
	MarginTop(1)

var activeButtonStyle = buttonStyle.Copy().
	Foreground(lipgloss.Color("#FFF7DB")).
	Background(lipgloss.Color("#F25D94")).
	MarginRight(2).
	Underline(true)

var boxStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderTop(true).
	BorderLeft(true).
	BorderRight(true).
	BorderBottom(true)

var subtleColor = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

var criticalColor = lipgloss.Color("#eb4034")
var warningColor = lipgloss.Color("#ebb734")
var okColor = lipgloss.Color("#13ad05")
