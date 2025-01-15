package entity

import(
	"gorm.io/gorm"
)

type Member struct{
	gorm.Model 
	UserName string `gorm:"uniqueIndex" valid:"required~UserName not null"`
	Password string `valid:"required~Password not null"`
	Email string `gorm:"uniqueIndex" valid:"required~Email not null , email~Email is invalid"`
} 