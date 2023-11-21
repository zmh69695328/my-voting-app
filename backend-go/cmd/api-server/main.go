package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zmh69695328/voting-backend-go/internal/handlers"
	"github.com/zmh69695328/voting-backend-go/internal/models/sqlite"
)

func main() {

	db := initDB("storage.db")
	migrate(db)

	// Echo instance
	e := echo.New()

	e.POST("/api/vote", handlers.PutVote(db))
	e.GET("/api/teams", handlers.GetTeams(db))
	e.POST("/api/teams", handlers.GetTeamsByGroup(db))
	e.POST("/api/ranking", handlers.GetRanking(db))
	e.POST("/api/teamvotes", handlers.GetTeamVotes(db))
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:9222"))
}

// Handler
func hello(c echo.Context) error {
	models.SQLiteDemo()
	return c.String(http.StatusOK, "Hello, World!")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { // if err is not nil
		panic(err)
	}

	if db == nil { // if db is nil
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS "team" (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"teamname" text,
		"workname" TEXT,
		"order" TEXT,
		"group" TEXT
	  );
	  CREATE TABLE IF NOT EXISTS "vote" (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"teamid" INTEGER NOT NULL,
		"score" INTEGER NOT NULL,
		"date" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	  );
	`
	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}

}
