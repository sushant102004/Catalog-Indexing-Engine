package catalog

type FashionItem struct {
	Name        string  `json:"name"`
	ItemType    string  `json:"item_type"` // Fashion
	Description string  `json:"description"`
	Category    string  `json:"category" example:"Clothing, Footware"`
	Color       string  `json:"color"`
	Price       float32 `json:"price"`
	Size        string  `json:"size" example:"Small, Medium, Large"`
	Material    string  `json:"material" example:"Cotton, Denim, Leather"`
}
