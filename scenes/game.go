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

	switch scene.musshi.GetCondition() {
	case musshi.TooHighBPM:
		heartColor = warningColor
	case musshi.TooLowBPM:
		heartColor = warningColor
	case musshi.VeryHighBPM:
		heartColor = criticalColor
	case musshi.VeryLowBPM:
		heartColor = criticalColor
	default:
		heartColor = okColor
	}

	var musshiDraw string
	switch scene.musshi.Activity() {
	case musshi.Sleeping:
		musshiDraw = sleepingMusshi
	case musshi.Playing:
		musshiDraw = playingMusshi
	case musshi.Reproducing:
		musshiDraw = reproducingMusshi
	case musshi.Dying:
		musshiDraw = deadMusshi
	default:
		musshiDraw = deadMusshi
	}

	lifetimeBox := lipgloss.NewStyle().Align(lipgloss.Right).Render(fmt.Sprintf("%d / %d seconds", int(scene.musshi.Age().Seconds()), int(scene.musshi.LifeTimeExpectancy.Seconds())))

	heartDraw := lipgloss.NewStyle().Align(lipgloss.Center).Foreground(heartColor).Render(heartComponent)
	currentBPM := lipgloss.NewStyle().Align(lipgloss.Left).PaddingTop(1).Foreground(heartColor).Render(fmt.Sprintf("BPM: %d", scene.musshi.Heart.BeatsPerMinute()))
	heartBox := lipgloss.JoinVertical(lipgloss.Center, heartDraw, currentBPM)

	musshiBox := lipgloss.JoinVertical(lipgloss.Center, musshiDraw, string(describeActivity(scene.musshi)))

	data := append(scene.musshi.Heart.Electrocardiogram(), -2, 5)
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

	tui := lipgloss.JoinVertical(lipgloss.Center, lifetimeBox, musshiBox, heartBox)

	doc.WriteString(lipgloss.Place(scene.width, 0,
		lipgloss.Center, lipgloss.Center,
		boxStyle.BorderForeground(lipgloss.Color(heartColor)).Render(tui),
		lipgloss.WithWhitespaceChars(" "),
	))

	return doc.String()
}
