package handlers

import (
	"database/sql"
	"fmt"
	"github.com/zmh69695328/voting-backend-go/internal/models/sqlite"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)
import "strconv"

// GetTasks endpoint
func GetVotes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTeams(db))
	}
}

// PutTask endpoint
func PutVote(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new task
		var vote models.Vote
		// Map imcoming JSON body to the new Task
		c.Bind(&vote)
		// Add a task using our new model
		// fmt.Println(vote, c)
		log.Println("投票记录：", vote.Score, vote.TeamID, vote.ID, vote.UserName)
		id, err := models.PutVote(db, vote.TeamID, vote.Score, vote.UserName)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

func GetRanking(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var rank models.Rank
		c.Bind(&rank)
		ranks := models.GetRanking(db, rank.Group)
		for i := 0; i < len(ranks.Ranks); i++ {
			ranks.Ranks[i].TotalScore, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", ranks.Ranks[i].TotalScore), 64)
		}
		// Return a JSON response if successful
		return c.JSON(http.StatusOK, ranks)
	}
}

func GetVotesByTeamNameAndUsername(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var vote models.VoteTeam
		c.Bind(&vote)
		votes, err := models.GetVotesByTeamNameAndUsername(db, vote.TeamName, vote.UserName, vote.IsDuplicate)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, votes)
		} else {
			return err
		}
	}
}

func GetVotesHistory(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// var vote models.VoteTeam
		// c.Bind(&vote)
		votes, err := models.GetVotesHistory(db)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusOK, votes)
		} else {
			return err
		}
	}
}
