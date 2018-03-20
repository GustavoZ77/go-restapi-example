package attacks

type AquaJet struct {
	Effectiveness int
}

func (aquaJet AquaJet) ToAttack() string {
	attk := "AquaJet attack "
	for i := 0; i < aquaJet.Effectiveness; i++ {
		attk += " ~~~ "
	}
	return attk
}
