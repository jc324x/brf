brf
===

[![Go Report Card](https://goreportcard.com/badge/github.com/jychri/brf)](https://goreportcard.com/report/github.com/jychri/brf) [![GoDoc](https://godoc.org/github.com/jychri/brf?status.svg)](https://godoc.org/github.com/jychri/brf)

Functions in package `brf` filter, reduce or transform strings and string slices. 

## Examples ##

```go
func After(s string, m string) (string, error)

func ExcludePrefix(esl []string, pfx string) (sl []string)

func ExcludeStrings(msl []string, esl []string) (sl []string)

func FirstLine(s string) (string, error)

func LowerKebab(s string) string

func NoBlanks(bsl []string) (sl []string)

func OnlyPrefix(psl []string, pfx string) (sl []string)

func Summary(sl []string, l int) string

sl := []string{"the", "sly", "brown", "fox", "jumped", "over", "the", "lazy"}
s := brf.Summary(sl, 15)
fmt.Println(s) // "the, sly, brown, fox..." 

func Trim(s string) string
```

## Projects ## 

- [git-in-sync](https://github.com/jychri/git-in-sync): Keep all of
  your git repos in sync across multiple computers
- [js2x](https://github.com/jychri/js2x): Convert Javascript to Markdown and beyond with js2x
- [wrangler](https://github.com/jychri/wrangler): Wrangle up macOS user accounts and groups
