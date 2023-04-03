package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Post struct {
	ID uuid.UUID `gorm:"type:char(36);primary_key"`
	UserId uint `json:"user_id" gorm:"not null"`
	CategoryId uint `json:"Category_id" gorm:"not null"`
	Category *Category
	Title string `json:"title" gorm:"type:varchar(50);not null"`
	HeadImg string `json:"head_img"`
	Content string `json:"content" gorm:"type:text;nt null"`
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"` //创建时默认插入
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"` //更新时默认更新
}

func (post *Post)  BeforeCreate(scope *gorm.Scope)error{
	return scope.SetColumn("ID",uuid.NewV4())
}