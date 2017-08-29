package utils

import "math/big"

// create an array filled with b
// ArrayOfBytes(3,0) => [0 0 0]
func ArrayOfBytes(i int, b byte) (p []byte) {
	for i != 0 {
		p = append(p, b)
		i--
	}
	return
}

// Join bigInt, and padding it when less than `expectedLen`.
func BigIntJoin(expectedLen int, bigs ...*big.Int) *big.Int {
	bs := []byte{}
	for i, b := range bigs {
		by := b.Bytes()
		dif := expectedLen - len(by)
		if dif > 0 && i != 0 {
			by = append(ArrayOfBytes(dif, 0), by...)
		}
		bs = append(bs, by...)
	}
	b := new(big.Int).SetBytes(bs)

	return b
}

// Split bigInt into `parts` parts.
func SplitBig(b *big.Int, parts int) []*big.Int {
	bs := b.Bytes()
	if len(bs)%2 != 0 {
		bs = append([]byte{0}, bs...)
	}
	l := len(bs) / parts
	as := make([]*big.Int, parts)
	for i := range as {
		as[i] = new(big.Int).SetBytes(bs[i*l : (i+1)*l])
	}
	return as

}
