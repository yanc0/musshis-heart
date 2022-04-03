package scenes

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/yanc0/musshis-heart/musshi"
)

type End struct {
	width  int
	height int
	musshi *musshi.Musshi
}

type EndParams struct {
	Width  int
	Height int
	Musshi *musshi.Musshi
}

func NewEnd(params EndParams) *End {
	return &End{
		width:  params.Width,
		height: params.Height,
		musshi: params.Musshi,
	}
}

func (scene End) Render() string {
	doc := strings.Builder{}

	okButton := activeButtonStyle.Render("Restart (r)")
	quitButton := buttonStyle.Render("Quit (q)")

	var personnalizedMessage string
	switch {
	case scene.musshi.LifeTimeExpectancy.Seconds() < 35:
		personnalizedMessage = "The heart was a complete mess."
		break
	case scene.musshi.LifeTimeExpectancy.Seconds() < 60:
		personnalizedMessage = "How sad to die so young."
		break
	case scene.musshi.LifeTimeExpectancy.Seconds() < 80:
		personnalizedMessage = "So close to love."
		break
	case scene.musshi.LifeTimeExpectancy.Seconds() < 105:
		personnalizedMessage = "He found love but died bruttally."
		break
	default:
		personnalizedMessage = "He found love and had a great life"
		break
	}

	ageAfterDeath := scene.musshi.DeadAt.Sub(scene.musshi.BornAt)
	endMessage := fmt.Sprintf("Your Musshi lived %d seconds.\n%s", int(ageAfterDeath.Seconds()), personnalizedMessage)

	message := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(endMessage)
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
