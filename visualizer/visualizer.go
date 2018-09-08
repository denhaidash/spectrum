package visualizer

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type RenderSettings struct {
	SizeX, SizeY int
	OutputPath   string
	Title        string
}

func Render(spectrum map[int]float64, settings RenderSettings) error {
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = settings.Title
	p.X.Label.Text = "Freq, Hz"
	p.Y.Label.Text = "An"

	err = plotutil.AddScatters(p, preparePlotPoints(spectrum))

	if err != nil {
		return err
	}

	return p.Save(vg.Length(settings.SizeX), vg.Length(settings.SizeY), settings.OutputPath)
}

func preparePlotPoints(spectrum map[int]float64) plotter.XYs {
	pts := make(plotter.XYs, 0, len(spectrum))

	for k, v := range spectrum {
		pts = append(pts, struct {
			X float64
			Y float64
		}{
			X: float64(k),
			Y: v,
		})
	}

	return pts
}
