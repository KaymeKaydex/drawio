package drawio

import (
	"image"
	"image/color"
	"sync"

	"github.com/fogleman/gg"
)

func Export(file MXFile) *Exporter {
	e := &Exporter{}
	e.f = file
	e.cellsMapWg = sync.WaitGroup{}

	// we need create map with id from wg for draw links faster
	e.cellsMapWg.Add(1)
	go func(e *Exporter) {
		e.cellsMap = createCellsMap(e)
		e.cellsMapWg.Done()
	}(e)

	return e
}

func createCellsMap(e *Exporter) map[string]MXCell {
	m := map[string]MXCell{}
	for _, cell := range e.f.Diagram.MXGraphModel.Root.MXCells {
		m[cell.ID] = cell
	}

	return m
}

type Exporter struct {
	f          MXFile
	cellsMap   map[string]MXCell
	cellsMapWg sync.WaitGroup
}

func (e *Exporter) ToImage() (image.Image, error) {
	f := e.f
	mxCells := f.Diagram.MXGraphModel.Root.MXCells

	dc := gg.NewContext(f.Diagram.MXGraphModel.PageWidth, f.Diagram.MXGraphModel.PageHeight)

	dc.DrawRectangle(0, 0, float64(f.Diagram.MXGraphModel.PageWidth), float64(f.Diagram.MXGraphModel.PageHeight))

	// set background

	// if background is empty
	if f.Diagram.MXGraphModel.Background == "" {
		dc.SetRGBA(0, 0, 0, 1)
	} else {
		color, err := parseHexColor(f.Diagram.MXGraphModel.Background)
		if err != nil {
			return nil, err
		}

		dc.SetColor(color)
	}
	dc.Fill()

	for _, cell := range mxCells {
		if cell.MXGeometry == nil {
			continue
		}

		dc.DrawRectangle(float64(cell.MXGeometry.X),
			float64(cell.MXGeometry.Y),
			float64(cell.MXGeometry.Width),
			float64(cell.MXGeometry.Height))

		if cell.FillColor == "" {
			dc.SetColor(color.White)
		} else {
			cellColor, err := parseHexColor(cell.FillColor)
			if err != nil {
				return nil, err
			}

			dc.SetColor(cellColor)
		}
		dc.Fill()
		// draw stroke
		// set black if empty
		if cell.StrokeColor == "" {
			dc.DrawRectangle(float64(cell.MXGeometry.X),
				float64(cell.MXGeometry.Y),
				float64(cell.MXGeometry.Width),
				float64(cell.MXGeometry.Height))

			dc.SetStrokeStyle(gg.NewSolidPattern(color.Black))
			dc.Stroke()
			dc.Fill()
		} // todo: if not empty?

		const S = 1024

		if cell.Value != "" {
			dc.SetColor(color.Black)
			// draw inline value if exists
			dc.DrawStringAnchored(cell.Value,
				float64(cell.MXGeometry.X)+float64(cell.MXGeometry.Width)/2,
				float64(cell.MXGeometry.Y)+float64(cell.MXGeometry.Height)/2,
				0.5, 0.5)
			dc.Fill()
		}

		// there is link
		if cell.Source != "" {
			if cell.Source != "" { // if there is link between 2 cells

			} else { // if there is link between cell and virtual point

			}
		}
	}

	return dc.Image(), nil
}
