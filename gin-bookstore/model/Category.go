package model

type Category struct {
	ID uint `json:"id" gorm:"primary_key"` //默认为主键
	Name string `json:"name" gorm:"type:varchar(50);not null;unique"`
	CreatedAt Time `json:"created_at" gorm:"type:timestamp"` //创建时默认插入
	UpdatedAt Time `json:"updated_at" gorm:"type:timestamp"` //更新时默认更新
}