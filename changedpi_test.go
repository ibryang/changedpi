package changedpi

import (
	"testing"
)

func TestChangeDpi(t *testing.T) {
	output, err := ChangeDpiByPath("image/go_72.jpeg", 300)
	if err != nil {
		t.Error(err)
	}
	err = SaveImage("image/go_300.jpeg", output)
	if err != nil {
		t.Error(err)
	}
}
