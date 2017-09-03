/*
 * John Adams
 * jna@retina.net
 * 8/30/2017
 *
 * slackwc: main.go
 * Secure word-counting microservice in Golang and Gin.
 *
 */

package main

import (
	"encoding/json"

	"slackwc/api"

	"github.com/gin-gonic/gin"
)

func WordCount(c *gin.Context) {
	// http://localhost:8080/api/v1/wc

	var iw api.WordRequest
	var wl api.WordList

	// Attempt to decode the incoming request
	decoder := json.NewDecoder(c.Request.Body)

	err := decoder.Decode(&iw)

	// reject if we fail the decode or if input is blank
	if err != nil {
		api.RespondWithError(400, "Bad Request", c)
		return
	}

	if iw.Input == "" {
		api.RespondWithError(400, "Bad Request", c)
		return
	}

	// count words
	counts := api.WordCounter(iw.Input)
	wl.Count = len(counts)
	wl.Words = counts

	// return our response
	c.JSON(200, wl)
}

func GetMainEngine() *gin.Engine {
	r := gin.Default()

	// all items in our API require authentication
	v1 := r.Group("api/v1", api.AuthRequired())
	{
		v1.POST("/wc", WordCount)
	}

	// handle page not found with JSON response
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	return r
}

func main() {
	GetMainEngine().Run(api.ListenPort)
}
