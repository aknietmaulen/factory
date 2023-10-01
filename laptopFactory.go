package main

import "fmt"

func getLaptop(laptopType string) (ILaptop, error) {
	if laptopType == "Dell" {
		return newDell(), nil
	}
	if laptopType == "HP" {
		return newHP(), nil
	}
	return nil, fmt.Errorf("there is no such laptop brand")
}
