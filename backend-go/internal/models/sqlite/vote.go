package models

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Team data
type Vote struct {
	ID       int    `json:"id"`
	Score    int    `json:"score"`
	TeamID   int    `json:"teamid"`
	UserName string `json:"username"`
	Date     string `json:"date"`
}

// TaskCollection is collection of Tasks
type VoteCollection struct {
	Teams []Team `json:"teams"`
}

type Rank struct {
	Leader     string `json:"leader"`
	WorkName   string `json:"workname"`
	Group      string `json:"group"`
	TeamName   string `json:"teamname"`
	TotalScore int    `json:"total_score"`
}

// TaskCollection is collection of Tasks
type RankCollection struct {
	Ranks []Rank `json:"ranks"`
}

type VoteTeam struct {
	Vote
	TeamName string `json:"teamname"`
	WorkName string `json:"workname"`
	Order    string `json:"order"`
	Group    string `json:"group"`
	Leader   string `json:"leader"`
}
type VoteTeamCollection struct {
	VoteTeam []VoteTeam `json:"VoteTeam"`
}

// GetRanking from the DB
func GetRanking(db *sql.DB, group string) RankCollection {
	sql := `WITH LatestVotes AS (
		SELECT
			v.teamid,
			v.score,
			v.date,
			v.username
		FROM
			vote v
		WHERE
			(v.teamid, v.username, v.date) IN (
				SELECT
					teamid,
					username,
					MAX(date) AS latest_date
				FROM
					vote
				GROUP BY
					teamid, username
			)
	)
	SELECT
			t.leader,
			t.workname,
		t."group",
		t.teamname,
		IFNULL(SUM(lv.score), 0) AS total_score
	FROM
		team t
	LEFT JOIN
		LatestVotes lv ON t.id = lv.teamid
	WHERE
		t."group" = ? -- 替换为您要筛选的组名
	GROUP BY
		t."group", t.teamname
	ORDER BY
		total_score DESC;
	`
	rows, err := db.Query(sql, group)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := RankCollection{}
	for rows.Next() {
		rank := Rank{}
		err2 := rows.Scan(&rank.Leader, &rank.WorkName, &rank.Group, &rank.TeamName, &rank.TotalScore)
		// Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Ranks = append(result.Ranks, rank)
	}
	return result
}

// PutTask into DB
func PutVote(db *sql.DB, teamid int, score int, username string) (int64, error) {
	// sql := "INSERT INTO tasks(name) VALUES(?)"
	fmt.Println(teamid, score)
	// score = 10
	sql := "INSERT INTO vote(teamid, score,username) VALUES(?, ?,?)"
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(teamid, score, username)
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

func GetVotesByTeamNameAndUsername(db *sql.DB, teamName, username string) (VoteTeamCollection, error) {
	query := `SELECT v.id, t.teamname, t.workname, t."group", t.leader, v.score, v.date, v.username " +
		"FROM vote v JOIN team t ON v.teamid = t.id WHERE 1`

	var args []interface{}

	if teamName != "" {
		query += " AND (t.teamname = ?)"
		args = append(args, teamName)
	}

	if username != "" {
		query += " AND (v.username = ?)"
		args = append(args, username)
	}
	var votes VoteTeamCollection
	rows, err := db.Query(query, args...)
	if err != nil {
		return votes, err
	}
	defer rows.Close()

	for rows.Next() {
		var vote VoteTeam
		err := rows.Scan(&vote.ID, &vote.TeamName, &vote.WorkName, &vote.Group, &vote.Leader, &vote.Score, &vote.Date, &vote.UserName)
		if err != nil {
			return votes, err
		}
		votes.VoteTeam = append(votes.VoteTeam, vote)
	}

	return votes, nil
}
