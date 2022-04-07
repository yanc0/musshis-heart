package scenes

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/guptarohit/asciigraph"
	"github.com/yanc0/musshis-heart/musshi"
)

type Game struct {
	width  int
	height int
	musshi *musshi.Musshi
}

type GameParams struct {
	Width  int
	Height int
	Musshi *musshi.Musshi
}

func NewGame(params GameParams) *Game {
	return &Game{
		width:  params.Width,
		height: params.Height,
		musshi: params.Musshi,
	}
}

func (scene Game) Render() string {
	doc := strings.Builder{}

	heartColor := okColor
	var warningMessage string

	switch scene.musshi.GetCondition() {
	case musshi.TooHighBPM:
		heartColor = warningColor
		warningMessage = "Your Musshi is having heart palpitations. Slow down."
	case musshi.TooLowBPM:
		heartColor = warningColor
		warningMessage = "Your Musshi needs a little more blood. Keep pumping."
	case musshi.VeryHighBPM:
		heartColor = criticalColor
		warningMessage = "Your Musshi is having a heart attack !"
	case musshi.VeryLowBPM:
		heartColor = criticalColor
		warningMessage = "Your Musshi is having a stroke !"

	default:
		heartColor = okColor
	}

	var musshiDraw string
	switch scene.musshi.Activity() {
	case musshi.Sleeping:
		musshiDraw = sleepingMusshi
	case musshi.Playing:
		musshiDraw = playingMusshi
	case musshi.Loving:
		musshiDraw = lovingMusshi
	case musshi.Dying:
		musshiDraw = sleepingMusshi
	default:
		musshiDraw = sleepingMusshi
	}

	lifetimeBox := lipgloss.NewStyle().Align(lipgloss.Right).Padding(1).Foreground(heartColor).Render(fmt.Sprintf("age: %ds / %ds life expectancy", int(scene.musshi.Age().Seconds()), int(scene.musshi.LifeTimeExpectancy.Seconds())))
	warningMessageBox := lipgloss.NewStyle().Align(lipgloss.Right).Padding(1).Foreground(heartColor).Render(warningMessage)

	heartDraw := lipgloss.NewStyle().Align(lipgloss.Center).Foreground(heartColor).Render(heartComponent)
	currentBPM := lipgloss.NewStyle().Align(lipgloss.Left).PaddingTop(1).Foreground(heartColor).Render(fmt.Sprintf("BPM: %d", scene.musshi.Heart.BeatsPerMinute()))
	heartBox := lipgloss.JoinVertical(lipgloss.Center, heartDraw, currentBPM)

	musshiBox := lipgloss.JoinVertical(lipgloss.Center, musshiDraw, string(describeActivity(scene.musshi)))

	data := append(scene.musshi.Heart.Electrocardiogram(), -1, 5)
	ecgPlot := asciigraph.Plot(data, asciigraph.Precision(0))

	ecgPlot = strings.ReplaceAll(ecgPlot, "┤", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "┼", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "-", " ")
	ecgPlot = strings.ReplaceAll(ecgPlot, "1", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "2", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "3", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "4", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "5", "")
	ecgPlot = strings.ReplaceAll(ecgPlot, "0", "")

	ecgBox := lipgloss.NewStyle().Align(lipgloss.Left).Render(ecgPlot)
	heartBox = lipgloss.JoinHorizontal(lipgloss.Center, ecgBox, heartBox)

	tui := lipgloss.JoinVertical(lipgloss.Center, lifetimeBox, warningMessageBox, musshiBox, heartBox)

	bxStyle := boxStyle.Copy().BorderForeground(lipgloss.Color(heartColor))
	doc.WriteString(lipgloss.Place(scene.width, 0,
		lipgloss.Center, lipgloss.Center,
		bxStyle.Render(tui),
		lipgloss.WithWhitespaceChars(" "),
	))

	return doc.String()
}
