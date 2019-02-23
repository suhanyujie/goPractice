package simplemath

import "testing"

func TestAdd(t *testing.T) {
	r := Add(1, 3)
	expectVal := 4
	if r != expectVal {
		t.Errorf("Add(1,2) failed.Got %d,expect %d.", r, expectVal)
	}
}
