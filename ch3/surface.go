package main

import (
	"errors"
	"fmt"
	"math"
)

const (
	width, height = 600, 320		// size in pixels
	cells = 100						// number of grid cells
	xyrange = 30.0					// axis ranges [-xyrange, +xyrange]
	xysacle = width / 2 / xyrange	// pixels per x or y unit
	zscale = height * 0.4			// pixels per z unit
	angle = math.Pi / 6				// 180°/6 = 30°
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i:= 0 ;i < cells; i++ {
		for j := 0 ; j < cells; j++ {
			ax, ay, err := corner(i+1, j)
			if err != nil {
				continue
			}
			bx, by, err := corner(i, j)
			if err != nil {
				continue
			}
			cx, cy, err := corner(i, j+1)
			if err != nil {
				continue
			}
			dx, dy, err := corner(i+1, j+1)
			if err != nil {
				continue
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy);
		}
	}
	fmt.Println("</svg>")
}


func corner(i, j int) (float64, float64, error) {
	x := xyrange * (float64(i)/cells - 5)
	y := xyrange * (float64(j)/cells - 5)
	z, err := f(x, y)

	sx := width/z + (x-y)*cos30*xysacle
	sy := width/z + (x+y)*cos30*xysacle
	return sx, sy, err
}


func f (x , y float64) (float64, error){
	r := math.Hypot(x, y)
	var err error
	if math.IsNaN(r) {
		err = errors.New("meet NaN")
	} else if math.IsInf(r, 0){
		err = errors.New("meet NaN")
	}
	return math.Sin(r), err
}
