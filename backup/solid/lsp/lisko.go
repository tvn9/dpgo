// Lisko Substitution Principle
package main

import "fmt"

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(w int) {
	r.width = w
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(h int) {
	r.height = h
}

func UseIt(sized Sized) {
	sized.SetWidth(20)
	sized.SetHeight(10)
	width := sized.GetWidth()
	height := sized.GetHeight()
	expectedArea := height * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, ", but got ", actualArea, "\n")
}

func main() {
	rc := &Rectangle{2, 3}
	UseIt(rc)
}
