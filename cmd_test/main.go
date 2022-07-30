package main

import (
	"GoORM"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := GoORM.NewEngine("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	defer engine.Close()
	s := engine.NewSession()
	_, _ = s.Raw("DROP TABLE IF EXISTS User;").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	_, _ = s.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := s.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}