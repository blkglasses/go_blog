package db

import (
	"go_blog/model"
	"testing"
	"time"
)

func init() {
	err := InitDB()
	if err != nil {
		panic(err)
	}
}
func TestInsertCategory(t *testing.T) {
	category := model.Category{
		Id:           0,
		CategoryName: "golang",
		CategoryNo:   0,
		Status:       0,
		CreateTime:   time.Now().String(),
	}
	id, err := InsertCategory(&category)
	if err != nil {
		panic(err)
	}
	t.Logf("插入数据成功，id为：%#v", id)
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(5)
	if err != nil {
		panic(err)
	}
	t.Logf("success:%#v", category)
}
func TestGetCategoryList(t *testing.T) {
	var ids []int64
	ids = append(ids, 1, 2, 3)
	categoryList, err := GetCategoryList(ids)
	if err != nil {
		panic(err)
	}
	t.Logf("success:%T\n", categoryList)
	for _, category := range categoryList {
		t.Logf("%#v\n", category)
	}
}

func TestGetAllCategory(t *testing.T) {
	List, err := GetAllCategory()
	if err != nil {
		panic(err)
	}
	t.Logf("success:%T\n", List)
	for _, category := range List {
		t.Logf("%#v\n", category)
	}
}
