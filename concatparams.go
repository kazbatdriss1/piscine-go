package piscine

func ConcatParams(args []string) string {
	var rus string
	for i := 0; i < len(args); i++ {
		rus += args[i]
		if i < len(args)-1 {
			rus += "\n"
		}
	}
	return rus
}
