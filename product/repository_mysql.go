package product

import (
	"api/config"
	"api/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

const (
	table          = "products"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Product, error) {

	var products []models.Product

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)

	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var product models.Product
		var createdAt, updatedAt string

		if err = rowQuery.Scan(&product.ID,
			&product.Name,
			&product.Qty,
			&product.Price,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		product.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		product.UpdateAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)
	}

	return products, nil
}

func Insert(ctx context.Context, prd models.Product)error{
	db, err := config.MySQL()

	if err != nil{
		log.Fatal("Can't connect to MySQL", err)
	}
	queryText := fmt.Sprintf("INSERT INTO %v (name, qty, price, createdat, updatedat) values('%v','%v','%v','%v','%v')", table,
		prd.Name,
		prd.Qty,
		prd.Price,
		time.Now().Format(layoutDateTime),
		time.Now().Format(layoutDateTime))

	_,err = db.ExecContext(ctx,queryText)
	if err!= nil{
		return err
	}
	return nil
}

func Update (ctx context.Context, prd models.Product)error{
	db, err := config.MySQL()

	if err != nil{
		log.Fatal("Can't connect to MySQL", err)
	}
	queryText := fmt.Sprintf("UPDATE %v set name = '%s', qty=%d, price = %d, updatedat = '%v' WHERE id = %d",
		table,
		prd.Name,
		prd.Qty,
		prd.Price,
		time.Now().Format(layoutDateTime),
		prd.ID,
	)
	fmt.Println(queryText)
	_,err = db.ExecContext(ctx, queryText)

	if err != nil{
		return err
	}
	return nil
}


// Delete
func Delete(ctx context.Context, prd models.Product) error {

	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = '%d'", table, prd.ID)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ditemukan")
	}

	return nil
}