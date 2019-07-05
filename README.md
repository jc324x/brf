# brf # 

A small utility package used in several of my open
source projects:

- [git-in-sync](https://github.com/jychri/git-in-sync): Keep all of
  your git repos in sync across multiple computers
- [js2x](https://github.com/jychri/js2x): Convert Javascript to Markdown and beyond with js2x
- [wrangler](https://github.com/jychri/wrangler): Wrangle up macOS user accounts and groups

It's mostly functions for working with string slices: 

```go
// NoBlanks returns a string slice with no blank entries.
func NoBlanks(bsl []string) (sl []string) {
	for i := range bsl {
		if bsl[i] != "" {
			sl = append(sl, bsl[i])
		}
	}
	return sl
}
```

Or for generating summaries of string slices:

```go
// Summary returns a set length string summarizing the contents of a string slice.
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
```
