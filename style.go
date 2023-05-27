package drawio

import "strings"

// Style example: ellipse;whiteSpace=wrap;html=1;aspect=fixed
type Style struct {
	IsEllipse bool
	FillColor string // only for ellipse
}

func ParseStyle(style string) Style {
	s := Style{}

	if strings.Contains(style, "ellipse") {
		s.IsEllipse = true
	}

	params := strings.Split(style, ";")
	for _, p := range params {
		if strings.HasPrefix(p, "fillColor") {
			s.FillColor = p[10:]
		}
	}
	return s
}
