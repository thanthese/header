package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var headingChars = []string{"=", "-"}
var maxLevel = len(headingChars)

const defaultWidth int = 75

func main() {
	level := flag.Int("d", -1, fmt.Sprintf("(optional) Header `depth` 0-%d, or -1 to derive from input.", maxLevel))
	width := flag.Int("w", defaultWidth, "(optional) Header `width`.")
	plain := flag.Bool("p", false, "(optional) Don't include banner -- \"`plain`\". (default with banners)")
	flag.Parse()

	if *level < -1 || *level > maxLevel {
		fmt.Printf("ERROR: level must be between -1 and %d\n", maxLevel)
		return
	}
	if *width < 5 {
		fmt.Println("ERROR: width must be greater than 5.")
		return
	}

	bytes, _ := ioutil.ReadAll(os.Stdin)
	fmt.Printf(BuildHeader(string(bytes), *level, *width, *plain))
}

// Main program. Takes (checked for correctness) parameters from user.
func BuildHeader(s string, level int, width int, plain bool) string {
	if level == -1 {
		level = len(regexp.MustCompile(`^#*`).FindString(s))
		if level > maxLevel {
			return fmt.Sprintf("ERROR: depth cannot exceed %d", maxLevel)
		}
	}
	nakedHeader := stripHeaders(s)
	return addHeaders(nakedHeader, level, width, plain)
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
	if level == 0 {
		return s
	}
	n := strings.Repeat("#", level) + " " + s
	if plain || len(n)+3 > width {
		return n
	}
	if len(s) == 0 {
		return n + strings.Repeat(headingChars[level-1], width-len(n))
	}
	return n + " " + strings.Repeat(headingChars[level-1], width-len(n)-1)
}
