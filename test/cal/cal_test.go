package cal

import (
	"testing"
)

func TestGetSum(t *testing.T) {
	res := add(10)
	if res != 55 {
		t.Fatalf("addUpper(10) is 55, error res=%v", res)
	}
	t.Logf("test GetSum success")
}
