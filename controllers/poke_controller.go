package controllers

import (
	"github.com/gin-gonic/gin"
)

type PokeController struct{}

func (pokeController PokeController) GetPokemon(c *gin.Context) {
	c.JSON(200, "Hello world")
}
