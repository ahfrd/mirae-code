package repository

import (
	"context"
	"database/sql"
	"fmt"
	"mirae-code/app/model/request"
	"mirae-code/app/model/response"
	"mirae-code/env"

	"github.com/gin-gonic/gin"
)

type stockRepository struct {
	di *env.Dependency
}

// ExampleRepository implements entity.stockRepository.

func NewStockRepository(di *env.Dependency) StockRepository {
	return &stockRepository{
		di: di,
	}
}

func (r *stockRepository) BeginTransaction() (*sql.Tx, error) {
	tx, err := r.di.DB.Begin()
	return tx, err
}

func (r *stockRepository) InsertStocks(ctx *gin.Context, request *request.AddStockRequest, tx *sql.Tx) error {
	q := `INSERT INTO stock (name,code,price,frequency,volume) values (?,?,?,?,?)`
	_, err := tx.ExecContext(ctx, q, request.Name, request.Code, request.Price, request.Frequency, request.Volume)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (r *stockRepository) CountStockData(ctx *gin.Context, request *request.GetAllStockRequest) (string, error) {
	var count string
	filter := r.filterQuery(request)
	q := fmt.Sprintf(`select count(*) as count from stock  
	 				  where 1 = 1 %s`, filter)
	fmt.Println(q)
	if err := r.di.DB.QueryRowContext(ctx, q).Scan(&count); err != nil {
		return "", err
	}
	return count, nil
}
func (r *stockRepository) SelectStock(ctx *gin.Context, request *request.GetAllStockRequest) ([]response.StockAllResponse, error) {
	var results []response.StockAllResponse
	where := r.filterQuery(request)
	q := fmt.Sprintf(`select id,name,code,price,frequency,volume
				      from stock 
					  where 1 = 1
					  %s
					  order by id desc 
					  limit %s,%s
					  `, where, request.PageNumber, request.PageSize)
	result, err := r.di.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer func() { _ = result.Close() }()
	for result.Next() {
		var dataStock response.StockAllResponse
		if err = result.Scan(&dataStock.Id, &dataStock.Name, &dataStock.Code,
			&dataStock.Price, &dataStock.Frequency, &dataStock.Volume); err != nil {
			return nil, err
		}
		results = append(results, dataStock)

	}

	return results, nil
}

func (r *stockRepository) filterQuery(request *request.GetAllStockRequest) string {
	var where string

	if request.Name != "" {
		where += fmt.Sprintf(" and name = '%s'", request.Name)
	}
	if request.Code != "" {
		where += fmt.Sprintf(" and code = '%s'", request.Code)

	}
	if request.Frequency != 0 {
		where += fmt.Sprintf(" and frequency = '%f'", request.Frequency)

	}
	if request.Volume != "" {
		where += fmt.Sprintf(" and volume = '%s'", request.Volume)

	}
	if request.Price != 0 {
		where += fmt.Sprintf(" and price = '%f'", request.Price)

	}

	return where
}

func (r *stockRepository) GetDetailStock(ctx *gin.Context, request *request.GetStockDetailRequest) (*response.StockAllResponse, error) {
	var (
		resultData response.StockAllResponse
		id         sql.NullInt64
		name       sql.NullString
		code       sql.NullString
		price      sql.NullFloat64
		volume     sql.NullString
		frequency  sql.NullFloat64
	)
	q := fmt.Sprintf(`select id,name,code,price,frequency,volume
					  from stock 
					  where id = %d`, request.Id)
	if err := r.di.DB.QueryRowContext(ctx, q).Scan(&id, &name, &code, &price, &frequency, &volume); err != nil {
		return nil, err
	}
	resultData.Id = int(id.Int64)
	resultData.Name = name.String
	resultData.Code = code.String
	resultData.Price = price.Float64
	resultData.Volume = volume.String
	resultData.Frequency = frequency.Float64
	return &resultData, nil
}

func (r *stockRepository) UpdateStock(ctx *gin.Context, request *request.EditStockRequest, tx *sql.Tx) error {
	q := fmt.Sprintf(`update stock set name = '%s', code = '%s', price = '%f', volume = '%s',frequency = '%f' where id = '%d' `,
		request.Name, request.Code, request.Price, request.Volume, request.Frequency, request.Id)

	fmt.Println(q)
	if _, err := tx.ExecContext(ctx, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func (r *stockRepository) DeleteStock(ctx *gin.Context, request *request.DeleteStockRequest, tx *sql.Tx) error {
	q := fmt.Sprintf(`DELETE FROM stock where id = %d`, request.Id)
	fmt.Println(q)
	if _, err := tx.ExecContext(ctx, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func (r *stockRepository) UpdateStockByCode(ctx *gin.Context, request *request.EditStockRequest, tx *sql.Tx) error {
	q := fmt.Sprintf(`update stock set name = '%s', code = '%s', price = '%f', volume = '%s',frequency = '%f' where code = '%s' `,
		request.Name, request.Code, request.Price, request.Volume, request.Frequency, request.Code)

	fmt.Println(q)
	if _, err := tx.ExecContext(ctx, q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}

func (r *stockRepository) UpdateStockWithScraping(request *request.ScrapingRequest) error {
	q := fmt.Sprintf(`update stock set name = '%s', code = '%s', price = '%f', volume = '%s',frequency = '%f' where code = '%s' `,
		request.Name, request.Code, request.Price, request.Volume, request.Frequency, request.Code)

	fmt.Println(q)
	if _, err := r.di.DB.ExecContext(context.Background(), q); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *stockRepository) InsertStocksWithScraping(request *request.ScrapingRequest) error {
	q := `INSERT INTO stock (name,code,price,frequency,volume) values (?,?,?,?,?)`
	_, err := r.di.DB.ExecContext(context.Background(), q, request.Name, request.Code, request.Price, request.Frequency, request.Volume)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *stockRepository) CountStockDataScrap(request *request.GetAllStockRequest) (string, error) {
	var count string
	filter := r.filterQuery(request)
	q := fmt.Sprintf(`select count(*) as count from stock  
	 				  where 1 = 1 %s`, filter)
	if err := r.di.DB.QueryRow(q).Scan(&count); err != nil {
		return "", err
	}
	return count, nil
}
