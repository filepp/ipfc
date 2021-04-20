package encoding

import (
	"math/big"
	"sort"
)

func Shirt(i int) *big.Int {
	b := big.NewInt(1)
	for ; i > 0; i-- {
		b = b.Mul(b, big.NewInt(2))
	}
	return b
}

func EncodeIndex(indexes []int) []byte {
	sort.Ints(indexes)
	a := big.NewInt(0)
	preIndex := 0
	preShirt := big.NewInt(1)
	for _, i := range indexes {
		b := big.NewInt(0).Mul(preShirt, Shirt(i-preIndex))
		a = a.Add(a, b)
		preIndex = i
		preShirt = b
	}
	return a.Bytes()
}

func DecodeIndex(b []byte) []int {
	a := make([]int, 0)
	for i := len(b) - 1; i >= 0; i-- {
		v := b[i]
		for j := 0; j < 8; j++ {
			if v&(1<<j) > 0 {
				a = append(a, (len(b)-i-1)*8+j)
			}
		}
	}
	return a
}
