package hre

import (
	"regexp/syntax"
)

// calcWidthRange calculates the maximum width of the given Regexp pattern.
//
// The return value -1 means than unlimited.
//
func calcWidthRange(expr string) [2]int {
	re, err := syntax.Parse(expr, syntax.Perl)

	if err != nil {
		// Failed to parse Regexp.
		return [2]int{0, 0}
	}

	return _regexpWidthRange(re)
}

// _regexpWidthRange returns the maximum width from a parsed Regexp.
func _regexpWidthRange(re *syntax.Regexp) [2]int {
	switch re.Op {

	case syntax.OpNoMatch, syntax.OpEmptyMatch:
		// matches no strings
		// matches empty string
		return [2]int{0, 0}

	case syntax.OpLiteral:
		// matches Runes sequence
		n := len(re.Rune)
		return [2]int{n, n}

	case syntax.OpCharClass, syntax.OpAnyCharNotNL, syntax.OpAnyChar:
		// matches Runes interpreted as range pair list
		// matches any character except newline
		// matches any character
		return [2]int{1, 1}

	case syntax.OpBeginLine, syntax.OpEndLine:
		// matches empty string at beginning of line
		// matches empty string at end of line
		fallthrough
	case syntax.OpBeginText, syntax.OpEndText:
		// matches empty string at beginning of text
		// matches empty string at end of text
		fallthrough
	case syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		// matches word boundary `\b`
		// matches word non-boundary `\B`
		return [2]int{0, 0}

	case syntax.OpCapture:
		// capturing subexpression with index Cap, optional name Name
		return _regexpWidthRange(re.Sub0[0])

	case syntax.OpStar:
		// matches Sub[0] zero or more times
		return [2]int{0, -1}

	case syntax.OpPlus:
		// matches Sub[0] one or more times return -1
		return [2]int{1, -1}

	case syntax.OpQuest:
		// matches Sub[0] zero or one times
		nn := _regexpWidthRange(re.Sub0[0])
		return [2]int{0, nn[1]}

	case syntax.OpRepeat:
		// matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		nn := _regexpWidthRange(re.Sub0[0])

		min := re.Min * nn[0]
		max := -1
		if re.Max != -1 {
			max = re.Max * nn[1]
		}

		return [2]int{min, max}

	case syntax.OpConcat:
		// matches concatenation of Subs
		var min, max int

		for _, sub := range re.Sub {
			nn := _regexpWidthRange(sub)
			min += nn[0]
			max += nn[1]
		}

		return [2]int{min, max}

	case syntax.OpAlternate:
		// matches alternation of Subs
		min := -1
		max := 0

		for _, sub := range re.Sub {
			nn := _regexpWidthRange(sub)
			if nn[0] < min || min == -1 {
				min = nn[0]
			}
			if nn[1] > max {
				max = nn[1]
			}
		}

		return [2]int{min, max}

	default:
		return [2]int{0, 0}
	}
}
