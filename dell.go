package main

type Dell struct {
	Laptop
}

func newDell() ILaptop {
	return &Dell{
		Laptop: Laptop{
			brand: "Dell",
			model: "Dell Vostro",
		},
	}
}
