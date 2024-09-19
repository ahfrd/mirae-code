package repository

import (
	"database/sql"
	"mirae-code/app/model/request"
	"mirae-code/app/model/response"

	"github.com/gin-gonic/gin"
)

type StockRepository interface {
	BeginTransaction() (*sql.Tx, error)
	InsertStocks(ctx *gin.Context, request *request.AddStockRequest, tx *sql.Tx) error
	SelectStock(ctx *gin.Context, request *request.GetAllStockRequest) ([]response.StockAllResponse, error)
	CountStockData(ctx *gin.Context, request *request.GetAllStockRequest) (string, error)
	GetDetailStock(ctx *gin.Context, request *request.GetStockDetailRequest) (*response.StockAllResponse, error)
	UpdateStock(ctx *gin.Context, request *request.EditStockRequest, tx *sql.Tx) error
	DeleteStock(ctx *gin.Context, request *request.DeleteStockRequest, tx *sql.Tx) error
	UpdateStockByCode(ctx *gin.Context, request *request.EditStockRequest, tx *sql.Tx) error
	UpdateStockWithScraping(request *request.ScrapingRequest) error
	InsertStocksWithScraping(request *request.ScrapingRequest) error
	CountStockDataScrap(request *request.GetAllStockRequest) (string, error)
}
