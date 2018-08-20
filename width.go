package hre

var reDup = re(`
--- zero or more
	\*
	|
--- one or more
	\+
	|
--- zero or one
	\?
`)

func hasDuplication(expr string) bool {
	for _, m := range reDup.FindAllStringIndex(expr, -1) {
		if m[0] == 0 || expr[m[0]-1] != '\\' {
			return true
		}
	}
	return false
}
