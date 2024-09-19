package response

type CountData struct {
	TotalRecord   string `json:"total_record"`
	TotalPage     string `json:"total_page"`
	RecordPerPage string `json:"record_per_page"`
	CurrentPage   string `json:"current_page"`
	StartRecord   string `json:"start_record"`
	FirstRecord   string `json:"first_record"`
}

type StockAllResponse struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Code      string  `json:"code"`
	Price     float64 `json:"price"`
	Frequency float64 `json:"frequency"`
	Volume    string  `json:"volume"`
}

type ListStockResponse struct {
	ListData   []StockAllResponse `json:"listData"`
	Pagination CountData          `json:"pagination"`
}
