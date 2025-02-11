package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"api-golang/database"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	s := &Server{router: gin.Default()}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.GET("/", s.handleRoot)
	s.router.GET("/ping", s.handlePing)
}

func (s *Server) handleRoot(c *gin.Context) {
	database.InsertView(c)
	tm, reqCount := database.GetTimeAndRequestCount(c)
	c.JSON(200, gin.H{
		"api":          "go",
		"currentTime":  tm,
		"requestCount": reqCount,
	})
}

func (s *Server) handlePing(c *gin.Context) {
	_, _ = database.GetTimeAndRequestCount(c)
	c.JSON(200, gin.H{"message": "pong"})
}

func main() {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		content, err := ioutil.ReadFile(os.Getenv("DATABASE_URL_FILE"))
		if err != nil {
			log.Fatalf("⛔ Unable to read database URL file: %v", err)
		}
		databaseUrl = string(content)
	}

	if err := database.InitDB(databaseUrl); err != nil {
		log.Fatalf("⛔ Unable to connect to database: %v", err)
	}
	log.Println("DATABASE CONNECTED 🥇")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	s := NewServer()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-quit
		log.Println("Shutting down server gracefully...")
		os.Exit(0)
	}()

	log.Printf("Starting server on :%s", port)
	s.router.Run(":" + port)
}
