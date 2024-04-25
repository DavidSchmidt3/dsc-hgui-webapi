package main

import (
	"log"
	"os"
	"strings"

	"context"
	"time"

	"github.com/DavidSchmidt3/dsc-hgui-webapi/api"
	"github.com/DavidSchmidt3/dsc-hgui-webapi/internal/db_service"
	hgui_api "github.com/DavidSchmidt3/dsc-hgui-webapi/internal/hgui-webapi"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("HGUI_API_PORT")
	if port == "" {
		port = "8080"
	}
	environment := os.Getenv("HGUI_API_ENVIRONMENT")
	if !strings.EqualFold(environment, "production") { // case insensitive comparison
		gin.SetMode(gin.DebugMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{""},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
	engine.Use(corsMiddleware)

	dbService := db_service.NewMongoService[hgui_api.Ambulance](db_service.MongoServiceConfig{})
	defer dbService.Disconnect(context.Background())
	engine.Use(func(ctx *gin.Context) {
		ctx.Set("db_service", dbService)
		ctx.Next()
	})

	// request routings
	hgui_api.AddRoutes(engine)
	engine.GET("/openapi", api.HandleOpenApi)
	engine.Run(":" + port)
}
