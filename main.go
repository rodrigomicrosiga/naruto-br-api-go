package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Clan string `json:"clan"`
}

var characters = []Character{
	{ID: 1, Name: "Naruto Uzumaki", Clan: "Uzumaki"},
	{ID: 2, Name: "Sasuke Uchiha", Clan: "Uchiha"},
	{ID: 3, Name: "Sakura Haruno", Clan: "Haruno"},
}

func main() {
	r := gin.Default()

	// Endpoint para listar personagens
	r.GET("/characters", func(c *gin.Context) {
		c.JSON(http.StatusOK, characters)
	})

	// Endpoint para buscar personagem por ID
	r.GET("/characters/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, char := range characters {
			if id == string(rune(char.ID)) {
				c.JSON(http.StatusOK, char)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": "Character not found"})
	})

	r.Run(":8080") // Executa o servidor na porta 8080
}
