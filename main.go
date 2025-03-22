package main

import (
	//"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type animal struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var animals = []animal{
	{ID: "1", Name: "Max"},
	{ID: "2", Name: "Lilya"},
}

func getAnimals(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, animals)
}

func main() {
	router := gin.Default()
	router.GET("/animals", getAnimals)
	router.Run("localhost:8080")
}
