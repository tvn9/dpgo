// Demonstration of OCP and Repository
// A type should only have one reason to change
// Separation of concerns - different types/packages handling different, independent tasks/problems
package main

import "fmt"

// Vehicle hold information about Vehicle being built
type Vehicle struct {
	BodyType    string
	Manufacture string
	YearBuilt   string
	VIN         string
	PowerTrain  string
	Color       string
}

// AddBodyType add, remove, or change the Vehicle Body Type
func (v *Vehicle) AddBodyType(bt string) {
	v.BodyType = bt
}

// AddManufacture adds the vehicle's manufacture name
func (v *Vehicle) AddManufacture(mf string) {
	v.Manufacture = mf
}

// AddYearBuilt adds the vehicle year built
func (v *Vehicle) AddYearBuilt(yb string) {
	v.YearBuilt = yb
}

// AddVIN adds the VIN info to the Vehicle object
func (v *Vehicle) AddVIN(vn string) {
	v.VIN = vn
}

// AddPowerTrain adds the vehicle's power train info
func (v *Vehicle) AddPowerTrain(pt string) {
	v.PowerTrain = pt
}

// AddColor add the vehicle's color
func (v *Vehicle) AddColor(cl string) {
	v.Color = cl
}

// Start the program
func main() {
	// Create new vehicle
	f150 := Vehicle{}

	f150.AddBodyType("Truck")
	f150.AddManufacture("Ford")
	f150.AddYearBuilt("2021")
	f150.AddVIN("2FADPJ9BM15693")
	fmt.Println(f150)

	// Remove year built
	f150.AddYearBuilt("")
	fmt.Println(f150)
}
