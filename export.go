package drawio

import (
	"image"
	"image/color"
	"image/draw"
)

func Export(file MXFile) *Exporter {
	return &Exporter{
		f: file,
	}
}

type Exporter struct {
	f MXFile
}

func (e *Exporter) ToPNG() (*image.RGBA, error) {
	f := e.f

	myimage := image.NewRGBA(
		image.Rect(
			f.Diagram.MXGraphModel.PageWidth,
			f.Diagram.MXGraphModel.PageWidth,
			f.Diagram.MXGraphModel.PageHeight,
			f.Diagram.MXGraphModel.PageHeight),
	) // x1,y1,  x2,y2 of background rectangle

	backHexColor := f.Diagram.MXGraphModel.Background
	if backHexColor == "" {
		// backfill entire background surface with color mygreen
		draw.Draw(myimage, myimage.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	} else {
		backgroundColor, err := parseHexColor(f.Diagram.MXGraphModel.Background)
		if err != nil {
			return nil, err
		}

		draw.Draw(myimage, myimage.Bounds(), &image.Uniform{C: backgroundColor}, image.Point{}, draw.Src)
	}

	return myimage, nil
}
