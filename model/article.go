package model

import "time"

type ArticleInfo struct {
	Id           int       `db:"id"`
	Title        string    `db:"title"`
	ViewCount    int       `db:"view_count"`
	CommentCount int       `db:"comment_count"`
	CategoryId   int       `db:"category_id"`
	UserId       int       `db:"user_id"`
	Status       int       `db:"status"`
	Summary      string    `db:"summary"`
	CreateTime   time.Time `db:"create_time"`
}

type ArticleDetail struct {
	ArticleInfo
	Content string `db:"content"`
	Category
}

// 用于文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
