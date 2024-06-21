package main

import (
	"database/sql"
	"fmt"
	"log"

	goora "github.com/sijms/go-ora/v2"
)

func main() {

	// Register the go-ora driver
	sql.Register("goora", goora.NewDriver())

	// Oracle 数据库连接信息
	dsn := goora.BuildUrl("10.251.16.185", 1521, "histdb", "SUTIE", "5Jxz6T^6$", nil)

	// 创建数据库连接
	db, err := sql.Open("goora", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 检查连接
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	var banner string
	querySQL := "SELECT banner FROM v$version"
	err = db.QueryRow(querySQL).Scan(&banner)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Oracle Version: %s\n", banner)
}
