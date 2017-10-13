// Generated code do not edit. Run `go generate volcano_example.go`.

// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate ./volcano

//+build ignore

package main

import (
	"image/color"

	"github.com/gonum/matrix/mat64"
	"github.com/skiesel/plot"
	"github.com/skiesel/plot/palette"
	"github.com/skiesel/plot/plotter"
	"github.com/skiesel/plot/vg"
	"github.com/skiesel/plot/vg/draw"
)

type deciGrid struct{ mat64.Matrix }

func (g deciGrid) Dims() (c, r int)   { r, c = g.Matrix.Dims(); return c, r }
func (g deciGrid) Z(c, r int) float64 { return g.Matrix.At(r, c) }
func (g deciGrid) X(c int) float64 {
	_, n := g.Matrix.Dims()
	if c < 0 || c >= n {
		panic("index out of range")
	}
	return 10 * float64(c)
}
func (g deciGrid) Y(r int) float64 {
	m, _ := g.Matrix.Dims()
	if r < 0 || r >= m {
		panic("index out of range")
	}
	return 10 * float64(r)
}

func main() {
	var levels []float64
	for l := 100.5; l < volcano.Matrix.(*mat64.Dense).Max(); l += 5 {
		levels = append(levels, l)
	}
	c := plotter.NewContour(volcano, levels, palette.Rainbow(len(levels), (palette.Yellow+palette.Red)/2, palette.Blue, 1, 1, 1))
	quarterStyle := draw.LineStyle{
		Color:  color.Black,
		Width:  vg.Points(0.5),
		Dashes: []vg.Length{0.2, 0.4},
	}
	halfStyle := draw.LineStyle{
		Color:  color.Black,
		Width:  vg.Points(0.5),
		Dashes: []vg.Length{5, 2, 1, 2},
	}
	c.LineStyles = append(c.LineStyles, quarterStyle, halfStyle, quarterStyle)

	h := plotter.NewHeatMap(volcano, palette.Heat(len(levels)*2, 1))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Maunga Whau Volcano"

	p.Add(h)
	p.Add(c)

	p.X.Padding = 0
	p.Y.Padding = 0
	_, p.X.Max, _, p.Y.Max = h.DataRange()

	name := "example_volcano"

	for _, ext := range []string{
		".eps",
		".pdf",
		".svg",
		".png",
		".tiff",
		".jpg",
	} {
		if err := p.Save(4, 4, name+ext); err != nil {
			panic(err)
		}
	}
}

