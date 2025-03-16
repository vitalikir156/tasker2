package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/vitalikir156/tasker2/db"
	"github.com/vitalikir156/tasker2/routes"
)

func main() {
	httpport, ok := os.LookupEnv("HTTPPORT")
	if !ok {
		httpport = "3000"
	}
	dbstring, ok := os.LookupEnv("DBSTRING")
	if !ok {
		dbstring = "postgres://testuser:p!ssword2717@db:5555/tasker"
	}
	dbconn, err := startdb(dbstring)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbconn.Close(context.Background())
	db.DB = dbconn
	starthttp(httpport)
}

func starthttp(p string) {
	app := fiber.New()
	routes.RegisterProductRoutes(app)
	fmt.Println(app.Listen(":" + p))
}

func startdb(s string) (*pgx.Conn, error) {
	var err error
	var dbc *pgx.Conn
	for i := 0; i < 5; i++ { // для повторных попыток прицепиться к DB
		dbc, err = db.Connect(s)
		if err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		fmt.Println(err)
	}
	return dbc, err
}
