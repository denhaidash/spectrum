package main

import (
	"os"

	"spectrum/audio"
	spectrum "spectrum/core"
	"spectrum/visualizer"
)

func main() {
	inputPath := os.Args[1]
	outputPath := os.Args[2]

	pcm, sampleRate, err := audio.ReadWav(inputPath)

	if err != nil {
		panic(err)
	}

	spectrumMap, err := spectrum.GetNormalizedFrequenciesSpectrum(pcm, sampleRate, 1)

	if err != nil {
		panic(err)
	}

	err = visualizer.Render(spectrumMap, visualizer.RenderSettings{
		SizeX:      600,
		SizeY:      500,
		OutputPath: outputPath,
		Title:      "Spectrum",
	})

	if err != nil {
		panic(err)
	}
}
