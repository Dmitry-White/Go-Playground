package data

type OldProduct struct {
	Id        int
	Name      string
	Inventory int
	Price     int
}

type Product struct {
	ID          int    `json:"id"`
	ProductCode string `json:"productCode"`
	Name        string `json:"name"`
	Inventory   int    `json:"inventory"`
	Price       int    `json:"price"`
	Status      string `json:"status"`
}
