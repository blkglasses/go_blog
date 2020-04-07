package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go_blog/model"
)

func InsertCategory(category *model.Category) (categoryId int, err error) {
	sqlstr := "insert into category(category_name,status,category_no) values(?,0,?)"
	result, err := DB.Exec(sqlstr, category.CategoryName, category.CategoryNo)
	if err != nil {
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		return
	}
	categoryId = int(id)
	return
}

// 获取单个分类信息
func GetCategoryById(id int64) (category *model.Category, err error) {
	category = &model.Category{}
	sqlstr := "select id,category_name,status,category_no,create_time from category where id=?"
	err = DB.Get(category, sqlstr, id)
	fmt.Printf("%#v", err)
	return
}

// 获取多个分类信息
func GetCategoryList(ids []int64) (categoryList []*model.Category, err error) {
	sqlstr, arr, err := sqlx.In("select id,category_name,category_no,status,create_time from category where id in(?)", ids)
	fmt.Println(sqlstr)
	fmt.Println(arr)
	if err != nil {
		return
	}
	err = DB.Select(&categoryList, sqlstr, arr...)
	return
}

// 获取所有分类信息
func GetAllCategory() (categoryList []*model.Category, err error) {
	sqlstr := "select id,category_name,category_no,status,create_time from category where id in(1,2,3) order by category_no asc"
	err = DB.Select(&categoryList, sqlstr)
	return

}
