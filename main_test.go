package main

import "testing"

func TestGetLevel(t *testing.T) {
	cases := []struct {
		in   string
		want int
	}{
		{"# test", 1},
		{"## test", 2},
		{"### test", 3},
		{"#### test", 4},
		{"##### test", -1},
		{"test", -1},
		{"", -1},
	}
	for _, c := range cases {
		got := getLevel(c.in)
		if got != c.want {
			t.Errorf("getLevel(%q) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestStripHeaders(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"test", "test"},
		{"te\n\nst", "test"},
		{"# test", "test"},
		{"## test", "test"},
		{"test ===", "test"},
		{"# test ==", "test"},
		{"## test =", "test"},
		{"test ---", "test"},
		{"# test --", "test"},
		{"## test -", "test"},
		{"test \"\"\"", "test"},
		{"# test \"\"", "test"},
		{"## test \"", "test"},
		{"test '''", "test"},
		{"test ''' ", "test"},
		{"# test ''", "test"},
		{"## test '", "test"},
		{"## \"test\" '", "\"test\""},
		{"## test ''''''''\n", "test"},
		{"##  test '", "test"},
		{"##  test  '", "test"},
		{"##  test  ' ", "test"},
		{"  test  ' ", "test"},
		{" test", "test"},
		{"  test", "test"},
		{"test ", "test"},
		{"test  ", "test"},
	}
	for _, c := range cases {
		got := stripHeaders(c.in)
		if got != c.want {
			t.Errorf("stripHeaders(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestAddHeaders(t *testing.T) {
	cases := []struct {
		in    string
		level int
		width int
		plain bool
		want  string
	}{
		{"", 1, 15, false, "# ============="},
		{"test", 1, 15, false, "# test ========"},
		{"test", 2, 15, false, "## test -------"},
		{"test", 3, 15, false, "### test \"\"\"\"\"\""},
		{"test", 4, 15, false, "#### test '''''"},
		{"test", 4, 12, false, "#### test ''"},
		{"test ttttt", 4, 15, false, "#### test ttttt"},
		{"test tttt", 4, 15, false, "#### test tttt"},
		{"test ttt", 4, 15, false, "#### test ttt"},
		{"test tt", 4, 15, false, "#### test tt ''"},
		{"test", 1, 15, true, "# test"},
		{"test", 2, 15, true, "## test"},
	}
	for _, c := range cases {
		got := addHeaders(c.in, c.level, c.width, c.plain)
		if got != c.want {
			t.Errorf("addHeaders(%q, %v, %v, %v) == %q, want %q",
				c.in, c.level, c.width, c.plain, got, c.want)
		}
	}
}
