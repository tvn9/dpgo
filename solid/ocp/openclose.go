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

type AndSpecification struct {
	first, second Specification
}

func (spec AndSpecification) IsSatisfied(i *Item) bool {
	return spec.first.IsSatisfied(i) && spec.second.IsSatisfied(i)
}

func (c ColorSpecification) IsSatisfied(i *Item) bool {
	return c.color == i.color
}
func (s SizeSpecification) IsSatisfied(i *Item) bool {
	return s.size == i.size
}

// NewFilter implement a new Filter that will take []Item, and the Specification
type NewFilter struct{}

func (nf *NewFilter) Filter(items []Item, spec Specification) []*Item {
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
	menShoe4 := Item{"Court Training 4", blue, medium}

	menShoes := []Item{menShoe1, menShoe2, menShoe3, menShoe4}

	// Looking for product in blue (Old)_
	fmt.Println("Search for items in Blue (Old)")
	f := Filter{}
	for _, v := range f.FilterByColor(menShoes, blue) {
		fmt.Printf(" - %s is blue\n", v.name)
	}

	// Search for product in BlackGreen (New)
	fmt.Println("Search for items in black and green (New)")
	nf := NewFilter{}
	bg := ColorSpecification{blackGreen}
	for _, v := range nf.Filter(menShoes, bg) {
		fmt.Printf(" - %s is black and green\n", v.name)
	}

	// Filter product in size large
	fmt.Println("Filter all product in large")
	lg := SizeSpecification{large}
	for _, v := range nf.Filter(menShoes, lg) {
		fmt.Printf(" - %s is large \n", v.name)
	}

	// Filter product in blue and medium
	fmt.Println("Blue and medium items")
	b := ColorSpecification{blue}
	m := SizeSpecification{medium}
	bm := AndSpecification{b, m}
	for _, v := range nf.Filter(menShoes, bm) {
		fmt.Printf(" - %s is blue and mideum\n", v.name)
	}

}
