package main

import (
	"fmt"
	"gorm.io/gorm"
    "gorm.io/driver/mysql"
    "os"
)

type User struct {
	Id      int    `gorm:"column:id;primaryKey"`
	Uid     int    `gorm:"column:uid"`
	Keyword string `gorm:"column:keywords"`
	Degree  string `gorm:"column:degree"`
	Gender  string `gorm:"column:gender"`
	City    string `gorm:"column:city"`
}

func (User) TableName() string{
	return "user"
}

func main(){
	dataSourceName:="jakehu:wave1995@tcp(127.0.0.1:3306)/test?charset=utf8&&parseTime=True"
	client, err:=gorm.Open(mysql.Open(dataSourceName), nil)
	checkerror(err)
    user:=read(client,"北京")

    if user != nil{
        fmt.Println("%+v\n", *user)
    }else{
        fmt.Println("无结果\n")
    }

}
func read(client *gorm.DB, city string) *User{
	var users []User
	client.Where("city=?", city).Find(&users)
	if len(users) > 0{
		return &users[0]
	}else{
		return nil
	}
} 

func checkerror(err error){
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}