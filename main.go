package main

// An example Bubble Tea server. This will put an ssh session into alt screen
// and continually print up to date terminal information.

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
	"github.com/yanc0/musshis-heart/musshi"
	"github.com/yanc0/musshis-heart/scenes"
)

var host = flag.String("host", "0.0.0.0", "host to listen on")
var port = flag.Int("port", 2222, "port to listen on")

func main() {
	flag.Parse()
	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", *host, *port)),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", *host, *port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		ticker := time.NewTicker(100 * time.Millisecond)
		for range ticker.C {
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}
	m := &gameState{
		term:    pty.Term,
		width:   pty.Window.Width,
		height:  pty.Window.Height,
		started: false,
	}
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

type gameState struct {
	term   string
	width  int
	height int

	started bool
	ended   bool

	musshi *musshi.Musshi
}

func (m gameState) Init() tea.Cmd {
	return tickCmd()
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Every(time.Millisecond*100, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m gameState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "r":
			startedAt := time.Now()
			m.musshi = musshi.NewMusshi()
			// create musshi with two beats
			m.musshi.Heart.Beat(startedAt.Add(time.Second * -1))
			m.musshi.Heart.Beat(startedAt)
			m.started = true
			m.ended = false
			return m, nil
		default:
			if m.ended {
				return m, nil
			}
			if !m.started {
				startedAt := time.Now()
				m.musshi = musshi.NewMusshi()
				// create musshi with two beats
				m.musshi.Heart.Beat(startedAt.Add(time.Second * -1))
				m.musshi.Heart.Beat(startedAt)
				m.started = true
				m.ended = false
				return m, nil
			}
			m.musshi.Heart.Beat(time.Now())
			if !m.musshi.Alive() {
				m.musshi.SetDeathTime(time.Now())
				m.ended = true
				m.started = false
			}
		}
	case tickMsg:
		if m.started {
			m.musshi.AlterLifeTimeExpectancy()
			if !m.musshi.Alive() {
				m.musshi.SetDeathTime(time.Now())
				m.ended = true
				m.started = false
			}
		}

		return m, tickCmd()
	}

	return m, nil
}

func (m gameState) View() string {
	if m.ended {
		return scenes.NewEnd(scenes.EndParams{
			Width:  m.width,
			Height: m.height,
			Musshi: m.musshi,
		}).Render()
	}
	if !m.started {
		return scenes.NewStart(scenes.StartParams{
			Width:  m.width,
			Height: m.height,
		}).Render()
	}

	return scenes.NewGame(scenes.GameParams{
		Width:  m.width,
		Height: m.height,
		Musshi: m.musshi,
	}).Render()

}
