package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type polyLine struct {
	ax, ay, bx, by, cx, cy, dx, dy, max, min float64
}

func main() {
	var polygon []polyLine
	var maxGlobal, minGlobal float64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, a := corner(i+1, j)
			bx, by, b := corner(i, j)
			cx, cy, c := corner(i, j+1)
			dx, dy, d := corner(i+1, j+1)
			max := maxf(maxf(maxf(a, b), c), d)
			maxGlobal = maxf(max, maxGlobal)
			min := minf(minf(minf(a, b), c), d)
			minGlobal = minf(min, minGlobal)
			if 	!math.IsNaN(ax) && !math.IsNaN(ay) &&
				!math.IsNaN(bx) && !math.IsNaN(by) &&
				!math.IsNaN(cx) && !math.IsNaN(cy) &&
				!math.IsNaN(dx) && !math.IsNaN(dy) {
				polygon = append(polygon, polyLine{ax, ay, bx, by, cx, cy, dx, dy, max, min})
			}
		}
	}
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < len(polygon); i++ {
		var color string
		if (polygon[i].max+polygon[i].min)/2 > 0 {
			colorInt := int(255 - (polygon[i].max / maxGlobal * 255))
			color = fmt.Sprintf("ff%02x%02[1]x", colorInt)
		} else {
			colorInt := int(255 - (polygon[i].min / minGlobal * 255))
			color = fmt.Sprintf("%02x%02[1]xff", colorInt)
		}
		fmt.Printf("<polygon stroke='#a0a0a0' fill='#%v' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			color, polygon[i].ax, polygon[i].ay, polygon[i].bx, polygon[i].by, polygon[i].cx, polygon[i].cy, polygon[i].dx, polygon[i].dy)
	}
	fmt.Println("</svg>")
}

func minf(a, b float64) float64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxf(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
