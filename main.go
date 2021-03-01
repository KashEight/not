package main

import (
	"github.com/KashEight/not/external"
	"github.com/KashEight/not/router"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"net/http"
)

func main() {
	engine := gin.Default()

	db := external.DBInit()

	router.Init(engine, db)

	_ = http.ListenAndServe("localhost:8080", engine)
}
