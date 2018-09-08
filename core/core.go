package core

import (
	"errors"
	"math/cmplx"

	"spectrum/fft"
	"spectrum/helper"
)

func GetNormalizedFrequenciesSpectrum(samples []int, sampleRate int, freqStep int) (map[int]float64, error) {
	return normalizeFrequenciesSpectrum(GetFrequenciesSpectrum(samples, sampleRate, freqStep))
}

func GetFrequenciesSpectrum(samples []int, sampleRate int, freqStep int) map[int]float64 {
	complexSamples := toComplexSamples(samples)

	fft.Transform(&complexSamples)

	freqSamplingStep := float64(sampleRate) / float64(len(complexSamples))
	scaleStep := float64(freqStep) / freqSamplingStep
	freqLimit := float64(sampleRate) / 2.0

	result := make(map[int]float64)

	for i := 0.0; i*freqSamplingStep <= freqLimit; i += scaleStep {
		result[int(i*freqSamplingStep)] = cmplx.Abs(complexSamples[int(i)])
	}

	return result
}

func toComplexSamples(samples []int) []complex128 {
	complexSamples := make([]complex128, helper.FindNearestBiggerPowerOf2(uint(len(samples))))

	for i, s := range samples {
		complexSamples[i] = complex(float64(s), 0)
	}

	return complexSamples
}

func normalizeFrequenciesSpectrum(spectrum map[int]float64) (map[int]float64, error) {
	maxValue := helper.FindMaxValue(spectrum)

	if maxValue == 0.0 {
		return nil, errors.New("Can't normalize spectrum. Max value is 0")
	}

	for k, v := range spectrum {
		spectrum[k] = v / maxValue
	}

	return spectrum, nil
}
