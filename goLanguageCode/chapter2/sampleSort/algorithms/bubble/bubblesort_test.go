package bubble

import "testing"

func TestBubbleSort(t *testing.T) {
	values := []int{65, 15, 97, 32, 41, 26}
	BubbleSort(values)
	if values[0] != 97 || values[1] != 65 {
		t.Error("BubbleSort() failed.Got", values, "Expected 97 65 41 32 26 15")
	}
}

func TestBubbleSort1(t *testing.T) {
	values := []int{65, 65, 97, 32, 41, 26}
	BubbleSort(values)
	if values[0] != 97 || values[1] != 65 {
		t.Error("BubbleSort() failed.Got", values, "Expected 97 65 65 41 32 26")
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{65}
	BubbleSort(values)
	if values[0] != 65 {
		t.Error("BubbleSort() failed.Got", values, "Expected 65")
	}
}
