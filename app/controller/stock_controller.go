package controller

import (
	"encoding/json"
	"fmt"
	"mirae-code/app/model/request"
	"mirae-code/app/model/response"
	"mirae-code/app/service"
	"mirae-code/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type StockController struct {
	StockService service.StockService
}

func NewStockController(StockService *service.StockService) StockController {
	return StockController{StockService: *StockService}
}

func (c *StockController) AddStock(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.AddStockRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.StockService.AddStock(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *StockController) GetAllStock(ctx *gin.Context) {
	requestId := guuid.New()
	price, _ := strconv.ParseFloat(ctx.Query("price"), 64)
	frequency, _ := strconv.ParseFloat(ctx.Query("frequency"), 64)

	bodyReq := &request.GetAllStockRequest{
		PageNumber: ctx.Query("pageNumber"),
		PageSize:   ctx.Query("pageSize"),
		Name:       ctx.Query("name"),
		Code:       ctx.Query("code"),
		Price:      price,
		Volume:     ctx.Query("volume"),
		Frequency:  frequency,
	}
	if err := ctx.ShouldBindQuery(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.FailedGettingDataErrorBindingRequest.Code(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorBindingRequest.Message(), err),
			Status:  response.FailedGettingDataErrorBindingRequest.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}
	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	response, err := c.StockService.GetAllStock(ctx, bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *StockController) GetDetailStock(ctx *gin.Context) {
	requestId := guuid.New()
	id, _ := strconv.Atoi(ctx.Param("id"))
	bodyReq := &request.GetStockDetailRequest{
		Id: id,
	}
	if err := ctx.ShouldBind(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.FailedGettingDataErrorBindingRequest.Code(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorBindingRequest.Message(), err),
			Status:  response.FailedGettingDataErrorBindingRequest.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}
	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)
	response, err := c.StockService.DetailStock(ctx, bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
func (c *StockController) EditStock(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.EditStockRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	bodyReq.Id = id
	if err := ctx.ShouldBindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.FailedGettingDataErrorBindingRequest.Code(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorBindingRequest.Message(), err),
			Status:  response.FailedGettingDataErrorBindingRequest.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.StockService.EditStock(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (c *StockController) DeleteStock(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.DeleteStockRequest
	id, _ := strconv.Atoi(ctx.Param("id"))
	bodyReq.Id = id
	if err := ctx.ShouldBind(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.FailedGettingDataErrorBindingRequest.Code(),
			Message: fmt.Sprintf(response.FailedGettingDataErrorBindingRequest.Message(), err),
			Status:  response.FailedGettingDataErrorBindingRequest.Status(),
		})
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.JSON(http.StatusBadRequest, response.GenericResponse{
			Code:    response.BadRequestMarshalError.Code(),
			Message: fmt.Sprintf(response.BadRequestMarshalError.Message(), err),
			Status:  response.BadRequestMarshalError.Status(),
		})
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.StockService.DeleteStock(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (c *StockController) ScrapingStock() {
	requestId := guuid.New()

	logStart := helpers.LogScrapStart("scraping stock", requestId.String())
	fmt.Println(logStart)

	response, err := c.StockService.ScrapingStock(requestId.String())
	if err != nil {
		helpers.LogScrapError(err.Error(), requestId.String())
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogScrapError(err.Error(), requestId.String())
		return
	}

	logStop := helpers.LogScrapEnd(string(responseData), requestId.String())
	fmt.Println(logStop)
}
