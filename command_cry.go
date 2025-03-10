package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/speaker"
	"github.com/gopxl/beep/vorbis"
)

var speakerOnce sync.Once

func commandCry(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New(usageError("cry"))
	}

	pokemon := args[0]

	pokemonInfo, exists := pokedexData.PokemonMap[pokemon]
	if !exists {
		return fmt.Errorf("you have yet to catch %s", pokemon)
	}

	cryURL := "https://raw.githubusercontent.com/PokeAPI/cries/main/cries/pokemon/latest/" + strconv.Itoa(pokemonInfo.ID) + ".ogg"

	done := make(chan struct{})
	errc := make(chan error, 1)

	go func() {
		fmt.Printf("%s lets out a cry...", pokemon)
		fmt.Println()

		err := downloadAndPlayCry(cryURL)
		if err != nil {
			errc <- err
		}
		close(done)
	}()

	select {
	case err := <-errc:
		return fmt.Errorf("error playing cry: %w", err)
	case <-done:
		return nil
	case <-time.After(5 * time.Second):
		return errors.New("error playing cry: timed out")
	}
}

func downloadAndPlayCry(url string) error {
	tempFile, err := os.CreateTemp("", "pokemon-cry-*.ogg")
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}
	defer os.Remove(tempFile.Name())

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading audio: %w", err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving audio: %w", err)
	}
	tempFile.Close()

	cry, err := os.Open(tempFile.Name())
	if err != nil {
		return fmt.Errorf("error opening audio: %w", err)
	}
	defer cry.Close()

	streamer, format, err := vorbis.Decode(cry)
	if err != nil {
		return fmt.Errorf("error decoding audio: %w", err)
	}
	defer streamer.Close()

	var initErr error
	speakerOnce.Do(func() {
		initErr = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	})
	if initErr != nil {
		return fmt.Errorf("error initializing speaker: %w", initErr)
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	select {
	case <-done:
		return nil
	case <-time.After(4 * time.Second): // Slightly shorter than the outer timeout
		return errors.New("error playing cry: timed out")
	}
}
