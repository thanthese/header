package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var headingChars = []string{"=", "-", "\"", "'"}

const defaultWidth int = 75

func main() {
	level := flag.Int("d", 0, "(optional) Header `depth` 1-4, or 0 to derive from input. (default 0)")
	width := flag.Int("w", defaultWidth, "(optional) Header `width`.")
	plain := flag.Bool("p", false, "(optional) Don't include banner -- \"`plain`\". (default with banners)")
	flag.Parse()

	if *level < 0 || *level > 4 {
		fmt.Println("Error: level must be between 0 and 4")
		return
	}
	if *width < 5 {
		fmt.Println("Error: width must be greater than 5.")
		return
	}

	bytes, _ := ioutil.ReadAll(os.Stdin)
	h := BuildHeader(string(bytes), *level, *width, *plain)
	fmt.Println(h)
}

// Main program. Takes (checked for correctness) parameters from user.
func BuildHeader(s string, level int, width int, plain bool) string {
	if level == 0 {
		level = getLevel(s)
		if level == -1 {
			level = 1
		}
	}
	nakedHeader := stripHeaders(s)
	return addHeaders(nakedHeader, level, width, plain)
}

// Return header level of string, -1 for none found.
func getLevel(s string) int {
	switch {
	case strings.HasPrefix(s, "# "):
		return 1
	case strings.HasPrefix(s, "## "):
		return 2
	case strings.HasPrefix(s, "### "):
		return 3
	case strings.HasPrefix(s, "#### "):
		return 4
	default:
		return -1
	}
}

// Strip all possible leading and trailing characters from input string, as
// well as all carriage returns.
func stripHeaders(s string) string {
	n := strings.Replace(s, "\n", "", -1)
	n = regexp.MustCompile("^#* *").ReplaceAllString(n, "")
	for _, ch := range headingChars {
		n = regexp.MustCompile(" +"+ch+"* *$").ReplaceAllString(n, "")
	}
	return n
}

// Given a string with no headers, return a header with the given level and
// width, and with banners for not.
//
// Don't add trailing markings if the header is long enough already, and don't
// add a single trailing mark. Also, if there's no input string don't double up
// on the spaces.
func addHeaders(s string, level int, width int, plain bool) string {
	n := strings.Repeat("#", level) + " " + s
	if plain {
		return n
	}
	if len(n)+3 > width {
		return n
	}
	if len(s) == 0 {
		return n + strings.Repeat(headingChars[level-1], width-len(n))
	}
	return n + " " + strings.Repeat(headingChars[level-1], width-len(n)-1)
}
