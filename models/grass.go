package models

import (
	"fmt"
	"rest-api-example-gin/attacks"
)

type Grass struct {
	TypeAttack []attacks.TypeAttack
}

func (grass *Grass) GrassAttack() {
	for _, attack := range grass.TypeAttack {
		fmt.Println(attack.ToAttack())
	}
}
