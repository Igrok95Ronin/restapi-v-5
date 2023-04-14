package main

import "fmt"

type Point struct {
	X, Y int
}

func moviePointPtr(p *Point, x, y int) {
	p.X += x
	p.Y += y
}

func (p *Point) movePointPtr2(x, y int) {
	p.X += x
	p.Y += y
}

func main() {
	p := Point{
		1,
		2,
	}
	fmt.Println(p)

	moviePointPtr(&p, 1, 1)
	p.movePointPtr2(1, 1)

	fmt.Println(p)
}
