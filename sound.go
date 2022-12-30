package main

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

type Sound struct {
	context *oto.Context
	cache   map[string]*mp3.Decoder
}

func (s *Sound) initCache() {
	s.cache = make(map[string]*mp3.Decoder)
}

func (s *Sound) initContext() {
	if s.context != nil {
		panic("More than one context created!")
	}

	c, ready, err := oto.NewContext(44100, 2, 2)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}

	<-ready

	s.context = c
}

func (s *Sound) open(path string) *mp3.Decoder {
	b, err := os.ReadFile(path)
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}

	reader := bytes.NewReader(b)

	decodedMp3, err := mp3.NewDecoder(reader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}

	return decodedMp3
}

func (s *Sound) play(what *mp3.Decoder) {
	player := s.context.NewPlayer(what)

	_, err := player.(io.Seeker).Seek(0, io.SeekStart)
	if err != nil {
		panic("player.Seek failed: " + err.Error())
	}

	player.Play()

	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	player.Close()
}
