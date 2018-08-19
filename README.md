# HRE

[![GoDoc](https://godoc.org/github.com/hangulize/hre?status.svg)](https://godoc.org/github.com/hangulize/hre)
[![Go Report Card](https://goreportcard.com/badge/github.com/hangulize/hre)](https://goreportcard.com/report/github.com/hangulize/hre)
[![Build Status](https://travis-ci.org/hangulize/hre.svg?branch=develop)](https://travis-ci.org/hangulize/hre)

RegExp dialect for Hangulize

- `^` - the beginning of word not line
- `$` - the end of word not line
- `^^` - the beginning of line (`^` in the standard)
- `$$` - the end of line (`$` in the standard)
- `cat{dog}` - "cat" before "dog" (positive lookahead)
- `{dog}cat` - "cat" after "dog" (positive lookbehind)
- `cat{~dog}` - "cat" before not "dog" (negative lookahead)
- `{~dog}cat` - "cat" after not "dog" (negative lookbehind)
- `<var>` - one of letters in the variable "var"
