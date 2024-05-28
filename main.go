package main

import (
	"vermouth-backend/configs"
	"vermouth-backend/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {

	db_record_pg := configs.Connect_tmrecord()
	fmt.Println("DB golang Postgresql = ", db_record_pg)
	defer db_record_pg.Close()

	db_analysis_pg := configs.Connect_tmanalysis()
	fmt.Println("DB golang Postgresql = ", db_analysis_pg)
	defer db_analysis_pg.Close()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	routes.Routes(router)
	log.Fatal(router.Run(":3020"))
}