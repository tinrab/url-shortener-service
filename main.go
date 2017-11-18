package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortID     string `json:"short_id"`
}

func main() {
	r := gin.Default()
	r.POST("/shorten/:url", func(c *gin.Context) {
		url := c.Param("url")
		c.JSON(http.StatusOK, ShortenResponse{
			OriginalURL: url,
			ShortID:     "",
		})
	})
	r.GET("/:id", func(c *gin.Context) {
		// TODO
	})
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
