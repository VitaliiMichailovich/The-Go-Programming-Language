package main

import (
	"math"
	"net/http"
	"log"
	"strconv"
	"fmt"
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
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write(poly())
}

func poly() []byte {
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
	reta := ("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='"+strconv.Itoa(width)+"' height='"+strconv.Itoa(height)+"'>")
	for i := 0; i < len(polygon); i++ {
		var color string
		if (polygon[i].max+polygon[i].min)/2 > 0 {
			colorInt := int(255 - (polygon[i].max / maxGlobal * 255))
			color = fmt.Sprintf("ff%02x%02[1]x", colorInt)
		} else {
			colorInt := int(255 - (polygon[i].min / minGlobal * 255))
			color = fmt.Sprintf("%02x%02[1]xff", colorInt)
		}
		reta += ("<polygon stroke='#a0a0a0' fill='#"+color+"' points='"+strconv.FormatFloat(polygon[i].ax, 'f', 6, 64)+","+strconv.FormatFloat(polygon[i].ay, 'f', 6, 64)+" "+strconv.FormatFloat(polygon[i].bx, 'f', 6, 64)+","+strconv.FormatFloat(polygon[i].by, 'f', 6, 64)+" "+strconv.FormatFloat(polygon[i].cx, 'f', 6, 64)+","+strconv.FormatFloat(polygon[i].cy, 'f', 6, 64)+" "+strconv.FormatFloat(polygon[i].dx, 'f', 6, 64)+","+strconv.FormatFloat(polygon[i].dy, 'f', 6, 64)+"'/>\n")
	}
	reta += "</svg>"
	return []byte(reta)
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
