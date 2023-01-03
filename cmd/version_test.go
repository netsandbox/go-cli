package cmd

import (
	"testing"
	"time"
)

func TestConvertStringToTime(t *testing.T) {
	cases := []struct {
		in   string
		want time.Time
	}{
		{"", time.Time{}},
		{"a", time.Time{}},
		{"0", time.Unix(0, 0).UTC()},
		{"1", time.Unix(1, 0).UTC()},
	}
	for _, c := range cases {
		got := convertStringToTime(&c.in)
		if got != c.want {
			t.Errorf("convertStringToTime(%s) == %s, want: %s", c.in, got, c.want)
		}
	}
}
