package main

import (
	"fmt"
	"strings"
)

func main() {
	// func Count(s, sep string) int
	fmt.Println(strings.Count("cheese", "e")) // 3

	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("cheese", "ch")) // true

	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("cheese", "se")) // true

	// func Index(s, sep string) int
	fmt.Println(strings.Index("cheese", "e")) // 2

	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-")) // a-b-c

	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("na", 4)) // nananana

	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2)) // oinky oinky oink

	// func Split(s, sep string) []string
	fmt.Printf("%q\n", strings.Split("a-b-c-d-e", "-")) // []string{"a", "b", "c", "d", "e"}

	// func ToLower(s string) string
	fmt.Println(strings.ToLower("CHEESE")) // cheese

	// func ToUpper(s string) string
	fmt.Println(strings.ToUpper("cheese")) // CHEESE

	// func Trim(s string, cutset string) string
	fmt.Println(strings.Trim(" !!! Achtung !!! ", " !")) // Achtung

	// func TrimSpace(s string) string
	fmt.Println(strings.TrimSpace(" \t\n Achtung \n\t\r ")) // Achtung

	// func Fields(s string) []string
	fmt.Printf("%q\n", strings.Fields("  foo bar  baz   ")) // []string{"foo", "bar", "baz"}

	// func Title(s string) string
	fmt.Println(strings.Title("herzlichen glückwunsch")) // Herzlichen Glückwunsch

	// func ToTitle(s string) string
	fmt.Println(strings.ToTitle("herzlichen glückwunsch")) // HERZLICHEN GLÜCKWUNSCH

	// func Compare(a, b string) int
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("b", "a")) // 1
	fmt.Println(strings.Compare("a", "a")) // 0

	// func Contains(s, substr string) bool
	fmt.Println(strings.Contains("chicken", "ken")) // true
	fmt.Println(strings.Contains("chicken", "dmr")) // false

	// func EqualFold(s, t string) bool
	fmt.Println(strings.EqualFold("Go", "go")) // true	

	// func Map(mapping func(rune) rune, s string) string
	fmt.Println(strings.Map(func(r rune) rune {
		if r == 'a' {
			return 'b'
		}
		return r
	}, "abcde")) // bbcde	

	// func SplitN(s, sep string, n int) []string
	fmt.Printf("%q\n", strings.SplitN("a-b-c-d-e", "-", 3)) // []string{"a", "b", "c-d-e"}

	// func SplitAfter(s, sep string) []string
	fmt.Printf("%q\n", strings.SplitAfter("a-b-c-d-e", "-")) // []string{"a-", "b-", "c-", "d-", "e"}

	// func SplitAfterN(s, sep string, n int) []string []string
	fmt.Printf("%q\n", strings.SplitAfterN("a-b-c-d-e", "-", 3)) // []string{"a-", "b-", "c-d-e"}

}
