package model

type ItemReplacement struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Category          string `json:"category"`
	PurchaseDate      string `json:"purchase_date"`
	TotalUsageDays    int    `json:"total_usage_days"`
	ReplacementRequired bool  `json:"replacement_required"`
}
