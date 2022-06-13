package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	g := gin.Default()
	g.GET("/:slug", shortUrlJumpHandle)
	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
