package main

import (
	"fmt"

	"github.com/ij4l/go-htmx/apps"
	"github.com/ij4l/go-htmx/apps/admin"
	"github.com/ij4l/go-htmx/apps/home"
	"github.com/ij4l/go-htmx/apps/shop"
	"github.com/jackc/pgx/v5"
)

func main() {
	// config, err := config.LoadConfig("app.env")
	// if err != nil {
	// 	log.Fatalf("Cannot load configuration: %v", err)
	// }

	// pgxConn, err := database.ConnectPostgreSql(config)
	// if err != nil {
	// 	log.Fatalf("Cannot connect to database: %v", err)
	// }

	// server, err := apps.NewServer(config, pgxConn)
	// if err != nil {
	// 	log.Fatalf("Cannot create server: %v", err)
	// }

	// Initialize(server, pgxConn)

	// if err := server.Start(); err != nil {
	// 	log.Fatalf("Cannot start server: %v", err)
	// }

	fmt.Println("Hello, World!")
}

func Initialize(server *apps.Server, pgxConn *pgx.Conn) {
	repo := apps.NewRepository(pgxConn)
	home.InitializeHomeHandler(server.Router, &repo)
	shop.InitializeShopHandler(server.Router, &repo)
	admin.InitializeAdminHandler(server.Router, &repo)
}
