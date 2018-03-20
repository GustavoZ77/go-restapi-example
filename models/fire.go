package models

import (
	"fmt"
	"rest-api-example-gin/attacks"
)

type Fire struct {
	TypeAttack []attacks.TypeAttack
}

func (fire *Fire) FireAttack() {
	for _, attack := range fire.TypeAttack {
		fmt.Println(attack.ToAttack())
	}
}
