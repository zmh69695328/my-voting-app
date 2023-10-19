package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Team data
type Vote struct {
	ID     int `json:"id"`
	Score  int `json:"score"`
	TeamID int `json:"teamid"`
}

// TaskCollection is collection of Tasks
type VoteCollection struct {
	Teams []Team `json:"teams"`
}

// GetTasks from the DB
// func GetTeams(db *sql.DB) TeamCollection {
// 	sql := "SELECT * FROM teams"
// 	rows, err := db.Query(sql)
// 	// Exit if the SQL doesn't work for some reason
// 	if err != nil {
// 		panic(err)
// 	}
// 	// make sure to cleanup when the program exits
// 	defer rows.Close()

// 	result := TeamCollection{}
// 	for rows.Next() {
// 		team := Team{}
// 		err2 := rows.Scan(&team.ID, &team.TeamName, &team.WorkName, &team.Order, &team.Group)
// 		// Exit if we get an error
// 		if err2 != nil {
// 			panic(err2)
// 		}
// 		result.Teams = append(result.Teams, team)
// 	}
// 	return result
// }

// PutTask into DB
func PutVote(db *sql.DB, teamid int, score int) (int64, error) {
	// sql := "INSERT INTO tasks(name) VALUES(?)"
	fmt.Println(teamid, score)
	// score = 10
	sql := "INSERT INTO vote(teamid, score) VALUES(?, ?)"
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(teamid, score)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask from DB
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
