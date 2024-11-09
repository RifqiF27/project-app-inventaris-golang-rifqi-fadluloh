package model


type Item struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	CategoryName     string  `json:"category"`
	Photo            string  `json:"photo_url"`
	Price            float64 `json:"price"`
	PurchaseDate     string  `json:"purchase_date"`
	UsageDays        int     `json:"usage_days"`
	DepreciationRate float64 `json:"depreciation_rate,omitempty"`
	CategoryID       int     `json:"-"`
}
