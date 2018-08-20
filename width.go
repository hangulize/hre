package hre

import (
	"regexp/syntax"
)

// calcMaxWidth calculates the maximum width of the given Regexp pattern.
//
// The return value -1 means than unlimited.
//
func calcMaxWidth(expr string) int {
	re, err := syntax.Parse(expr, syntax.Perl)

	if err != nil {
		// Failed to parse Regexp.
		return 0
	}

	return _regexpMaxWidth(re)
}

// _regexpMaxWidth returns the maximum width from a parsed Regexp.
func _regexpMaxWidth(re *syntax.Regexp) int {
	switch re.Op {

	case syntax.OpNoMatch, syntax.OpEmptyMatch:
		// matches no strings
		// matches empty string
		return 0

	case syntax.OpLiteral:
		// matches Runes sequence
		return len(re.Rune)

	case syntax.OpCharClass, syntax.OpAnyCharNotNL, syntax.OpAnyChar:
		// matches Runes interpreted as range pair list
		// matches any character except newline
		// matches any character
		return 1

	case syntax.OpBeginLine, syntax.OpEndLine:
		// matches empty string at beginning of line
		// matches empty string at end of line
		return 0

	case syntax.OpBeginText, syntax.OpEndText:
		// matches empty string at beginning of text
		// matches empty string at end of text
		return 0

	case syntax.OpWordBoundary, syntax.OpNoWordBoundary:
		// matches word boundary `\b`
		// matches word non-boundary `\B`
		return 0

	case syntax.OpCapture:
		// capturing subexpression with index Cap, optional name Name
		return _regexpMaxWidth(re.Sub0[0])

	case syntax.OpStar, syntax.OpPlus:
		// matches Sub[0] zero or more times
		// matches Sub[0] one or more times return -1
		return -1

	case syntax.OpQuest:
		// matches Sub[0] zero or one times
		return _regexpMaxWidth(re.Sub0[0])

	case syntax.OpRepeat:
		// matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		if re.Max == -1 {
			return -1
		}
		return re.Max * _regexpMaxWidth(re.Sub0[0])

	case syntax.OpConcat:
		// matches concatenation of Subs
		total := 0
		for _, sub := range re.Sub {
			width := _regexpMaxWidth(sub)
			total += width
		}
		return total

	case syntax.OpAlternate:
		// matches alternation of Subs
		max := 0
		for _, sub := range re.Sub {
			width := _regexpMaxWidth(sub)
			if width > max {
				max = width
			}
		}
		return max

	default:
		return -1
	}
}
