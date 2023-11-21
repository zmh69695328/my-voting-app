package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Team data
type Team struct {
	ID       int    `json:"id"`
	TeamName string `json:"teamname"`
	WorkName string `json:"workname"`
	Order    string `json:"order"`
	Group    string `json:"group"`
	Leader   string `json:"leader"`
	ImageUrl string
}

// TaskCollection is collection of Tasks
type TeamCollection struct {
	Teams []Team `json:"teams"`
}

// Task is a struct containing Team data
type TeamVotes struct {
	ID       int     `json:"id"`
	TeamName string  `json:"teamname"`
	UserName *string `json:"username"`
	Score    *int    `json:"score"`
	Date     *string `json:"date"`
}

// TaskCollection is collection of Tasks
type TeamVotesCollection struct {
	TeamVotes []TeamVotes `json:"teamvotes"`
}

// GetTasks from the DB
func GetTeams(db *sql.DB) TeamCollection {
	sql := "SELECT * FROM team"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TeamCollection{}
	for rows.Next() {
		team := Team{}
		err2 := rows.Scan(&team.ID, &team.TeamName, &team.WorkName, &team.Order, &team.Group, &team.Leader)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Teams = append(result.Teams, team)
	}
	return result
}

// get teams by group
func GetTeamsByGroup(db *sql.DB, group string) TeamCollection {
	sql := `SELECT * FROM team AS t where t."group" = ?`
	fmt.Println(group)
	rows, err := db.Query(sql, group)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		fmt.Println(group)
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TeamCollection{}
	for rows.Next() {
		team := Team{}
		err2 := rows.Scan(&team.ID, &team.TeamName, &team.WorkName, &team.Order, &team.Group, &team.Leader)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Teams = append(result.Teams, team)
	}
	return result
}

func GetTeamsVotesByUsername(db *sql.DB, username *string) TeamVotesCollection {
	sql := `WITH LatestVotes AS (
		SELECT
		  v.teamid,
		  v.username,
		  MAX(v.date) AS latest_date
		FROM vote AS v
		WHERE v.username = ?
		GROUP BY v.teamid, v.username
	  )
	  SELECT
		  t.id,
		t.teamname,
		lv.username,
		v.date,
		v.score
	  FROM team AS t
	  LEFT JOIN LatestVotes AS lv ON t.id = lv.teamid
	  LEFT JOIN vote AS v ON lv.teamid = v.teamid
						 AND lv.username = v.username
						 AND lv.latest_date = v.date
	  ORDER BY t.id;
	  `
	rows, err := db.Query(sql, username)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := TeamVotesCollection{}
	for rows.Next() {
		teamvotes := TeamVotes{}
		err2 := rows.Scan(&teamvotes.ID, &teamvotes.TeamName, &teamvotes.UserName, &teamvotes.Date, &teamvotes.Score)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.TeamVotes = append(result.TeamVotes, teamvotes)
	}
	return result
}
