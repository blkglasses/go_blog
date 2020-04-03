package model

type Category struct {
	Id           int64     `db:"id"`
	CategoryName string    `db:"category_name"`
	CategoryNo   int       `db:"category_no"`
	Status       int       `db:"status"`
	CreateTime   string `db:"create_time"`
}