/*
	Main file for Idea.Bin
	author: CSB

	Responsible for running the command server
		- Serve homepage
		- Serve profile
*/

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	// Routes that our API will provide
	//
	// /jokes - which will retrieve a list of jokes a user can see
	// /jokes/like/:jokeID - which will capture likes sent to a particular joke
	api.GET("/jokes", JokeHandler)
	api.POST("/jokes/like/:jokeID", LikeJoke)

	// Start and run the server
	router.Run(":3000")
}

// Let's create our Jokes struct. This will contain information about a Joke

// Joke contains information about a single Joke
type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"likes"`
	Joke  string `json:"joke" binding:"required"`
}

// We'll create a list of jokes
var jokes = []Joke{
	Joke{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	Joke{2, 0, "What do you call a fake noodle? An Impasta."},
	Joke{3, 0, "How many apples grow on a tree? All of them."},
	Joke{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	Joke{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	Joke{6, 0, "Why did the coffee file a police report? It got mugged."},
	Joke{7, 0, "How does a penguin build it's house? Igloos it together."},
}

// JokeHandler retrieves a list of available jokes
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, jokes)
}

// Eventually, we need to store the changes within the database through SQLite
// Until then, we can just return pointer to the local variable to test and ensure accuracy

// LikeJoke increments the likes of a particular joke Item
func LikeJoke(c *gin.Context) {
	// confirm Joke ID sent is valid
	// remember to import the `strconv` package
	if jokeid, err := strconv.Atoi(c.Param("jokeID")); err == nil {
		// find joke, and increment likes
		for i := 0; i < len(jokes); i++ {
			if jokes[i].ID == jokeid {
				jokes[i].Likes++
				c.JSON(http.StatusOK, &jokes)
				return
			}
		}

		// return a pointer to the updated jokes list

	} else {
		// Joke ID is invalid
		c.AbortWithStatus(http.StatusNotFound)
	}
}
