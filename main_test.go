package main

import "testing"

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
		{"test", 1, 12, false, "# test ====="},
		{"test ttttt", 2, 13, false, "## test ttttt"},
		{"test tttt", 2, 13, false, "## test tttt"},
		{"test ttt", 2, 13, false, "## test ttt"},
		{"test tt", 2, 13, false, "## test tt --"},
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
