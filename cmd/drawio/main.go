package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"github.com/KaymeKaydex/drawio.git"
)

func main() {

	data, err := os.ReadFile("testdata/simple/colored-background.drawio")
	if err != nil {
		log.Fatal(err)
	}

	mxfile, err := drawio.Unmarshal(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mxfile)

	image, err := drawio.Export(*mxfile).ToPNG()
	if err != nil {
		log.Fatal(err)
	}

	output := "/Users/maxim-konovalov/MyProj/drawio/test.png" // output image will live here
	myfile, err := os.Create(output)                          // ... now lets save output image
	if err != nil {
		panic(err)
	}
	defer myfile.Close()
	png.Encode(myfile, image) // output file /tmp/two_rectangles.png
}