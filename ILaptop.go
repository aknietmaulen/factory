package main

type ILaptop interface {
	setBrand(brand string)
	getBrand() string
	setModel(model string)
	getModel() string
}
