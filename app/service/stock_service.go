package service

import (
	"mirae-code/app/model/request"
	"mirae-code/app/model/response"

	"github.com/gin-gonic/gin"
)

type StockService interface {
	AddStock(ctx *gin.Context, request *request.AddStockRequest, uid string) (*response.GenericResponse, error)
	GetAllStock(ctx *gin.Context, request *request.GetAllStockRequest, uid string) (*response.GenericResponse, error)
	DetailStock(ctx *gin.Context, request *request.GetStockDetailRequest, uid string) (*response.GenericResponse, error)
	EditStock(ctx *gin.Context, request *request.EditStockRequest, uid string) (*response.GenericResponse, error)
	DeleteStock(ctx *gin.Context, request *request.DeleteStockRequest, uid string) (*response.GenericResponse, error)
	ScrapingStock(uid string) (*response.GenericResponse, error)
}
