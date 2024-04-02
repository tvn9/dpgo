package main

import "fmt"

type Person struct {
	// address
	StreetAddress, Postcode, City string

	// Job
	CompanyName, Position string

	// Salary
	AnnualIncome int
}

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*p}
}

func (p *PersonBuilder) Build() *Person {
	return p.person
}

func (p *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*p}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(income int) *PersonJobBuilder {
	pjb.person.AnnualIncome = income
	return pjb
}

type PersonAddressBuilder struct {
	PersonBuilder
}

func (pab *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	pab.person.StreetAddress = streetAddress
	return pab
}

func (pab *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pab.person.City = city
	return pab
}

func (pab *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	pab.person.Postcode = postcode
	return pab
}

func main() {
	pb := NewPersonBuilder()
	pb.
		Lives().At("2356 Singer Blvd").In("Saigon").WithPostcode("P1K3Q6").
		Works().At("TTHN").AsA("Programmer").Earning(25000)
	person := pb.Build()
	fmt.Println(*person)
}
