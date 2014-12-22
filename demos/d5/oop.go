package main

import "fmt"

type Shap struct {
	x, y float64
}

func (shap *Shap) draw() {
	fmt.Printf("Shap(%f,%f)\n", shap.x, shap.y)
}

type Line struct {
	Shap
	endx, endy float64
}

func (line *Line) draw() {
	fmt.Printf("Line(%f,%f->%f,%f)\n", line.x, line.y, line.endx, line.endy)
}

func main() {
	shap := new(Line)
	shap.draw()
}
