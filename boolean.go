package faker

func Boolean() bool {
	return IntBetween(0, 100) > 50
}
