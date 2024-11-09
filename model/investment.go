package model

type Investment struct {
    TotalInvestment    float64 `json:"total_investment"`
    DepreciatedValue   float64 `json:"depreciated_value"`
}

type ItemInvestment struct {
    ItemID            int     `json:"item_id"`
    Name              string  `json:"name"`
    InitialPrice      float64 `json:"initial_price"`
    DepreciatedValue  float64 `json:"depreciated_value"`
    DepreciationRate  float64 `json:"depreciation_rate"`
}
