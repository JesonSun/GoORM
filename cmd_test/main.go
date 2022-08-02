package main

import (
	"GoORM"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func main() {
	engine, _ := GoORM.NewEngine("mysql", "root:123456@(localhost)/douyin?charset=utf8mb4&parseTime=True&loc=Local")
	defer engine.Close()
	s := engine.NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	if err1 != nil || err2 != nil {
		fmt.Println("error12")
	}
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		fmt.Println("error")
	}
}
