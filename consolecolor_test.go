package golg

import "testing"

func TestNewColorHSV(t *testing.T) {
	expected := NewColorRGB(0x31, 0x70, 0x21)
	actual := NewColorHSV(108, 71, 44)

	if *expected != *actual {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
}
