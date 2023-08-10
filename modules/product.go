package modules

type Product struct {
	ID           string        `gqlgen:"id"`
	Manufacturer *Manufacturer `gqlgen:"manufacturer"`
	Upc          string        `gqlgen:"upc"`
	Name         string        `gqlgen:"name"`
	Price        int           `gqlgen:"price"`
}

type Manufacturer struct {
	ID   string `gqlgen:"id"`
	Name string `gqlgen:"name"`
}

