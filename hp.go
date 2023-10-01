package main

type HP struct {
	Laptop
}

func newHP() ILaptop {
	return &HP{
		Laptop: Laptop{
			brand: "HP",
			model: "HP Pavilion",
		},
	}
}
