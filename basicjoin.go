package piscine

func BasicJoin(elems []string) string {
	var rus string
	for i := 0; i < len(elems); i++ {
		rus += elems[i]
	}
	return rus
}
