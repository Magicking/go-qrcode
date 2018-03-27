package qrcode

import (
	"bytes"
	"text/template"
)

var svgTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<svg version="1.1" baseProfile="tiny" xmlns="http://www.w3.org/2000/svg" width="{{.Width}}" height="{{.Height}}">
<rect shape-rendering="optimizeSpeed"  x="0" y="0" width="{{.Width}}" height="{{.Height}}" fill="white" />
{{range .Bitmap}}<rect shape-rendering="optimizeSpeed"  x="{{.X}}" y="{{.Y}}" width="{{.PixWidth}}" height="{{.PixHeight}}" fill="black" />
{{end}}
</svg>`

func SVG(content string, level RecoveryLevel, size int) (string, error) {
	var q *QRCode

	q, err := New(content, level)

	if err != nil {
		return "", err
	}

	var bitmap []interface{}
	qbitmap := q.Bitmap()
	pixWidth := 5
	pixHeight := 5
	for x, line := range qbitmap {
		for y, e := range line {
			if e {
				bitmap = append(bitmap, map[string]interface{}{
					"X":         2 + x*pixWidth,
					"Y":         2 + y*pixWidth,
					"PixWidth":  5,
					"PixHeight": 5,
				})
			}
		}
	}
	Width := pixWidth * len(qbitmap[0]) //TODO check order
	Height := pixHeight * len(qbitmap)
	tmpl, err := template.New("SVG").Parse(svgTemplate)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	tmpl.Execute(&b, map[string]interface{}{
		"PixWidth":  pixWidth,
		"PixHeight": pixHeight,
		"Width":     Width,
		"Height":    Height,
		"Bitmap":    bitmap,
	})
	return b.String(), nil
}
