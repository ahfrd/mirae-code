package request

type AddStockRequest struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	Frequency float64 `json:"frequency"`
	Volume    string  `json:"volume"`
}

type GetAllStockRequest struct {
	PageNumber string  `json:"pageNumber" binding:"required"`
	PageSize   string  `json:"pageSize" binding:"required"`
	Name       string  `json:"name"`
	Code       string  `json:"code"`
	Price      float64 `json:"price"`
	Frequency  float64 `json:"frequency"`
	Volume     string  `json:"volume"`
}

type GetStockDetailRequest struct {
	Id int `json:"id"`
}

type EditStockRequest struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	Frequency float64 `json:"frequency"`
	Volume    string  `json:"volume"`
}

type DeleteStockRequest struct {
	Id int `json:"id"`
}

type ScrapingRequest struct {
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	Frequency float64 `json:"frequency"`
	Volume    string  `json:"volume"`
}
