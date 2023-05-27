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

	dc.Push()
	dc.DrawRectangle(0, 0, float64(f.Diagram.MXGraphModel.PageWidth), float64(f.Diagram.MXGraphModel.PageHeight))

	// set background

	// if background is empty
	if f.Diagram.MXGraphModel.Background != "" {
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

		style := ParseStyle(cell.Style)

		if style.IsEllipse {
			cell.FillColor = style.FillColor
			cell.StrokeColor = style.StrokeColor

			err := e.drawCircle(dc, cell)
			if err != nil {
				return nil, err
			}
		} else {
			err := e.drawRectangle(dc, cell)
			if err != nil {
				return nil, err
			}
		}

		if cell.Value != "" {
			dc.SetColor(color.Black)
			// draw inline value if exists
			dc.DrawStringAnchored(cell.Value,
				float64(cell.MXGeometry.X)+float64(cell.MXGeometry.Width)/2,
				float64(cell.MXGeometry.Y)+float64(cell.MXGeometry.Height)/2,
				0.5, 0.5)
			dc.Fill()
		}
		dc.Stroke()

		// there is link
		if cell.Source != "" {
			if cell.Target != "" { // if there is link between 2 cells
				e.cellsMapWg.Wait()      // wait if map not done
				dc.SetColor(color.Black) // todo

				sourceGeometry := e.cellsMap[cell.Source].MXGeometry
				targetGeometry := e.cellsMap[cell.Target].MXGeometry

				dc.DrawLine(
					float64(sourceGeometry.X)+float64(sourceGeometry.Width),
					float64(sourceGeometry.Y)+float64(sourceGeometry.Height)/2,
					float64(targetGeometry.X)+float64(targetGeometry.Width),
					float64(targetGeometry.Y)+float64(targetGeometry.Height)/2)
				dc.Stroke()
				dc.Fill()
			} else { // if there is link between cell and virtual point

			}
		}
	}

	dc.Pop()

	return dc.Image(), nil
}

func (e *Exporter) drawCircle(dc *gg.Context, cell MXCell) error {
	radius := cell.MXGeometry.Width / 2

	dc.DrawCircle(cell.MXGeometry.X+radius, cell.MXGeometry.Y+radius, radius)

	// set color
	if cell.FillColor == "" {
		dc.SetColor(color.White)
	} else {
		cellColor, err := parseHexColor(cell.FillColor)
		if err != nil {
			return err
		}

		dc.SetColor(cellColor)
	}
	dc.Fill()

	if cell.StrokeColor == "" {
		dc.DrawCircle(cell.MXGeometry.X+radius, cell.MXGeometry.Y+radius, radius)

		dc.SetStrokeStyle(gg.NewSolidPattern(color.Black))
		dc.Stroke()
		dc.Fill()
	} // todo: if not empty?

	return nil
}

func (e *Exporter) drawRectangle(dc *gg.Context, cell MXCell) error {
	dc.DrawRectangle(cell.MXGeometry.X,
		cell.MXGeometry.Y,
		cell.MXGeometry.Width,
		cell.MXGeometry.Height)

	if cell.FillColor == "" {
		dc.SetColor(color.White)
	} else {
		cellColor, err := parseHexColor(cell.FillColor)
		if err != nil {
			return err
		}

		dc.SetColor(cellColor)
	}
	dc.Fill()
	// draw stroke
	// set black if empty
	if cell.StrokeColor == "" {
		dc.DrawRectangle(cell.MXGeometry.X,
			cell.MXGeometry.Y,
			cell.MXGeometry.Width,
			cell.MXGeometry.Height)

		dc.SetStrokeStyle(gg.NewSolidPattern(color.Black))
		dc.Stroke()
		dc.Fill()
	} // todo: if not empty?

	return nil
}