// Data extracted from RDatasets volcano data for the Maunga Whau volcano topographic data.
var volcano = deciGrid{mat64.NewDense(87, 61, []float64{
	100, 100, 101, 101, 101, 101, 101, 100, 100, 100, 101, 101, 102, 102, 102, 102, 103, 104, 103, 102, 101, 101, 102, 103, 104, 104, 105, 107, 107, 107, 108, 108, 110, 110, 110, 110, 110, 110, 110, 110, 108, 108, 108, 107, 107, 108, 108, 108, 108, 108, 107, 107, 107, 107, 106, 106, 105, 105, 104, 104, 103,
	101, 101, 102, 102, 102, 102, 102, 101, 101, 101, 102, 102, 103, 103, 103, 103, 104, 105, 104, 103, 102, 102, 103, 105, 106, 106, 107, 109, 110, 110, 110, 110, 111, 112, 113, 114, 116, 115, 114, 112, 110, 110, 110, 109, 108, 109, 109, 109, 109, 108, 108, 108, 108, 107, 107, 106, 106, 105, 105, 104, 104,
	102, 102, 103, 103, 103, 103, 103, 102, 102, 102, 103, 103, 104, 104, 104, 104, 105, 106, 105, 104, 104, 105, 106, 107, 108, 110, 111, 113, 114, 115, 114, 115, 116, 118, 119, 119, 121, 121, 120, 118, 116, 114, 112, 111, 110, 110, 110, 110, 109, 109, 109, 109, 108, 108, 107, 107, 106, 106, 105, 105, 104,
	103, 103, 104, 104, 104, 104, 104, 103, 103, 103, 103, 104, 104, 104, 105, 105, 106, 107, 106, 106, 106, 107, 108, 110, 111, 114, 117, 118, 117, 119, 120, 121, 122, 124, 125, 126, 127, 127, 126, 124, 122, 120, 117, 116, 113, 111, 110, 110, 110, 109, 109, 109, 109, 108, 108, 107, 107, 106, 106, 105, 105,
	104, 104, 105, 105, 105, 105, 105, 104, 104, 103, 104, 104, 105, 105, 105, 106, 107, 108, 108, 108, 109, 110, 112, 114, 115, 118, 121, 122, 121, 123, 128, 131, 129, 130, 131, 131, 132, 132, 131, 130, 128, 126, 122, 119, 115, 114, 112, 110, 110, 110, 110, 110, 109, 109, 108, 107, 107, 107, 106, 106, 105,
	105, 105, 105, 106, 106, 106, 106, 105, 105, 104, 104, 105, 105, 106, 106, 107, 109, 110, 110, 112, 113, 115, 116, 118, 119, 121, 124, 126, 126, 129, 134, 137, 137, 136, 136, 135, 136, 136, 136, 135, 133, 129, 126, 122, 118, 116, 115, 113, 111, 110, 110, 110, 110, 109, 108, 108, 108, 107, 107, 106, 106,
	105, 106, 106, 107, 107, 107, 107, 106, 106, 105, 105, 106, 106, 107, 108, 109, 111, 113, 114, 116, 118, 120, 121, 122, 123, 125, 127, 129, 130, 135, 140, 142, 142, 142, 141, 140, 140, 140, 140, 139, 137, 134, 129, 125, 121, 118, 116, 114, 112, 110, 110, 110, 111, 110, 109, 109, 108, 108, 107, 107, 106,
	106, 107, 107, 108, 108, 108, 108, 107, 107, 106, 106, 107, 108, 108, 110, 113, 115, 117, 118, 120, 122, 124, 125, 127, 128, 129, 131, 134, 135, 141, 146, 147, 146, 146, 145, 144, 144, 144, 143, 142, 141, 139, 135, 130, 126, 122, 118, 116, 114, 112, 112, 113, 112, 110, 110, 109, 109, 108, 108, 107, 106,
	107, 108, 108, 109, 109, 109, 109, 108, 108, 107, 108, 108, 110, 111, 113, 116, 118, 120, 123, 125, 127, 129, 130, 132, 134, 135, 137, 139, 142, 146, 152, 152, 151, 151, 150, 149, 148, 148, 146, 145, 143, 142, 139, 135, 131, 127, 122, 119, 117, 115, 115, 115, 114, 112, 110, 110, 109, 109, 108, 107, 107,
	108, 109, 109, 110, 110, 110, 110, 109, 109, 108, 110, 110, 113, 116, 118, 120, 122, 125, 127, 129, 133, 136, 138, 140, 141, 142, 148, 150, 151, 156, 158, 159, 158, 157, 158, 158, 154, 151, 149, 148, 146, 144, 141, 137, 134, 130, 125, 122, 120, 118, 117, 117, 115, 113, 111, 110, 110, 109, 108, 107, 107,
	109, 110, 110, 111, 111, 111, 111, 110, 110, 110, 112, 114, 118, 121, 123, 125, 127, 129, 133, 137, 141, 143, 145, 146, 148, 150, 154, 156, 159, 161, 162, 163, 164, 163, 164, 164, 160, 157, 154, 151, 149, 146, 144, 140, 137, 133, 129, 126, 124, 121, 119, 118, 116, 114, 112, 111, 110, 109, 108, 107, 106,
	110, 110, 111, 113, 112, 111, 113, 112, 112, 114, 116, 119, 121, 124, 127, 129, 133, 138, 143, 146, 149, 149, 151, 153, 154, 157, 159, 160, 163, 165, 166, 167, 168, 168, 168, 168, 166, 162, 159, 157, 154, 152, 149, 144, 140, 136, 133, 131, 128, 125, 122, 119, 117, 115, 113, 111, 110, 109, 108, 107, 106,
	110, 111, 113, 115, 114, 113, 114, 114, 115, 117, 119, 121, 124, 126, 129, 133, 140, 145, 150, 154, 155, 155, 157, 159, 161, 162, 164, 165, 167, 168, 169, 170, 172, 174, 172, 172, 171, 169, 166, 163, 161, 158, 153, 148, 143, 140, 137, 134, 131, 128, 125, 120, 118, 116, 114, 112, 110, 109, 108, 107, 105,
	111, 113, 115, 117, 116, 115, 116, 117, 117, 119, 121, 124, 126, 128, 132, 137, 143, 151, 156, 161, 161, 162, 163, 165, 166, 167, 168, 170, 171, 173, 175, 177, 179, 178, 177, 176, 176, 174, 171, 169, 165, 161, 156, 152, 148, 144, 140, 138, 135, 131, 127, 123, 119, 117, 115, 113, 111, 110, 108, 106, 105,
	114, 115, 117, 117, 117, 118, 119, 119, 120, 121, 124, 126, 128, 131, 137, 143, 150, 156, 160, 163, 165, 168, 170, 171, 172, 173, 174, 175, 177, 179, 180, 182, 183, 183, 183, 183, 180, 178, 177, 172, 168, 164, 160, 156, 152, 148, 144, 141, 138, 134, 130, 126, 121, 117, 114, 112, 110, 110, 108, 106, 104,
	116, 118, 118, 118, 120, 121, 121, 122, 122, 123, 125, 128, 130, 134, 141, 147, 152, 156, 160, 165, 168, 170, 174, 176, 179, 180, 181, 181, 182, 182, 183, 184, 186, 187, 187, 184, 184, 181, 180, 176, 172, 168, 165, 161, 157, 153, 149, 145, 142, 138, 133, 129, 125, 120, 115, 111, 110, 110, 108, 106, 104,
	118, 120, 120, 121, 122, 123, 124, 124, 125, 126, 127, 129, 132, 135, 142, 149, 153, 157, 161, 166, 170, 174, 178, 180, 182, 183, 184, 184, 185, 186, 186, 187, 189, 189, 189, 189, 189, 186, 182, 179, 175, 171, 168, 165, 162, 157, 152, 149, 145, 141, 137, 131, 125, 120, 116, 111, 110, 110, 108, 106, 104,
	120, 121, 122, 123, 124, 125, 126, 127, 127, 128, 130, 132, 134, 137, 142, 151, 155, 158, 162, 169, 172, 176, 181, 183, 184, 186, 187, 188, 189, 189, 189, 189, 190, 190, 191, 190, 190, 188, 186, 183, 180, 175, 171, 168, 165, 161, 157, 152, 149, 145, 141, 134, 127, 121, 116, 112, 110, 110, 108, 106, 104,
	120, 122, 125, 126, 126, 127, 128, 129, 130, 130, 132, 134, 136, 139, 145, 152, 157, 160, 167, 172, 175, 178, 181, 185, 186, 188, 190, 191, 192, 193, 193, 192, 192, 191, 192, 191, 191, 190, 190, 187, 184, 181, 177, 172, 169, 165, 161, 156, 152, 147, 143, 139, 131, 123, 119, 115, 111, 110, 108, 106, 105,
	121, 124, 126, 128, 129, 129, 130, 131, 132, 133, 135, 137, 139, 143, 150, 154, 159, 164, 170, 173, 176, 179, 184, 186, 189, 190, 191, 192, 193, 194, 195, 194, 193, 192, 191, 191, 191, 191, 190, 190, 188, 184, 181, 177, 173, 169, 165, 160, 155, 149, 145, 142, 136, 129, 123, 118, 114, 110, 108, 108, 107,
	122, 125, 127, 130, 130, 131, 133, 134, 135, 136, 137, 140, 143, 147, 154, 158, 162, 166, 171, 174, 177, 181, 186, 189, 190, 190, 191, 192, 191, 191, 190, 189, 188, 189, 190, 190, 191, 190, 190, 190, 189, 186, 184, 181, 177, 173, 169, 164, 158, 152, 148, 144, 140, 134, 125, 118, 115, 111, 110, 108, 107,
	122, 125, 128, 130, 132, 133, 135, 136, 137, 139, 140, 143, 147, 152, 157, 161, 164, 168, 172, 175, 179, 182, 186, 190, 190, 190, 190, 189, 187, 184, 184, 183, 182, 182, 183, 183, 183, 184, 185, 186, 187, 186, 185, 184, 181, 177, 173, 169, 163, 157, 149, 145, 141, 136, 130, 119, 116, 112, 110, 108, 106,
	123, 126, 129, 131, 133, 135, 137, 138, 139, 141, 143, 147, 150, 156, 161, 164, 167, 170, 173, 177, 181, 184, 187, 188, 190, 189, 187, 185, 183, 179, 176, 174, 174, 174, 174, 174, 176, 177, 179, 180, 182, 183, 182, 181, 181, 180, 176, 171, 166, 160, 152, 147, 142, 138, 133, 126, 121, 115, 110, 106, 105,
	124, 127, 130, 132, 135, 137, 138, 140, 142, 144, 147, 149, 154, 157, 161, 165, 168, 171, 175, 178, 181, 184, 186, 187, 187, 184, 184, 181, 179, 175, 171, 169, 168, 168, 168, 169, 170, 172, 174, 177, 178, 179, 180, 181, 181, 180, 179, 174, 167, 161, 155, 148, 144, 139, 134, 128, 121, 115, 110, 106, 105,
	123, 128, 131, 133, 136, 138, 140, 142, 144, 146, 149, 151, 154, 157, 160, 164, 168, 172, 175, 178, 181, 183, 184, 184, 185, 183, 180, 177, 174, 170, 167, 165, 164, 164, 164, 165, 166, 168, 171, 175, 176, 178, 180, 181, 180, 180, 179, 177, 170, 163, 157, 150, 144, 139, 134, 128, 121, 115, 110, 108, 107,
	123, 127, 131, 134, 136, 138, 140, 142, 144, 147, 149, 151, 154, 157, 160, 164, 168, 171, 174, 178, 180, 181, 181, 182, 183, 181, 178, 173, 169, 166, 163, 161, 161, 160, 160, 161, 163, 165, 168, 173, 176, 178, 179, 180, 181, 180, 180, 175, 173, 166, 159, 152, 145, 139, 134, 127, 121, 115, 110, 109, 108,
	120, 124, 128, 131, 134, 137, 139, 142, 144, 146, 149, 151, 153, 156, 160, 163, 167, 171, 174, 178, 180, 180, 180, 180, 180, 180, 175, 171, 167, 162, 160, 158, 157, 157, 157, 158, 159, 162, 166, 170, 175, 177, 178, 180, 181, 181, 180, 178, 175, 169, 160, 154, 148, 140, 134, 128, 121, 115, 110, 110, 109,
	118, 121, 125, 129, 132, 134, 137, 140, 142, 145, 147, 149, 151, 155, 159, 163, 166, 169, 173, 177, 179, 180, 180, 180, 180, 179, 174, 169, 166, 161, 158, 156, 154, 153, 153, 154, 156, 159, 163, 169, 173, 175, 178, 180, 181, 180, 180, 179, 175, 170, 160, 154, 149, 142, 135, 128, 122, 116, 111, 110, 110,
	117, 120, 121, 125, 129, 132, 135, 138, 140, 143, 145, 147, 149, 153, 157, 160, 163, 166, 171, 174, 177, 179, 180, 180, 180, 179, 172, 168, 164, 160, 157, 154, 151, 149, 150, 150, 154, 158, 164, 169, 174, 178, 180, 180, 180, 180, 178, 177, 175, 170, 161, 153, 148, 142, 135, 129, 123, 116, 113, 112, 110,
	115, 118, 120, 122, 126, 130, 133, 136, 138, 141, 143, 145, 148, 151, 154, 157, 160, 163, 168, 171, 174, 177, 179, 179, 179, 176, 171, 167, 164, 160, 156, 153, 149, 148, 149, 151, 155, 158, 163, 170, 173, 177, 179, 180, 180, 180, 178, 175, 173, 171, 162, 154, 147, 141, 136, 130, 124, 117, 115, 112, 110,
	114, 116, 118, 120, 122, 127, 131, 133, 136, 138, 141, 143, 146, 148, 151, 154, 157, 160, 164, 168, 171, 174, 178, 178, 179, 177, 173, 169, 165, 161, 157, 154, 151, 149, 150, 152, 155, 159, 166, 171, 175, 177, 179, 180, 180, 179, 176, 174, 171, 168, 159, 151, 146, 141, 135, 129, 124, 119, 116, 113, 110,
	115, 114, 116, 118, 120, 122, 127, 129, 132, 136, 139, 141, 143, 146, 148, 151, 153, 156, 160, 164, 167, 172, 174, 176, 177, 176, 173, 170, 166, 162, 159, 157, 154, 153, 154, 155, 158, 161, 169, 172, 174, 176, 178, 178, 178, 178, 175, 172, 169, 162, 156, 149, 144, 140, 134, 128, 123, 118, 115, 112, 110,
	113, 113, 114, 116, 118, 120, 122, 125, 129, 133, 136, 138, 141, 143, 146, 149, 150, 153, 156, 160, 165, 170, 173, 176, 176, 176, 173, 172, 169, 165, 163, 160, 158, 157, 158, 159, 161, 166, 170, 170, 173, 175, 176, 178, 176, 173, 171, 168, 164, 158, 153, 146, 140, 137, 132, 127, 121, 117, 113, 111, 110,
	111, 112, 113, 114, 116, 118, 120, 122, 126, 130, 133, 136, 139, 142, 145, 147, 148, 151, 155, 158, 163, 168, 173, 176, 177, 177, 176, 174, 171, 169, 166, 164, 161, 161, 162, 164, 165, 167, 170, 170, 171, 173, 173, 173, 170, 168, 165, 163, 160, 155, 149, 143, 138, 134, 130, 125, 119, 116, 112, 110, 109,
	110, 112, 113, 113, 114, 116, 118, 120, 123, 127, 131, 134, 137, 141, 143, 145, 148, 150, 154, 157, 161, 166, 171, 176, 178, 178, 178, 176, 174, 172, 170, 167, 167, 167, 166, 168, 170, 169, 168, 167, 168, 168, 168, 168, 167, 165, 163, 160, 156, 152, 146, 140, 136, 131, 128, 122, 118, 114, 110, 110, 109,
	109, 110, 111, 112, 114, 116, 118, 119, 120, 124, 128, 131, 136, 140, 142, 145, 147, 150, 153, 157, 160, 165, 170, 174, 178, 179, 179, 178, 178, 176, 174, 171, 170, 170, 170, 168, 167, 166, 164, 163, 161, 162, 163, 163, 163, 161, 160, 157, 153, 148, 142, 136, 130, 127, 124, 120, 117, 113, 110, 110, 109,
	108, 109, 111, 112, 114, 116, 117, 118, 120, 121, 125, 128, 132, 138, 142, 144, 147, 149, 153, 156, 160, 164, 170, 174, 178, 180, 180, 179, 179, 178, 176, 172, 170, 170, 170, 168, 166, 164, 162, 160, 157, 156, 157, 158, 158, 156, 153, 151, 149, 144, 139, 130, 127, 124, 121, 118, 115, 112, 110, 110, 109,
	108, 109, 111, 113, 114, 116, 117, 118, 119, 120, 122, 126, 130, 135, 139, 143, 147, 149, 152, 156, 160, 164, 169, 173, 177, 180, 180, 180, 180, 179, 178, 174, 170, 170, 168, 167, 165, 163, 161, 157, 154, 153, 152, 152, 152, 149, 148, 147, 144, 140, 134, 128, 125, 122, 119, 117, 114, 110, 110, 109, 109,
	107, 108, 111, 112, 114, 115, 116, 117, 119, 120, 121, 124, 128, 133, 137, 141, 145, 149, 152, 156, 160, 164, 168, 172, 176, 179, 180, 180, 180, 179, 178, 174, 170, 168, 166, 165, 163, 161, 158, 154, 150, 149, 148, 146, 145, 143, 143, 143, 140, 136, 130, 126, 123, 120, 118, 115, 112, 110, 110, 109, 109,
	107, 108, 110, 112, 113, 113, 115, 116, 118, 120, 122, 125, 128, 132, 136, 140, 145, 148, 150, 155, 160, 164, 167, 170, 174, 177, 179, 179, 178, 176, 176, 173, 169, 166, 164, 163, 161, 159, 155, 152, 148, 145, 143, 141, 140, 139, 139, 138, 136, 132, 128, 124, 121, 118, 116, 114, 111, 110, 110, 109, 108,
	107, 108, 109, 111, 113, 114, 116, 117, 119, 120, 122, 125, 128, 132, 137, 141, 144, 146, 149, 152, 157, 162, 166, 168, 171, 173, 175, 175, 173, 172, 172, 171, 168, 165, 162, 160, 158, 156, 153, 149, 145, 142, 139, 138, 137, 136, 135, 133, 131, 129, 126, 122, 119, 117, 114, 112, 110, 110, 109, 108, 107,
	108, 109, 110, 112, 114, 115, 116, 117, 119, 120, 122, 126, 129, 133, 137, 141, 143, 146, 148, 151, 155, 160, 164, 167, 168, 169, 170, 170, 169, 168, 167, 168, 166, 163, 160, 158, 155, 153, 150, 147, 143, 140, 137, 136, 134, 133, 132, 130, 129, 127, 125, 121, 118, 115, 112, 110, 110, 110, 108, 107, 107,
	109, 110, 111, 113, 115, 116, 117, 118, 120, 121, 123, 126, 129, 133, 138, 141, 143, 146, 148, 150, 155, 159, 163, 165, 166, 167, 168, 168, 166, 165, 164, 161, 160, 159, 158, 155, 152, 149, 147, 144, 141, 138, 135, 134, 132, 130, 129, 128, 126, 124, 122, 120, 117, 113, 111, 110, 110, 110, 108, 107, 107,
	110, 111, 112, 113, 116, 117, 118, 119, 120, 122, 125, 127, 130, 133, 138, 141, 143, 146, 148, 150, 154, 159, 162, 163, 164, 166, 166, 166, 165, 163, 161, 159, 157, 156, 155, 153, 150, 146, 143, 140, 138, 136, 133, 132, 130, 129, 128, 125, 124, 122, 120, 119, 117, 114, 111, 110, 110, 109, 108, 107, 107,
	111, 112, 113, 114, 116, 117, 118, 119, 120, 123, 125, 128, 130, 134, 139, 141, 144, 146, 148, 151, 154, 158, 161, 164, 166, 167, 168, 166, 165, 163, 161, 158, 156, 154, 152, 150, 146, 142, 139, 137, 135, 133, 131, 130, 129, 128, 127, 125, 123, 121, 120, 118, 116, 113, 111, 110, 110, 109, 108, 107, 106,
	111, 112, 113, 115, 117, 118, 118, 120, 121, 124, 126, 128, 131, 135, 139, 142, 144, 146, 148, 151, 155, 160, 164, 165, 168, 169, 169, 168, 166, 163, 160, 158, 156, 153, 151, 148, 145, 142, 139, 137, 135, 132, 130, 129, 127, 126, 125, 124, 123, 120, 120, 117, 116, 114, 112, 110, 110, 108, 107, 106, 106,
	112, 113, 114, 116, 117, 118, 119, 120, 122, 124, 127, 129, 132, 135, 139, 142, 144, 146, 149, 152, 157, 162, 167, 169, 170, 170, 170, 168, 165, 163, 161, 159, 157, 155, 151, 148, 145, 141, 139, 136, 134, 132, 130, 128, 127, 126, 124, 123, 122, 120, 119, 117, 116, 114, 112, 111, 109, 107, 106, 106, 105,
	113, 114, 115, 116, 117, 119, 119, 120, 122, 125, 127, 129, 132, 135, 139, 142, 144, 147, 149, 154, 159, 164, 169, 170, 170, 170, 170, 170, 168, 165, 163, 161, 158, 155, 151, 148, 145, 142, 139, 137, 135, 132, 131, 128, 126, 125, 124, 122, 121, 120, 119, 117, 115, 113, 111, 110, 109, 106, 105, 105, 104,
	113, 114, 115, 117, 118, 119, 120, 121, 123, 125, 127, 130, 132, 135, 139, 142, 145, 148, 150, 156, 161, 166, 170, 170, 170, 170, 170, 170, 169, 166, 163, 161, 159, 155, 151, 148, 146, 143, 140, 138, 135, 134, 132, 130, 127, 125, 123, 121, 120, 120, 119, 116, 114, 112, 110, 110, 108, 106, 105, 104, 104,
	114, 115, 116, 117, 118, 119, 120, 121, 123, 126, 128, 130, 133, 136, 139, 142, 145, 148, 152, 157, 161, 166, 168, 170, 170, 170, 170, 168, 166, 164, 163, 160, 159, 155, 151, 148, 146, 143, 141, 138, 136, 134, 132, 130, 128, 125, 123, 121, 120, 120, 118, 116, 113, 111, 110, 110, 109, 106, 105, 104, 104,
	115, 116, 117, 118, 119, 120, 121, 121, 123, 126, 128, 131, 134, 136, 139, 142, 145, 149, 152, 157, 161, 163, 164, 166, 168, 167, 166, 164, 163, 161, 160, 158, 156, 152, 149, 147, 144, 143, 141, 139, 136, 134, 132, 130, 128, 125, 122, 120, 120, 119, 117, 115, 113, 110, 110, 109, 107, 106, 105, 104, 104,
	115, 116, 117, 118, 119, 120, 121, 122, 123, 125, 128, 131, 134, 137, 139, 142, 145, 149, 152, 156, 159, 159, 160, 162, 162, 161, 161, 160, 159, 158, 157, 155, 153, 150, 148, 146, 145, 143, 142, 140, 137, 134, 131, 129, 126, 124, 122, 120, 119, 117, 115, 113, 111, 110, 109, 109, 107, 106, 105, 104, 104,
	114, 115, 116, 116, 118, 119, 120, 121, 122, 126, 129, 132, 135, 137, 140, 143, 146, 149, 152, 155, 156, 157, 158, 159, 159, 159, 158, 158, 157, 155, 153, 151, 150, 149, 147, 146, 145, 144, 142, 141, 138, 135, 132, 128, 125, 122, 120, 118, 117, 115, 113, 112, 110, 109, 108, 108, 106, 105, 105, 104, 104,
	113, 114, 115, 116, 117, 118, 119, 120, 123, 126, 129, 132, 135, 138, 140, 143, 146, 148, 151, 153, 154, 156, 157, 157, 157, 157, 156, 155, 154, 152, 150, 149, 148, 147, 146, 145, 144, 142, 141, 140, 139, 136, 132, 129, 125, 121, 118, 116, 115, 113, 111, 110, 109, 108, 108, 107, 106, 105, 104, 104, 104,
	112, 113, 114, 115, 116, 117, 119, 120, 122, 126, 130, 133, 136, 138, 141, 143, 146, 148, 150, 152, 154, 155, 155, 155, 155, 155, 154, 152, 152, 150, 148, 147, 146, 145, 145, 143, 142, 141, 140, 140, 140, 137, 133, 129, 125, 120, 117, 115, 111, 110, 110, 109, 108, 107, 107, 106, 105, 105, 104, 104, 103,
	111, 112, 114, 115, 116, 117, 118, 120, 122, 125, 131, 134, 137, 139, 142, 144, 146, 148, 150, 152, 153, 153, 153, 153, 153, 153, 153, 151, 149, 147, 146, 144, 144, 143, 143, 142, 141, 140, 140, 140, 140, 138, 134, 130, 123, 120, 118, 111, 110, 110, 110, 108, 107, 106, 108, 105, 105, 104, 104, 103, 103,
	111, 112, 113, 115, 115, 116, 117, 119, 121, 126, 131, 135, 138, 140, 142, 144, 146, 148, 150, 151, 151, 151, 151, 151, 151, 151, 151, 150, 148, 146, 144, 142, 141, 141, 142, 141, 140, 140, 140, 140, 140, 140, 136, 132, 126, 120, 115, 110, 110, 110, 109, 107, 106, 105, 107, 105, 104, 104, 104, 103, 103,
	112, 113, 113, 114, 115, 116, 117, 119, 122, 127, 132, 135, 139, 141, 143, 145, 147, 149, 150, 150, 150, 150, 150, 150, 150, 150, 150, 149, 147, 144, 142, 141, 140, 140, 140, 140, 140, 140, 140, 140, 140, 140, 137, 133, 128, 120, 117, 110, 110, 110, 108, 106, 105, 105, 106, 105, 104, 104, 103, 103, 103,
	112, 113, 114, 114, 116, 117, 118, 120, 122, 128, 132, 136, 139, 141, 144, 146, 147, 149, 150, 150, 150, 150, 150, 150, 150, 150, 150, 149, 146, 143, 141, 140, 140, 139, 139, 139, 140, 140, 140, 140, 140, 140, 137, 133, 129, 121, 118, 110, 110, 109, 107, 106, 105, 105, 105, 104, 104, 103, 103, 103, 102,
	112, 114, 114, 115, 116, 117, 119, 120, 122, 128, 133, 136, 140, 142, 144, 146, 148, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 148, 145, 142, 140, 138, 138, 138, 137, 138, 140, 140, 140, 140, 140, 140, 137, 134, 130, 122, 118, 110, 110, 108, 106, 105, 103, 104, 104, 104, 104, 103, 103, 102, 102,
	113, 114, 115, 116, 116, 117, 118, 120, 123, 129, 133, 137, 140, 142, 144, 146, 149, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 147, 143, 141, 139, 137, 136, 136, 135, 136, 138, 140, 140, 140, 140, 139, 136, 134, 130, 123, 119, 113, 109, 108, 106, 104, 103, 104, 104, 104, 103, 103, 102, 102, 101,
	114, 115, 115, 116, 117, 118, 118, 120, 123, 129, 133, 137, 140, 143, 145, 147, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 148, 145, 142, 139, 138, 136, 135, 134, 134, 134, 136, 138, 137, 138, 139, 137, 134, 132, 125, 122, 117, 114, 109, 107, 105, 103, 102, 104, 104, 103, 103, 102, 102, 101, 101,
	114, 115, 116, 117, 117, 119, 118, 120, 123, 128, 132, 136, 139, 142, 145, 148, 150, 150, 150, 150, 150, 150, 150, 150, 150, 150, 147, 144, 141, 139, 136, 135, 134, 133, 132, 132, 134, 134, 134, 134, 135, 133, 131, 128, 124, 120, 116, 113, 110, 107, 104, 102, 102, 103, 103, 103, 102, 102, 102, 101, 100,
	115, 116, 116, 117, 118, 119, 119, 120, 124, 128, 132, 136, 139, 142, 145, 148, 150, 150, 150, 150, 150, 150, 150, 150, 150, 149, 146, 143, 140, 138, 135, 134, 133, 131, 131, 131, 131, 131, 131, 131, 130, 127, 124, 122, 119, 117, 115, 112, 109, 106, 104, 101, 102, 103, 103, 102, 102, 102, 101, 100, 100,
	115, 116, 117, 118, 118, 119, 120, 123, 125, 128, 131, 135, 138, 141, 145, 148, 150, 150, 150, 150, 150, 150, 150, 150, 150, 147, 145, 142, 139, 137, 134, 132, 131, 130, 129, 128, 128, 128, 128, 128, 126, 123, 121, 119, 116, 114, 112, 110, 108, 105, 103, 101, 103, 103, 103, 102, 102, 101, 100, 100, 100,
	116, 117, 118, 118, 119, 120, 122, 123, 125, 128, 131, 134, 137, 141, 145, 148, 149, 150, 150, 150, 150, 150, 150, 150, 148, 145, 143, 141, 138, 135, 133, 130, 129, 128, 127, 126, 125, 125, 125, 124, 123, 120, 118, 116, 114, 111, 109, 107, 106, 104, 102, 100, 101, 101, 102, 102, 101, 100, 100, 100, 100,
	116, 117, 118, 119, 120, 121, 123, 124, 126, 128, 130, 133, 137, 140, 144, 145, 147, 148, 149, 150, 149, 149, 147, 146, 144, 141, 139, 136, 133, 131, 129, 128, 127, 126, 125, 124, 123, 123, 122, 121, 120, 118, 116, 114, 112, 108, 107, 105, 103, 102, 100, 100, 100, 100, 101, 101, 100, 100, 100, 100, 100,
	117, 118, 119, 119, 120, 121, 123, 124, 126, 128, 129, 131, 135, 139, 142, 143, 145, 146, 147, 147, 147, 146, 144, 142, 140, 138, 135, 133, 130, 128, 127, 126, 125, 124, 123, 122, 121, 120, 119, 118, 117, 115, 114, 112, 110, 106, 105, 102, 101, 100, 100, 100, 100, 100, 100, 100, 100, 99, 99, 99, 99,
	117, 118, 119, 120, 120, 121, 123, 124, 125, 126, 128, 129, 132, 137, 140, 142, 143, 143, 144, 144, 144, 143, 141, 139, 137, 135, 133, 130, 128, 127, 126, 125, 123, 122, 121, 120, 119, 117, 116, 115, 114, 112, 111, 108, 107, 105, 100, 100, 100, 100, 100, 100, 100, 99, 99, 99, 99, 99, 99, 99, 98,
	116, 117, 118, 120, 120, 121, 122, 123, 124, 125, 126, 128, 130, 134, 139, 140, 141, 141, 141, 141, 141, 140, 138, 136, 134, 133, 131, 129, 127, 125, 124, 123, 122, 120, 119, 118, 117, 116, 114, 112, 111, 108, 109, 106, 106, 100, 100, 100, 100, 100, 99, 99, 99, 99, 99, 99, 99, 98, 98, 98, 97,
	114, 115, 116, 117, 119, 119, 120, 121, 122, 123, 125, 127, 129, 133, 136, 134, 134, 136, 138, 138, 137, 137, 135, 133, 132, 130, 129, 127, 125, 124, 122, 121, 120, 119, 117, 116, 115, 114, 112, 110, 109, 108, 107, 105, 105, 100, 100, 100, 100, 99, 99, 99, 98, 98, 98, 98, 98, 97, 97, 97, 97,
	112, 113, 114, 115, 116, 116, 117, 119, 120, 122, 124, 126, 127, 129, 129, 128, 127, 129, 132, 133, 133, 133, 133, 131, 129, 127, 126, 125, 124, 122, 121, 119, 118, 117, 116, 114, 113, 112, 110, 109, 108, 106, 106, 105, 100, 100, 100, 98, 98, 98, 98, 98, 98, 97, 97, 97, 97, 97, 97, 97, 96,
	109, 111, 112, 112, 113, 113, 113, 114, 116, 119, 121, 123, 124, 125, 124, 123, 123, 123, 125, 127, 129, 129, 128, 128, 127, 125, 124, 123, 122, 121, 119, 118, 117, 116, 114, 113, 112, 110, 109, 108, 107, 106, 105, 100, 100, 100, 97, 97, 97, 97, 97, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96,
	106, 107, 108, 108, 109, 110, 110, 112, 113, 114, 117, 119, 120, 121, 119, 117, 117, 117, 118, 120, 123, 124, 125, 125, 125, 123, 121, 120, 120, 119, 118, 117, 116, 115, 114, 113, 111, 109, 109, 107, 106, 105, 100, 100, 100, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96,
	104, 105, 105, 106, 106, 107, 108, 108, 109, 109, 111, 115, 116, 114, 113, 112, 111, 110, 111, 113, 116, 119, 122, 122, 122, 121, 120, 119, 118, 118, 117, 116, 115, 114, 113, 112, 111, 108, 108, 106, 105, 100, 100, 100, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96,
	102, 103, 103, 104, 104, 105, 106, 106, 107, 108, 109, 111, 112, 110, 109, 108, 108, 108, 108, 109, 110, 112, 116, 117, 117, 118, 118, 118, 117, 116, 116, 115, 114, 113, 112, 111, 110, 107, 107, 105, 100, 100, 100, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96,
	101, 102, 103, 103, 104, 105, 105, 106, 106, 107, 108, 109, 109, 107, 106, 106, 105, 105, 105, 106, 107, 108, 109, 110, 111, 113, 114, 115, 115, 115, 114, 113, 112, 111, 110, 108, 108, 106, 105, 100, 100, 100, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96,
	100, 101, 102, 102, 103, 103, 104, 104, 105, 106, 106, 107, 106, 106, 106, 105, 105, 104, 103, 103, 104, 105, 107, 108, 110, 111, 111, 112, 112, 113, 113, 112, 111, 110, 108, 107, 106, 105, 100, 100, 100, 98, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96,
	100, 101, 101, 102, 102, 103, 103, 104, 104, 105, 105, 105, 105, 106, 105, 105, 104, 103, 102, 101, 102, 103, 104, 106, 107, 110, 111, 111, 111, 112, 112, 112, 110, 107, 107, 106, 105, 102, 100, 100, 99, 98, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 95,
	99, 100, 101, 102, 102, 103, 103, 103, 104, 104, 104, 104, 103, 104, 104, 104, 104, 102, 101, 101, 102, 103, 104, 105, 107, 110, 111, 111, 111, 111, 111, 111, 108, 106, 105, 105, 102, 101, 100, 99, 99, 98, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 95, 95,
	99, 100, 100, 101, 101, 102, 102, 102, 103, 103, 103, 103, 102, 103, 103, 104, 103, 102, 101, 101, 101, 102, 103, 104, 106, 109, 110, 111, 111, 111, 110, 110, 107, 105, 103, 104, 100, 100, 99, 99, 98, 98, 97, 97, 96, 96, 96, 96, 96, 96, 96, 96, 96, 96, 95, 95, 95, 95, 95, 95, 95,
	99, 100, 100, 100, 101, 101, 101, 102, 102, 103, 102, 102, 101, 102, 102, 103, 103, 101, 101, 100, 101, 101, 102, 103, 105, 109, 110, 110, 111, 110, 110, 109, 106, 105, 100, 102, 100, 99, 99, 99, 98, 98, 97, 97, 96, 96, 96, 96, 96, 96, 95, 95, 95, 95, 95, 95, 95, 95, 95, 95, 94,
	99, 99, 99, 99, 100, 100, 101, 101, 102, 102, 101, 101, 101, 101, 101, 102, 102, 101, 100, 100, 101, 101, 101, 103, 104, 107, 109, 109, 110, 110, 109, 108, 105, 102, 100, 100, 99, 99, 99, 98, 98, 98, 97, 96, 96, 96, 96, 96, 95, 95, 95, 95, 95, 95, 95, 94, 94, 94, 94, 94, 94,
	98, 99, 99, 99, 99, 100, 100, 101, 101, 102, 101, 100, 100, 100, 101, 101, 101, 100, 100, 100, 100, 101, 101, 101, 103, 106, 107, 109, 109, 109, 109, 107, 104, 101, 100, 99, 99, 99, 98, 98, 98, 97, 96, 96, 96, 96, 95, 95, 95, 95, 95, 95, 95, 94, 94, 94, 94, 94, 94, 94, 94,
	98, 98, 98, 99, 99, 99, 100, 100, 101, 101, 100, 100, 99, 99, 100, 100, 100, 100, 100, 100, 100, 101, 101, 101, 102, 105, 106, 109, 108, 109, 107, 105, 102, 100, 100, 99, 99, 98, 98, 98, 97, 96, 96, 96, 96, 95, 95, 95, 95, 95, 95, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94,
	97, 98, 98, 98, 99, 99, 99, 100, 100, 100, 100, 100, 99, 99, 99, 100, 100, 100, 100, 100, 100, 100, 101, 101, 101, 103, 104, 105, 106, 105, 104, 101, 100, 100, 99, 99, 98, 98, 97, 97, 97, 96, 96, 96, 95, 95, 95, 95, 95, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94,
	97, 97, 97, 98, 98, 99, 99, 99, 100, 100, 100, 99, 99, 99, 99, 99, 100, 100, 100, 100, 100, 100, 101, 101, 100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 99, 99, 98, 97, 97, 97, 96, 96, 96, 95, 95, 95, 95, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94, 94,
})}