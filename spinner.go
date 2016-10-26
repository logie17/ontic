package main

import (
	"fmt"
	"sync"
	"time"
)

type Spinner struct {
	Charset     []rune
	FrameRate   time.Duration
	stopRunning chan struct{}
	stopOnce    sync.Once
	Title       string
}

func NewSpinner(title string) *Spinner {
	return &Spinner{
		Charset:     []rune(`⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏`),
		FrameRate:   time.Millisecond * 150,
		stopRunning: make(chan struct{}),
		Title:       title,
	}
}

func (s *Spinner) Start() *Spinner {
	go s.writer()
	return s
}

func (s *Spinner) Stop() {
	s.stopOnce.Do(func() {
		close(s.stopRunning)
	})
}

func (s *Spinner) animate() {
	for i := 0; i < len(s.Charset); i++ {
		fmt.Printf("\r  \033[36m%s\033[m %s\r", s.Title, string(s.Charset[i]))
		time.Sleep(s.FrameRate)
	}
}

func (s *Spinner) writer() {
	s.animate()
	for {
		select {
		case <-s.stopRunning:
			return
		default:
			s.animate()
		}
	}
}
