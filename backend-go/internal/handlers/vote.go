package handlers

import (
	"database/sql"
	"github.com/zmh69695328/voting-backend-go/internal/models/sqlite"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
		id, err := models.PutVote(db, vote.TeamID, vote.Score)
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
