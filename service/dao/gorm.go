package dao

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"user_centor/service/model"
)

var DB *gorm.DB
func init(){
	dsn := "root:123456@tcp(127.0.0.1:3306)/grpc?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err!=nil{
		fmt.Println("打开数据库出错：",err)
		panic(err.Error())
	}
	DB=db
	err=db.AutoMigrate(&model.User{})
	if err!=nil{
		fmt.Println("automigrate err",err)

	}
}

