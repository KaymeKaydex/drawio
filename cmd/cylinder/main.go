package main

import (
	"image/color"

	"github.com/fogleman/gg"
)

func main() {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetLineWidth(float64(5))

	dc.Push()
	dc.DrawEllipse(S/2, S/4, S*7/32, S/16)
	dc.SetStrokeStyle(gg.NewSolidPattern(color.Black))
	dc.Stroke()
	dc.Fill()
	dc.Pop()

	// draw left line
	dc.Push()
	dc.DrawLine(S/2-S*7/32, S/4, S/2-S*7/32, S/2)
	dc.SetColor(color.Black)
	dc.Stroke()
	dc.Fill()
	dc.Pop()

	// draw right line
	dc.Push()
	dc.DrawLine(S/2+S*7/32, S/4, S/2+S*7/32, S/2)
	dc.SetColor(color.Black)
	dc.Stroke()
	dc.Fill()
	dc.Pop()

	// draw 2nd ellipse
	dc.Push()
	dc.DrawEllipse(S/2, S/2, S*7/32, S/16)
	dc.SetStrokeStyle(gg.NewSolidPattern(color.Black))
	dc.Stroke()
	dc.Fill()
	dc.Pop()
	dc.SavePNG("out.png")
}
