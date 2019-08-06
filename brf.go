// Package brf collects functions with brief output.
package brf

import (
	"bytes"
	"errors"
	"strings"
)

// NoBlanks returns a string slice with no blank entries.
func NoBlanks(bsl []string) (sl []string) {
	for i := range bsl {
		if bsl[i] != "" {
			sl = append(sl, bsl[i])
		}
	}
	return sl
}

// ExcludeStrings returns a string slice, excluding the entries from a second slice.
func ExcludeStrings(msl []string, esl []string) (sl []string) {
	for _, s := range msl {
	inner:
		for i, c := range esl {
			max := len(esl) - 1
			switch {
			case s == c:
				break inner
			case s != c && i == max:
				sl = append(sl, s)
			}
		}
	}
	return sl
}

// ExcludePrefix returns a string slice, excluding entries with the prefix.
func ExcludePrefix(esl []string, pfx string) (sl []string) {
	for _, s := range esl {
		if strings.HasPrefix(s, pfx) {
			continue
		} else {
			sl = append(sl, s)
		}
	}
	return sl
}

// OnlyPrefix returns a string slice, with only entries with the matching prefix.
func OnlyPrefix(psl []string, pfx string) (sl []string) {
	for _, s := range psl {
		if strings.HasPrefix(s, pfx) {
			sl = append(sl, s)
		} else {
			continue
		}
	}

	return sl
}

// TODO: rename to ReduceStrings

// Reduce returns a string slice with no duplicate entries.
func Reduce(dsl []string) (sl []string) {
	smap := make(map[string]bool)

	for i := range dsl {
		if smap[dsl[i]] == true {
		} else {
			smap[dsl[i]] = true
			sl = append(sl, dsl[i])
		}
	}

	return sl
}

// Summary returns a string $l characters long, summarizing the contents of a string slice.
func Summary(sl []string, l int) string {

	if len(sl) == 0 {
		return ""
	}

	var csl []string // check slice
	var b bytes.Buffer

	for _, s := range sl {
		lc := len(strings.Join(csl, ", ")) // (l)ength(c)heck
		switch {
		case lc <= l-10 && len(s) <= 20: //
			csl = append(csl, s)
		case lc <= l && len(s) <= 12:
			csl = append(csl, s)
		}
	}

	b.WriteString(strings.Join(csl, ", "))

	if len(sl) != len(csl) {
		b.WriteString("...")
	}

	return b.String()
}

// First returns the first line from a multi line string.
func First(s string) (string, error) {
	lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")

	switch {
	case len(lines) == 0:
		return "", errors.New("Len == 0")
	case len(lines) >= 1:
		return lines[0], nil
	default:
		return "", errors.New("Unacceptable lines len")
	}
}

// After returns the substring after a matching substring.
func After(s string, m string) (string, error) {
	m = strings.TrimSpace(m)
	s = strings.TrimSpace(s)

	if !strings.Contains(s, m) {
		return "", errors.New("No match found")
	}

	s = strings.TrimPrefix(s, m)
	s = strings.TrimSpace(s)
	return s, nil
}

// TODO: ADD TEST

// LowerKebab converts a string to lower kebab case
func LowerKebab(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	return s
}

// Trim trims trailing spaces and "\n" markings from a string...
func Trim(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "\n")
	return s
}
