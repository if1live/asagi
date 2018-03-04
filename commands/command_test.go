package commands

import "testing"

func Test_splitCommand(t *testing.T) {
	cases := []struct {
		text string
		cmd  string
		arg  string
	}{
		{"", "", ""},
		{" ", "", ""},
		{"/", "/", ""},
		{"/foo", "/foo", ""},
		{"/한글", "/한글", ""},

		// has whitespace
		{" /foo", "/foo", ""},
		{"/foo ", "/foo", ""},
		{" /foo ", "/foo", ""},

		// parameters
		{"/foo bar", "/foo", "bar"},
		{"/foo\tbar", "/foo", "bar"},
		{"/foo bar spam", "/foo", "bar spam"},
	}
	for _, c := range cases {
		cmd, arg := splitCommand(c.text)
		if c.cmd != cmd {
			t.Errorf("expected %#v, got %#v", c.cmd, cmd)
		}
		if c.arg != arg {
			t.Errorf("expected %#v, got %#v", c.arg, arg)
		}
	}
}
