package brf

import (
	"reflect"
	"testing"
)

func TestExcludeStrings(t *testing.T) {
	sl1 := []string{"the", "extra", "sly", "brown", "fox", "jumped", "over", "the", "lazy", "hairy", "dog"}
	sl2 := []string{"extra", "hairy"}
	want := []string{"the", "sly", "brown", "fox", "jumped", "over", "the", "lazy", "dog"}

	got := ExcludeStrings(sl1, sl2)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("ExcludeStrings: != DeepEqual got: %v != want: %v)", got, want)
	}
}

func TestExcludePrefix(t *testing.T) {
	for _, tr := range []struct {
		in, want []string
	}{
		{[]string{"a", "b", "c", "a_", "b_", "c_"}, []string{"a", "b", "c", "a_", "b_", "c_"}},
		{[]string{"_x", "_y", "_z", "a", "b", "c"}, []string{"a", "b", "c"}},
	} {

		got := ExcludePrefix(tr.in, "_")

		if !reflect.DeepEqual(got, tr.want) {
			t.Errorf("ExcludePrefix: != DeepEqual got: %v != want: %v", got, tr.want)
		}
	}
}

func TestOnlyPrefix(t *testing.T) {
	for _, tr := range []struct {
		in, want []string
	}{
		{[]string{"_x", "_y", "_z", "a", "b", "c"}, []string{"_x", "_y", "_z"}},
	} {

		got := OnlyPrefix(tr.in, "_")

		if !reflect.DeepEqual(got, tr.want) {
			t.Errorf("OnlyPrefix: != DeepEqual got: %v != want: %v", got, tr.want)
		}
	}

}

func TestReduce(t *testing.T) {

	for _, tr := range []struct {
		in, want []string
	}{
		{[]string{"a", "b", "b", "c", "c", "c", "c"}, []string{"a", "b", "c"}},
		{[]string{"x", "y", "z", "z", "z", "z"}, []string{"x", "y", "z"}},
	} {

		got := Reduce(tr.in)

		if !reflect.DeepEqual(got, tr.want) {
			t.Errorf("Single: != DeepEqual (%v -> %v != %v)", tr.in, got, tr.want)
		}
	}
}

func TestSummary(t *testing.T) {

	sl1 := []string{"the", "sly", "brown", "jumped", "over", "the", "lazy", "dog"}
	sl2 := []string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", ".", "Maecenas"}
	sl3 := []string{"a", "b", "c", "d", "e", "f", "g"}

	for _, tr := range []struct {
		sl   []string
		l    int
		want string
	}{
		{sl1, 15, "the, sly, brown, jumped..."},
		{sl2, 25, "Lorem, ipsum, dolor, sit, amet..."},
		{sl3, 20, "a, b, c, d, e, f, g"},
	} {

		got := Summary(tr.sl, tr.l)

		if len(got) >= (len(got) + 12) {
			t.Errorf("Summary: != len(got) (%v >= %v)", len(got), (len(got) + 12))
		}

		if got != tr.want {
			t.Errorf("Summary: (%v != %v)", got, tr.want)
		}

	}
}

func TestFirst(t *testing.T) {

	s1 := "First line.\n Second line. \n Third line.\n"
	s2 := "Only one line."
	s3 := ""

	for _, tr := range []struct {
		in, want string
	}{
		{s1, "First line."},
		{s2, "Only one line."},
		{s3, ""},
	} {

		got := First(tr.in)

		if got != tr.want {
			t.Errorf("First: (%v != %v)", got, tr.want)
		}
	}
}

func TestMatchLine(t *testing.T) {

	s1 := "- user: jychri "
	s2 := "  oath_token: 324\n"

	for _, tr := range []struct {
		in, pfx, want string
	}{
		{s1, "- user:", "jychri"},
		{s1, " - user:", "jychri"},
		{s2, "oath_token:", "324"},
	} {

		got := Match(tr.in, tr.pfx)

		if got != tr.want {
			t.Errorf("MatchLine: ('%v' != '%v')", got, tr.want)
		}
	}
}
