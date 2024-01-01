// Demonstration of Open Close Principle (OCP), and Specification (Enterprise Pattern)
package main

import "fmt"

type Color byte

const (
	red Color = iota
	green
	blue
	blackGreen
	whiteBlue
)

type Size byte

const (
	small Size = iota
	medium
	large
	xLarge
)

type Item struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
	//
}

//----------------- Using the Old way to filter -----------------

// FilterByColor compares the item's color to the targeted color for a match
func (f *Filter) FilterByColor(items []Item, color Color) []*Item {
	filteredItems := make([]*Item, 0)
	for i, v := range items {
		if v.color == color {
			filteredItems = append(filteredItems, &items[i])
		}
	}
	return filteredItems
}

//-------------- Using the specification pattern (New) -----------------

// Specification pattern interface
type Specification interface {
	IsSatisfied(i *Item) bool
}

type ColorSpecification struct {
	color Color
}

type SizeSpecification struct {
	size Size
}

func (c ColorSpecification) IsSatisfied(i *Item) bool {
	return c.color == i.color
}
func (s SizeSpecification) IsSatisfied(i *Item) bool {
	return s.size == i.size
}

// NewFilter implement a new Filter that will take []Item, and the Specification
type NewFilter struct{}

func (nf *NewFilter) NFilter(items []Item, spec Specification) []*Item {
	output := make([]*Item, 0)
	for i, v := range items {
		if spec.IsSatisfied(&v) {
			output = append(output, &items[i])
		}
	}
	return output
}

func main() {
	// Create new items
	menShoe1 := Item{"Court Training 1", blackGreen, large}
	menShoe2 := Item{"Court Training 2", whiteBlue, large}
	menShoe3 := Item{"Court Training 3", blue, medium}

	menShoes := []Item{menShoe1, menShoe2, menShoe3}

	// Looking for product in blue (Old)_
	fmt.Println("Search for items in Blue (Old)")
	f := Filter{}
	for i, v := range f.FilterByColor(menShoes, blue) {
		fmt.Printf("%d#, %s is blue\n", i+1, v.name)
	}

	// Search for product in BlackGreen (New)
	fmt.Println("Search for items in black and green (New)")
	nf := NewFilter{}
	bg := ColorSpecification{blackGreen}
	for i, v := range nf.NFilter(menShoes, bg) {
		fmt.Printf("%d# - %s is black and green\n", i+1, v.name)
	}
}
