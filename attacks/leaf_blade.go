package attacks

type LeafBlade struct {
	Effectiveness int
}

func (LeafBlade LeafBlade) ToAttack() string {
	attk := "LeafBlade attack power"
	for i := 0; i < LeafBlade.Effectiveness; i++ {
		attk += " ▓▓▓ "
	}
	return attk
}
