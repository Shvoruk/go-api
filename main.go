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

func addAnimal(c *gin.Context) {
	var newAnimal animal
	if err := c.BindJSON(&newAnimal); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	animals = append(animals, newAnimal)
	c.IndentedJSON(http.StatusCreated, newAnimal)
}

func main() {
	router := gin.Default()
	router.GET("/animals", getAnimals)
	router.POST("/animals", addAnimal)
	router.Run("localhost:8080")
}
