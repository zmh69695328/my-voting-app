package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/zmh69695328/voting-backend-go/internal/models/sqlite"

	"github.com/labstack/echo/v4"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)

type H map[string]interface{}

// GetTasks endpoint
func GetTeams(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var teams models.TeamCollection = models.GetTeams(db)
		for i := 0; i < len(teams.Teams); i++ {
			teams.Teams[i].ImageUrl = getImageUrl(teams.Teams[i].ID)
		}
		return c.JSON(http.StatusOK, teams)
	}
}

func GetTeamVotes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var vote models.TeamVotes
		// Map imcoming JSON body to the new Task
		c.Bind(&vote)
		var teams models.TeamVotesCollection = models.GetTeamsVotesByUsername(db, vote.UserName)
		return c.JSON(http.StatusOK, teams)
	}
}

var (
	accessKey = "bILga2f-6hoM1r4ZdMoSGRBbvIUtFGMfH7QtTZT8"
	secretKey = "VUD4a0ht_bdEK8MKNaWZ_IORcwPf2s7T23WnM6kO"
	bucket    = "zmhtmp"
)

func getImageUrl(id int) string {
	// const downloadUrl = "http://s2nuoazh3.bkt.clouddn.com/work" + fmt.Sprint(id) + ".png"
	mac := auth.New(accessKey, secretKey)

	// 公开空间访问
	// domain := "http://s2nuoazh3.bkt.clouddn.com"
	// key := "work"+fmt.Sprint(id) + ".png"
	// publicAccessURL := storage.MakePublicURL(domain, key)
	// fmt.Println(publicAccessURL)

	// 私有空间访问
	domain := "https://www.qiniu.zhuminhao.top"
	key := "work" + fmt.Sprint(id) + ".png"
	deadline := time.Now().Add(time.Second * 86400).Unix() //1天有效期
	privateAccessURL := storage.MakePrivateURL(mac, domain, key, deadline)
	fmt.Println(privateAccessURL)
	return privateAccessURL
}

// PutTask endpoint
// func PutTask(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// Instantiate a new task
// 		var task models.Task
// 		// Map imcoming JSON body to the new Task
// 		c.Bind(&task)
// 		// Add a task using our new model
// 		id, err := models.PutTask(db, task.Name)
// 		// Return a JSON response if successful
// 		if err == nil {
// 			return c.JSON(http.StatusCreated, H{
// 				"created": id,
// 			})
// 			// Handle any errors
// 		} else {
// 			return err
// 		}
// 	}
// }

// DeleteTask endpoint
// func DeleteTask(db *sql.DB) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, _ := strconv.Atoi(c.Param("id"))
// 		// Use our new model to delete a task
// 		_, err := models.DeleteTask(db, id)
// 		// Return a JSON response on success
// 		if err == nil {
// 			return c.JSON(http.StatusOK, H{
// 				"deleted": id,
// 			})
// 			// Handle errors
// 		} else {
// 			return err
// 		}
// 	}
// }
