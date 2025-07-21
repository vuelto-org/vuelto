/*
 * Copyright (C) 2025 vuelto-org
 *
 * This file is part of the Vuelto project, licensed under the VL-Cv1.1 License.
 * Primary License: GNU GPLv3 or later (see <https://www.gnu.org/licenses/>).
 * If unmaintained, this software defaults to the MIT License as per Vuelto License V1.1,
 * at which point the copyright no longer applies.
 *
 * Distributed WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 */

package vuelto

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

type AudioPlayer struct {
	Streamer beep.StreamSeekCloser
	Format   beep.Format
	File     *os.File
	Done     chan bool
}

// Opens a audio file. It supports two file formats: WAV and MP3.
// Plays the audio file using Start function. Stop the audio file using Stop function.
// Close the audio file using Close function.
func OpenAudioFile(filePath string) (*AudioPlayer, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening audio file: %s", filePath)
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	switch strings.ToLower(filepath.Ext(filePath)) {
	case ".wav":
		streamer, format, err = wav.Decode(f)
	case ".mp3":
		streamer, format, err = mp3.Decode(f)
	default:
		f.Close()
		return nil, fmt.Errorf("Unsupported audio format: %s", filepath.Ext(filePath))
	}

	if err != nil {
		f.Close()
		return nil, fmt.Errorf("Error decoding audio file: %s", filePath)
	}

	return &AudioPlayer{
		Streamer: streamer,
		Format:   format,
		File:     f,
		Done:     make(chan bool),
	}, nil
}

// Starts playing the audio file.
func (a *AudioPlayer) Start() {
	speaker.Init(a.Format.SampleRate, a.Format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(a.Streamer, beep.Callback(func() {
		a.Done <- true
	})))
}

// Stops playing the audio file.
func (a *AudioPlayer) Stop() {
	speaker.Clear()
	a.Streamer.Seek(0)
}

// Closes the audio file.
func (a *AudioPlayer) Close() {
	a.File.Close()
}
