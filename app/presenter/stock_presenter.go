package presenter

import (
	"fmt"
	"mirae-code/app/controller"
	"mirae-code/app/repository"
	"mirae-code/app/service"
	"mirae-code/env"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func mainStockPresenter(di *env.Dependency) *controller.StockController {
	StockRepository := repository.NewStockRepository(di)
	StockService := service.NewStockService(&StockRepository)
	StockController := controller.NewStockController(&StockService)
	return &StockController
}

func AddStock(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainStockPresenter(env).AddStock(ctx)
	}
}

func GetAllStock(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainStockPresenter(env).GetAllStock(ctx)
	}
}

func GetDetailStock(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainStockPresenter(env).GetDetailStock(ctx)
	}
}

func EditStock(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainStockPresenter(env).EditStock(ctx)
	}
}

func DeleteStock(env *env.Dependency) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mainStockPresenter(env).DeleteStock(ctx)
	}
}

func ScrappingStock(env *env.Dependency) {
	stockController := mainStockPresenter(env)
	c := cron.New()
	fmt.Println(env.Params.Schedular.JobExecTime)
	c.AddFunc(env.Params.Schedular.JobExecTime, func() {
		stockController.ScrapingStock()
	})

	c.Start()

	select {}
}
