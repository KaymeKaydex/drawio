// https://blog.golang.org/go-imagedraw-package

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {

	new_png_file := "/Users/maxim-konovalov/MyProj/drawio/two_rectangles.png" // output image will live here

	myimage := image.NewRGBA(image.Rect(0, 0, 220, 220)) // x1,y1,  x2,y2 of background rectangle
	mygreen := color.RGBA{0, 100, 0, 255}                //  R, G, B, Alpha

	// backfill entire background surface with color mygreen
	draw.Draw(myimage, myimage.Bounds(), &image.Uniform{mygreen}, image.ZP, draw.Src)

	red_rect := image.Rect(60, 80, 120, 160) //  geometry of 2nd rectangle which we draw atop above rectangle
	myred := color.RGBA{200, 0, 0, 255}

	// create a red rectangle atop the green surface
	draw.Draw(myimage, red_rect, &image.Uniform{myred}, image.ZP, draw.Src)

	myfile, err := os.Create(new_png_file) // ... now lets save output image
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, myimage) // output file /tmp/two_rectangles.png
}
