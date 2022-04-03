package scenes

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type Start struct {
	width  int
	height int
}

type StartParams struct {
	Width  int
	Height int
}

func NewStart(params StartParams) *Start {
	return &Start{
		width:  params.Width,
		height: params.Height,
	}
}

func (scene Start) Render() string {
	doc := strings.Builder{}

	okButton := activeButtonStyle.Render("Begin (Enter)")
	quitButton := buttonStyle.Render("Quit (q)")

	welcomeMessage := `You are the heart of a Musshi.
	Send heart ❤ beat with any key.`

	message := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(welcomeMessage)
	button := lipgloss.JoinHorizontal(lipgloss.Top, okButton, quitButton)
	ui := lipgloss.JoinVertical(lipgloss.Center, message, button)

	dialog := lipgloss.Place(scene.width, scene.height,
		lipgloss.Center, lipgloss.Center,
		boxStyle.BorderForeground(lipgloss.Color("#874BFD")).Render(ui),
		lipgloss.WithWhitespaceChars("°"),
		lipgloss.WithWhitespaceForeground(subtleColor),
	)

	doc.WriteString(dialog)
	return doc.String()
}
