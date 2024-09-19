package service

import (
	"fmt"
	"log"
	"math"
	"mirae-code/app/model/request"
	"mirae-code/app/model/response"
	"mirae-code/app/repository"
	"mirae-code/helpers"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type StockServiceImpl struct {
	StockRepository repository.StockRepository
}

func NewStockService(StockRepository *repository.StockRepository) StockService {
	return &StockServiceImpl{
		StockRepository: *StockRepository,
	}
}

func (s *StockServiceImpl) AddStock(ctx *gin.Context, params *request.AddStockRequest, uid string) (*response.GenericResponse, error) {
	var (
		resData response.GenericResponse
		err     error
		message string
	)

	txBegin, err := s.StockRepository.BeginTransaction()
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err),
		}, nil
	}
	countDataStock, err := s.StockRepository.CountStockData(ctx, &request.GetAllStockRequest{
		Code: params.Code,
	})
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}

	if countDataStock > "0" {
		err = s.StockRepository.UpdateStockByCode(ctx, &request.EditStockRequest{
			Code:      params.Code,
			Name:      params.Name,
			Price:     params.Price,
			Frequency: params.Frequency,
			Volume:    params.Volume,
		}, txBegin)
		message = "update stock"
	} else {
		err = s.StockRepository.InsertStocks(ctx, params, txBegin)
		message = "create stock"
	}
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		_ = txBegin.Rollback()
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err.Error()),
		}, nil
	}
	_ = txBegin.Commit()
	resData.Status = response.Success.Status()
	resData.Code = response.Success.Code()
	resData.Message = fmt.Sprintf(response.Success.Message(), message)
	return &resData, nil
}

func (s *StockServiceImpl) GetAllStock(ctx *gin.Context, request *request.GetAllStockRequest, uid string) (*response.GenericResponse, error) {
	var (
		data    response.ListStockResponse
		resData response.GenericResponse
	)
	countDataStock, err := s.StockRepository.CountStockData(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}

	floTotalRow, _ := strconv.ParseFloat(countDataStock, 64)
	floRecordPerPage, _ := strconv.ParseFloat(request.PageSize, 64)
	floCurrentPage, _ := strconv.ParseFloat(request.PageNumber, 64)
	totalPage := math.Ceil(floTotalRow / floRecordPerPage)
	currentPage := request.PageNumber
	firstRecord := (floCurrentPage - 1) * floRecordPerPage
	startRecord := firstRecord + 1
	countData := response.CountData{
		TotalRecord:   countDataStock,
		TotalPage:     strconv.FormatFloat(totalPage, 'f', 0, 64),
		RecordPerPage: request.PageSize,
		CurrentPage:   currentPage,
		StartRecord:   strconv.FormatFloat(startRecord, 'f', 0, 64),
		FirstRecord:   strconv.FormatFloat(firstRecord, 'f', 0, 64),
	}
	request.PageNumber = strconv.FormatFloat(firstRecord, 'f', 0, 64)
	selectDataStock, err := s.StockRepository.SelectStock(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}
	data.ListData = selectDataStock
	data.Pagination = countData

	resData.Status = response.Success.Status()
	resData.Code = response.Success.Code()
	resData.Message = fmt.Sprintf(response.Success.Message(), " get all stock")
	resData.Data = data
	return &resData, nil
}
func (s *StockServiceImpl) DetailStock(ctx *gin.Context, request *request.GetStockDetailRequest, uid string) (*response.GenericResponse, error) {
	var resData response.GenericResponse

	detailData, err := s.StockRepository.GetDetailStock(ctx, request)
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)

		return &response.GenericResponse{
			Code:    response.FailedGettingDataErrorDB.Code(),
			Status:  response.FailedGettingDataErrorDB.Status(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorDB.Message(), err),
		}, nil
	}
	resData.Status = response.Success.Status()
	resData.Code = response.Success.Code()
	resData.Message = fmt.Sprintf(response.Success.Message(), " get detail stock")
	resData.Data = detailData
	return &resData, nil
}

func (s *StockServiceImpl) EditStock(ctx *gin.Context, request *request.EditStockRequest, uid string) (*response.GenericResponse, error) {
	var resData response.GenericResponse
	txBegin, err := s.StockRepository.BeginTransaction()
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err),
		}, nil
	}
	if err := s.StockRepository.UpdateStock(ctx, request, txBegin); err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		_ = txBegin.Rollback()
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err.Error()),
		}, nil
	}
	_ = txBegin.Commit()
	resData.Status = response.Success.Status()
	resData.Code = response.Success.Code()
	resData.Message = fmt.Sprintf(response.Success.Message(), " update stock")
	return &resData, err
}

