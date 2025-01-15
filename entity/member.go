package entity
import (
	//"github.com/onsi/gomega"
	//"time"
	"gorm.io/gorm")

type Member struct{
	gorm.Model
	UserName string `gorm:"uniqueIndex" valid:"required~UserName not null"`
	
	Password string
	Email string `gorm:"uniqueIndex" valid:"required~Email not null ,email~Email is invalid"`
}