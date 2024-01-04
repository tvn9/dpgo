// Interface Segregation Principle
package main

type Document struct {
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Faxer interface {
	Fax(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {
}

type Photocopier struct{}

func (p Phtocopier) Scan(d Document) {
}

func (p Photocopier) Print(d Document) {
}

type MyFax struct {
}

func (m MyFax) Fax(d Document) {}

type AllInOnePrinter interface {
	Printer
	Scanner
	Faxer
}

func main() {

}
