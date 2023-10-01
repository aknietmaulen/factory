package main

import "fmt"

func main() {
	dellLaptop, _ := getLaptop("Dell")
	hpLaptop, _ := getLaptop("HP")

	printLaptopDetails(dellLaptop)
	printLaptopDetails(hpLaptop)
}

func printLaptopDetails(l ILaptop) {
	fmt.Printf("Laptop Brand: %s", l.getBrand())
	fmt.Println()
	fmt.Printf("Laptop Model: %s", l.getModel())
	fmt.Println()
}
