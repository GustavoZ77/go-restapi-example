package server

import (
	"github.com/gin-gonic/gin"
	"rest-api-example-gin/controllers"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	pokeController := new(controllers.PokeController)
	r.GET("/pokecontroller", pokeController.GetPokemon)
	return r
}
