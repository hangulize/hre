package hre

var reRepeatition = re(`
--- zero or more
	\*
	|
--- one or more
	\+
	|
--- zero or one
	\?
`)

func hasRepeatition(expr string) bool {
	for _, m := range reRepeatition.FindAllStringIndex(expr, -1) {
		if m[0] == 0 || expr[m[0]-1] != '\\' {
			return true
		}
	}
	return false
}
