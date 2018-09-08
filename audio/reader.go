package audio

import (
	"os"

	"github.com/go-audio/wav"
)

func ReadWav(path string) ([]int, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	d := wav.NewDecoder(file)
	pcm, err := d.FullPCMBuffer()

	if err != nil {
		return nil, 0, err
	}

	return pcm.Data, pcm.Format.SampleRate, nil
}
