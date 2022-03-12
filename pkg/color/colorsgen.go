// The following directive is necessary to make the package coherent:

//go:build ignore
// +build ignore

// This program generates xtermcolors.go. It can be invoked by running
// go generate
package main

import (
	"fmt"
	"image/color"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var customFuncs = template.FuncMap{
	"inc": func(i int) int {
		return i + 16
	},
	"hex": func(c color.Color) string {
		r, g, b, _ := c.RGBA()
		return fmt.Sprintf("#%02x%02x%02x", uint8(r), uint8(g), uint8(b))
	},
	"colon": func(i int) string {
		return strconv.Itoa(i) + ":"
	},
}

var packageTemplate = template.Must(template.New("").Funcs(customFuncs).Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
package color

var colors = []string{
	// ANSI colors
{{- range $index, $element := .AnsiColors }}
	{{ printf "%-4s%q" (colon $index) (hex $element)}},
{{- end }}

	// XTERM colors
{{- range $index, $element := .XtermColors }}
	{{ printf "%-5s%q" (colon (inc $index)) $element}},
{{- end }}
}
`))

var ansiColors = []color.Color{
	color.Black,                              // Black
	color.RGBA{0x00CD, 0x00, 0x00, 0x00},     // Red
	color.RGBA{0x00, 0x00CD, 0x00, 0x00},     // Green
	color.RGBA{0x00CD, 0x00CD, 0x00, 0x00},   // Yellow
	color.RGBA{0x00, 0x00, 0x00EE, 0x00},     // Blue
	color.RGBA{0x00CD, 0x00, 0x00CD, 0x00},   // Magent
	color.RGBA{0x00, 0x00CD, 0x00CD, 0x00},   // Cyan
	color.RGBA{0x00E5, 0x00E5, 0x00E5, 0x00}, // Grey

	color.RGBA{0x007F, 0x007F, 0x007F, 0x00}, // Dark Grey
	color.RGBA{0x00FF, 0x00, 0x00, 0x00},     // Light Red
	color.RGBA{0x00, 0x00FF, 0x00, 0x00},     // Light Green
	color.RGBA{0x00FF, 0x00FF, 0x00, 0x00},   // Light Yellow
	color.RGBA{0x005C, 0x005C, 0x00FF, 0x00}, // Light Blue
	color.RGBA{0x00FF, 0x00, 0x00FF, 0x00},   // Light Magent
	color.RGBA{0x00, 0x00FF, 0x00FF, 0x00},   // Light Cyan
	color.White,                              // White
}

func main() {
	f, err := os.Create("colors.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	colored := []int{0}
	for i := 95; i < 256; i += 40 {
		colored = append(colored, i)
	}

	colorPalette := []string{}

	for _, r := range colored {
		for _, g := range colored {
			for _, b := range colored {
				colorPalette = append(colorPalette, fmt.Sprintf("#%02x%02x%02x", r, g, b))
			}
		}
	}

	grayscale := []int{}
	for i := 8; i < 240; i += 10 {
		grayscale = append(grayscale, i)
	}

	grayscalePalette := []string{}
	for _, rgb := range grayscale {
		grayscalePalette = append(grayscalePalette, fmt.Sprintf("#%02x%02x%02x", rgb, rgb, rgb))
	}

	colors := append(colorPalette, grayscalePalette...)

	packageTemplate.Execute(f, struct {
		Timestamp   time.Time
		XtermColors []string
		AnsiColors  []color.Color
	}{
		Timestamp:   time.Now(),
		XtermColors: colors,
		AnsiColors:  ansiColors,
	})
}
