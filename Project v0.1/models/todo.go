package models

type Todo struct {
	Id    int64  `json:"id" gorm:"column:id;primaryKey"  `
	Title string `json:"title" binding:"required" gorm:"column:title"`
}
