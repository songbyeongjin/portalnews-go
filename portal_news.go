package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"io"
	"os"
	"portal_news/infra/gin_route"
	"portal_news/infra/oauth"
)

func main() {
	// *** Set Oauth

	oauth.SetOauthGoogleConfig()

	// Set Oauth ***

	// *** Set Router

	f, _ := os.Create("./server.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin_route.SetRouter()

	// Set Router ***

	// Server Start
	router.Run(":8080")
}