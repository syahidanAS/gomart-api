package models

import "time"

type (
	//Product
	Product struct{
		ID int 					`json:"id"`
		Name string 			`json:"name"`
		Qty	int					`json:"qty"`
		Price int				`json:"price"`
		CreatedAt time.Time 	`json:"createdAt"`
		UpdateAt time.Time 		`json:"updatedAt"`
	}
)
