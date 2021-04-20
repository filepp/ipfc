package encoding

import "testing"


func TestEncodeIndex(t *testing.T) {
	indexes := []int{2, 1, 14, 7, 8, 15, 100}
	data := EncodeIndex(indexes)
	dIndexes := DecodeIndex(data)

	t.Logf("%v\n", indexes)
	t.Logf("%v\n", dIndexes)

	for i, v := range dIndexes {
		if v != indexes[i] {
			t.Errorf("aaaaaaaaaa")
		}
	}
}
