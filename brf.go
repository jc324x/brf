// Package brf collects functions with brief output.
package brf

import (
	"bytes"
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

// ExcludeStrings returns a string slice, without the excluded strings.
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

// ExcludePrefix returns a string slice, without the excluded prefix strings.
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

// OnlyPrefix returns a string slice, with only matching prefix strings.
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

// NoDuplicates returns a string slice with no duplicate entries.
func NoDuplicates(dsl []string) (sl []string) {
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

// Summary returns string that summarizes the contents of a string slice.
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

// FirstLine returns the first line from a multi line string.
func FirstLine(s string) string {
	lines := strings.Split(strings.TrimSuffix(s, "\n"), "\n")

	switch {
	case len(lines) == 0:
		return ""
	case len(lines) >= 1:
		return lines[0]
	default:
		return ""
	}
}

// After returns the substring following a matching substring.
func After(s string, m string) string {
	m = strings.TrimSpace(m)
	s = strings.TrimSpace(s)

	if !strings.Contains(s, m) {
		return ""
	}

	s = strings.TrimPrefix(s, m)
	s = strings.TrimSpace(s)
	return s
}

// LowerKebab returns a string in lower kebab case.
func LowerKebab(s string) string {
	s = strings.ToLower(s)
	s = strings.Replace(s, " ", "-", -1)
	s = strings.Replace(s, "_", "-", -1)
	return s
}

// Trim returns a string without trailing spaces and line breaks.
func Trim(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, "\n")
	return s
}
