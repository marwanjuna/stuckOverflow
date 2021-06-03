package main

import (
	"os"

	"github.com/Excellent-Echo/stuckOverflow/API/API/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// r.Use(handler.CorsMiddleware())

	routes.UserRoute(r)
	routes.QuestionRoute(r)
	routes.CategoryRoute(r)
	routes.AnswerRoute(r)
	routes.JobRoute(r)

	port := os.Getenv("PORT")
	r.Run(":" + port)
}
