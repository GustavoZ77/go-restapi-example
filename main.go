package main

import (
	"rest-api-example-gin/attacks"
	"rest-api-example-gin/models"
	"rest-api-example-gin/server"
)

func main() {
	leafBlade := attacks.LeafBlade{20}

	var arrGrassAttackas []attacks.TypeAttack
	arrGrassAttackas = append(arrGrassAttackas, leafBlade)

	flameThrower := attacks.FlameThrower{20}

	var arrFireAttackas []attacks.TypeAttack
	arrFireAttackas = append(arrFireAttackas, flameThrower)

	myPokemon := models.MyPokemon{models.Grass{arrGrassAttackas}, models.Fire{arrFireAttackas}}
	myPokemon.GrassAttack()
	myPokemon.FireAttack()
	server.Init()
}
