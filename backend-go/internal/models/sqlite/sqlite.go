package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func SQLiteDemo() {
	var start = time.Now().UnixMicro()
	database, err := sql.Open("sqlite3", fmt.Sprintf("./alpha%d.db", start))
	if nil != err {
		fmt.Println(err)
	}
	stmt, _ := database.Prepare("create table if not exists user(id integer primary key, firstname text, lastname text)")
	stmt.Exec()

	stmt, _ = database.Prepare("insert into user( firstname, lastname) values(?,?)")
	stmt.Exec("Jack", "Chen")

	var id int
	var firstname string
	var lastname string
	rows, err := database.Query("select id, firstname, lastname from user")
	if nil != err {
		fmt.Println(err)
	}
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	var end = time.Now().UnixMicro()
	fmt.Printf("SQLiteDemo: %d\n", end-start)

}
