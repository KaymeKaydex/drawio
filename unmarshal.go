package drawio

import (
	"encoding/xml"
	"time"
)

type MXFile struct {
	Host     string    `xml:"host,attr"`
	Modified time.Time `xml:"modified,attr"`
	Agent    string    `xml:"agent,attr"`
	Version  string    `xml:"version,attr"`
	Etag     string    `xml:"etag,attr"`
	Type     string    `xml:"type,attr"`

	Diagram Diagram `xml:"diagram"`
}

type Diagram struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr"`

	MXGraphModel MXGraphModel `xml:"mxGraphModel"`
}

type MXGraphModel struct {
	DX         int    `xml:"dx,attr"`
	DY         int    `xml:"dy,attr"`
	Grid       bool   `xml:"grid,attr"`
	PageWidth  int    `xml:"pageWidth,attr"`
	PageHeight int    `xml:"pageHeight,attr"`
	Background string `xml:"background,attr,omitempty"` // hex color

	Root Root `xml:"root"`
}

type Root struct {
	MXCells []MXCell `xml:"mxCell,omitempty"`
}

type MXCell struct {
	ID          string `xml:"id,attr"`
	Parent      string `xml:"parent,attr,omitempty"`
	Value       string `xml:"value,attr,omitempty"`
	Style       string `xml:"style,attr,omitempty"`
	Vertex      bool   `xml:"vertex,attr,omitempty"`
	Source      string `xml:"source,attr,omitempty"`      // только для связующих линий
	Target      string `xml:"target,attr,omitempty"`      // только для связующих линий
	FillColor   string `xml:"fillColor,attr,omitempty"`   // цвет фона
	StrokeColor string `xml:"strokeColor,attr,omitempty"` // цвет рамки

	MXGeometry *MXGeometry `xml:"mxGeometry,omitempty"`
}

type MXGeometry struct {
	// x="340" y="230" width="200" height="110"
	X       float64  `xml:"x,attr"`
	Y       float64  `xml:"y,attr"`
	Width   float64  `xml:"width,attr"`
	Height  float64  `xml:"height,attr"`
	As      string   `xml:"as,attr"`
	MXPoint *MXPoint `xml:"mxPoint,omitempty"` // for link without target; for them target is some point
}

type MXPoint struct {
	X float64 `xml:"x,attr"`
	Y float64 `xml:"y,attr"`
}

func Unmarshal(data []byte) (*MXFile, error) {
	f := &MXFile{}

	err := xml.Unmarshal(data, f)
	if err != nil {
		return nil, err
	}

	return f, nil
}
