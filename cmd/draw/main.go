// https://blog.golang.org/go-imagedraw-package

package main

import "github.com/fogleman/gg"

func main() {

	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(100, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.Image()
	dc.SavePNG("out.png")
}
