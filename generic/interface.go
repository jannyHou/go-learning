package generic

import (
	"fmt"
	"math"
	"strings"
)

// Differ
type Differ[T any] interface {
	fmt.Stringer
	Diff(T) float64
}

// Point2D
type Point2D struct {
	x, y int
}

func (p *Point2D) String() string {
	return fmt.Sprintf("Type Point2D, with x = %d, y = %d", p.x, p.y)
}

// c for compare
func (p *Point2D) Diff(c *Point2D) float64 {
	xd := p.x - c.x
	yd := p.y - c.y
	d := math.Sqrt(float64(xd*xd) + float64(yd*yd))
	return d
}

// Point3D
type Point3D struct {
	x, y, z int
}

func (p *Point3D) String() string {
	return fmt.Sprintf("Type Point3D, with x = %d, y = %d, z= %d", p.x, p.y, p.z)
}

// c for compare
func (p *Point3D) Diff(c *Point3D) float64 {
	xd := p.x - c.x
	yd := p.y - c.y
	zd := p.z - c.z
	d := math.Sqrt(float64(xd*xd) + float64(yd*yd) + float64(zd*zd))
	return d
}

type Pair[D Differ[D]] struct {
	val1 D
	val2 D
}

func (p *Pair[D]) Distance() float64 {
	return p.val1.Diff(p.val2)
}

func (p *Pair[D]) String() string {
	s1 := p.val1.String()
	s2 := p.val2.String()
	var sb strings.Builder
	sb.WriteString(s1)
	sb.WriteString(".  ")
	sb.WriteString(s2)
	return sb.String()
}

func FindCloser[D Differ[D]](p1 *Pair[D], p2 *Pair[D]) (*Pair[D], float64) {
	d1 := p1.Distance()
	d2 := p2.Distance()
	if d1 > d2 {
		return p2, d2
	}
	return p1, d1

}
func Test10() {
	a := &Point2D{x: 3, y: 4}
	fmt.Println(a)
	b := &Point2D{x: 6, y: 7}

	dist := a.Diff(b)
	fmt.Println(dist)

	p1 := &Pair[*Point2D]{a, b}
	p2 := &Pair[*Point2D]{&Point2D{x: 3, y: 4}, &Point2D{x: 3, y: 4}}
	fmt.Println(p1.Distance())
	fmt.Println(p2.Distance())

	p, d := FindCloser(p1, p2)
	fmt.Println(p, d)

	g := &Point3D{x: 3, y: 4, z: 5}
	fmt.Println(g)
	h := &Point3D{x: 4, y: 5, z: 6}
	fmt.Println(h)

	p3 := &Pair[*Point3D]{&Point3D{x: 9, y: 10, z: 11}, &Point3D{x: 12, y: 13, z: 14}}
	pp, dd := FindCloser(p3, p3)
	fmt.Println(pp, dd)
}
