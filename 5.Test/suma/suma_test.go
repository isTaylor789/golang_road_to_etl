package suma

import "testing"

func TestSum(t *testing.T) {
	want := 44

	got := Add(12, 32)

	if got != int32(want) {
		t.Logf("Se esperaba: %d, se obtuvo: %d", want, got)
		t.Fail()
	}

}

func TestSumList(t *testing.T) {
	want := 55

	got := AddList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if got != want {
		t.Logf("Se esperaba: %d, se obtuvo: %d", want, got)
		t.Fail()
	}

}
