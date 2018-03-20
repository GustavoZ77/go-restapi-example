package attacks

type FlameThrower struct {
	Effectiveness int
}

func (flameThrower FlameThrower) ToAttack() string {
	attk := "FlameThrower attack"
	for i := 0; i < flameThrower.Effectiveness; i++ {
		attk += " ░░░ "
	}
	return attk
}
