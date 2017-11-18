package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/teris-io/shortid"
	"gopkg.in/mgo.v2"
)

type shortenResponse struct {
	OriginalURL string `json:"original_url"`
	ShortID     string `json:"short_id"`
	ShortURL    string `json:"short_url"`
}

type link struct {
	ID  string `bson:"_id"`
	URL string `bson:"url"`
}

var links *mgo.Collection

func main() {
	// Connect to mongo
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	// Get links collection
	links = session.DB("app").C("links")

	r := gin.Default()
	// Set up routes
	api := r.Group("/api")
	{
		api.POST("/shorten/:url", shortenEndpoint)
	}
	// Run HTTP server
	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

func shortenEndpoint(c *gin.Context) {
	url := c.Param("url")
	id := shortid.MustGenerate()

	if err := links.Insert(link{URL: url, ID: id}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, shortenResponse{
		OriginalURL: url,
		ShortID:     id,
		ShortURL:    fmt.Sprintf("%s/%s", c.Request.Host, id),
	})
}
