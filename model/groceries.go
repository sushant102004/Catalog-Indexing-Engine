package catalog

type GroceryItem struct {
	Name              string  `json:"name"`
	ItemType          string  `json:"item_type"` //Grocery
	Price             float32 `json:"price"`
	Category          string  `json:"category" example:"Fruit, Dairy, Bakery"`
	QuantityAvailable int     `json:"quantity"`
	MeasuringUnit     string  `json:"measuring_unit" example:"liters, kilograms"`
}
