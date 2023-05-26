package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Employee struct {
	ID            uint
	Name          string
	Address       string
	Age           uint
	Birthdate     string
	Level         string
	Id_department uint
}

type Department struct {
	ID   uint
	Name string
}

func main()  {
	psqlInfo := "host=localhost port=5432 user=postgres password=kutabumi20 dbname=database_ruangguru sslmode=disable"

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Employee{}, &Department{})

	fmt.Println("Migration Success")

}