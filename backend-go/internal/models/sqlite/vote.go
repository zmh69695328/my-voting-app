package models

import (
	"database/sql"
	"log"
	"time"

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
	Leader     string  `json:"leader"`
	WorkName   string  `json:"workname"`
	Group      string  `json:"group"`
	TeamName   string  `json:"teamname"`
	TotalScore float64 `json:"total_score"`
}

// TaskCollection is collection of Tasks
type RankCollection struct {
	Ranks []Rank `json:"ranks"`
}

type VoteTeam struct {
	Vote
	TeamName    string `json:"teamname"`
	WorkName    string `json:"workname"`
	Order       string `json:"order"`
	Group       string `json:"group"`
	Leader      string `json:"leader"`
	IsDuplicate bool   `json:"isduplicate"`
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
			MAX(v.date),
			v.username
		FROM
			vote v
		GROUP BY v.teamid,v.username
	)
	SELECT
		t.leader,
		t.workname,
		t."group",
		t.teamname,
		IFNULL(AVG(lv.score), 0) AS total_score
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
	// fmt.Println(teamid, score)
	// score = 10
	// 获取北京时间
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Println(err)
	}
	beijingTime := time.Now().In(loc).Format("2006-01-02 15:04:05.999")
	sql := "INSERT INTO vote(teamid, score,username,date) VALUES(?, ?,?,?)"
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(teamid, score, username, beijingTime)
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

func GetVotesByTeamNameAndUsername(db *sql.DB, teamName, username string, isDuplicate bool) (VoteTeamCollection, error) {
	var query string
	if isDuplicate {
		query = `SELECT v.id, t.teamname, t.workname, t."group", t.leader, v.score, MAX(v.date), v.username " +
		"FROM vote v JOIN team t ON v.teamid = t.id WHERE 1`

	} else {
		query = `SELECT v.id, t.teamname, t.workname, t."group", t.leader, v.score, v.date, v.username " +
		"FROM vote v JOIN team t ON v.teamid = t.id WHERE 1`
	}

	var args []interface{}

	if teamName != "" {
		query += " AND (t.teamname = ?)"
		args = append(args, teamName)
	}

	if username != "" {
		query += " AND (v.username = ?)"
		args = append(args, username)
	}
	if isDuplicate {
		query += " GROUP BY v.teamid,v.username"
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

type VoteTeamHistory struct {
	Vote
	TeamName  string         `json:"teamname"`
	WorkName  string         `json:"workname"`
	VoteCount int            `json:"votecount`
	Voters    sql.NullString `json:"voters"`
}
type VoteTeamHistoryCollection struct {
	Round1 []VoteTeamHistory `json:"round1"`
	Round2 []VoteTeamHistory `json:"round2"`
}

func GetVotesHistory(db *sql.DB) (VoteTeamHistoryCollection, error) {
	query := `SELECT 
	team.id, 
	team.teamname, 
	team.workname, 
	COUNT(vote.id) as votecount, 
	GROUP_CONCAT(vote.username, ',') as voters
		FROM team
		LEFT JOIN (
			SELECT 
				vote.id, 
				vote.teamid, 
				vote.score, 
				MAX(vote.date), 
				vote.username
			FROM vote
			GROUP BY vote.teamid,vote.username
		) AS vote ON team.id = vote.teamid
		WHERE team."group" = '最佳落地奖'
		GROUP BY team.id`

	var args []interface{}
	var votes VoteTeamHistoryCollection
	rows, err := db.Query(query, args...)
	if err != nil {
		return votes, err
	}
	defer rows.Close()

	for rows.Next() {
		var vote VoteTeamHistory
		err := rows.Scan(&vote.ID, &vote.TeamName, &vote.WorkName, &vote.VoteCount, &vote.Voters)
		if err != nil {
			return votes, err
		}
		votes.Round1 = append(votes.Round1, vote)
	}

	query = `SELECT 
	team.id, 
	team.teamname, 
	team.workname, 
	COUNT(vote.id) as vote_count, 
	GROUP_CONCAT(vote.username, ',') as voters
FROM team
LEFT JOIN (
	SELECT 
		vote.id, 
		vote.teamid, 
		vote.score, 
		MAX(vote.date), 
		vote.username
	FROM vote
	GROUP BY vote.teamid,vote.username
) AS vote ON team.id = vote.teamid
WHERE team."group" = '人工智能' OR team."group" = '数据赋能' OR team."group" = '融合创新'
GROUP BY team.id`
	rows, err = db.Query(query, args...)
	if err != nil {
		return votes, err
	}
	defer rows.Close()

	for rows.Next() {
		var vote VoteTeamHistory
		err := rows.Scan(&vote.ID, &vote.TeamName, &vote.WorkName, &vote.VoteCount, &vote.Voters)
		if err != nil {
			return votes, err
		}
		votes.Round2 = append(votes.Round2, vote)
	}

	return votes, nil
}