func (s *StockServiceImpl) DeleteStock(ctx *gin.Context, request *request.DeleteStockRequest, uid string) (*response.GenericResponse, error) {
	var resData response.GenericResponse
	txBegin, err := s.StockRepository.BeginTransaction()
	if err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err),
		}, nil
	}
	if err := s.StockRepository.DeleteStock(ctx, request, txBegin); err != nil {
		helpers.LogError(ctx, err.Error(), uid)
		_ = txBegin.Rollback()
		return &response.GenericResponse{
			Code:    response.FailedCreatingDataErrorInternal.Code(),
			Status:  response.FailedCreatingDataErrorInternal.Status(),
			Message: fmt.Sprintf(response.FailedCreatingDataErrorInternal.Message(), err.Error()),
		}, nil
	}
	_ = txBegin.Commit()
	resData.Status = response.Success.Status()
	resData.Code = response.Success.Code()
	resData.Message = fmt.Sprintf(response.Success.Message(), " delete stock")
	return &resData, err
}

func (s *StockServiceImpl) ScrapingStock(uid string) (*response.GenericResponse, error) {
	scrapingData := []request.ScrapingRequest{}
	count := 25
	var start int
	var totalItems int

	c := colly.NewCollector()
	// currentPage := 1

	c.OnHTML("table tbody tr", func(ch *colly.HTMLElement) {
		name := ch.ChildText(".longName")
		code := ch.ChildText(".symbol")
		price := ch.ChildText("fin-streamer[data-field='regularMarketPrice']")
		volume := ch.ChildText("fin-streamer[data-field='regularMarketVolume']")
		frequency := ch.ChildText("fin-streamer[data-field='regularMarketChange']")

		// Append the scraped data to the slice
		if name != "" && code != "" && price != "" && volume != "" && frequency != "" {
			priceFloat, _ := strconv.ParseFloat(price, 64)
			frequencyFloat, _ := strconv.ParseFloat(frequency, 64)

			financeYahoo := request.ScrapingRequest{
				Name:      name,
				Code:      code,
				Price:     priceFloat,
				Volume:    volume,
				Frequency: frequencyFloat,
			}
			countDataStock, err := s.StockRepository.CountStockDataScrap(&request.GetAllStockRequest{
				Code: financeYahoo.Code,
			})
			if err != nil {
				helpers.LogScrapError(err.Error(), uid)
			}
			if countDataStock > "0" {
				err = s.StockRepository.UpdateStockWithScraping(&request.ScrapingRequest{
					Code:      financeYahoo.Code,
					Name:      financeYahoo.Name,
					Price:     financeYahoo.Price,
					Frequency: financeYahoo.Frequency,
					Volume:    financeYahoo.Volume,
				})
			} else {
				err = s.StockRepository.InsertStocksWithScraping(&request.ScrapingRequest{
					Code:      financeYahoo.Code,
					Name:      financeYahoo.Name,
					Price:     financeYahoo.Price,
					Frequency: financeYahoo.Frequency,
					Volume:    financeYahoo.Volume,
				})
			}
			if err != nil {
				helpers.LogScrapError(err.Error(), uid)

			}
			scrapingData = append(scrapingData, financeYahoo)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div.total", func(h *colly.HTMLElement) {
		totalText := h.Text
		re := regexp.MustCompile(`of (\d+)`)
		matches := re.FindStringSubmatch(totalText)
		if len(matches) > 1 {
			totalItems, _ = strconv.Atoi(matches[1])
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, "scraped!")
		// Handle pagination
		if start+count < totalItems {
			start += count
			url := fmt.Sprintf("https://finance.yahoo.com/markets/stocks/gainers/?start=%d&count=%d", start, count)
			fmt.Println("Visiting next page:", url)
			err := c.Visit(url)
			if err != nil {
				helpers.LogScrapError(err.Error(), uid)
			}
		} else {
			// Print results once all pages are scraped
			fmt.Println("Scraped Data:")
			fmt.Println("scraped")
		}

	})

	// Start the scraping with the initial page
	startURL := fmt.Sprintf("https://finance.yahoo.com/markets/stocks/gainers/?start=0&count=%d", count)
	fmt.Println("Starting URL:", startURL)
	err := c.Visit(startURL)
	if err != nil {
		log.Println("Error visiting URL:", err)
	}
	return &response.GenericResponse{
		Code:    response.Success.Code(),
		Status:  response.Success.Status(),
		Message: fmt.Sprintf(response.Success.Message(), "scrapSuccses"),
		Data:    scrapingData,
	}, nil
}
