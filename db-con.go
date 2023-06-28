package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DataBase *gorm.DB
var err error
var urlDns="root:@Ram81718@tcp(127.0.0.1:3306)/ramdb"

func DataMigration() {
	DataBase,err=gorm.Open(mysql.Open(urlDns),&gorm.Config{})
	if err!=nil{
		fmt.Print(err.Error())
		panic("connection failed.......Sorry")
	}
	DataBase.AutoMigrate(&Employee{})

}