package main

import (
	"fmt"

	easyorm "github.com/Whitea029/easy-orm"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := easyorm.NewEngine("sqlite3", "whitea.db")
	defer engine.Close()
	session := engine.NewSession()
	session.Raw("DROP TABLE IF EXISTS User;").Exec()
	session.Raw("CREATE TABLE User(Name text);").Exec()
	session.Raw("CREATE TABLE User(Name text);").Exec()
	result, _ := session.Raw("INSERT INTO User(`Name`) values (?), (?)", "Tom", "Sam").Exec()
	count, _ := result.RowsAffected()
	fmt.Printf("Exec success, %d affected\n", count)
}
