package main

type Laptop struct {
	brand string
	model string
}

func (l *Laptop) setBrand(brand string) {
	l.brand = brand
}

func (l *Laptop) getBrand() string {
	return l.brand
}
func (l *Laptop) setModel(model string) {
	l.model = model
}

func (l *Laptop) getModel() string {
	return l.model
}
