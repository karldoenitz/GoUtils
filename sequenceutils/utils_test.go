package sequenceutils

import (
	"testing"
)

func TestIsIn(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	// a[-8:]
	result := Slice(a, -8).([]string)
	t.Logf("%#v", result)
	// a[3:]
	result = Slice(a, 3).([]string)
	t.Logf("%#v", result)
	// a[:-1][::-1]
	result = Slice(a, 0, -1, -1).([]string)
	t.Logf("%#v", result)
	b := "123456789"
	// b[1:-2][::-2]
	rs := Slice(b, 1, -2, -2).(string)
	t.Logf("%#v", rs)
	// b[1:-2][::3]
	rs = Slice(b, 1, -2, 3).(string)
	t.Logf("%#v", rs)
}
